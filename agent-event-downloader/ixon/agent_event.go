package ixon

import (
	"agent-event-downloader/util"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Event struct {
	Level     string   `json:"level"`
	Actor     Resource `json:"actor"`
	TimeStamp string   `json:"time"`
	PublicId  string   `json:"publicId"`
	Message   struct {
		TemplateLocalized string     `json:"templateLocalised"`
		TextLocalized     string     `json:"textLocalised"`
		Template          string     `json:"template"`
		Parameters        []Resource `json:"parameters"`
		Text              string     `json:"text"`
	} `json:"message"`
}

type EventListResponse struct {
	ApiResponse
	Data []Event
}

func (s *AgentService) GetEventList(agentId string, afterDate time.Time, beforeDate time.Time) ([]map[string]string, error) {

	parsedEventData := make([]map[string]string, 0)

	// Current date used as "before", will change during loop
	currentBeforeDate := beforeDate.Format(util.ISO8601)

	// Until break...
GetAllEvents:
	for {

		log.Infof("Getting events from %s and before", currentBeforeDate)

		// Retrieve event data
		eventData, err := s.getEventsBefore(agentId, currentBeforeDate)

		if err != nil {
			return nil, err
		}

		// No response
		if len(eventData.Data) == 0 {
			break
		}

		// For every retrieved event
		for _, event := range eventData.Data {

			// Parse time of event
			eventTime, _ := time.Parse(util.ATOM, event.TimeStamp)

			// If event is earlier than the start time, stop processing events
			if eventTime.Sub(afterDate) < 0 {
				break GetAllEvents
			}

			// Append event and (hardcoded) set values
			currDate := make(map[string]string)
			currDate["Level"] = event.Level
			currDate["Actor_Name"] = event.Actor.Name
			currDate["Actor_PublicId"] = event.Actor.PublicId
			currDate["Actor_Type"] = event.Actor.Type
			currDate["TimeStamp"] = event.TimeStamp
			currDate["PublicId"] = event.PublicId
			currDate["Message"] = event.Message.TextLocalized
			parsedEventData = append(parsedEventData, currDate)

			currentBeforeDate = event.TimeStamp
		}

		// Wait a bit to not spam the API too much
		time.Sleep(time.Millisecond * 200)

		// Done processing events for this request, continue with new request with new before date
	}

	return parsedEventData, nil
}

func (s *AgentService) getEventsBefore(agentId string, beforeDate string) (*EventListResponse, error) {

	replacements := map[string]string{"publicId": agentId}
	endpoint, err := s.client.ParseEndpoint("AgentEventList", replacements)

	if err != nil {
		return nil, err
	}

	queryParams := map[string]string{"before": beforeDate}

	res, err := s.client.MakeRequest(http.MethodGet, endpoint, nil, nil, queryParams)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("something went wrong while fetching agent events")
	}

	var listResponse EventListResponse

	err = json.NewDecoder(res.Body).Decode(&listResponse)

	if err != nil {
		return nil, err
	}

	return &listResponse, nil
}
