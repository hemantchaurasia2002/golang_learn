package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	evilNinjas := []string{"Shreder", "Cassey", "Dragoz", "Pimp"}

	for _, evilNinjas := range evilNinjas {
		go attack(evilNinjas)
	}

	time.Sleep(time.Second*2)
}

func attack(target string) {
	fmt.Println("Thowing ninja stars at", target)
	time.Sleep(time.Second)
}