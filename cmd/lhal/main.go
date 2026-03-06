package main

import (
	"cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}