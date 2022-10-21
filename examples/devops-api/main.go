package main

import(
	"os"
	"fmt"
	"net/http"
	"crypto/rand"
	"encoding/base32"
)

func main() {
	protocol := "http://"
	url := protocol + "www.google.com"
	response, err := http.Get(url)
	if err != nil { //if theres an error in GET request
		fmt.Printf("error in GET request: %s\nErrors: %s\n",url, err)
		os.Exit(1)
	}
	fmt.Println(response)

	// random number generator	
	//rand.Seed(time.Now().UnixNano())
	//min := 10
	//max := 30


	//fmt.Println(rand.Intn(max - min + 1) + min)

	randId := getRandId(15)
	fmt.Println("Random ID: ", randId)

}

func getRandId(length int) string {
    randomBytes := make([]byte, 32)
    _, err := rand.Read(randomBytes)
    if err != nil {
        panic(err)
    }
    return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}
