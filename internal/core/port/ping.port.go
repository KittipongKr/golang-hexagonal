package port

import "context"

type PingRepo interface {
	GetPongRepo() (string, error)
}

type PingServ interface {
	GetPongService(ctx context.Context) (string, error)
	GetJsonplaceholderService(ctx context.Context) error
}
