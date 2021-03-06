package upcrm

import (
	"flag"
	"os"
	"testing"

	"go-common/app/admin/main/up/conf"
	"go-common/library/net/http/blademaster"

	"gopkg.in/h2non/gock.v1"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	if os.Getenv("DEPLOY_ENV") != "" {
		flag.Set("app_id", "main.archive.up-admin")
		flag.Set("conf_token", "930697bb7def4df0713ef8080596b863")
		flag.Set("tree_id", "36438")
		flag.Set("conf_version", "1")
		flag.Set("deploy_env", "uat")
		flag.Set("conf_host", "config.bilibili.co")
		flag.Set("conf_path", "/tmp")
		flag.Set("region", "sh")
		flag.Set("zone", "sh001")
	} else {
		flag.Set("conf", "../../cmd/up-admin.toml")
	}
	if os.Getenv("UT_LOCAL_TEST") != "" {
		flag.Set("conf", "../../cmd/up-admin.toml")
	}
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
	d.httpClient = blademaster.NewClient(conf.Conf.HTTPClient.Normal)
	d.httpClient.SetTransport(gock.DefaultTransport)
	os.Exit(m.Run())
}
