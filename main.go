package main

import (
	"fmt"
	"net/http"
)

var (
	websites []string
)

func init() {
	websites = append(websites, "https://www.blic.rs")
	websites = append(websites, "https://www.kupujemprodajem.com")
	websites = append(websites, "https://www.mondo.rs")
	websites = append(websites, "hta.com")
	websites = append(websites, "https://www.krstarica.com")
	websites = append(websites, "https://bloxico.com")
	websites = append(websites, "https://www.lossajt")
	websites = append(websites, "https://www.n1info.rs")
	websites = append(websites, "https://go.dev/tour/concurrency/2")
	websites = append(websites, "https://www.tehnomanija123456789.rs/")
}

func main() {
	ch := make(chan string)
	go checkWebsite(ch)
	for valid := range ch {
		fmt.Println(valid)
	}

}

func checkWebsite(ch chan string) {
	for _, website := range websites {
		res, err := http.Get(website)
		if err != nil {
			ch <- fmt.Sprintf("%v is unavailable, %v", website, err)
		} else {
			ch <- fmt.Sprintf("%v is available, status %v", website, res.Status)
		}
	}
	close(ch)
}
