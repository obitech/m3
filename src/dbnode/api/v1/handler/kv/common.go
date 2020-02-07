package kv

import (
	clusterkv "github.com/m3db/m3/src/cluster/kv"
	"github.com/m3db/m3/src/x/instrument"
)

const (
	KVPathName = "kv"
)

type Handler struct {
	Store          clusterkv.Store
	InstrumentOpts instrument.Options
}
