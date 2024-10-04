package main

import (
	"fmt"
	"log"
	"os"
	"simple-template/server/initial"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	g := initial.InitGin()
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	var port uint16
	parsed, err := strconv.ParseUint(os.Getenv("BIND_PORT"), 10, 16)
	if err != nil {
		port = 3000
		log.Printf("Error during get app port: %v", err)
	} else {
		port = uint16(parsed)
	}

	bindAddress := fmt.Sprintf("127.0.0.1:%v", port)
	g.Run(bindAddress)
}
