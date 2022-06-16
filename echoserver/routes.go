package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func getContacts(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func getContact(c echo.Context) error {
	// get params id
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func saveContact(c echo.Context) error {
	contact := new(ContactModel)
	client := Client()
	coll := client.Database("wrapme-local").Collection("contacts")
	doc := bson.D{{"name", contact.Name}, {"age", contact.Age}}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusCreated, result)
}

func updateContact(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func deleteContact(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
