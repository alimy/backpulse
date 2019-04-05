package serve

import (
	"github.com/backpulse/dashboard/dist"
	"net/http"
)

func newAssets() http.Handler {
	return http.FileServer(dist.AssetFile())
}
