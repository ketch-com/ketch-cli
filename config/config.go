package config

import (
	"context"
	"go.ketch.com/lib/orlop"
)

type Config struct {
	URL   string
	Host  string
	Port  int32
	TLS   orlop.TLSConfig
	Vault orlop.VaultConfig
}

var marker = &Config{}

func AddToContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, marker, cfg)
}

func GetFromContext(ctx context.Context) *Config {
	if x := ctx.Value(marker); x != nil {
		if v, ok := x.(*Config); ok {
			return v
		}
	}
	return nil
}
