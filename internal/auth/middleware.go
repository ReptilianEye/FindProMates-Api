package auth

import (
	"context"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/pkg/jwt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}
		userId, err := jwt.ParseToken(token)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		userIDObj, _ := primitive.ObjectIDFromHex(userId)
		user, err := app.App.Users.FindById(userIDObj)
		if err != nil || user == nil {
			next.ServeHTTP(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), userCtxKey, userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}
