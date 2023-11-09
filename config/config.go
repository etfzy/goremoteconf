package config

import (
	"errors"
	"strconv"
	"sync"

	"github.com/Unknwon/goconfig"
)

type Config interface {
	GetIntValue(section, key string, df int) (int, error)
	GetStringValue(section, key string, df string) (string, error)
	GetBoolValue(section, key string, df bool) (bool, error)
	Publish(data []byte) error
}

type ConfigInfos struct {
	lock sync.RWMutex
	cf   *goconfig.ConfigFile
}

func New() Config {
	return &ConfigInfos{
		lock: sync.RWMutex{},
		cf:   nil,
	}
}

func (c *ConfigInfos) Publish(data []byte) error {
	cf, err := goconfig.LoadFromData(data)
	if err != nil {
		return err
	}

	if cf != nil {
		c.lock.Lock()
		c.cf = cf
		c.lock.Unlock()
		return nil
	} else {
		return errors.New("config load failed!")
	}
	return nil
}

func (c *ConfigInfos) GetIntValue(section, key string, df int) (int, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.cf == nil {
		return df, errors.New("config is nil")
	}

	vs, err := c.cf.GetValue(section, key)

	if err != nil {
		return df, err
	}

	vi, err := strconv.Atoi(vs)
	if err != nil {
		return df, err
	}

	return vi, nil
}

func (c *ConfigInfos) GetStringValue(section, key string, df string) (string, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.cf == nil {
		return df, errors.New("config is nil")
	}

	vs, err := c.cf.GetValue(section, key)

	if err != nil {
		return df, err
	}

	if vs == "" {
		return df, errors.New("config is Null!")
	}
	return vs, nil
}

func (c *ConfigInfos) GetBoolValue(section, key string, df bool) (bool, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.cf == nil {
		return df, errors.New("config is nil")
	}

	vs, err := c.cf.GetValue(section, key)

	if err != nil {
		return df, err
	}

	if vs == "true" {
		return true, nil
	} else if vs == "false" {
		return false, nil
	}

	return df, errors.New("config is unkonow!")
}
