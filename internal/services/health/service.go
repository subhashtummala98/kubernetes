package health

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

const healthEndpoint = "/health"
const livenessEndpoint = "/live"
const readinessEndpoint = "/ready"

type healthService struct {
	// db is a client that holds the connection to database
	db *gorm.DB
}

// NewHealthService Returns new healthService
func NewHealthService(db *gorm.DB) *healthService {
	return &healthService{
		db: db,
	}
}

// RegisterBanterServiceEndpoints registers endpoint and handlers for banter service
func (c *healthService) RegisterHealthServiceEndpoints() {
	http.HandleFunc(healthEndpoint+readinessEndpoint, c.readinessHandler)
	http.HandleFunc(healthEndpoint+livenessEndpoint, c.livenessHandler)
}

// readiness handler returns 200 OK if the database is connected
func (c *healthService) readiness(w http.ResponseWriter, r *http.Request) {

	err := c.db.DB().Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Database is not connected"))
		if err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		log.Println(err)
	}
}

// liveness handler returns 200 if the endpoint is accessible
func (c *healthService) liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Println(err)
	}
}
