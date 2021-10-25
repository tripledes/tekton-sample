package types

import "go.mongodb.org/mongo-driver/bson/primitive"

//Quote represents a quote object
type Quote struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Author string             `json:"author" bson:"author,omitempty"`
	Text   string             `json:"text" bson:"text,omitempty"`
	Tags   []string           `json:"tags" bson:"tags,omitempty"`
}
