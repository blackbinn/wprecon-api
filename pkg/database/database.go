package database

import (
	"context"
	"fmt"
	"time"

	"github.com/blackcrw/wprecon-api/pkg/config"
	"github.com/blackcrw/wprecon-api/pkg/models"
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

// GetVulnerabilities ::
func GetVulnerabilities(plugin string, version string) (*models.Plugin, error) {
	collection, err := collectionPluginVuln()

	defer collection.Database().Client().Disconnect(context.TODO())

	if err != nil {
		return &models.Plugin{}, err
	}

	var vulnerability models.Plugin

	filter := bson.M{"PluginName": plugin, "Vulnerabilities.Version": version}
	options := options.FindOne().SetProjection(bson.M{"Vulnerabilities.$": 1})

	err = collection.FindOne(context.TODO(), filter, options).Decode(&vulnerability)

	if err != nil {
		return &models.Plugin{}, err
	}

	return &vulnerability, nil
}
