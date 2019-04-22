package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Hello world")
	out, err := exec.Command("bash", "-c","go run main.go").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
	fmt.Printf("The date is %s\n", string(out))

}


