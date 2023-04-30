package main

import (
	"github.com/chat/model"
	"github.com/chat/router"
	"github.com/chat/service"
)

func main() {
	model.Setup()
	go service.Manager.Start()
	router.RouterSetup()
}
