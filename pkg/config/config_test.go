package config_test

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/jomoespe/webc-spikes/pkg/config"

	"testing"
)

func TestCanLoadToml(t *testing.T) {
	tests := [...]struct {
		filename       string
		name           string
		numberOfRoutes int
	}{
		{"../../testdata/config/example-1.toml", "This is a configuration", 2},
	}

	for _, tt := range tests {
		doc, _ := ioutil.ReadFile(tt.filename)

		c := config.NewConfig()
		err := toml.Unmarshal(doc, &c)
		if err != nil {
			t.Fatalf("unmarshalling error: %v", err)
		}

		if c.Name != tt.name {
			t.Fatalf("wrong name. expected: %s, got=%s", tt.name, c.Name)
		}
		if len(c.Routes) != tt.numberOfRoutes {
			t.Fatalf("wrong number of routes. expected: %d, got=%d", tt.numberOfRoutes, len(c.Routes))
		}

		// TODO make more checks (overriden values, default values if not set, set values in file not in struct,...)

		/*
			// iterate config fields
			if "http://mpp.zooplus.com" != c.Routes["/my-pet-profile"].To {
				t.Fatalf("wrong target for route %s. expected: %s, got=%s", "/my-pet-profile", "http://mpp.zooplus.com", c.Routes["/my-pet-profile"].To)
			}
			for routeName, route := range config.Routes {
				t.Errorf("config.%s.target=%s", routeName, route.Target)
				t.Errorf("config.%s.timeout=%s", routeName, route.Timeout)
			}
		*/

	}
}