package mongodb

import "errors"

type MongoHelper struct {
	database string
	host     string
	port     string
	user     string
	pass     string
}

func NewMongoHelper(database, host, port, user, pass string) *MongoHelper {
	return &MongoHelper{
		database: database,
		host:     host,
		port:     port,
		user:     user,
		pass:     pass,
	}
}

func (helper *MongoHelper) Connect() error {
	if err := helper.validateConnectionVars(); err != nil {
		return err
	}

	return nil
}

func (helper *MongoHelper) validateConnectionVars() error {
	if helper.atLeastOneVarEmpty() {
		return errors.New("Insuficient informations provided for a database connection")
	}
	return nil
}

func (helper *MongoHelper) atLeastOneVarEmpty() bool {
	return helper.host == "" || helper.port == "" || helper.user == "" || helper.pass == ""
}
