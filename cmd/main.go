package main

import (
	"fmt"

	"dddemo/server"
)

func main() {
	fmt.Println("i am alive")

	srv := server.NewServer()
	srv.Run()

}
