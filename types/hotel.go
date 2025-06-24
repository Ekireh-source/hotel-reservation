package types

import "go.mongodb.org/mongo-driver/bson/primitive"


type Hotel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Location	string `json:"location" bson:"location"`
	Room 	    []primitive.ObjectID `json:"room" bson:"room"`
}



type RoomType int

const (

	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSideRoomType

	DeluxeRoomType
)


type Room struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
    Type 	  RoomType `json:"type" bson:"type"`
	BasePrice   float64   `json:"basePrice" bson:"basePrice"`
	HotelID     primitive.ObjectID `json:"hotelID" bson:"hotelID"`
	Price 	  float64   `json:"price" bson:"price"`

}