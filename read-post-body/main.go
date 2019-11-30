package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type timeoutHandler struct{}

func (h timeoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bodyBytes := make([]byte, 0)
	z := bytes.NewBuffer(bodyBytes)
	for {
		_, err := io.CopyN(z, r.Body, 256)
		// bodyBytes = append(bodyBytes, z.Bytes())

		if err == io.EOF {
			break
		} else if err != nil {
			break
		}
	}

	fmt.Printf("%v; %q; %T; %v;\n", z, z, z, bodyBytes)
}

func main() {
	h := timeoutHandler{}
	s := &http.Server{
		ReadHeaderTimeout: 20 * time.Second,
		Handler:           h,
		Addr:              ":8080",
	}
	if err := s.ListenAndServe(); err != nil {
		fmt.Println("🤷")
	}
}
