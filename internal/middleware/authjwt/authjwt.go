package authjwt

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type httpError struct {
	Response string `json:"response"`
}

/*
	@TODO
	- move firebase server app into application interface
*/

// Verify - Middleware function to verify firebase jwt
func Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Header.Get("Authorization") != "" {
			header := strings.Split(r.Header.Get("Authorization"), "Bearer ")
			jwt := header[1]
			opt := option.WithCredentialsFile("./firebase-admin.json")
			ctx := context.Background()

			app, err := firebase.NewApp(ctx, nil, opt)
			if err != nil {
				log.Fatalf("error initializing app: %v\n", err)
			}

			client, err := app.Auth(ctx)
			if err != nil {
				log.Fatalf("error getting Auth client: %v\n", err)
			}

			token, err := client.VerifyIDToken(ctx, jwt)
			if err != nil {
				w.WriteHeader(401)
				jsonErr := httpError{Response: "Unauthorized, Invalid JWT"}
				json.NewEncoder(w).Encode(jsonErr)
			}

			if token != nil {
				// Call the next handler, which can be another middleware in the chain, or the final handler.
				next.ServeHTTP(w, r)
			}
		} else {
			jsonErr := httpError{Response: "Authorization Header does not exist"}
			json.NewEncoder(w).Encode(jsonErr)
		}
	})
}
