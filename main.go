package main

import (
	"github.com/kennedysantiagoo/handler"

	_ "github.com/lib/pq"
)

func main() {
	handler.HandleRequests()
}
