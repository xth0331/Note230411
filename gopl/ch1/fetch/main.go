package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//	func main() {
//		for _, url := range os.Args[1:] {
//			resp, err := http.Get(url)
//			if err != nil {
//				fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
//				os.Exit(1)
//			}
//			if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
//				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
//				os.Exit(1)
//			}
//		}
//	}
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = strings.Join([]string{"http://", url}, "")
			fmt.Println("Modified string:", url)
			fmt.Println()
		}
		resp, err := http.Get(url)
		fmt.Println("状态码为:", resp.StatusCode)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
