package remote

import "github.com/etfzy/goremoteconf/config"

type Remote interface {
	Watch(watch interface{}, conf config.Config) error
	GetConfig(watch interface{}) ([]byte, error)
}
