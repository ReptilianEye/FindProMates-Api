package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	C *mongo.Collection
}

// func (m *UserModel) All() ([]models.User, error) {
// 	ctx := context.TODO()
// 	// users := []models.{}

// 	cursor, err := m.C.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = cursor.All(ctx, &users)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil

// }
