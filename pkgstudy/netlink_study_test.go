package pkgstudy

import (
	"log"
	"os"
	"testing"
)

func Test(t *testing.T) {
	if os.Getenv(EnvName) == "1" {
		IsHostA = true
	}
	ns := SetupNetNamespace()
	bridge := SetupBridge()
	SetupVEthPeer(bridge, ns)
	SetupNsDefaultRoute()
	SetupIPTables()
	SetupRouteNs2Ns()
	log.Println("Config Finished.")
}
