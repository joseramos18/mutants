package repository

import (
	"context"
	"fmt"
	"math"
	"mutants/interfaces"
	"mutants/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MutantRepository struct {
	client            *mongo.Client
	mutantsCollection *mongo.Collection
}

func New(client *mongo.Client) interfaces.IMutantRepository {
	return &MutantRepository{
		client:            client,
		mutantsCollection: client.Database("secondApp").Collection("mutants"),
	}
}

func (repository *MutantRepository) SaveDna(dna models.DNA, ctx context.Context) error {
	fmt.Println("4")
	result, err := repository.mutantsCollection.InsertOne(ctx, dna)
	fmt.Println(result)
	fmt.Println(err)
	return err
}

func (repository *MutantRepository) GetStatistics(ctx context.Context) (models.Statistics,  error){
	var statics models.Statistics
	cursor, err := repository.mutantsCollection.Find(ctx, bson.D{})
	if err != nil {
		for cursor.Next(context.TODO()) {
			var dna models.DNA
			err := cursor.Decode(&dna)
			if err != nil {
				return statics, err
			}
			if dna.IsMutant {
				statics.CountMutant = statics.CountMutant + 1
			} else {
				statics.CountHuman = statics.CountHuman + 1
			}
		}
		statics.Ratio = math.Round(float64(statics.CountHuman)/float64(statics.CountMutant)*100)/100
		return statics, nil
	}
	return statics, err
	
}
