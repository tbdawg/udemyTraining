package main

import "fmt"

func main() {
	n := hashBucket("Go", 12)
	fmt.Println(n)
}

// get the bucket that the word belongs in via the
// modulus of the number of bucket available
func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
