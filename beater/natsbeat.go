package beater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/nfvsap/natsbeat/config"
)

// Natsbeat configuration.
type Natsbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of natsbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Natsbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts natsbeat.
func (bt *Natsbeat) Run(b *beat.Beat) error {
	logp.Info("natsbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		uri := bt.config.URIs[0]
		body, err := GetJson(
			fmt.Sprintf(
				"http://%s:%d/%s",
				bt.config.NATShost, bt.config.NATSmport, uri))
		if err != nil {
			fmt.Errorf("failed to get NATS monitoring data from (%s): %v", uri, err)
			continue
		}

		data := make(map[string]interface{})
		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Errorf("failed to unmarshal response: %v", err)
			continue
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type": b.Info.Name,
				"varz": data,
			},
		}
		bt.client.Publish(event)
		logp.Info("Event sent")
	}
}

// Stop stops natsbeat.
func (bt *Natsbeat) Stop() {
	bt.client.Close()
	close(bt.done)
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
