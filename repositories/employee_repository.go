package repositories

import (
	"context"
	"employee-management-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepository interface {
	FindAll() ([]models.Employee, error)
	FindByID(id string) (*models.Employee, error)
	Create(employee *models.Employee) (string, error)
	Update(id string, employee *models.Employee) error
	Delete(id string) error
}

type employeeRepo struct {
	collection *mongo.Collection
}

func NewEmployeeRepository(collection *mongo.Collection) EmployeeRepository {
	return &employeeRepo{collection: collection}
}

func (r *employeeRepo) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var employee models.Employee
		if err := cursor.Decode(&employee); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (r *employeeRepo) FindByID(id string) (*models.Employee, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	var employee models.Employee
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&employee)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepo) Create(employee *models.Employee) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, employee)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *employeeRepo) Update(id string, employee *models.Employee) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": employee})
	return err
}

func (r *employeeRepo) Delete(id string) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
