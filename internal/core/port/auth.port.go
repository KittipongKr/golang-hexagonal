package port

import (
	d "csat-servay/internal/core/domain"
)

// NOTE: auth service
type AuthServ interface {
	// NOTE: login service
	LoginOneIdService(user *d.OneAuthResb) (*d.AuthResp, error)
}
