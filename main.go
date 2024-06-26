package main

import "api/publicFileApi"

func main() {
	server := publicFileApi.NewApiServer(":8080")
	server.Run()
}