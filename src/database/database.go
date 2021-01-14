package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/blackcrw/wprecon-api/src/config"
	"github.com/blackcrw/wprecon-api/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect ::
func Connect() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MONGOURI))

	if err != nil {
		return nil, fmt.Errorf("An error occurred while trying to create a new instance of mongodb. (%s)", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err = client.Connect(ctx); err != nil {
		return nil, fmt.Errorf("An error occurred while trying to connect to mongodb. (%s)", err)
	}

	return client, nil
}

func collectionPluginVuln() (*mongo.Collection, error) {
	client, err := Connect()

	if err != nil {
		return nil, err
	}

	database := client.Database("wprecon")
	collection := database.Collection("vulnerabilitys")

	return collection, nil
}

func GetVulnerabilities(plugin string, version string) *models.Plugin {
	var vulnerability models.Plugin

	collection, err := collectionPluginVuln()

	if err != nil {
		log.Println(err)
	}

	defer collection.Database().Client().Disconnect(context.TODO())

	filter := bson.M{"PluginName": plugin, "Vulnerabilities.Version": version}
	options := options.FindOne().SetProjection(bson.M{"Vulnerabilities.$": 1})

	err = collection.FindOne(context.TODO(), filter, options).Decode(&vulnerability)

	switch err {
	case mongo.ErrNoDocuments:
	default:
		log.Println(err)
	}

	fmt.Println(vulnerability)

	return &vulnerability
}
