package main

import (
	"goproject/app"
	"goproject/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8116")
}
