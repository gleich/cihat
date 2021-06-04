package main

import (
	"time"

	"github.com/gleich/cihat/pkg/api"
	"github.com/gleich/cihat/pkg/lights"
)

func main() {
	client := api.CreateClient()

	for {
		checks := api.GetChecks(client)
		lights.UpdateLights(checks)
		time.Sleep(time.Second * 2)
	}
}
