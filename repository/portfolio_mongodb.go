package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (

	// portfolioRepositoryMongodb is Adapter
	portfolioRepositoryMongodb struct {
		client   *mongo.Client
		database *mongo.Database
	}
)

func NewPortfolioRepositoryMongodb(client *mongo.Client, dbname string) portfolioRepositoryMongodb {
	return portfolioRepositoryMongodb{client: client, database: client.Database(dbname)}
}

func (r portfolioRepositoryMongodb) GetAll(userid int) (datas Portfolios, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := primitive.M{"userid": userid}

	var cur *mongo.Cursor

	if cur, err = r.database.Collection("portfolio").Find(ctx, filter); err != nil {
		return
	}

	for cur.Next(ctx) {
		var row Portfolio
		if err = cur.Decode(&row); err == nil {
			datas = append(datas, row)
		}
	}

	return
}

func (r portfolioRepositoryMongodb) GetBySymbol(userid int, symbol string) (data *Portfolio, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := primitive.M{"symbol": symbol, "userid": userid}

	log.Printf("filter: %#v", filter)
	res := r.database.Collection("portfolio").FindOne(ctx, filter)
	if res.Err() != nil {
		err = res.Err()
		return
	}

	res.Decode(&data)

	return
}
