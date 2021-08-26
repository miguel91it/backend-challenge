package goliothMongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Connect() error
	CreateDocuments(dbname string, collectionName string, documents []map[string]interface{}) (interface{}, error)
	GetDocsByFilter(dbname string, collectionName string, filters bson.M) ([]bson.M, error)
	DropCollection(dbName string, collectionName string) error
}

type MongoClient struct {
	client *mongo.Client
	IP     string
	Port   int
}

func NewMongoClient(ip string, port int) *MongoClient {
	return &MongoClient{IP: ip, Port: port}
}

func (mc *MongoClient) Connect() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	dbHost := fmt.Sprintf("mongodb://%s:%v", mc.IP, mc.Port)

	clientOptions := options.Client().ApplyURI(dbHost)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return err
	}

	mc.client = client

	fmt.Println("Application connected to the database.")

	return nil
}

func (db MongoClient) CreateDocuments(dbname string, collectionName string, documents []map[string]interface{}) (interface{}, error) {

	var data []interface{}

	for _, doc := range documents {
		data = append(data, doc)
	}

	collection := db.client.Database(dbname).Collection(collectionName)

	res, err := collection.InsertMany(context.Background(), data)
	if err != nil {
		return nil, err
	}

	ids := res.InsertedIDs

	return ids, nil

}

func (db MongoClient) GetDocsByFilter(dbname string, collectionName string, filters bson.M) ([]bson.M, error) {
	collection := db.client.Database(dbname).Collection(collectionName)
	// opts := options.Find().SetSort(bson.D{) // Ordenar dados

	cursor, err := collection.Find(context.Background(), filters)

	if err != nil {
		fmt.Println("error getting db: ", err)
		return []bson.M{}, err
	}

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (db MongoClient) DropCollection(dbName string, collectionName string) error {

	collection := db.client.Database(dbName).Collection(collectionName)

	err := collection.Drop(context.TODO())

	if err != nil {
		return err
	}

	return nil
}
