package main

import (
	"fmt"

	"github.com/mrdomino/bzlmod-issue/api"
)

func main() {
	req := &api.HelloRequest{
		Name: "World",
	}
	fmt.Printf("Hello %s", req.Name)
}
