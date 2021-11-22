package repositories

import (
	"context"
	"time"

	"github.com/jonioliveira/go-template/internals/core/ports"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MongoClientTimeout = 5
)

type UserRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

var _ ports.UserRepository = (*UserRepository)(nil)

func NewUserRepository(conn string) (*UserRepository, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		conn,
	))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return &UserRepository{
		client:     client,
		database:   client.Database("goTemplateHttp"),
		collection: client.Database("goTemplateHttp").Collection("users"),
	}, nil
}

func (r *UserRepository) Login(email string, password string) error {
	// Here your code for login in mongo database
	return nil
}

func (r *UserRepository) Register(email string, password string) error {
	// Here your code for save in mongo database
	return nil
}
