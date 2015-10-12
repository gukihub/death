package main

import (
	"fmt"
)

func worker1() {
	var i int
	for {
		i += 1
	}

}

func worker2() {
	for {
	}
}

func main() {
	go worker1()
	worker2()
	fmt.Println("done")
}
