package healthcheck

import (
	"fmt"
	"kirby/httputil"
	"net/http"

	"github.com/jinzhu/gorm"
)

// CheckDatabaseConnection verifies successful database connection
func CheckDatabaseConnection(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var result DatabaseCheckResult
		db.Raw("SELECT current_database()").Scan(&result)
		if result != (DatabaseCheckResult{}) {
			httputil.RespondWithJSON(w, http.StatusOK, Healthcheck{
				Status:  "ok",
				Message: fmt.Sprintf("Database connection to '%s' succeeded", result.CurrentDatabase),
			})
			return
		}
		httputil.RespondWithJSON(w, http.StatusOK, Healthcheck{
			Status:  "failed",
			Message: "Database connection failed",
		})
	}
	return handler
}
