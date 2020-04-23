package rpc

import (
	"context"
	"github.com/benka-me/laruche/go-pkg/laruche"
)

func (c *App) GetBee(ctx context.Context, req *laruche.BeeReq) (*laruche.Bee, error) {
	return c.DB.GetBee(laruche.Namespace(req.BeeName))
}

func (c *App) GetBees(ctx context.Context, req *laruche.BeesReq) (*laruche.Bees, error) {
	bees, err := c.DB.GetAllBees(ctx)
	return &laruche.Bees{
		Bees:          bees,
		StatusMessage: "",
	}, err
}
