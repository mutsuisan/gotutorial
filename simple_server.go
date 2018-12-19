package main

import "net/http"


func main() {
  http.HandleFunc("/", someFunc)
  http.HandleFunc("/test/", otherFunc)
  http.ListenAndServe(":8080", nil)
} 

func someFunc(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("Hello World!!\n This is a first Go app!"))
}

func otherFunc(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("Hello I'm in second page!"))
}
