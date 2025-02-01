package main

import (
	"github.com/judegiordano/gogetem/pkg/logger"
	"github.com/judegiordano/gogetem/pkg/mongo"
	"github.com/judegiordano/sst_template/internal/models"
)

func main() {
	model := mongo.IndexModel{
		Keys: mongo.Bson{"message": 1},
	}
	idx, err := mongo.CreateIndex[models.Health](model)
	if err != nil {
		logger.Fatal("[ERROR CREATING INDEX]", err)
	}
	logger.Info("[CREATED INDEX]", *idx)
}
