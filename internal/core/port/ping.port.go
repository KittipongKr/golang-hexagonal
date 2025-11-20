package port

type PingRepo interface {
	GetPongRepo() (string, error)
}

type PingServ interface {
	GetPongService() (string, error)
	GetJsonplaceholderService() error
}
