package main

import (
	"os"

	"github.com/yodfhafx/go-ecommerce/config"
	"github.com/yodfhafx/go-ecommerce/modules/servers"
	"github.com/yodfhafx/go-ecommerce/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())
	// fmt.Println(cfg.App())
	// fmt.Println(cfg.Db())
	// fmt.Println(cfg.Jwt())
	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
