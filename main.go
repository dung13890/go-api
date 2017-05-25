package main

import (
	"go-api/routers"
)

func main() {
	routers.InitRoutes().Run(":8000")
}
