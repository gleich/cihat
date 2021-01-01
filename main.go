package main

import (
	"fmt"

	"github.com/Matt-Gleich/cihat/pkg/api"
)

func main() {
	client := api.CreateClient()
	out := api.GetChecks(client)
	fmt.Println(out)
}
