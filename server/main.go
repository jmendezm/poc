package server

import (
	"technical_test/api_server"
	"technical_test/postgres_connection"
)

func Start() {
	postgres_connection.Init()
	api_server.Init()
}
