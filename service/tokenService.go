package service

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/JcksonMCD/golang-jwt/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Uid        string
	User_Type  string
	jwt.StandardClaims
}

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_Name: firstName,
		Last_Name:  lastName,
		Uid:        uid,
		User_Type:  userType,
		StandardClaims: jwt.StandardClaims{
			// Expires after 24 hours
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			// Extended expiration
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	// Tokens encoded
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	//Create a context with a timeout to ensure the operation doesn't run indefinitely.
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	//prepare the fields to be updated in the database.
	updateFields := bson.D{
		{"token", signedToken},
		{"refresh_token", signedRefreshToken},
		{"updated_at", time.Now()},
	}

	// define the filter to locate the specific user in the database using id field.
	filter := bson.M{"user_id": userId}

	// Step 4: Set up the update options, specifically enabling upsert (create the document if it doesn't exist).
	upsert := true
	options := options.Update().SetUpsert(upsert) // Enable upsert to create a new document if one doesn't exist.

	_, err := UserCollection.UpdateOne(
		ctx, // Pass the context to ensure the operation respects the timeout.
		filter,
		bson.D{
			{"$set", updateFields}, // Use the $set operator to update the specified fields.
		},
		options,
	)

	if err != nil {
		log.Printf("Failed to update tokens for user %s: %v", userId, err)
		return
	}

	return
}

func ValidateToken() {

}
