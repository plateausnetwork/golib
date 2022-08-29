package mongodb

import (
	"time"
)

type (
	Options struct {
		URI          string
		DatabaseName string
		CtxTimeout   time.Duration
		IsReader     bool
	}
)

func (o *Options) SetDefaults() {
	if o.CtxTimeout == 0 {
		o.CtxTimeout = 10 * time.Second
	}
	if o.URI == "" {
		o.URI = "mongodb://localhost:27017"
	}
}
