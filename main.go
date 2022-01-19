package main

import (
	"fmt"
	_ "gocloud.dev/secrets"
	_ "gocloud.dev/secrets/hashivault"
)

func main() {
	fmt.Println("Hello world!")
}
