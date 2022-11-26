package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoHelper struct {
	credentials *credentials
}

type credentials struct {
	database string
	host     string
	port     string
	user     string
	pass     string
}

func NewMongoHelper(database, host, port, user, pass string) *MongoHelper {
	return &MongoHelper{
		credentials: &credentials{
			database: database,
			host:     host,
			port:     port,
			user:     user,
			pass:     pass,
		},
	}
}

func (helper *MongoHelper) Connect() error {
	if err := helper.validateConnectionVars(); err != nil {
		return err
	}

	_, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(helper.getConnectionURL()))

	return err
}

func (helper *MongoHelper) validateConnectionVars() error {
	if helper.atLeastOneVarEmpty() {
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
	return fmt.Sprintf("mongodb://%s:%s@%s:%s",
		helper.credentials.user,
		helper.credentials.pass,
		helper.credentials.host,
		helper.credentials.port,
	)
}
