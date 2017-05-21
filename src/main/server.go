package main

import (
	"net/http"
	"time"
)

func InitServer() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("0.0.0.0:9595", nil)
	if err != nil {
		log.Fatal(" ", err.Error())
		panic("")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Debugf("[ %s ]\t%s", time.Now().Format(time.RFC1123Z),r.RequestURI)
	w.Write([]byte{ 0x60})
}

