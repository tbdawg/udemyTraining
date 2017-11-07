package main

import (
	"fmt"
	"net/http"
	"bufio"
	"os"
	"log"
)

func main() {
	result, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(result.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = "\tdef: "
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0

	for k, v := range words { // k, _ redundant
		fmt.Print(k)
		fmt.Println(v)
		if i == 200 {
			break
		}
		i++
	}
}
