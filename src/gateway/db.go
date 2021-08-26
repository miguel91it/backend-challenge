package gateway

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	Connect() error
	CreateDocuments(dbname string, collectionName string, documents []map[string]interface{}) (interface{}, error)
	GetDocsByFilter(dbname string, collectionName string, filters bson.M) ([]bson.M, error)
	DropCollection(dbName string, collectionName string) error
}
