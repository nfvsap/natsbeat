package beater

import (
	"encoding/json"
	"fmt"
	"github.com/nfvsap/natsbeat/config"
	"io/ioutil"
	"net/http"
)

func RetrieveData(config config.Config, uri string) (*map[string]interface{}, error){

	body, err := GetJson(
		fmt.Sprintf(
			"http://%s:%d/%s",
			config.NATShost, config.NATSmport, uri))
	if err != nil {
		return nil, fmt.Errorf("failed to get NATS monitoring data from (%s): %v", uri, err)
	}

	data := make(map[string]interface{})
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &data, nil
}


func GetJson(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http get failed: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body %v", err)
	}

	return body, nil
}
