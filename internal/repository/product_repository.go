package repository

import (
	"context"

	"github.com/jufergom/products-api/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ProductRepository struct {
	Collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{Collection: db.Collection("products")}
}

func (r *ProductRepository) FindAll() ([]model.Product, error) {
	cur, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var products []model.Product
	for cur.Next(context.TODO()) {
		var p model.Product
		if err := cur.Decode(&p); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) FindByID(id string) (*model.Product, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var product model.Product
	err = r.Collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
