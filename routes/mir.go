package routes

import (
	"github.com/backpulse/core/api/client"
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/backpulse/core/api/admin"
	"github.com/backpulse/core/database"
	"github.com/backpulse/core/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// mirChain return middleware use in mux
func mirChain() []mux.MiddlewareFunc {
	config := utils.GetConfig()
	return []mux.MiddlewareFunc{
		jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(config.Secret), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		}).Handler,
		mux.MiddlewareFunc(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				id := utils.GetUserObjectID(r)
				_, err := database.GetUserByID(id)
				if err != nil {
					utils.RespondWithJSON(w, http.StatusUnauthorized, "unauthorized", nil)
					return
				}
				h.ServeHTTP(w, r)
			})
		}),
	}
}

// mirEntries get all entries that used to register to Mir
func mirEntries() []interface{} {
	return []interface{}{
		&admin.About{Chain: mirChain()},
		&admin.Albums{},
		&admin.Articles{},
		&admin.Constants{},
		&admin.Contact{},
		&admin.Files{},
		&admin.Galleries{},
		&admin.Photos{},
		&admin.Project{},
		&admin.Sites{},
		&admin.Tracks{},
		&admin.Users{},
		&admin.VideoGroups{},
		&admin.Videos{},

		&client.Client{},
	}
}
