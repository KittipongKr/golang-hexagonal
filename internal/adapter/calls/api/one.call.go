package call_api

import (
	envCfgs "csat-servay/configs/env"
	m "csat-servay/internal/adapter/calls/models"
	d "csat-servay/internal/core/domain"
	p "csat-servay/internal/core/port"
	"csat-servay/pkg/logs"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

type oneCall struct {
	resty  *resty.Client
	oneCfg envCfgs.OneConfig
}

func NewOneCall(
	resty *resty.Client,
	oneCfg envCfgs.OneConfig,
) p.OneCall {
	return oneCall{
		resty:  resty,
		oneCfg: oneCfg,
	}
}

func (tp oneCall) GetPwd(req *d.OneAuthResb) (*d.AuthResp, error) {
	reqb := m.OneAuthReq{
		GrantType:    tp.oneCfg.GrantType,
		ClientID:     tp.oneCfg.ClientID,
		ClientSecret: tp.oneCfg.ClientSecret,
		Username:     req.Username,
		Password:     req.Password,
	}

	resp := m.OneAuthRes{}
	response, err := tp.resty.R().
		SetBody(reqb).
		SetResult(&resp).
		Post(tp.oneCfg.URL + "/api/oauth/getpwd")
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	if response.StatusCode() != 200 {
		logs.Dubug(response)
		errMsg := fmt.Sprintf("/api/oauth/getpwd {status - %d}", response.StatusCode())
		return nil, errors.New(errMsg)
	}

	res := d.AuthResp{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		AccountID:    resp.AccountID,
	}

	return &res, nil
}

func (tp oneCall) GetAccount(token string) (*d.AuthAccountResp, error) {
	resp := m.OneAccountResp{}
	response, err := tp.resty.R().
		SetHeader("Authorization", "Bearer "+token).
		SetResult(&resp).
		Get(tp.oneCfg.URL + "/api/account")
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	if response.StatusCode() != 200 {
		logs.Dubug(response)
		errMsg := fmt.Sprintf("/api/account {status - %d}", response.StatusCode())
		return nil, errors.New(errMsg)
	}

	res := d.AuthAccountResp{}
	if err := copier.Copy(&res, &resp); err != nil {
		logs.Error(err)
		return nil, err
	}

	return &res, nil
}
