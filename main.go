package main

import (
	"LobbyService/route"
)

func main() {
	engine := route.RegisterRouter()
	engine.Run(":8888")
}
