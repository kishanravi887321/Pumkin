package repository

import (
	"context"
	"fmt"

	"github.com/kishanravi887321/pumkin/server-side/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(
	collection *mongo.Collection,
) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	var user models.User

	err := r.collection.FindOne(
		ctx,
		bson.D{{Key: "email", Value: email}},
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(
	ctx context.Context,
	user *models.User,
) error {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	fmt.Println("Inserted user :", result)
	// if id, ok := result.InsertedID.(bson.ObjectID); ok {
	// 	user.UserID = id.Hex()
	// }

	return nil
}
