package gateway

type Repository interface {
	Connect() error
	CreateDocuments(dbname string, collectionName string, documents []map[string]interface{}) (interface{}, error)
	DropCollection(dbName string, collectionName string) error
}
