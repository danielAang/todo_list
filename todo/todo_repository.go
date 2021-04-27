package todo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type todoRepo struct {
	db *mongo.Database
}

func NewTodoRepo(db *mongo.Database) *todoRepo {
	return &todoRepo{
		db: db,
	}
}

func (h *todoRepo) FindById(ID string) (*Todo, error) {
	var t Todo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := bson.M{"_id": ID}
	cursor := h.db.Collection("todo").FindOne(ctx, query)
	err := cursor.Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (h *todoRepo) FindAll(skip, limit int64) ([]Todo, error) {
	var todos []Todo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := bson.D{}
	options := options.Find()
	options.SetSort(map[string]int{"title": -1})
	options.SetSkip(skip)
	options.SetLimit(limit)
	cursor, err := h.db.Collection("todo").Find(ctx, query, options)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var t Todo
		if err = cursor.Decode(&t); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (h *todoRepo) Save(todo *Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var hex primitive.ObjectID
	var err error
	defer cancel()
	if todo.Id != "" {
		hex, err = primitive.ObjectIDFromHex(todo.Id)
		if err != nil {
			return err
		}
		todo.Id = hex.Hex()
	} else {
		todo.Id = primitive.NewObjectID().Hex()
	}
	update := bson.M{
		"$set": todo,
	}
	upsert := true
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	log.Println("id: ", hex)
	_, err = h.db.Collection("todo").UpdateByID(ctx, todo.Id, update, &opt)
	if err != nil {
		return err
	}
	return nil
}

func (h *todoRepo) Delete(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if hex, err := primitive.ObjectIDFromHex(ID); err != nil {
		return err
	} else {
		query := bson.M{"_id": hex}
		if _, err = h.db.Collection("todo").DeleteOne(ctx, query); err != nil {
			return err
		} else {
			return nil
		}
	}
}
