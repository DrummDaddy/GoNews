package mongo

import (
	"GoNews/pkg/storage"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store представляет собой тип хранилища для MongoDB.
type Store struct {
	client *mongo.Client
	db     *mongo.Database
}

// New создаёт новое хранилище.
func New(connectionString string, dbName string) (*Store, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	return &Store{client: client, db: db}, nil
}

func (s *Store) Posts() ([]storage.Post, error) {
	collection := s.db.Collection("posts")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var posts []storage.Post
	for cursor.Next(context.Background()) {
		var post storage.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (s *Store) AddPost(post storage.Post) error {
	collection := s.db.Collection("posts")
	_, err := collection.InsertOne(context.Background(), post)
	return err
}

func (s *Store) UpdatePost(post storage.Post) error {
	collection := s.db.Collection("posts")
	_, err := collection.UpdateOne(context.Background(), bson.M{"id": post.ID}, bson.M{"$set": post})
	return err
}

func (s *Store) DeletePost(post storage.Post) error {
	collection := s.db.Collection("posts")
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": post.ID})
	return err
}
