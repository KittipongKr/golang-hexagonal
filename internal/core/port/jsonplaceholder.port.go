package port

import "context"

type JsonplaceholderApi interface {
	TestGetEndpoint(ctx context.Context) error
}
