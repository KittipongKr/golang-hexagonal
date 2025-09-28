package port

import (
	d "csat-servay/internal/core/domain"
)

type OneCall interface {
	GetPwd(req *d.OneAuthResb) (*d.AuthResp, error)
	GetAccount(token string) (*d.AuthAccountResp, error)
}
