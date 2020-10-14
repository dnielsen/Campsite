package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (api *API) GetSpeakerById(id string) (*Speaker, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v/%v", api.c.Service.Speaker.Host, api.c.Service.Speaker.Port, id), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		return nil, err
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

func (api *API) GetAllSpeakers() (*[]Speaker, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v", api.c.Service.Speaker.Host, api.c.Service.Speaker.Port), nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	// Make the request.
	res, err := api.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		return nil, err
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

func (api *API) CreateSpeaker(i SpeakerInput) (*Speaker, error) {
	// Marshal the speaker input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal speaker input: %v", err)
		return nil, err
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
		return nil, err
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

func (api *API) EditSpeakerById(id string, i SpeakerInput) error {
	// Marshal the speaker input.
	b, err := json.Marshal(i)
	if err != nil {
		log.Printf("Failed to marshal speaker input: %v", err)
		return err
	}
	// Create a request.
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%v:%v/%v", api.c.Service.Speaker.Host, api.c.Service.Speaker.Port, id), bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return err
	}

	// Make the request.
	if _, err = api.client.Do(req); err != nil {
		log.Printf("Failed to do request: %v", err)
		return err
	}
	return nil
}

func (api *API) DeleteSpeakerById(id string) error {
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