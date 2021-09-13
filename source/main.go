package main

import "Go-API-Gin/source/application"

func main() {
	application.Route().Listen(":4000")
}
