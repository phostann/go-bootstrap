package mysql

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"shopping-mono/pkg/configs"
	"shopping-mono/platform/database/mysql/ent"
)

type Queries struct {
	DB *ent.Client
}

func New(cfg *configs.Config) (*Queries, func()) {
	client, err := ent.Open(cfg.Database.Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DB))
	if err != nil {
		log.Fatalf("connect db failed: %v", err)
	}

	// auto migration
	if cfg.Database.Migration == "auto" {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	cleanup := func() {
		err := client.Close()
		if err != nil {
			log.Printf("close db connection failed: %v", err)
		}
	}

	return &Queries{
		DB: client,
	}, cleanup
}
