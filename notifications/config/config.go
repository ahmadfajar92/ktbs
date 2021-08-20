package config

import (
	"os"
	"strconv"
)

type (
	// Config interface
	Config interface {
		Set(k string, v interface{})
		Get(k string) interface{}
	}

	config struct {
		vars map[string]interface{}
	}
)

// Load func
func Load() Config {
	cfg := new(config)
	cfg.vars = make(map[string]interface{})

	cfg.vars["PORT"] = os.Getenv("PORT")
	cfg.vars["DEBUG"], _ = strconv.ParseBool(os.Getenv("DEBUG"))
	cfg.vars["TIMEZONE"] = os.Getenv("TIMEZONE")

	return cfg
}

func (cfg *config) Get(k string) interface{} {
	return cfg.vars[k]
}

func (cfg *config) Set(k string, v interface{}) {
	cfg.vars[k] = v
}
