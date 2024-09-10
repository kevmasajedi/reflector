package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)


func handler(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	fmt.Fprintf(w, ip)
}
func main() {
	http.HandleFunc("/", handler)
	port := "8011"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	fmt.Printf("Reflector is running on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
