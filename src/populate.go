package main

import (
	"fmt"
	"my"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	i, _ := strconv.Atoi(args[0])

	fmt.Println(args)

	client := my.RedisClient("localhost:6379")

	for n := 0; n <= i; n++ {
		fmt.Println("- ", n)
		err := client.Set(my.RandStringRunes(10), "VALUE"+my.RandStringRunes(10), 0).Err()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("DONE")
}
