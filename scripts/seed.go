package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Ekireh-source/hotel-reservation/db"
	"github.com/Ekireh-source/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func main() {
	fmt.Println("Seeding the database...")
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DBName)

	hotel := types.Hotel{
		Name : "Grand Hotel",
		Location: "New York",
	}

	room := types.Room{

		Type: types.SingleRoomType,
		BasePrice: 100.0,
	}

	_ = room
	insertedHotel, err := hotelStore.InsertHotel(context.TODO(), &hotel)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted hotel: %+v\n", insertedHotel)
}