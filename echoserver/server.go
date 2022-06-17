package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
)

type (
	ContactModel struct {
		Id   interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
		Name string      `json:"name" form:"name" query:"name" validate:"required"`
		Age  int         `json:"age" form:"age" query:"age" validate:"required"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("N .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental ariable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	col := client.Database("wrapme-local").Collection("contacts")
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// routes
	e.GET("/contacts", func(ctx echo.Context) error {
		filterCursor, err := col.Find(context.TODO(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		var contacts []ContactModel
		if err = filterCursor.All(context.TODO(), &contacts); err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(http.StatusOK, contacts)
	})
	//e.GET("/contacts", getContacts)
	//e.GET("/contacts/:id", getContact)

	// create contact
	e.POST("/contacts", func(c echo.Context) error {
		contact := new(ContactModel)
		err := c.Bind(contact)
		if err != nil {
			return err
		}
		if err = c.Validate(contact); err != nil {
			return err
		}
		doc := bson.D{{"name", contact.Name}, {"age", contact.Age}}
		result, err := col.InsertOne(context.TODO(), doc)
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusCreated, result)
	})

	e.PUT("/contacts/:id", func(c echo.Context) error {
		param := c.Param("id")
		fmt.Println("ddddd", param)
		id, _ := primitive.ObjectIDFromHex(param)
		filter := bson.D{{"_id", id}}
		contact := new(ContactModel)
		err := c.Bind(contact)
		if err != nil {
			return err
		}
		if err = c.Validate(contact); err != nil {
			return err
		}
		doc := bson.D{{"$set", bson.D{{"name", contact.Name}, {"age", contact.Age}}}}
		result, err := col.UpdateOne(context.TODO(), filter, doc)
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusCreated, result)
	})

	// start server
	e.Logger.Fatal(e.Start(":3333"))
}
