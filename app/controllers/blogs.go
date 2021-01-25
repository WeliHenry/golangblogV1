package controllers

import (
	"context"
	"encoding/json"
	"github.com/WeliHenry/golangblogV1/app/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

// PostCreateBlog is for creating blog
func CreateBlog(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	blog := models.Blog{}
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid JSON request body "))
		return
	}
	blog.ID = primitive.NewObjectID().Hex()
	_, err = db.Collection("blogs").InsertOne(context.TODO(), blog)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("An error occured while creating blog post ! "))
		return
	}
	json.NewEncoder(res).Encode(blog)
}

// GetBlogs is for getting all blogs
func GetBlogs(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var blogs []*models.Blog
	cur, err := db.Collection("blogs").Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem models.Blog
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		blogs = append(blogs, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	respondJSON(res, http.StatusOK, blogs)
}

// GetBlog is for getting a single blog
func GetBlog(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	blog := models.Blog{}
	filter := bson.M{"_id": params["_id"]}
	decoder := db.Collection("blogs").FindOne(context.TODO(), filter)
	err := decoder.Decode(&blog)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("An error occured while returning blog post !"))
		return
	}
	respondJSON(res, http.StatusOK, blog)
}

// DeleteBlog is for deleting a single blog
func DeleteBlog(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	filter := bson.M{"_id": params["_id"]}
	_, err := db.Collection("blogs").DeleteOne(context.TODO(), filter)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("An error occured while deleting blog post "))
		return
	}
	deleteResult := models.Response{
		Message: "Blog deleted successfully",
	}
	respondJSON(res, http.StatusOK, deleteResult)
}

//PutUpdateBlog is for updating a single blog
func UpdateBlogs(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	blog := models.Blog{}
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid JSON request body "))
		return
	}
	filter := bson.M{"_id": params["_id"]}
	update := bson.M{"$set": blog}
	db.Collection("blogs").FindOneAndUpdate(context.TODO(), filter, update)
	updateResponse := models.Response{}
	updateResponse.Message = "Blog updated successfully!"
	json.NewEncoder(res).Encode(updateResponse)

}
