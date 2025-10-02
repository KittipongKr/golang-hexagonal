package call_api

import (
	"errors"
	"strconv"

	cfgs "csat-servay/configs/env"
	h "csat-servay/internal/adapter/calls/helpers"
	m "csat-servay/internal/adapter/calls/models"
	d "csat-servay/internal/core/domain"
	p "csat-servay/internal/core/port"
	logs "csat-servay/pkg/logs"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
)

type raCall struct {
	resty *resty.Client
	raCfg cfgs.RaConfig
}

func NewRaCall(
	resty *resty.Client,
	raCfg cfgs.RaConfig,
) p.RaCall {
	return raCall{
		resty: resty,
		raCfg: raCfg,
	}
}

func (tp raCall) GetsearchAllUser() ([]d.User, error) {

	resp := m.RaSearchAllUserRes{}

	apiPath := "/api/vMonk/externalApi/searchAllUser"

	response, err := tp.resty.R().
		SetHeader("Authorization", "Bearer "+tp.raCfg.TokenSearch).
		SetResult(&resp).
		Post(tp.raCfg.URL + apiPath)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	if response.StatusCode() != 200 {
		spew.Dump(response)
		err := errors.New(apiPath + " : status - " + strconv.Itoa(response.StatusCode()))
		logs.Error(err)
		return nil, err
	}

	return h.ConvRaUserToUser(resp.Result), nil
}
