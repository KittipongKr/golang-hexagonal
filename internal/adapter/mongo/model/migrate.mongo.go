package mongo_model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseModel struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time         `json:"daaleted_at" bson:"daaleted_at"`
}

func (b *BaseModel) SetInsertMeta() {
	now := time.Now()
	b.Id = primitive.NewObjectID()
	b.CreatedAt = now
	b.UpdatedAt = now
	b.DeletedAt = nil

}

func (b *BaseModel) SetUpdateMeta() {
	now := time.Now()
	b.UpdatedAt = now
}
func (b *BaseModel) SetDeleteMeta() {
	now := time.Now()
	b.DeletedAt = &now
}

type MongoCollections struct {
	Users *mongo.Collection
}
