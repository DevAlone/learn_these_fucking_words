package main

import (
	"./initialization"
	"./models"
	"./scheduler"
	"./server"
	"flag"
	"fmt"
)

func main() {
	if err := models.InitDb(); err != nil {
		panic(err)
	}

	var initFlag string
	flag.StringVar(&initFlag, "init", "", "[db] - run initialization script")

	flag.Parse()

	fmt.Println(flag.NFlag())

	if len(initFlag) > 0 {
		fmt.Println("start initialization")

		if err := initialization.InitByFlag(initFlag); err != nil {
			panic(err)
		}
	}

	if flag.NFlag() > 0 || flag.NArg() > 0 {
		return
	}

	if err := scheduler.Init(); err != nil {
		panic(err)
	}

	if err := server.Run(); err != nil {
		panic(err)
	}

	if err := models.DB.Close(); err != nil {
		panic(err)
	}
}
