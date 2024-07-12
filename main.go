package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	// URL ของ MongoDB Atlas
	clientOptions := options.Client().ApplyURI("apikey")

	// สร้าง client และเชื่อมต่อไปยัง MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// ตรวจสอบการเชื่อมต่อ
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// เลือกฐานข้อมูลและคอลเลคชัน
	collection := client.Database("test").Collection("numbers")

	// สร้าง context สำหรับการทำงานกับ MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ตัวอย่างการเพิ่มข้อมูลลงในคอลเลคชัน
	doc := bson.D{{"name", "pi"}, {"value", 3.14159}}
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	// ปิดการเชื่อมต่อ
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
