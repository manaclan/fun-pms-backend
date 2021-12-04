package hotel

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	Client *mongo.Client
}

func (services Services) AddHotel(hotelName string, location Location) (int, string) {

	fmt.Println("Service: AddHotel")
	var hotels []Hotel
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	hotelsCollection := services.Client.Database("pms-backend").Collection("hotels")

	cursor, err := hotelsCollection.Find(
		context.Background(), bson.M{"location": location})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(ctx, &hotels); err != nil {
		panic(err)
	}
	if len(hotels) > 0 {
		return 404, "Hotel existed!"
	}
	newHotel := Hotel{Name: hotelName, Location: location}
	if err != nil {
		panic(err)
	}
	_, err = hotelsCollection.InsertOne(ctx, newHotel)
	if err != nil {
		panic(err)
	}
	return 200, "Succeed!"
}

// func (services Services) ValidateUser(username string, password string) bool {
// 	fmt.Println("Service: ValidateUser")
// 	var users []User
// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	defer cancel()
// 	currentUsersCollection := database.GetDB().Database("users").Collection("current-users")

// 	cursor, err := currentUsersCollection.Find(
// 		context.Background(), bson.M{"username": username})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cursor.Close(context.Background())
// 	if err = cursor.All(ctx, &users); err != nil {
// 		panic(err)
// 	}
// 	if len(users) == 0 {
// 		return false
// 	}
// 	err = CheckPasswordHash(users[0].Password, password)
// 	return err == nil
// }
