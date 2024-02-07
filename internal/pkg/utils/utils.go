package utils

import (
	"math/rand"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func MapTo[E, V any](arr []E, mapper func(E) V) []V {
	mapped := make([]V, len(arr))
	for i, v := range arr {
		mapped[i] = mapper(v)
	}
	return mapped
}
func Any[E any](arr []E, predicate func(E) bool) bool {
	for _, v := range arr {
		if predicate(v) {
			return true
		}
	}
	return false
}
func All[E any](arr []E, predicate func(E) bool) bool {
	for _, v := range arr {
		if !predicate(v) {
			return false
		}
	}
	return true
}
func Ternary[E any](condition bool, a, b E) E {
	if condition {
		return a
	}
	return b
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
func CreateUsername(firstName, lastName string) string {
	return strings.ToLower(firstName) + "_" + strings.ToLower(lastName) + randomNumbersSufix(3)
}
func randomNumbersSufix(length int) string {
	res := ""
	for i := 0; i < length; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}

// func readFromDB() {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(context.Background())
// 	db := client.Database(os.Getenv("DB_NAME"))
// 	usersColl := db.Collection("users")
// 	projectsColl := db.Collection("projects")
// 	file, err := os.ReadFile("dbschema/db.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var db_sample map[string]interface{}
// 	err = json.Unmarshal(file, &db_sample)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, user := range db_sample["users"].([]interface{}) {
// 		_, err := usersColl.InsertOne(context.Background(), user)
// 		if err != nil {
// 			fmt.Println(err)
// 			// log.Fatal(err)
// 		}
// 	}
// 	for _, project := range db_sample["projects"].([]interface{}) {

// 		res, err := projectsColl.InsertOne(context.Background(), project)
// 		if err != nil {
// 			fmt.Println(err)
// 			// log.Fatal(err)
// 		} else {
// 			fmt.Println(res.InsertedID)
// 		}
// 	}
// }
