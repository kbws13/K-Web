package config

import (
	"KWeb/framework/contract"
	"KWeb/tests"
	. "github.com/smartystreets/goconvey/convey"
	"path/filepath"
	"testing"
)

func TestKConfig_GetInt(t *testing.T) {
	Convey("test KWeb env normal case", t, func() {
		basePath := tests.BasePath
		folder := filepath.Join(basePath, "config")
		serv, err := NewKConfig(folder, map[string]string{}, contract.EnvDevelopment)
		So(err, ShouldBeNil)
		conf := serv.(*KConfig)
		timeout := conf.GetInt("database.mysql.timeout")
		So(timeout, ShouldEqual, 1)
	})
}
