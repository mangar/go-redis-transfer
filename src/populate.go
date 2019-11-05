package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

func main() {

	args := os.Args[1:]
	i, _ := strconv.Atoi(args[0])

	fmt.Println(args)

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	for n := 0; n <= i; n++ {
		fmt.Println("- ", n)
		err := client.Set(RandStringRunes(10), "VALUE"+RandStringRunes(10), 0).Err()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("DONE")
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
