package main

import (
		"fmt"
		"log"
		"net/http"
		"os"
)

const(
	usage = "usage: ./main v1|v2"
	v1 = "v1"
	v2 = "v2"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal(usage)
	}
	version := os.Args[1]
	switch version {
	case v1: {
		http.HandleFunc("/", HelloServerV1)
	}
	case v2: {
		http.HandleFunc("/", HelloServerV2)
	}
	default: {
		log.Fatal(usage)
	}
	}
	http.ListenAndServe(":8080", nil)
}

func HelloServerV1(w http.ResponseWriter, r *http.Request) {
	version := 1
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "pretty cool app -- version %d\n", version)
	fmt.Fprintf(w, "pod: %s\n", hostname)
}

func HelloServerV2(w http.ResponseWriter, r *http.Request) {
	version := 2
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "coolest app ever -- version %d\n", version)
	fmt.Fprintf(w, "pod: %s\n", hostname)
}
