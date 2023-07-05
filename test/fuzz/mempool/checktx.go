package checktx

import (
	"github.com/PikeEcosystem/tendermint/abci/example/kvstore"
	"github.com/PikeEcosystem/tendermint/config"
	mempl "github.com/PikeEcosystem/tendermint/mempool"
	"github.com/PikeEcosystem/tendermint/proxy"
)

var mempool mempl.Mempool

func init() {
	app := kvstore.NewApplication()
	cc := proxy.NewLocalClientCreator(app)
	appConnMem, _ := cc.NewABCIClient()
	err := appConnMem.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	mempool = mempl.NewCListMempool(cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	_, err := mempool.CheckTxSync(data, mempl.TxInfo{})
	if err != nil {
		return 0
	}

	return 1
}
