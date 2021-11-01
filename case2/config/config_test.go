package config_test

import (
	"io/ioutil"
	"os"
	"testing"

	. "case2/config"

	"github.com/stretchr/testify/assert"
)

var cfg = []byte(`
db:
  host: "localhost"
  port: "3306"
  username: "test"
  password: "testing1234"
  dbname: "name"
omdb:
  apikey: "apiTest"
`)

func TestLoadConfig(t *testing.T) {
	err := ioutil.WriteFile("/tmp/config.yaml", cfg, 0644)
	if err != nil {
		t.Fail()
		t.Logf("Can't prepare temporary file for test. Got `%v`", err)
	}

	t.Run("Test read config from a file", func(t *testing.T) {
		conf, _ := LoadConfig("/tmp/config.yaml")
		assert.Equal(t, "localhost", conf.DB.Host)
		assert.Equal(t, "3306", conf.DB.Port)
		assert.Equal(t, "test", conf.DB.UserName)
		assert.Equal(t, "testing1234", conf.DB.Password)
		assert.Equal(t, "name", conf.DB.DBName)
		assert.Equal(t, "apiTest", conf.Omdb.ApiKey)
	})

	err = os.Remove("/tmp/config.yaml")
	if err != nil {
		t.Fail()
		t.Logf("Can't clean up test resource")
	}
}
