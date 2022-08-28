//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongodbImpl struct {
	client   *mongo.Client
	database *mongo.Database
}

func New(opts Options) (db Database, err error) {
	opts.SetDefaults()
	ctx, cancel := context.WithTimeout(context.Background(), opts.CtxTimeout)
	defer cancel()
	clientOpts := options.Client().ApplyURI(opts.URI)
	if opts.IsReader {
		clientOpts.SetReadPreference(readpref.SecondaryPreferred())
	}
	var dbImpl mongodbImpl
	if dbImpl.client, err = mongo.Connect(ctx, clientOpts); err != nil {
		return nil, ErrCouldNotConnect(err)
	}
	if err = dbImpl.client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, ErrCantPing(err)
	}
	dbImpl.database = dbImpl.client.Database(opts.DatabaseName)
	return dbImpl, nil
}

func (i mongodbImpl) Name() string {
	return i.database.Name()
}

func (i mongodbImpl) Collection(name string) Collection {
	return i.database.Collection(name)
}
