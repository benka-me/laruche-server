package rpc

import (
	"context"
	"github.com/benka-me/laruche/go-pkg/laruche"
)

func (c *App) PublishBee(ctx context.Context, req *laruche.PushBee) (*laruche.PushBeeRes, error) {
	req.Bee.Public = true
	nps, err := c.DB.InsertBee(req.Bee)
	return &laruche.PushBeeRes{Id: nps}, err
}

func (c *App) PublishHive(ctx context.Context, req *laruche.PushHive) (*laruche.PushHiveRes, error) {
	req.Hive.Public = true
	nps, err := c.DB.InsertHive(req.Hive)
	return &laruche.PushHiveRes{Id: nps}, err
}

func (c *App) PrivatizeBee(ctx context.Context, req *laruche.PushBee) (*laruche.PushBeeRes, error) {
	req.Bee.Public = false
	nps, err := c.DB.InsertBee(req.Bee)
	return &laruche.PushBeeRes{Id: nps}, err
}

func (c *App) PrivatizeHive(ctx context.Context, req *laruche.PushHive) (*laruche.PushHiveRes, error) {
	req.Hive.Public = false
	nps, err := c.DB.InsertHive(req.Hive)
	return &laruche.PushHiveRes{Id: nps}, err
}
