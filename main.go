package main

import (
	"Migration/config"
	"Migration/router"
)

func main() {
	config.Init()

	r := router.SetupRouter()
	r.Run(":8080")
}
