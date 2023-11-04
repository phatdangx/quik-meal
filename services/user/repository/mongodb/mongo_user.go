package mongodb

import (
	"context"
	c "user/constants"
	"user/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepo struct {
	col *mongo.Collection
}

func NewUserMongoRepo(client *mongo.Client) *mongoUserRepo {
	return &mongoUserRepo{
		client.Database(c.MongoDatabaseName).Collection(c.MongoUserCollection),
	}
}

func (m *mongoUserRepo) CreateNewUser(ctx context.Context, user *models.User) error {
	_, err := m.col.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
