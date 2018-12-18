// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period time.Duration `config:"period"`
	URIs []string `config:"uris"`
	NATShost string `config:"natshost"`
	NATSmport int `config:"natsmport"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
	URIs: []string{},
	NATShost: "localhost",
	NATSmport: 8222,
}
