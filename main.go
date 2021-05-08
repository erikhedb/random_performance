package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	start := time.Now()

	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := make(map[int]int)
	for len(m) < (10 * 1000 * 1000) {
		var randValue = rand.Intn(20 * 1000 * 1000)
		m[randValue] += 1
	}

	elapsed := time.Since(start)
	fmt.Printf("Random generation took %s\n", elapsed)

	for key := range m {
		_, err2 := f.WriteString(fmt.Sprintf("%d,", key))
		if err2 != nil {
			log.Fatal(err2)
		}

	}

	elapsed = time.Since(start)
	fmt.Printf("File operation took %s\n", elapsed)

}
