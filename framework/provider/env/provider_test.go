package env

import (
	"KWeb/framework"
	"KWeb/framework/contract"
	"KWeb/framework/provider/app"
	"KWeb/tests"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestKEnvProvider(t *testing.T) {
	Convey("test KWeb env normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewKContainer()
		sp := &app.KAppProvider{BaseFolder: basePath}

		err := c.Bind(sp)
		So(err, ShouldBeNil)

		sp2 := &KEnvProvider{}
		err = c.Bind(sp2)
		So(err, ShouldBeNil)

		envServ := c.MustMake(contract.EnvKey).(contract.Env)
		So(envServ.AppEnv(), ShouldEqual, "development")
		// So(envServ.Get("DB_HOST"), ShouldEqual, "127.0.0.1")
		// So(envServ.AppDebug(), ShouldBeTrue)
	})
}
