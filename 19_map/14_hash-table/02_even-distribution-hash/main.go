package main

import (
	"fmt"
	"net/http"
	"log"
	"bufio"
)

func main() {
	// get the book Moby Dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	//scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Create slice to hold counts
	buckets := make([]int, 12)

	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), len(buckets))
		buckets[n]++
	}

	fmt.Println(buckets)
}

func HashBucket(word string, buckets int) int {
	var sum int

	for _, v := range word {
		sum += int(v)
	}

	return sum % buckets
}
