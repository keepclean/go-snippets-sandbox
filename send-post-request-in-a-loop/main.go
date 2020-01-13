package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "http://localhost/", "url to send request to")
	var contentType string
	flag.StringVar(&contentType, "content-type", "text/plain", "payload's content type")
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", time.Second*10, "timeout for sending request")
	var filePath string
	flag.StringVar(&filePath, "file", "", "file to send")

	flag.Parse()

	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10,
		},
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// var payload io.ReadSeeker = file

	for {
		sendRequest(client, &url, &contentType, file)
		if _, err := file.Seek(0, 0); err != nil {
			log.Println(err)
			break
		}
		randomSleep(3, 100)
	}
}

func sendRequest(c *http.Client, url, contentType *string, payload io.Reader) {
	body := ioutil.NopCloser(payload)
	request, err := http.NewRequest("POST", *url, body)
	if err != nil {
		log.Println(err)
		return
	}

	request.Header.Set("Content-type", *contentType)

	response, err := c.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if _, err := io.Copy(ioutil.Discard, response.Body); err != nil {
			log.Println(err)
		}
		response.Body.Close()
	}()
}

func randomSleep(t, jitter int) {
	time.Sleep(
		time.Second*time.Duration(rand.Intn(t)) +
			time.Millisecond*time.Duration(rand.Intn(jitter)))
}
