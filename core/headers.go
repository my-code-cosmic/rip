package core

import (
	"regexp"
	"sync"
)

// This will match <first-group>:<second-group> of a header
// Test it here https://regex101.com/r/cn0QvY/1
const (
	PATTERN = "([a-zA-Z0-9-]+)[^*](.*)"
)

// Pre define the compiler
var pattern *regexp.Regexp = regexp.MustCompile(PATTERN)

type headers struct {
	entries map[string]string
	mu      sync.Mutex
}

// Holder for our singelton instance
var headerInstance *headers

var once sync.Once

func ParseHeaders(args Arguments) *headers {
	// This will run only once
	once.Do(func() {
		headerInstance = &headers{
			entries: make(map[string]string),
		}

		headerInstance.mu.Lock()
		defer headerInstance.mu.Unlock()

		for _, line := range args.Headers() {
			header := pattern.FindStringSubmatch(line)

			headerInstance.Add(header[1], header[2])
		}

		if args.IsJSONPayload() {
			headerInstance.Add("Content-Type", "application/json; charset=UTF-8")
		}

		headerInstance.Add("User-Agent", "Rest In Peace")
	})

	return headerInstance
}

func HeaderInstance() *headers {
	return headerInstance
}

func (h *headers) Add(key string, value string) {
	h.entries[key] = value
}

func (h *headers) Headers() map[string]string {
	return h.entries
}
