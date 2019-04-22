package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Hello world")
	out, err := exec.Command("go run ./main.go").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}


