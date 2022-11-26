package mongodbtest

import (
	"github.com/benweissmann/memongo"
)

func SetupMemoryServer() (*memongo.Server, error) {
	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		return nil, err
	}
	return mongoServer, nil
}
