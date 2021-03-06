package app

import (
	"context"
	"fmt"
	"github.com/WeliHenry/golangblogV1/app/controllers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

// App has Router and Database instances
type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

// Initialize the App with predefined configurations
func (app *App) Initialize(MongoURI string) {
	clientOptions := options.Client().ApplyURI(MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	db := client.Database("GOBLOG")
	app.DB = db

	app.Router = mux.NewRouter()
	app.setRouters()
}

func (app *App) setRouters() {
	app.Post("/api/blogs/", app.handleRequest(controllers.CreateBlog))
	app.Get("/api/blogs/", app.handleRequest(controllers.GetBlogs))
	app.Get("/api/blogs/{_id}", app.handleRequest(controllers.GetBlog))
	app.Put("/api/blogs/{_id}", app.handleRequest(controllers.UpdateBlogs))
	app.Delete("/api/blogs/{_id}", app.handleRequest(controllers.DeleteBlog))
}

// Get wraps the router for GET method
func (app *App) Get(path string, f func(res http.ResponseWriter, req *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (app *App) Post(path string, f func(res http.ResponseWriter, req *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

// Delete wraps the router for DELETE method
func (app *App) Delete(path string, f func(res http.ResponseWriter, req *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("DELETE")
}

// Put wraps the router for PUT method
func (app *App) Put(path string, f func(res http.ResponseWriter, req *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("PUT")
}

// RequestHandlerFunction handles the requests
type RequestHandlerFunction func(db *mongo.Database, w http.ResponseWriter, r *http.Request)

func (app *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		handler(app.DB, resp, req)
	}
}

// Run listens to the port provided to it
func (app *App) Run(port string) {
	http.ListenAndServe(port, app.Router)
}
