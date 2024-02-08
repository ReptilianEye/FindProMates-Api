package utils

import (
	"log"
	"math/rand"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Set map[string]bool

func ToSet[E any](arr []E, mapper func(E) string) Set {
	set := make(Set)
	for _, v := range arr {
		set[mapper(v)] = true
	}
	return set
}
func ToSlice[E any](s Set, mapper func(string) E, firstInOrder ...string) []E {
	slice := make([]E, 0, len(s))
	for _, k := range firstInOrder {
		if _, ok := s[k]; !ok {
			log.Fatal("Key not found in set")
		}
		slice = append(slice, mapper(k))
		delete(s, k)
	}
	for k := range s {
		slice = append(slice, mapper(k))
	}
	return slice
}
func (s Set) Add(keys ...string) {
	for _, k := range keys {
		s[k] = true
	}
}
func (s Set) Union(other Set) Set {
	union := make(Set)
	for k := range s {
		union[k] = true
	}
	for k := range other {
		union[k] = true
	}
	return union
}
func (s Set) Intersection(other Set) Set {
	intersection := make(Set)
	if len(s) > len(other) {
		s, other = other, s
	}
	for k := range s {
		if other[k] {
			intersection[k] = true
		}
	}
	return intersection
}
func (s Set) Contains(key string) bool {
	return s[key]
}
func Identity[E any](e E) E {
	return e
}

// MergeSlices merges two slices and returns a new slice with unique elements from both slices.
// The base slice contains elements of type E, while the new slice contains elements of type string.
// The mapperEToString function is used to convert elements of type E to strings,
// and the mapperStringToE function is used to convert strings back to elements of type E.
// The anchors parameter is an optional variadic parameter that specifies additional elements to include in the result.
// The function returns a new slice of type []E with unique elements from both input slices and the anchors.
func MergeSlices[E any](base []E, new []string, mapperEToString func(E) string, mapperStringToE func(string) E, anchors ...string) []E {
	baseS := ToSet(base, mapperEToString)
	newS := ToSet(new, Identity)
	union := baseS.Union(newS)
	union.Add(anchors...)
	return ToSlice(union, mapperStringToE, anchors...)
}
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
func Ternary(condition bool, a, b any) any {
	if condition {
		return a
	}
	return b
}
func Elivis[E any](a *E, b E) E {
	if a != nil {
		return *a
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
