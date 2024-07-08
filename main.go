package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://somenode-ecommerse-server.onrender.com",
		"https://aapan.shop/",
		"https://deleted-files-90991.onrender.com",
	}

	for _, url := range urls {
		go func(url string) {
			for {
				GetReq(url)
				fmt.Println("I will send it every 5 minutes")
				time.Sleep(5 * time.Minute)
			}
		}(url)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is healthy")
	})

	http.ListenAndServe(":8080", nil)
}

func GetReq(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
