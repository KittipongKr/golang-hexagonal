package service

import (
	envCfgs "csat-servay/configs/env"
	d "csat-servay/internal/core/domain"
	p "csat-servay/internal/core/port"
	"csat-servay/pkg/jwt"
	"csat-servay/pkg/logs"
)

type authServ struct {
	appConfig envCfgs.AppConfig
	jwtConfig envCfgs.JwtConfig
	userRepo  p.UserRepo
	oneCall   p.OneCall
}

func NewAuthServ(
	appConfig envCfgs.AppConfig,
	jwtConfig envCfgs.JwtConfig,
	userRepo p.UserRepo,
	oneCall p.OneCall,
) p.AuthServ {
	return authServ{
		appConfig: appConfig,
		jwtConfig: jwtConfig,
		userRepo:  userRepo,
		oneCall:   oneCall,
	}
}

func (a authServ) LoginOneIdService(user *d.OneAuthResb) (*d.AuthResp, error) {
	onePwd, err := a.oneCall.GetPwd(user)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	account, err := a.oneCall.GetAccount(onePwd.AccessToken)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	_ = account

	accessPayload := map[string]interface{}{
		"sub":  account.ID,
		"role": []string{"admin", "user"},
		"iss":  a.appConfig.AppName,
	}
	accessToken, err := jwt.SignTokenHMAC(a.jwtConfig.HMACSecret, accessPayload, a.jwtConfig.AccessExp)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	refreshPayload := map[string]interface{}{
		"sub": account.ID,
		"iss": a.appConfig.AppName,
	}
	refreshToken, err := jwt.SignTokenHMAC(a.jwtConfig.HMACSecret, refreshPayload, a.jwtConfig.RefreshExp)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	resp := d.AuthResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AccountID:    account.ID,
	}

	return &resp, nil
}
