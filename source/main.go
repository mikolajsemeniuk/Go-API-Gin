package main

import application "Go-API-Gin/source/startup"

func main() {
	application.Route().Listen(":4000")
}
