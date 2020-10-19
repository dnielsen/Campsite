package service

import (
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/model"
	"gorm.io/gorm"
	"net/http"
)

type API struct {
	db     *gorm.DB
	client HttpClient
	c      *config.Config
}

func NewAPI(db *gorm.DB, c *config.Config) *API {
	// We define our own HttpClient to enable mocking (for easier testing).
	// We don't have tests yet, however, it's a common practice to do that
	// for this reason.
	var client HttpClient = http.DefaultClient
	return &API{db, client, c}
}

// We define our own interface so that we can mock it,
// and therefore test our fetch functions.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type EventAPI interface {
	GetAllEvents() (*[]model.Event, error)
	CreateEvent(i model.EventInput) (*model.Event, error)
	GetEventById(id string) (*model.Event, error)
	EditEventById(id string, i model.EventInput) error
	DeleteEventById(id string) error
}
