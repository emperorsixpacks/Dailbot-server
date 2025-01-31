package main

import (
	"fmt"

	"github.com/emperorsixpacks/dailbot/server/config"
)

func main() {
  fmt.Println(config.GetConfig())
}
