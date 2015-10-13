package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Printf("usage: %s nb_workers\n", os.Args[0])
}

func worker() {
	for {
	}
}

func main() {

	var nb_workers int
	var err error

	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	if nb_workers, err = strconv.Atoi(os.Args[1]); err != nil {
		usage()
		os.Exit(1)
	}

	for i := 0; i < nb_workers; i++ {
		go worker()
	}
	for {
		time.Sleep(1 * time.Second)
	}

	fmt.Println("done")

}
