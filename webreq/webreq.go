package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// First, let's define the myGetUrl I'm going to work with. This will be
	// a small site I own.
	myGetUrl := "https://uac.agrohacksstuff.io"

	// We'll make the Get request. Super simple.
	resp, err := http.Get(myGetUrl)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}

	// Assuming we made a valid response, we'll want to make sure that we
	// sever the connection after we're done using it.
	defer resp.Body.Close()

	// Read in the response body using io's ReadAll()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
	}

	// now we can just print the response.
	fmt.Println(string(body))

	// Of course, we can also search for a particular string to see if it
	// exists within the request.

	searchString := "UAC Lookup script"
	if strings.Contains(string(body), searchString) {
		fmt.Printf("%s was found in the body response!\n", searchString)
	} else {
		fmt.Printf("%s was NOT found in the body response.\n", searchString)
	}

	// Now if we want to make a POST request, things get a little more involved.
	// Using the url module, we can create a datatype specifically for the post
	// values and send them along using the http.NewRequest() value, which has
	// a few more options than just straight-up http.Get().
	myPostUrl := "https://uac.agrohacksstuff.io"

	// We can create the form data now using the url.Values() datatype.
	formData := url.Values{
		"key1": {"value1"},
		"key2": {"value2"},
	}

	// Now URL encode the above values to send along the wire
	encodedFormData := formData.Encode()

	// And now create the POST request. Note that this only formulates the request,
	// it doesn't actually make the request.
	postReq, err := http.NewRequest("POST", myPostUrl, strings.NewReader(encodedFormData))
	if err != nil {
		fmt.Printf("Error creating POST request: %v\n", err)
		return
	}

	// Sets the content-type header, which is fairly important for POST requests.
	// Sometimes this will instead be set to application/json
	postReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Now to make the request. First let's establish the client handler.
	client := &http.Client{}

	// Now tell the client to...y'know, do stuff.
	postResp, err := client.Do(postReq)
	if err != nil {
		fmt.Printf("Error making POST request: %v\n", err)
		return
	}
	// don't forget to sever the connection on exit
	defer resp.Body.Close()

	// Now we can read the response body.
	postBody, err := io.ReadAll(postResp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
	}

	fmt.Println(string(postBody))

}
