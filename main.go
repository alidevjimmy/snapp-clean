package main

import "github.com/alidevjimmy/snapp-clean/cmd"

func main() {
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
