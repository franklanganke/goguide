package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.de")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}
	c, _ := io.Copy(lw, resp.Body)
	fmt.Println("", c)
}

func (logWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	return len(b), nil
}
