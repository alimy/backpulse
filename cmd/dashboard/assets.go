package serve

import (
	"github.com/alimy/backpulse-dashboard/dist"
	"net/http"
)

func newAssets() http.Handler {
	return http.FileServer(dist.AssetFile())
}
