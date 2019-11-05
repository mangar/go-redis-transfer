package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func main() {

	args := os.Args[1:]
	portFrom := args[0]
	portTo := args[1]

	fmt.Println("Params: ", args)

	fromServer := "localhost:"
	fromServer += portFrom

	clientFrom := redis.NewClient(&redis.Options{
		Addr:     fromServer,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	toServer := "localhost:"
	toServer += portTo

	clientTo := redis.NewClient(&redis.Options{
		Addr:     toServer,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	keys, _ := clientFrom.Keys("*").Result()
	l := len(keys)

	fmt.Println("Tamanho: ", l)

	timer1 := time.NewTimer(time.Second)

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		value, _ := clientFrom.Get(key).Result()
		fmt.Println("- ", key, "::", value)

		err := clientTo.Set(key, value, 0).Err()
		if err != nil {
			panic(err)
		}
	}

	timer2 := time.NewTimer(time.Second)

	// fmt.Println(keys)
	fmt.Println("From:", fromServer, " To:", toServer)
	fmt.Println("Tamanho: ", l)
	//

	fmt.Println("DONE")
}
