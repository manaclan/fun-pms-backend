package hotels

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	Client *mongo.Client
}

func (service Services) Search(keyword string) (int, []Hotel) {
	var hotels []Hotel

	return 200, hotels
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
