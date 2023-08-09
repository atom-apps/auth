package auth

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		door *casdoorsdk.Client,
		micro contracts.MicroService,
	) (*UserService, error) {
		obj := &UserService{
			door:  door,
			micro: micro,
		}
		if err := obj.Prepare(); err != nil {
			return nil, err
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
