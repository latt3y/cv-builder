package main

import (
  "fmt"
  "net/http"
  "io"
  "errors"
  "os"
)

func get_root(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("got / request");
  io.WriteString(w, "This is my website!\n");
}

func get_hello(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("get /hello request");
  io.WriteString(w, "Hello HTTP");
}

func main() {
  http.HandleFunc("/", get_root);
  http.HandleFunc("/hello", get_hello);

  err := http.ListenAndServe(":3333", nil);

  if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
