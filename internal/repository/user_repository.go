package repository

import (
	"context"

	"github.com/jufergom/products-api/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CustomerRepository struct {
	Collection *mongo.Collection
}

func NewCustomerRepository(db *mongo.Database) *CustomerRepository {
	return &CustomerRepository{Collection: db.Collection("customers")}
}

func (r *CustomerRepository) FindAll() ([]model.Customer, error) {
	var customers []model.Customer
	cur, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var u model.Customer
		if err := cur.Decode(&u); err != nil {
			return nil, err
		}
		customers = append(customers, u)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerRepository) FindByID(id string) (*model.Customer, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var customer model.Customer
	err = r.Collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
