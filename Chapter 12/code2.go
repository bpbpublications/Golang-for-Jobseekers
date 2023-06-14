package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	zzz = make(chan string)
)

func generateData(interval int, ch chan string) {
	for {
		time.Sleep(time.Duration(interval) * time.Second)
		val := rand.Int()
		ch <- fmt.Sprintf("time: %v :: interval: %d :: num: %v", time.Now(), interval, val)
	}
}

func printer(ch chan string) {
	f, _ := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	bw := bufio.NewWriter(f)
	i := 0
	for {
		bw.WriteString(<-ch + "\n")
		i = i + 1
		if i >= 10 {
			bw.Flush()
			i = 0
			fmt.Println("data flushed into file")
		}
	}
}

func main() {
	go generateData(1, zzz)
	go generateData(2, zzz)

	printer(zzz)
}
