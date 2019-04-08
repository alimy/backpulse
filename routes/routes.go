package routes

import (
	"github.com/gorilla/mux"

	mirE "github.com/alimy/mir/module/mux"
)

//NewRouter creates router with route handlers
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	entries := mirEntries()
	mirE.Register(r, entries...)

	return r
}
