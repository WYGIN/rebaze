package main

import (
	"github.com/WYGIN/rebaze/cli/internal/bazer"
)

func main() {
	if err := bazer.ReBaze(); err != nil {
		panic(err)
	}
}
