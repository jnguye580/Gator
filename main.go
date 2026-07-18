package main

import (
	"fmt"
	"log"

	"github.com/jnguye580/GATOR-PROJECT/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("Jeremi Nguyen")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
