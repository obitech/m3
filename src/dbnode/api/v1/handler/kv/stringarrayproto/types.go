package stringarrayproto

import (
	"github.com/m3db/m3/src/dbnode/storage/bootstrap/bootstrapper/commitlog"
	"github.com/m3db/m3/src/dbnode/storage/bootstrap/bootstrapper/fs"
	"github.com/m3db/m3/src/dbnode/storage/bootstrap/bootstrapper/peers"
	"github.com/m3db/m3/src/dbnode/storage/bootstrap/bootstrapper/uninitialized"
)

var (
	AllowedBootstrappers = []string{
		peers.PeersBootstrapperName,
		commitlog.CommitLogBootstrapperName,
		fs.FileSystemBootstrapperName,
		uninitialized.UninitializedTopologyBootstrapperName,
	}
)

type Bootstrapper struct {
	name string `json:"name"`
}

// BootstrapperList indicates the order of Bootstrappers.
type BootstrapperList []Bootstrapper
