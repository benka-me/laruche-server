package rpc

import (
	"context"
	"github.com/benka-me/laruche/go-pkg/laruche"
)

func (c *App) SetBee(ctx context.Context, req *laruche.PushBee) (*laruche.PushBeeRes, error) {
	err := c.DB.SetBee(req.Bee)
	return &laruche.PushBeeRes{}, err
}

func (c *App) SetHive(ctx context.Context, req *laruche.PushHive) (*laruche.PushHiveRes, error) {
	err := c.DB.SetHive(req.Hive)
	return &laruche.PushHiveRes{}, err
}
