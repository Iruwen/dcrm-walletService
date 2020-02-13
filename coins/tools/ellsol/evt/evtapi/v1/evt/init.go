package evt

import (
	"fmt"
	"github.com/fsn-dev/dcrm-walletService/coins/tools/ellsol/evt/evtapi/client"
	"github.com/fsn-dev/dcrm-walletService/coins/tools/ellsol/evt/evtconfig"
)

type Instance struct {
	client *client.Instance
	config *evtconfig.Instance
}

func New(config *evtconfig.Instance, client *client.Instance) *Instance {
	return &Instance{
		client: client,
		config: config,
	}
}

func (it *Instance) path(method string) string {
	return fmt.Sprintf("evt/%v", method)
}
