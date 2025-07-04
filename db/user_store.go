package db

import (
	"context"
	"fmt"

	"github.com/Ekireh-source/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



const userColl = "users"

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
 	UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error
}

type MongoUserStore struct {
	client *mongo.Client
	dbName string
	coll *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, dbname string) *MongoUserStore {
	
	return &MongoUserStore{
		client: client,
		coll : client.Database(dbname).Collection(userColl),
	}
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("------Dropping user collection------")
	return s.coll.Drop(ctx)
}


func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error {
	update := bson.D{
		{Key: "$set", Value: params.TOBSON()},
	}

	_, err := s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
		}
	return nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	//TODO: Handle if user was not deleted
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err

	}

	return nil
}


func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}


func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {

	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*types.User

	if err := cur.All(ctx, &users); err != nil {
		return []*types.User{}, nil
	}
	return users, nil
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	

	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	
	return &user, nil
}