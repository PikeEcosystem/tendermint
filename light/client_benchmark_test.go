package light_test

import (
	"context"
	"testing"
	"time"

	"github.com/PikeEcosystem/tendermint/libs/sync"

	dbm "github.com/tendermint/tm-db"

	"github.com/PikeEcosystem/tendermint/libs/log"
	"github.com/PikeEcosystem/tendermint/light"
	"github.com/PikeEcosystem/tendermint/light/provider"
	mockp "github.com/PikeEcosystem/tendermint/light/provider/mock"
	dbs "github.com/PikeEcosystem/tendermint/light/store/db"
	"github.com/PikeEcosystem/tendermint/types"
)

// NOTE: block is produced every minute. Make sure the verification time
// provided in the function call is correct for the size of the blockchain. The
// benchmarking may take some time hence it can be more useful to set the time
// or the amount of iterations use the flag -benchtime t -> i.e. -benchtime 5m
// or -benchtime 100x.
//
// Remember that none of these benchmarks account for network latency.
var (
	mu sync.Mutex
	// Shouldn't initialize variables here since affecting test (this is for benchmark test)
	benchmarkFullNode *mockp.Mock
	genesisBlock      *types.LightBlock
)

func setupData() {
	mu.Lock()
	defer mu.Unlock()
	if benchmarkFullNode == nil || genesisBlock == nil {
		benchmarkFullNode = mockp.New(genMockNode(chainID, 1000, 100, 1, bTime))
		genesisBlock, _ = benchmarkFullNode.LightBlock(context.Background(), 1)
	}
}

func BenchmarkSequence(b *testing.B) {
	setupData()
	c, err := light.NewClient(
		context.Background(),
		chainID,
		light.TrustOptions{
			Period: 24 * time.Hour,
			Height: 1,
			Hash:   genesisBlock.Hash(),
		},
		benchmarkFullNode,
		[]provider.Provider{benchmarkFullNode},
		dbs.New(dbm.NewMemDB(), chainID),
		light.Logger(log.TestingLogger()),
		light.SequentialVerification(),
	)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err = c.VerifyLightBlockAtHeight(context.Background(), 1000, bTime.Add(1000*time.Minute))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBisection(b *testing.B) {
	setupData()
	c, err := light.NewClient(
		context.Background(),
		chainID,
		light.TrustOptions{
			Period: 24 * time.Hour,
			Height: 1,
			Hash:   genesisBlock.Hash(),
		},
		benchmarkFullNode,
		[]provider.Provider{benchmarkFullNode},
		dbs.New(dbm.NewMemDB(), chainID),
		light.Logger(log.TestingLogger()),
	)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err = c.VerifyLightBlockAtHeight(context.Background(), 1000, bTime.Add(1000*time.Minute))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBackwards(b *testing.B) {
	setupData()
	trustedBlock, _ := benchmarkFullNode.LightBlock(context.Background(), 0)
	c, err := light.NewClient(
		context.Background(),
		chainID,
		light.TrustOptions{
			Period: 24 * time.Hour,
			Height: trustedBlock.Height,
			Hash:   trustedBlock.Hash(),
		},
		benchmarkFullNode,
		[]provider.Provider{benchmarkFullNode},
		dbs.New(dbm.NewMemDB(), chainID),
		light.Logger(log.TestingLogger()),
	)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err = c.VerifyLightBlockAtHeight(context.Background(), 1, bTime)
		if err != nil {
			b.Fatal(err)
		}
	}
}
