package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
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
