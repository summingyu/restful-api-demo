package conf_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/summingyu/restful-api-demo/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		fmt.Println(conf.C().App.Name)
		should.Equal("demo", conf.C().App.Name)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	os.Setenv("MYSQL_DATABASE", "unit_test")
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("unit_test", conf.C().MySQL.Database)
	}

}

func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		conf.C().MySQL.GetDB()

	}

}
