package main

import "server/initialize"

func main() {
	router := initialize.Routers()
	router.Run(":9124")
}
