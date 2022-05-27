package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mon MongoInstance

const dbName = "hrms-fiber"
const mongoURI = "mongodb://developer:developer123@localhost:27017/" + dbName + "?ssl=false&authMechanism=SCRAM-SHA-256&authSource=admin"

type Employee struct {
	ID       string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string  `json:"name"`
	Salary   float64 `json:"salary"`
	Age      int     `json:"age"`
	Position string  `json:"position"`
}

func ConnectDb() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		log.Fatal(err)
	}

	mon = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func main() {
	app := fiber.New()

	if err := ConnectDb(); err != nil {
		log.Fatal(err)
	}

	app.Get("/employee", func(ctx *fiber.Ctx) error {
		query := bson.D{{}}
		cursor, err := mon.Db.Collection("employees").Find(ctx.Context(), query)
		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		var employees = make([]Employee, 0)

		if err := cursor.All(ctx.Context(), &employees); err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		return ctx.JSON(employees)
	})

	app.Post("/employee", func(ctx *fiber.Ctx) error {
		collection := mon.Db.Collection("employees")

		employee := new(Employee)

		if err := ctx.BodyParser(employee); err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		employee.ID = ""

		result, err := collection.InsertOne(ctx.Context(), employee)
		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: result.InsertedID}}
		record := collection.FindOne(ctx.Context(), &filter)

		createdRecord := &Employee{}
		if err := record.Decode(createdRecord); err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		return ctx.Status(201).JSON(createdRecord)
	})

	app.Put("/employee/:id", func(ctx *fiber.Ctx) error {
		collection := mon.Db.Collection("employees")
		id := ctx.Params("id")

		employeeId, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return ctx.Status(400).SendString(err.Error())
		}

		employee := new(Employee)

		if err := ctx.BodyParser(employee); err != nil {
			return ctx.Status(400).SendString(err.Error())
		}

		query := bson.D{{Key: "_id", Value: employeeId}}
		updateQuery := bson.D{{Key: "$set", Value: bson.D{
			{Key: "name", Value: employee.Name},
			{Key: "age", Value: employee.Age},
			{Key: "salary", Value: employee.Salary},
			{Key: "position", Value: employee.Position},
		}}}

		if err := collection.FindOneAndUpdate(ctx.Context(), &query, &updateQuery).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				return ctx.SendStatus(400)
			}
			return ctx.SendStatus(500)
		}

		employee.ID = id

		return ctx.Status(200).JSON(employee)
	})

	app.Delete("/employee/:id", func(ctx *fiber.Ctx) error {
		collection := mon.Db.Collection("employees")

		employeeId, err := primitive.ObjectIDFromHex(ctx.Params("id"))

		if err != nil {
			return ctx.Status(404).SendString(err.Error())
		}

		query := bson.D{{Key: "_id", Value: employeeId}}

		result, err := collection.DeleteOne(ctx.Context(), &query)

		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		if result.DeletedCount < 1 {
			return ctx.SendStatus(404)
		}

		return ctx.Status(200).SendString("record deleted")
	})

	log.Fatal(app.Listen(":8080"))
}
