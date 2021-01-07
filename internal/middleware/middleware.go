package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/JSONhilder/overseer_api/internal/utils"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// VerifyJwt - Middleware function to verify firebase jwt
func VerifyJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// @TODO - Remove in production
		if r.Header.Get("Authorization") == "DEV" {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
			return
		}

		if r.Header.Get("Authorization") != "" {
			header := strings.Split(r.Header.Get("Authorization"), "Bearer ")
			jwt := header[1]
			opt := option.WithCredentialsFile("./firebase-admin.json")
			ctx := context.Background()

			app, err := firebase.NewApp(ctx, nil, opt)
			if err != nil {
				str := fmt.Sprint("error initializing app: %v\n", err)
				utils.JSONError(w, str, 500)
			}

			client, err := app.Auth(ctx)
			if err != nil {
				str := fmt.Sprint("error getting Auth client: %v\n", err)
				utils.JSONError(w, str, 500)
			}

			token, err := client.VerifyIDToken(ctx, jwt)
			if err != nil {
				utils.JSONError(w, "Unauthorized, Invalid JWT", 401)
			}

			if token != nil {
				// Set content header to json
				w.Header().Set("Content-Type", "application/json")
				// Call the next handler, which can be another middleware in the chain, or the final handler.
				next.ServeHTTP(w, r)
			}
		} else {
			utils.JSONError(w, "Authorization Header does not exist", 404)
		}
	})
}
