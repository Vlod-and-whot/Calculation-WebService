package main

import (
	"github.com/Vlod-and-whot/Calculation-WebService/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}
