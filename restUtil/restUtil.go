package restUtil

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// CreateHTTPClient - creates http client
func CreateHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &http.Client{Transport: tr}
}

// Get - HTTP Get for the given URL
func Get(client http.Client, url string, data interface{}) {
	resp, err := client.Get(url)
	checkError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	err = json.Unmarshal(body, &data)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
