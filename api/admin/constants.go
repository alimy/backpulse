package admin

import (
	"net/http"

	"github.com/alimy/mir"
	"github.com/backpulse/core/constants"
	"github.com/backpulse/core/utils"
)

// Constants indicator constants handler
type Constants struct {
	group        mir.Group `mir:"admin"`
	getLanguages mir.Get   `mir:"/constants/languages"`
}

// GetLanguages : return array of languages
func (c *Constants) GetLanguages(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, "success", constants.Languages)
}
