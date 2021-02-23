package main

import (
	"fmt"
	"github.com/docker/distribution/uuid"
)

func main() {
	UUID := uuid.Generate()
	fmt.Println(UUID.String())
}
