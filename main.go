package main

import (
	"context"
	"example/FindProMates-Api/internal/models/mongodb"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Application struct {
	Projects *mongodb.ProjectModel
	Users    *mongodb.UserModel
}

var sandboxHTML = []byte(`
<!DOCTYPE html>
<html lang="en">
<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
<div id="sandbox" style="height:100vh; width:100vw;"></div>
<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
<script>
 new window.EmbeddedSandbox({
   target: "#sandbox",
   // Pass through your server href if you are embedding on an endpoint.
   // Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
   initialEndpoint: "http://localhost:8080/graphql",
 });
 // advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
</script>
</body>

</html>`)
var app *Application

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	db_name := os.Getenv("DB_NAME")
	if db_name == "" {
		log.Fatal("DB_NAME is not set")
	}
	app = &Application{
		Projects: &mongodb.ProjectModel{
			C: client.Database(db_name).Collection("projects"),
		},
		Users: &mongodb.UserModel{
			C: client.Database(db_name).Collection("users"),
		},
	}
	usersHandler := handler.New(&handler.Config{
		Schema: app.GetUsersSchema(),
		Pretty: true,
	})
	http.Handle("/users", usersHandler)
	http.Handle("/sandbox", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(sandboxHTML)
	}))
	http.ListenAndServe(":8080", nil)

}
