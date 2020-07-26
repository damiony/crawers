package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0")
	resp, err := http.DefaultClient.Do(req)
	// resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	return result, err
}
