package main

import (
	"fmt"
	"time"

	_ "github.com/google/go-cmp/cmp" // Add this line
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

type Controller struct {
	kubeclientset kubernetes.Interface
	workqueue     workqueue.RateLimitingInterface
	//Needs a resource specific lister
	//Needs a resource specific informer
	//Pod Example:
	// podsLister    v1.PodLister
	// podsInformer  cache.SharedIndexInformer
}

func NewController(
	kubeclientset kubernetes.Interface,
	podsInformer cache.SharedIndexInformer) *Controller {

	workqueue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "YourResourceHere")

	controller := &Controller{
		kubeclientset: kubeclientset,
		workqueue:     workqueue,
		//Needs a resource specific lister
		//Needs a resource specific informer
		//Pod Example:
		// podsLister:    v1.NewPodLister(podsInformer.GetIndexer()),
		// podsInformer:  podsInformer,
	}

	podsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueue,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueue(new)
		},
		DeleteFunc: controller.enqueue,
	})

	return controller
}

func (c *Controller) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}

func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	err := func(obj interface{}) error {
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		if key, ok = obj.(string); !ok {
			c.workqueue.Forget(obj)
			klog.Errorf("expected string in workqueue but got %#v", obj)
			return nil
		}
		if err := c.syncHandler(key); err != nil {
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		c.workqueue.Forget(key)
		klog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		klog.Errorf("error processing item: %v", err)
		return true
	}

	return true
}

func (c *Controller) syncHandler(key string) error {
	//resource-specific logic for object processing
	return nil
}

func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer runtime.HandleCrash()
	defer c.workqueue.ShutDown()

	klog.Info("Starting Pod controller")

	//Need to run informer as a goroutine here
	//Example for pods:
	//go c.podsInformer.Run(stopCh)

	// if !cache.WaitForCacheSync(stopCh, c.podsInformer.HasSynced) {
	// 	return fmt.Errorf("failed to wait for caches to sync")
	// }

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

func main() {
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeclientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	//Set up informer factory and informer instance
	//Make controller instance
	//Pod Example:

	// informerFactory := informers.NewSharedInformerFactory(kubeclientset, time.Second*30)
	// podsInformer := informerFactory.Core().V1().Pods().Informer()

	// controller := NewController(kubeclientset, podsInformer)

	// stopCh := make(chan struct{})
	// defer close(stopCh)

	// go informerFactory.Start(stopCh)

	// if err = controller.Run(2, stopCh); err != nil {
	// 	klog.Fatalf("Error running controller: %s", err.Error())
	// }

}
