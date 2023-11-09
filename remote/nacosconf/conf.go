package nacosconf

import (
	"github.com/etfzy/goremoteconf/config"
	"github.com/etfzy/goremoteconf/remote"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Nacos struct {
	Client config_client.IConfigClient
}

type Conf struct {
	Namespace string
	Username  string
	Password  string
	Servers   []string
	Port      uint64
}

type Event struct {
	DataID    string
	Group     string
	EventCall EventCall
}

type EventCall func(string, string, string, string, config.Config)

func (n *Nacos) Watch(event interface{}, conf config.Config) error {
	w := event.(*Event)
	err := n.Client.ListenConfig(vo.ConfigParam{
		DataId: w.DataID,
		Group:  w.Group,
		OnChange: func(namespace, group, dataId, data string) {
			w.EventCall(namespace, group, dataId, data, conf)
		},
	})
	return err
}

func (n *Nacos) GetConfig(watch interface{}) ([]byte, error) {
	w := watch.(*Event)
	content, err := n.Client.GetConfig(vo.ConfigParam{
		DataId: w.DataID,
		Group:  w.Group})
	return []byte(content), err
}

func New(conf interface{}) (remote.Remote, error) {
	temp := &Nacos{}
	c := conf.(*Conf)
	// 创建一个Nacos客户端
	clientConfig := constant.ClientConfig{
		NamespaceId:         c.Namespace,
		NotLoadCacheAtStart: true,
		Username:            c.Username,
		Password:            c.Password,
		LogDir:              "./nacoscache/register/log",
		CacheDir:            "./nacoscache/register/cache",
		LogLevel:            "error",
		OpenKMS:             false,
	}
	serverConfigs := []constant.ServerConfig{}
	for _, v := range c.Servers {
		serverConfigs = append(serverConfigs, constant.ServerConfig{
			Scheme:      "http",
			ContextPath: "/nacos",
			IpAddr:      v,
			Port:        c.Port,
		})
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		return nil, err
	}
	temp.Client = configClient

	return temp, nil

}
