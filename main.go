package main

import (
	"fmt"
	"log"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	plays := Plays{}
	plays, err := GetPlaysFromUser("Sebki")
	checkerr(err)
	fmt.Println(plays)
}
