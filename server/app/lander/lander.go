package lander

import (
	"errors"
)

//数据库通用接口：
	//1：增
	//2：删
type LanderInterface interface {
	Land(m Models) error
	Remove(m Models) error

	Run()
	StopSafe()
}

var DefaultLander LanderInterface

func SetLander(l LanderInterface) {
	DefaultLander = l
}

func GetLander() LanderInterface {
	return DefaultLander
}

func Lander(m Models) error {
	if DefaultLander == nil {
		return errors.New("default lander is nil")
	}

	return DefaultLander.Land(m)
}

func Remove(m Models) error {
	if DefaultLander == nil {
		return errors.New("default lander is nil")
	}

	return DefaultLander.Remove(m)
}