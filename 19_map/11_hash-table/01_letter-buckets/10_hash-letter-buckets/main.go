package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	buckets := make([]int, 200)
	for scanner.Scan() {
		n := hashbucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
}

func hashbucket(word string) int {
	return int(word[0])
}
