package controllers

import (
	"mongodb-rest-api/entities"
	"mongodb-rest-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpsertProduct(c *fiber.Ctx) error {

	var res entities.Result
	var filter interface{}
	var products interface{}

	body := new(entities.ReqBody)

	if err := c.BodyParser(body); err != nil {

		res.Products = nil
		res.Message = "Invalid body!!"

		return c.Status(400).JSON(res)
	}

	for i := range body.Products {

		filter = bson.D{{Key: "Code", Value: body.Products[i].Code}}

		products = bson.D{{Key: "$set", Value: bson.D{
			{Key: "Name", Value: body.Products[i].Name},
			{Key: "Code", Value: body.Products[i].Code},
			{Key: "Barcode", Value: body.Products[i].Barcode},
			{Key: "Active", Value: body.Products[i].Active},
			{Key: "UpsertedAt", Value: time.Now()},
		}}}

		if !models.UpsertDocument(products, filter, "Products") {

			res.Products = append(res.Products, entities.ProductResponse{
				Code:    body.Products[i].Code,
				Success: false,
			})

		} else {

			res.Products = append(res.Products, entities.ProductResponse{
				Code:    body.Products[i].Code,
				Success: true,
			})
		}
	}

	res.Message = "Upsert finished!!"

	return c.Status(200).JSON(res)
}

func DeleteProduct(c *fiber.Ctx) error {

	var res entities.Result
	var filter interface{}

	body := new(entities.ReqBody)

	if err := c.BodyParser(body); err != nil {

		res.Products = nil
		res.Message = "Invalid body!!"

		return c.Status(400).JSON(res)
	}

	for i := range body.Products {

		filter = bson.M{"Code": body.Products[i].Code}

		deleted, rows := models.DeleteDocument(filter, "Products")

		if !deleted || rows == 0 {

			res.Products = append(res.Products, entities.ProductResponse{
				Code:    body.Products[i].Code,
				Success: false,
			})

		} else {
			res.Products = append(res.Products, entities.ProductResponse{
				Code:    body.Products[i].Code,
				Success: true,
			})
		}
	}

	res.Message = "Delete finished!!"

	return c.Status(200).JSON(res)
}

func FindProduct(c *fiber.Ctx) error {

	var products = make(map[string]interface{})
	success, documents := models.FindDocument("Products")

	products["Products"] = documents

	if !success {
		return c.Status(500).JSON(products)
	}

	if len(documents) == 0 {
		return c.Status(404).JSON(products)
	}

	return c.Status(200).JSON(products)
}
