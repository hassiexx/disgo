package main

import (
	"context"
	"fmt"

	"github.com/hassieswift621/disgo/gateway"
)

func main() {
	s := gateway.NewSession(1, 1, "")
	err := s.Open(context.Background())
	fmt.Println(err)
	select {}
}
