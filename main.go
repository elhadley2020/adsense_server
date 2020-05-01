package main

import (
	"net/http"
	"io"
	"log"
)

func main(){
	home := func(w http.ResponseWriter, req *http.Request){
		io.WriteString(w,
		"<!doctype html><html><head><title>Title</title><script data-ad-client='ca-pub-8840478853854273' async src='https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js'></script></head><body><h1>Header 1</h1><p>Paragraph 1</p></body></html>")
		}
	
	http.HandleFunc("/", home)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
	
	
}
