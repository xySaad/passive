package parser

import (
	"bufio"
	"io"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"strings"
)

// ParseRequest reads an .md file containing a raw HTTP request
// and converts it into an *http.Request.
func ParseRequest(mdPath string) (*http.Request, error) {
	// Read whole file
	data, err := os.ReadFile(mdPath)
	if err != nil {
		return nil, err
	}

	text := string(data)

	// Find beginning of raw request (e.g. "POST /...")
	start := strings.Index(text, "POST ")
	if start == -1 {
		return nil, io.EOF
	}
	raw := text[start:]

	reader := bufio.NewReader(strings.NewReader(raw))
	tp := textproto.NewReader(reader)

	// Request line (e.g. "POST /path HTTP/2")
	line, err := tp.ReadLine()
	if err != nil {
		return nil, err
	}
	parts := strings.SplitN(line, " ", 3)
	if len(parts) < 3 {
		return nil, io.ErrUnexpectedEOF
	}
	method, requestURI, proto := parts[0], parts[1], parts[2]

	// Headers
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		return nil, err
	}
	headers := http.Header(mimeHeader)

	// Read the remaining body
	bodyBytes, _ := io.ReadAll(reader)
	bodyStr := strings.TrimSpace(string(bodyBytes))

	// Build request
	reqURL := "https://web.facebook.com" + requestURI
	parsedURL, err := url.Parse(reqURL)
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method: method,
		URL:    parsedURL,
		Proto:  proto,
		Header: headers,
		Body:   io.NopCloser(strings.NewReader(bodyStr)),
	}

	return req, nil
}
