package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)


type Doggie struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main(){
	// This function will use the  Dog breeds API to specifically call a random dog
	// Was originally just using fmt.print instead of log, but after consultation with edwin, got on the right track
	fmt.Println("Test that API form dog workds")

	// We need first the url that we are going to be hitting
	url := "https://dog.ceo/api/breed/hound/images/random Fetch!"

	// Do we have any added configurations to this request such as a timeout? No then we don't need a client and can execute from our http
	requestGetResponse, err := http.Get(url)

	// Still not a fan of how errors are handled....but ill take it!
	if err != nil {
		log.Fatal(err)
	}

//	 GOING FOR HELP WT=ITH RESPONSE BODIES

}