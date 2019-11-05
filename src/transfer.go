package main

import (
	"fmt"
	"my"
	"os"
)

func main() {

	args := os.Args[1:]
	portFrom := args[0]
	portTo := args[1]

	fmt.Println("Params: ", args)

	clientFrom := my.RedisClient("localhost:" + portFrom)
	keys, _ := clientFrom.Keys("*").Result()
	l := len(keys)

	clientTo := my.RedisClient("localhost:" + portTo)

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		value, _ := clientFrom.Get(key).Result()
		fmt.Println("- ", key, "::", value)

		err := clientTo.Set(key, value, 0).Err()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("From:", portFrom, " To:", portTo)
	fmt.Println("Tamanho: ", l)
	//

	fmt.Println("DONE")
}
