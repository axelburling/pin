/*
Copyright © 2023 Axel Burling axel.burling@gmail.com
*/
package main

import (
	"math/rand"
	"time"

	"github.com/axelburling/pin/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}
