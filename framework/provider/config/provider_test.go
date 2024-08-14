package config

import (
	"KWeb/framework"
	"KWeb/framework/contract"
	"KWeb/framework/provider/app"
	"KWeb/framework/provider/env"
	"KWeb/tests"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestKConfig_Normal(t *testing.T) {
	Convey("test KWeb config normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewKContainer()
		c.Bind(&app.KAppProvider{BaseFolder: basePath})
		c.Bind(&env.KEnvProvider{})

		err := c.Bind(&KConfigProvider{})
		So(err, ShouldBeNil)

		conf := c.MustMake(contract.ConfigKey).(contract.Config)
		So(conf.GetString("database.mysql.hostname"), ShouldEqual, "127.0.0.1")
		So(conf.GetInt("database.mysql.timeout"), ShouldEqual, 1)
		So(conf.GetFloat64("database.mysql.readtime"), ShouldEqual, 2.3)
		// So(conf.GetString("database.mysql.password"), ShouldEqual, "mypassword")

		maps := conf.GetStringMap("database.mysql")
		So(maps, ShouldContainKey, "hostname")
		So(maps["timeout"], ShouldEqual, 1)

		maps2 := conf.GetStringMapString("databse.mysql")
		So(maps2["timeout"], ShouldEqual, "")

		type Mysql struct {
			Hostname string
			Username string
		}
		ms := &Mysql{}
		err = conf.Load("database.mysql", ms)
		Println(ms)
		So(err, ShouldBeNil)
		So(ms.Hostname, ShouldEqual, "127.0.0.1")
	})
}
