package main

import (
	"github.com/ANANDHA007/LocaSync/config"
	"github.com/ANANDHA007/LocaSync/core"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	//Test Data

	config := config.Config{
		Store: "InMemoryStore",
	}

	LocaSync, err := core.New(config)

	if err != nil {
		log.Info(err)
	}

	LocaSync.Set("Ak", "Anandh the programmer", "123")
	LocaSync.Get("Ak")

}
