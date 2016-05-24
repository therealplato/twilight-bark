package main

import (
	"os"
	"time"
)

func main() {
	tone(60)
}

func tone(hz int64) {
	// Timer needs period
	// F := fmt.sprintf("%ds", hz)
	t := 1000000000 * time.Nanosecond // 1s
	c := time.Tick(t)
	for {
		<-c
		os.Stdout.Write([]byte("."))
	}
}
