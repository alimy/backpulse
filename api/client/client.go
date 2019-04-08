package client

import (
	"net/http"

	"github.com/alimy/mir"
	"github.com/backpulse/core/utils"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	group             mir.Group `mir:"{name}"`
	greetings         mir.Get   `mir:"/"`
	getContact        mir.Get   `mir:"/contact"`
	getAbout          mir.Get   `mir:"/about"`
	getDefaultGallery mir.Get   `mir:"/galleries/default"`
	getGalleries      mir.Get   `mir:"/galleries"`
	getGallery        mir.Get   `mir:"/gallery/{short_id}"`
	getProjects       mir.Get   `mir:"/projects"`
	getProject        mir.Get   `mir:"/projects/{short_id}"`
	getArticles       mir.Get   `mir:"/articles"`
	getArticle        mir.Get   `mir:"/articles/{short_id}"`
	getVideoGroups    mir.Get   `mir:"/videogroups"`
	getVideoGroup     mir.Get   `mir:"/videogroups/{short_id}"`
	getVideo          mir.Get   `mir:"/videos/{short_id}"`
	getAlbums         mir.Get   `mir:"/albums"`
	getAlbum          mir.Get   `mir:"/albums/{short_id}"`
	getTrack          mir.Get   `mir:"/albums/{short_id}"`
}

// Greetings : return greetings content
func (c *Client) Greetings(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, "Welcome to the API", bson.M{
		"wrapper": "https://github.com/backpulse/wrapper",
	})
}
