package gzip

import (
	json "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/spiral/roadrunner/plugins"
	rrhttp "github.com/spiral/roadrunner/plugins/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCfg struct {
	gzip    string
	httpCfg string
	target  string
}

func (cfg *testCfg) Get(name string) plugins.Config {
	if name == rrhttp.ID {
		return &testCfg{target: cfg.httpCfg}
	}

	if name == ID {
		return &testCfg{target: cfg.gzip}
	}
	return nil
}
func (cfg *testCfg) Unmarshal(out interface{}) error {
	j := json.ConfigCompatibleWithStandardLibrary
	return j.Unmarshal([]byte(cfg.target), out)
}

func Test_Disabled(t *testing.T) {
	logger, _ := test.NewNullLogger()
	logger.SetLevel(logrus.DebugLevel)

	c := plugins.NewContainer(logger)
	c.Register(ID, &Service{cfg: &Config{Enable: true}})

	assert.NoError(t, c.Init(&testCfg{
		httpCfg: `{
			"address": ":6029",
			"workers":{
				"command": "php ../../tests/http/client.php echo pipes",
			}
	}`,
		gzip: `{"enable":false}`,
	}))

	s, st := c.Get(ID)
	assert.NotNil(t, s)
	assert.Equal(t, plugins.StatusInactive, st)
}

// TEST bug #275
func Test_Bug275(t *testing.T) {
	logger, _ := test.NewNullLogger()
	logger.SetLevel(logrus.DebugLevel)

	c := plugins.NewContainer(logger)
	c.Register(ID, &Service{})

	assert.Error(t, c.Init(&testCfg{
		httpCfg: "",
		gzip:    `{"enable":true}`,
	}))

	s, st := c.Get(ID)
	assert.NotNil(t, s)
	assert.Equal(t, plugins.StatusInactive, st)
}
