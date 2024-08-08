package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/JcksonMCD/golang-jwt/database"
	"github.com/JcksonMCD/golang-jwt/models"
	helper "github.com/JcksonMCD/golang-jwt/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() {

}

func VerifyPassword() {

}

// Signup returns a handler function that registers a new user.
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the user data.
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the email already exists in the database.
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Println("Error occurred while checking email:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while checking email."})
			return
		}

		if count > 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "This email is already being used by another user account."})
			return
		}

		// Check if the phone number already exists in the database.
		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if err != nil {
			log.Println("Error occurred while checking phone number:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while checking phone number."})
			return
		}

		if count > 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "This phone number is already being used by another user account."})
			return
		}

		// saving logic to go here....

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	}
}

func Login() {

}

func GetUsers() {

}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		err := helper.MatchUserTypeToID(c, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err = UserCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
