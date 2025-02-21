package services

import (
	"fmt"
	"net/url"
	"time"
)

// Convert provided date to epoch milliseconds
func getDateInEpochMilliseconds(date string) (int64, error) {
	var layouts = []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
	}
	for _, layout := range layouts {
		parsedTime, err := time.Parse(layout, date)
		if err == nil {
			return parsedTime.UnixNano() / int64(time.Millisecond), nil
		}
	}

	return 0, fmt.Errorf("failed to parse date '%s' with any of the supported formats: %v", date, layouts)
}

// Add query params to an endpoint
func addParamsToEndpoint(endpoint string, params map[string]string) (string, error) {
	if len(params) == 0 {
		return endpoint, nil
	}

	processedParams := make(map[string]string)
	for k, v := range params {
		if k == "from" || k == "to" {
			epoch, err := getDateInEpochMilliseconds(v)
			if err != nil {
				return "", fmt.Errorf("params handling error %s: %w", k, err)
			}
			processedParams[k] = fmt.Sprintf("%d", epoch)
		} else {
			processedParams[k] = v
		}
	}

	urlParams := url.Values{}
	for k, v := range processedParams {
		urlParams.Set(k, v)
	}

	return fmt.Sprintf("%s?%s", endpoint, urlParams.Encode()), nil
}
