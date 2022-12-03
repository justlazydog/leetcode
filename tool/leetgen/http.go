package main

import (
	"io"
	"net/http"
)

// graphqlRequest does a http request to leetcode
func graphqlRequest(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "__cfduid=dae082e425ee3916c04a5170b832e268e1524142659; _ga=GA1.2.1432146910.1524142661; _gid=GA1.2.650843898.1529736240; csrftoken=iSKedVXxGDkBBXbP9chsyXhbIrRedF7iw2EMRZiemtzKD8vjHSWZJKkKQVIwZKp7; __atuvc=2%7C25; __atuvs=5b2ded02313c83c4001; _gat=1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:74.0) Gecko/20100101 Firefox/74.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bts, nil
}
