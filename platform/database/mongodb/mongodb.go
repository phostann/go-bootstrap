package mongodb

import (
	"context"
	"fmt"
	"log"
	"shopping-mono/pkg/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	DB *mongo.Client
}

func New(cfg configs.Config) (*MongoDB, func()) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?maxPoolSize=%d", cfg.MongoDB.User, cfg.MongoDB.Password, cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.MaxPoolSize)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("connect to mongodb failed: %v", err)
	}
	log.Println("connected to mongodb!!!")
	cleanup := func() {
		client.Disconnect(context.TODO())
	}
	return &MongoDB{
		DB: client,
	}, cleanup

}
