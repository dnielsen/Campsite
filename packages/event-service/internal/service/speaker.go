package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Speaker struct {
	ID         string    `json:"id" gorm:"type:uuid"`
	Name       string    `json:"name" gorm:"not null"`
	Bio        string    `json:"bio" gorm:"not null"`
	Headline   string    `json:"headline" gorm:"not null"`
	Photo      string    `json:"photo" gorm:"not null"`
	Sessions   []Session `json:"sessions,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}

type SpeakerInput struct {
	// Name is a required field with a minimum and maximum length of 2 and 50 respectively.
	Name     string `json:"name,omitempty" validate:"required,min=2,max=50"`
	Bio      string `json:"bio,omitempty" validate:"required,min=20,max=2000"`
	Headline string `json:"headline,omitempty" validate:"required,min=2,max=30"`
	Photo    string `json:"photo,omitempty" validate:"required,min=10,max=150"`
}

func (api *api) GetSpeakerById(id string) (*Speaker, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v/%v", api.c.Service.Speaker.Host, api.c.Service.Speaker.Port, id), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
	}

	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	// Unmarshal the received body bytes
	var speaker Speaker
	if err = json.Unmarshal(readBytes, &speaker); err != nil {
		log.Printf("Failed to unmarshal speaker body: %v", err)
		return nil, err
	}

	return &speaker, nil
}

func (api *api) GetAllSpeakers() (*[]Speaker, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v", api.c.Service.Speaker.Host, api.c.Service.Speaker.Port), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
	}

	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	// Unmarshal the received body bytes
	var speakers []Speaker
	if err = json.Unmarshal(readBytes, &speakers); err != nil {
		log.Printf("Failed to unmarshal speaker body: %v", err)
		return nil, err
	}

	return &speakers, nil
}

func (api *api) CreateSpeaker(i SpeakerInput) (*Speaker, error) {
	// Marshal the speaker input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal speaker input: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v", api.c.Service.Speaker.Host, api.c.Service.Speaker.Port), bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
	}

	// Read the response body.
	readBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	// Unmarshal the received body bytes.
	var speaker Speaker
	if err = json.Unmarshal(readBytes, &speaker); err != nil {
		log.Printf("Failed to unmarshal speaker body: %v", err)
		return nil, err
	}

	return &speaker, nil
}



func (api *api) DeleteSpeakerById(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%v:%v/%v", api.c.Service.Speaker.Host, api.c.Service.Speaker.Port, id), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return err
	}

	// Make the request.
	if _, err := api.client.Do(req); err != nil {
		log.Printf("Failed to do request: %v", err)
		return err
	}

	return nil
}