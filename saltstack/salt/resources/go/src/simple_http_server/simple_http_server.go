package main

import (
    "fmt"
    "net/http"
    "strings"
)

func handler_old(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handler(w http.ResponseWriter, r *http.Request) {
 // Create return string
 var request []string

 // Add the request string
 url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
 request = append(request, url)

 // Add the host
 request = append(request, fmt.Sprintf("Host: %v", r.Host))

 // Loop through headers
 for name, headers := range r.Header {
   name = strings.ToLower(name)
   for _, h := range headers {
     request = append(request, fmt.Sprintf("%v: %v", name, h))
   }
 }
 
 // If this is a POST, add post data
 if r.Method == "POST" {
    r.ParseForm()
    request = append(request, "\n")
    request = append(request, r.Form.Encode())
 }      
     fmt.Fprintf(w, strings.Join(request, "\n") , r.URL.Path[1:])
}



func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":9000", nil)
}
