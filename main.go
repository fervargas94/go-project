package main

import (
	handlers "github.com/fervargas94/proxy-app/api/handlers"
	"github.com/fervargas94/proxy-app/api/utils"
	"github.com/fervargas94/proxy-app/api/server"
)

func main() {
	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
	
}