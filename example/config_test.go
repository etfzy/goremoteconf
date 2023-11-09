package example

import (
	"fmt"
	"testing"
	"time"

	"github.com/etfzy/goremoteconf"
	"github.com/etfzy/goremoteconf/config"
	"github.com/etfzy/goremoteconf/remote/nacosconf"
)

/*

type Conf struct {
	Namespace string
	Username  string
	Password  string
	Servers   []string
	Port      uint64
}

type Watch struct {
	DataID string
	Group  string
}
*/

func callnacos(namespace, group, dataId, data string, conf config.Config) {
	conf.Publish([]byte(data))
}

func TestConfig(t *testing.T) {
	t.Run("config", func(t *testing.T) {
		remote, err := nacosconf.New(&nacosconf.Conf{
			Namespace: "",
			Username:  "",
			Password:  "",
			Servers:   []string{""},
			Port:      8848,
		})
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}

		watch := &nacosconf.Event{
			DataID:    "",
			Group:     "",
			EventCall: callnacos,
		}

		remoteConf, err := goconf.New(watch, callnacos, remote)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}

		go func() {
			for {
				fmt.Println(remoteConf.GetConfig().GetIntValue("", "", 0))
				time.Sleep(time.Duration(5) * time.Second)
			}
		}()
		time.Sleep(time.Duration(3600) * time.Second)
	})
}
