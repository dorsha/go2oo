package restUtil

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	defaultBodyType = "application/json"
	csrfTokenHeader = "X-CSRF-TOKEN"
)

type cookieJar struct {
	jar map[string][]*http.Cookie
}

func (p *cookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	p.jar[u.Host] = cookies
}

func (p *cookieJar) Cookies(u *url.URL) []*http.Cookie {
	return p.jar[u.Host]
}

// CreateHTTPClient - creates http client
func CreateHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &http.Client{Transport: tr}
}

// Get - HTTP GET for the given URL
func Get(client http.Client, url string, data interface{}, user string,
	password string) {
	resp, err := client.Get(url)
	checkError(err)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	err = json.Unmarshal(respBody, &data)
	checkError(err)
}

// Post - HTTP POST for the given URL
func Post(client http.Client, url string, data interface{}, respData interface{},
	user string, password string) {

	body, err := json.Marshal(&data)
	checkError(err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	checkError(err)

	if len(user) > 0 && len(password) > 0 {
		req.SetBasicAuth(user, password)
	}
	req.Header.Set("Accept", defaultBodyType)
	req.Header.Set("Content-Type", defaultBodyType)
	req.Header.Set(csrfTokenHeader, getCsrfToken(&client, url)) // CSRF token

	resp, err := client.Do(req)
	checkError(err)

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	err = json.Unmarshal(respBody, &respData)

	checkError(err)
}

func getCsrfToken(client *http.Client, url string) string {
	// CSRF support
	jar := &cookieJar{}
	jar.jar = make(map[string][]*http.Cookie)
	client.Jar = jar
	headResponse, err := client.Head(url)
	checkError(err)
	token := headResponse.Header.Get(csrfTokenHeader)
	return token
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
