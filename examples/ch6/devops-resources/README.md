# DevOps resource struct module

#### This module is used and consumed by the devops rest api and custom terraform provider for bootcamp section 6.1.4. Its purpose is to organize the structures of the resources defined in this bootcamp section in order to have less confusion with how to set up the client-side communication for the terraform custom provider.

## How to import module for the resources built in go:

#### Here is an example of how to include the module inside of the go.mod within the terraform provider:
```
require github.com/liatrio/devops-bootcamp/examples/ch6/devops-resources devops-api(TODO: this will be changed to the default branch before being pushed up)    
```

#### Now that you have the right path targeted for the import you can now use this module for any code corresponding to communicating the terraform provider with the devops api.

#### Example in how the import is call in your go code:
```
import "github.com/liatrio/devops-bootcamp/examples/ch6/devops-resources"
```

## How to call package devops_resource:

#### To access struct type that is defined in the module you call it like this:

##### Engineer Struct:
```
devops_resource.Engineer
```

##### Dev Struct:
```
devops_resource.Dev
```

##### Ops Struct:
```
devops_resource.Ops
```

##### DevOps Struct:
```
devops_resource.DevOps
```

## Notes: 
- The structs use references of another struct object to manage lists.
    - This allows easier management of the lists when updating a resource.

#### Struct Example:
```
type DevOps struct {
        Id  string `json:"id"`
        Devs []*Dev `json:"dev"`
        Ops []*Ops `json:"ops"`
}
```

##### DevOps Struct is a good example of this since it manages a list of type Dev and Ops.
