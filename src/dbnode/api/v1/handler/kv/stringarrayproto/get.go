package stringarrayproto

import (
	"net/http"
	"path"

	clusterkv "github.com/m3db/m3/src/cluster/kv"
	"github.com/m3db/m3/src/dbnode/api/v1/handler"
	"github.com/m3db/m3/src/dbnode/api/v1/handler/kv"
	"github.com/m3db/m3/src/dbnode/kvconfig"
	"github.com/m3db/m3/src/x/instrument"
)

var (
	// BootstrapperGetURL is the URL for retrieving bootstrappers from the kv store.
	BootstrapperURL = path.Join(handler.RoutePrefixV1, kv.KVPathName, kvconfig.BootstrapperKey)

	BootstrapperGetHTTPMethod = http.MethodGet
)

// BootstrapperGetHandler is the handler for retrieving stringarrayprotos from the  kv
// store.
type BootstrapperGetHandler kv.Handler

// NewGetHandler returns a new instance of BootstrapperGetHandler.
func NewBootstrapperGetHandler(
	instrumentOpts instrument.Options,
	store clusterkv.Store) *BootstrapperGetHandler {
	return &BootstrapperGetHandler{
		Store:          store,
		InstrumentOpts: instrumentOpts,
	}
}

func (h *BootstrapperGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Accept payload

	// Parse payload

	// Validate bootstrappers
}
