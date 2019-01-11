package main

import (
	"LoadBalance/route"
)

func main() {
	engine := route.RegisterRouter()
	engine.RunTLS("tws2sg1.dreamtech8.com", "./STAR.dreamtech8.com.crt", "./STAR.dreamtech8.com.key")
}
