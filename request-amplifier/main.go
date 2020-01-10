package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:              ":8888",
		Handler:           Multiplier{},
		ReadHeaderTimeout: 10 * time.Second,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGINT)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		log.Println("HTTP server is shuting down ...")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

type Multiplier struct{}

func (m Multiplier) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if _, err := io.Copy(ioutil.Discard, r.Body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(http.StatusInternalServerError, err)
		return
	}

	for i := range []int{0, 1, 2} {
		go sendRequest(r, i)
	}

	w.WriteHeader(http.StatusOK)
}

func sendRequest(r *http.Request, i int) {
	httpClient := &http.Client{Timeout: time.Second * 10}
	log.Println("request #", i)
	u := fmt.Sprintf("http://%s%s", r.Host, r.URL.String())
	req, err := http.NewRequest(r.Method, u, nil)
	if err != nil {
		log.Println(http.StatusInternalServerError, err)
		return
	}
	req.Header = r.Header

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if _, err := io.CopyN(ioutil.Discard, resp.Body, 4096); err != nil {
			log.Println("drain response body:", err)
		}
		resp.Body.Close()
	}()

	log.Println("recieved response for attempt #", i, ":", resp.Status)
}
