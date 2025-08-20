package facebook

import (
	"errors"
	"io"
	"net/http"
)

func Tokens() (values map[string]string, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://web.facebook.com/qsdqs5d4", nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:141.0) Gecko/20100101 Firefox/141.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Cookie", "datr=kVakaLiedE6WzTZhtNgKYeCf; fr=0kdUhDIbHw9HybHJp..BopFaR..AAA.0.0.BopFak.AWfZb9y5gjge_a5qr7LX0Jlv-3c; sb=kVakaOa0EvHQIbTG1yuNXrMm; wd=856x963; ps_l=1; ps_n=1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("TE", "trailers")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	text := string(body)
	values = make(map[string]string)
	for name, rgx := range CompiledRegex {
		matches := rgx.FindStringSubmatch(text)
		if matches == nil {
			err = errors.New("failed to scrap fb token")
			return
		}
		values[name] = matches[1]

	}
	return
}
