package main

import (
			"fmt"
			"bufio"
			"net/http"
			"log"
)

func main() {

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([][]string, 12)

	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, len(buckets))
		buckets[n] = append(buckets[n], word)
	}

	for i := range buckets {
		fmt.Println(len(buckets[i]))
	}
}

func HashBucket(word string, bucketSize int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return (sum + len(word)) % bucketSize
}