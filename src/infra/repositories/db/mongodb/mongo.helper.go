package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Helper *MongoHelper

type MongoHelper struct {
	credentials *credentials
	client      *mongo.Client
}

type credentials struct {
	database string
	host     string
	port     string
	user     string
	pass     string
	uri      string // just for test porpuses
}

func (helper *MongoHelper) SetConnectionUri(uri string) {
	helper.credentials.uri = uri
}

func (helper *MongoHelper) GetCient() *mongo.Client {
	return helper.client
}

func NewMongoHelper(database, host, port, user, pass string) *MongoHelper {
	Helper = &MongoHelper{
		credentials: &credentials{
			database: database,
			host:     host,
			port:     port,
			user:     user,
			pass:     pass,
		},
	}
	return Helper
}

func (helper *MongoHelper) Connect() error {
	if err := helper.validateConnectionVars(); err != nil {
		return err
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(helper.getConnectionURL()))
	helper.client = client

	return err
}

func (helper *MongoHelper) Disconnect() {
	if helper.client != nil {
		helper.client.Disconnect(context.TODO())
	}
}

func (helper *MongoHelper) validateConnectionVars() error {
	if helper.atLeastOneVarEmpty() && helper.credentials.uri == "" {
		return errors.New("Insuficient informations provided for a database connection")
	}
	return nil
}

func (helper *MongoHelper) atLeastOneVarEmpty() bool {
	return helper.credentials.host == "" ||
		helper.credentials.port == "" ||
		helper.credentials.user == "" ||
		helper.credentials.pass == ""
}

func (helper *MongoHelper) getConnectionURL() string {
	if helper.credentials.uri != "" {
		return helper.credentials.uri
	}

	return fmt.Sprintf("mongodb://%s:%s@%s:%s",
		helper.credentials.user,
		helper.credentials.pass,
		helper.credentials.host,
		helper.credentials.port,
	)
}

func (helper *MongoHelper) GetCollection(collectionName string) *mongo.Collection {
	return helper.client.Database(helper.credentials.database).Collection(collectionName)
}

func (helper *MongoHelper) IsConnectionInvalid() bool {
	if Helper.GetCient() == nil {
		return true
	}

	if err := Helper.GetCient().Ping(context.TODO(), nil); err != nil {
		return true
	}

	return false
}
