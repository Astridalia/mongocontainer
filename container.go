package mongocontainer

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo contains the mongo client and database instances
type Mongo struct {
	client *mongo.Client
	db     *mongo.Database
}

// MongoImpl defines the interface for interacting with MongoDB
type MongoImpl interface {
	FindOne(collection string, filter interface{}) *mongo.SingleResult
	Upsert(collection string, filter interface{}, update interface{}) *mongo.SingleResult
	FindOneAndDelete(collection string, filter interface{}) *mongo.SingleResult
	Disconnect() error
}

// FindOne finds a single document that matches the filter in the given collection
func (m *Mongo) FindOne(collection string, filter interface{}) *mongo.SingleResult {
	return m.db.Collection(collection).FindOne(context.Background(), filter)
}

// Upsert updates a document if it already exists or inserts it otherwise
func (m *Mongo) Upsert(collection string, filter interface{}, update interface{}) *mongo.SingleResult {
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)
	return m.db.Collection(collection).FindOneAndUpdate(context.Background(), filter, update, opts)
}

// FindOneAndDelete deletes a document that matches the filter in the given collection
func (m *Mongo) FindOneAndDelete(collection string, filter interface{}) *mongo.SingleResult {
	return m.db.Collection(collection).FindOneAndDelete(context.Background(), filter)
}

// Disconnect disconnects the mongo client
func (m *Mongo) Disconnect() error {
	return m.client.Disconnect(context.Background())
}

// Setup connects to the MongoDB database and returns a Mongo instance
func Setup(uri string, database string) (MongoImpl, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("error while creating mongo client: %v", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	db := client.Database(database)
	return &Mongo{client: client, db: db}, nil
}
