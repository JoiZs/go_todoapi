package main

import (
	"fmt"
	"os"

	"github.com/JoiZs/go_todoapi/api"
	"github.com/JoiZs/go_todoapi/initializer"
)

func main() {
	initializer.Init()
	api := api.NewAPI(fmt.Sprintf(":%s", os.Getenv("PORT")))

	api.Run()
}
