package healthcheck

import (
	"kirby/httputil"
	"net/http"
)

// DatabaseCheckResult struct
type DatabaseCheckResult struct {
	CurrentDatabase string
}

// CheckHTTPConnection verifies that the server is accepting HTTP connections
func CheckHTTPConnection(w http.ResponseWriter, r *http.Request) {
	httputil.RespondWithJSON(w, http.StatusOK, Healthcheck{
		Status:  "ok",
		Message: "Server is accepting connections",
	})
}
