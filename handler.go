package pluginapi

import "net/http"

// Plugin is the interface every plugin must implement.
type Plugin interface {
	RegisterRoutes(mux *http.ServeMux)
}
