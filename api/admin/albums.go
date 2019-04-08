package admin

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alimy/mir"
	"github.com/backpulse/core/database"
	"github.com/backpulse/core/models"
	"github.com/backpulse/core/utils"
	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
	"gopkg.in/mgo.v2/bson"
)

// Albums indicator albums handler
type Albums struct {
	group               mir.Group  `mir:"admin"`
	updateAlbumsIndexes mir.Put    `mir:"/albums/{name}/indexes"`
	createAlbum         mir.Post   `mir:"/albums/{name}"`
	getAlbums           mir.Get    `mir:"/albums/{name}"`
	getAlbum            mir.Get    `mir:"/albums/{name}/{id}"`
	updateAlbum         mir.Put    `mir:"/albums/{name}/{id}"`
	deleteAlbum         mir.Delete `mir:"/albums/{name}/{id}"`
}

// CreateAlbum : create new album
func (a *Albums) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	site, _ := database.GetSiteByName(name)
	user, _ := database.GetUserByID(utils.GetUserObjectID(r))

	if !utils.IsAuthorized(site, user) {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	var album models.Album
	/* Parse json to models.Album */
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotAcceptable, "error", nil)
		return
	}

	album.ShortID, _ = shortid.Generate()
	album.SiteID = site.ID
	album.OwnerID = site.OwnerID
	album.ID = bson.NewObjectId()

	albums, _ := database.GetAlbums(site.ID)
	album.Index = len(albums)

	err = database.AddAlbum(album)
	if err != nil {
		log.Print(err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, "error", nil)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "success", nil)
	return
}

// GetAlbums : return albums of site
func (a *Albums) GetAlbums(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	site, _ := database.GetSiteByName(name)
	user, _ := database.GetUserByID(utils.GetUserObjectID(r))

	if !utils.IsAuthorized(site, user) {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	albums, err := database.GetAlbums(site.ID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "error", nil)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "success", albums)
	return
}

// GetAlbum : return specific album
func (a *Albums) GetAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	id := vars["id"]

	site, _ := database.GetSiteByName(name)
	user, _ := database.GetUserByID(utils.GetUserObjectID(r))

	if !utils.IsAuthorized(site, user) {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	album, err := database.GetAlbum(bson.ObjectIdHex(id))
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "error", nil)
		return
	}

	if album.SiteID != site.ID {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	tracks, _ := database.GetAlbumTracks(album.ID)
	album.Tracks = tracks

	utils.RespondWithJSON(w, http.StatusOK, "success", album)
	return
}

// DeleteAlbum : remove album from db
func (a *Albums) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	id := vars["id"]

	site, _ := database.GetSiteByName(name)
	user, _ := database.GetUserByID(utils.GetUserObjectID(r))

	if !utils.IsAuthorized(site, user) {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	album, err := database.GetAlbum(bson.ObjectIdHex(id))
	if album.SiteID != site.ID {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	for _, track := range album.Tracks {
		database.RemoveTrack(track.ID)
	}

	err = database.RemoveAlbum(album.ID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "error", nil)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "success", nil)
	return
}

// UpdateAlbum : rename & change image
func (a *Albums) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	siteName := vars["name"]
	id := vars["id"]

	site, _ := database.GetSiteByName(siteName)
	user, _ := database.GetUserByID(utils.GetUserObjectID(r))

	if !utils.IsAuthorized(site, user) {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	var album models.Album
	/* Parse json to models.Album */
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotAcceptable, "error", nil)
		return
	}

	na, err := database.GetAlbum(bson.ObjectIdHex(id))
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotFound, "not_found", nil)
		return
	}
	if na.SiteID != site.ID {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	err = database.UpdateAlbum(bson.ObjectIdHex(id), album)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "error", nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "success", nil)
	return
}

// UpdateAlbumsIndxes : update order of albums
func (a *Albums) UpdateAlbumsIndexes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	siteName := vars["name"]

	site, _ := database.GetSiteByName(siteName)
	user, _ := database.GetUserByID(utils.GetUserObjectID(r))

	if !utils.IsAuthorized(site, user) {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	var albums []models.Album
	/* Parse json to models.Album */
	err := json.NewDecoder(r.Body).Decode(&albums)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotAcceptable, "error", nil)
		return
	}

	err = database.UpdateAlbumsIndexes(site.ID, albums)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotAcceptable, "error", nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "success", nil)
	return
}
