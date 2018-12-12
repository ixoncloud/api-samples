package ixon

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	DiscoveryUrl = "https://api.ixon.net"
)

type DiscoveryService service

func (s *DiscoveryService) DiscoverApiEndpoints() error {

	// Make discovery request
	res, err := s.client.MakeRequest(
		http.MethodGet,
		DiscoveryUrl,
		nil,
		nil,
		nil,
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		s.client.LogApiError(res, "Could not discover API endpoints")
		return errors.New("could not discover API endpoints")
	}

	var discoveryResponse ApiResponse

	err = json.NewDecoder(res.Body).Decode(&discoveryResponse)

	if err != nil {
		return err
	}

	s.client.apiEndpoints = make(map[string]string)

	for _, link := range discoveryResponse.Links {
		s.client.apiEndpoints[link.Rel] = link.Href
	}

	return nil
}
