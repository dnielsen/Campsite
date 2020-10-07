package service

import (
	"campsite/packages/event-service/internal/config"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
	"os"
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
	GetAllEvents() (*[]Event, error)
	CreateEvent(i EventInput) (*Event, error)
	GetEventById(id string) (*Event, error)
	EditEventById(id string, i EventInput) error
	DeleteEventById(id string) error
}

type SessionAPI interface {
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput) (*Session, error)
	GetSessionById(id string) (*Session, error)
	EditSessionById(id string, i SessionInput) error
	DeleteSessionById(id string) error
}

type SpeakerAPI interface {
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	EditSpeakerById(id string, i SpeakerInput) error
	DeleteSpeakerById(id string) error
}

type ImageAPI interface {
	GetImage(filename string) (*os.File, error)
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader, host string) (*Upload, error)
}