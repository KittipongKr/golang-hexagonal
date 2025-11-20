package repository

import (
	p "csat-servay/internal/core/port"
)

type pingRepo struct {
}

func NewPingRepo() p.PingRepo {
	return pingRepo{}
}

func (r pingRepo) GetPongRepo() (string, error) {
	return "pong", nil
}
