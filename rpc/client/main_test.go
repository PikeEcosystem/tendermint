package client_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/PikeEcosystem/tendermint/abci/example/kvstore"
	nm "github.com/PikeEcosystem/tendermint/node"
	rpctest "github.com/PikeEcosystem/tendermint/rpc/test"
)

var node *nm.Node

func TestMain(m *testing.M) {
	// start an tendermint node (and kvstore) in the background to test against
	dir, err := ioutil.TempDir("/tmp", "rpc-client-test")
	if err != nil {
		panic(err)
	}

	app := kvstore.NewPersistentKVStoreApplication(dir)
	node = rpctest.Starttendermint(app)

	code := m.Run()

	// and shut down proper at the end
	rpctest.Stoptendermint(node)
	_ = os.RemoveAll(dir)
	os.Exit(code)
}
