package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	h := &Multiplier{ControlFile: "./control.file"}
	srv := &http.Server{
		Addr:              ":8888",
		Handler:           h,
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

	go func() {
		for {
			h.Control()
			time.Sleep(time.Second * 10)
		}
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

type Multiplier struct {
	ControlFile        string
	AmplificationLevel int
}

func (m *Multiplier) Control() {
	fd, err := os.Open(m.ControlFile)
	if err != nil {
		log.Println("can not open control file:", err)
		return
	}
	defer fd.Close()

	b := bufio.NewScanner(fd)
	b.Scan()
	l, err := strconv.Atoi(b.Text())
	if err != nil {
		log.Println("can not read control file:", err)
		return
	}
	m.AmplificationLevel = l
}

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

	for i := 0; i < m.AmplificationLevel; i++ {
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
