package main

import (
	"fmt"

	"github.com/smartbot/storage/api"
	"github.com/smartbot/storage/pkg/config"
)

func main() {
	config.LoadConfig()
	r := api.RegisterRoutes()
	r.Run(fmt.Sprintf("%s%d", ":", config.Config.Port))

}
