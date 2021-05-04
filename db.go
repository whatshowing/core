package core

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func ParseObjectId(id string) (primitive.ObjectID, error) {
	IdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Failed to parse id")
		log.Println(err)
		return IdObject, status.Error(codes.InvalidArgument, "the id is invalid")
	}
	return IdObject, nil
}
