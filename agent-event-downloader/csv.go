package main

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"os"
	"sort"
)

func writeHeaders(csvWriter *csv.Writer, headers []string) ([]string, error) {

	// Copy headers to new slice and sort them
	sortedHeaders := make([]string, len(headers))
	copy(sortedHeaders, headers)
	sort.Strings(sortedHeaders)

	// Write the headers to the file
	err := csvWriter.Write(sortedHeaders)

	if err != nil {
		return nil, err
	}

	csvWriter.Flush()
	return sortedHeaders, nil
}

func writeData(csvWriter *csv.Writer, headers []string, eventData []map[string]string) error {
	// Every event
	for _, event := range eventData {
		// Slice for storing all values
		values := make([]string, 0, 1+len(event))

		// Every value of this event
		for _, k := range headers {

			// Get value by header
			v := event[k]
			values = append(values, v)
		}

		// Write to buffer
		err := csvWriter.Write(values)

		if err != nil {
			return err
		}
	}

	// Flush writer and get errors
	csvWriter.Flush()
	err := csvWriter.Error()

	if err != nil {
		return err
	}

	return nil
}

func openOrCreateFile(path string) (*os.File, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Info("Creating file")
		createdFile, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		return createdFile, nil
	}
	openedFile, err := os.OpenFile(path, os.O_RDWR, 0660)

	if err != nil {
		return nil, err
	}

	return openedFile, err
}
