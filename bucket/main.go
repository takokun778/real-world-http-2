package main

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	BucketSize := 10
	ctx := context.Background()
	e := rate.Every(time.Second / 10)
	l := rate.NewLimiter(e, BucketSize)
	for _, task := range tasks {
		err := l.Wait(ctx)
		if err != nil {
			panic(err)
		}
		// TODO: task
	}
}
