// https://medium.com/@matryer/make-ctrl-c-cancel-the-context-context-bd006a8ad6ff
package main

import (
	"context"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		select {
		case <-sig:
			cancel()
		case <-ctx.Done():
		}
		signal.Stop(sig)
	}()

	// go func() {
	// 	time.Sleep(time.Second)
	// 	cancel()
	// }()

	boredSleep(ctx)
}

func boredSleep(ctx context.Context) {
	time.Sleep(time.Second * 10)
}
