package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	go tone(391, "/")
	go tone(60, " ")
	go tone(2000, `\`)
	quit := new(chan int)
	<-*quit
}

func tone(F_hz int, z string) {
	// F = 1/T    F=frequency T=period
	// 10hz = 1cycle/100ms (10cycles/1s)
	s := int64(1000000000)
	T_ns := int64(float64(s) / float64(F_hz))
	F_ns, err := time.ParseDuration(fmt.Sprintf("%dns", T_ns))
	if err != nil {
		panic(err)
	}
	d := time.Tick(F_ns)
	go func(d <-chan time.Time, z string) {
		for {
			<-d
			os.Stdout.Write([]byte(z))
		}
	}(d, z)
}
