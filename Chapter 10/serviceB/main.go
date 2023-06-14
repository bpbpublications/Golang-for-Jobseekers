package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	serviceAURL := os.Getenv("SERVICE_A_URL")
	loopInterval := os.Getenv("LOOP_INTERVAL")
	loopIntervalInt, err := strconv.Atoi(loopInterval)
	if err != nil {
		loopIntervalInt = 5
	}

	for {
		data := map[string]float64{
			"revenue": rand.Float64() * 100.0,
		}
		raw, _ := json.Marshal(data)
		fmt.Printf("sending the following data: %v\n", string(raw))
		resp, err := http.Post(serviceAURL, "application/json", bytes.NewBuffer(raw))
		if err != nil {
			fmt.Println("unable to contact service a url")
		}
		if resp.StatusCode != 200 {
			fmt.Println("unable to save record")
		}
		time.Sleep(time.Duration(loopIntervalInt) * time.Second)
	}
}
