package audit

import (
	"encoding/json"
	"fmt"
	"io"
)

type Item struct {
	ID           string   `json:"id"`
	Score        *float32 `json:"score"`
	NumericValue *float32 `json:"numericValue"`
}

type Environment struct {
	NetworkUserAgent string `json:"networkUserAgent"`
	HostUserAgent    string `json:"hostUserAgent"`
	BenchmarkIndex   int    `json:"benchmarkIndex"`
}

type Report struct {
	LighthouseVersion string          `json:"lighthouseVersion"`
	RequestedURL      string          `json:"requestedUrl"`
	MainDocumentURL   string          `json:"mainDocumentUrl"`
	FinalDisplayedURL string          `json:"finalDisplayedUrl"`
	FinalURL          string          `json:"finalUrl"`
	FetchTime         string          `json:"fetchTime"`
	GatherMode        string          `json:"gatherMode"`
	UserAgent         string          `json:"userAgent"`
	Environment       Environment     `json:"environment"`
	Audits            map[string]Item `json:"audits"`
}

// NewFromFile creates Report from the provided file reference
func NewFromReader(file io.Reader) (Report, error) {
	var test Report

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return test, fmt.Errorf("could not read file: %s", err)
	}

	err = json.Unmarshal(byteValue, &test)
	if err != nil {
		return test, fmt.Errorf("could not unmarshal struct: %s", err)
	}

	return test, nil
}
