package main

import (
	"fmt"
	"go_web_carpentry/router"
)

func main() {
	// Router
	r := router.Router()

	fmt.Println("Start Echo server...")

	// Start server
	r.Logger.Fatal(r.Start(":3100"))

}