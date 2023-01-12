package mongodbtest

import (
	"github.com/benweissmann/memongo"
	"github.com/claudiomozer/starwarsapi/src/infra/repositories/db/mongodb"
)

func SetupMemoryServer() (*memongo.Server, error) {
	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		return nil, err
	}
	return mongoServer, nil
}

func ConnectWithMongoServer() error {
	memServer, err := SetupMemoryServer()

	if err != nil {
		return err
	}

	sut := mongodb.NewMongoHelper("dbtest", "", "", "", "")
	sut.SetConnectionUri(memServer.URI())
	err = sut.Connect()

	return err
}
