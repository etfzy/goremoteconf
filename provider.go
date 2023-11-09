package goconf

import (
	"github.com/etfzy/goremoteconf/config"
	"github.com/etfzy/goremoteconf/remote"
)

type RemoteConf interface {
	GetConfig() config.Config
}

type RemoteConfInfos struct {
	conf   config.Config
	remote remote.Remote
}

func New(watch interface{}, remote remote.Remote) (RemoteConf, error) {
	temp := &RemoteConfInfos{
		conf:   config.New(),
		remote: remote,
	}

	content, err := temp.remote.GetConfig(watch)

	if err != nil {
		return nil, err
	}

	temp.conf.Publish([]byte(content))

	err = temp.remote.Watch(watch, temp.conf)
	if err != nil {
		return nil, err
	}

	return temp, err
}

func (r *RemoteConfInfos) GetConfig() config.Config {
	return r.conf
}
