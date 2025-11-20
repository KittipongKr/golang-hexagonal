package service

import "csat-servay/internal/core/port"

type pingServ struct {
	pingRepo           port.PingRepo
	jsonplaceholderApi port.JsonplaceholderApi
}

func NewPingServ(
	pingRepo port.PingRepo,
	jsonplaceholderApi port.JsonplaceholderApi,
) port.PingServ {
	return pingServ{
		pingRepo:           pingRepo,
		jsonplaceholderApi: jsonplaceholderApi,
	}
}

func (s pingServ) GetPongService() (string, error) {
	resp, err := s.pingRepo.GetPongRepo()
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (s pingServ) GetJsonplaceholderService() error {
	if err := s.jsonplaceholderApi.TestGetEndpoint(); err != nil {
		return err
	}

	return nil
}
