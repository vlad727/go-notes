package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/rand"
	"time"
)

func main() {

	for {
		time.Sleep(5 * time.Second)
		randString := rand.String(6)
		fmt.Println(randString)
	}
}
