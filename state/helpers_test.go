package state_test

import (
	"bytes"
	"fmt"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	tmabci "github.com/PikeEcosystem/tendermint/abci/types"
	"github.com/PikeEcosystem/tendermint/crypto"
	"github.com/PikeEcosystem/tendermint/crypto/ed25519"
	tmrand "github.com/PikeEcosystem/tendermint/libs/rand"
	tmstate "github.com/PikeEcosystem/tendermint/proto/tendermint/state"
	"github.com/PikeEcosystem/tendermint/proxy"
	sm "github.com/PikeEcosystem/tendermint/state"
	"github.com/PikeEcosystem/tendermint/types"
	tmtime "github.com/PikeEcosystem/tendermint/types/time"
)

type paramsChangeTestCase struct {
	height int64
	params tmproto.ConsensusParams
}

func newTestApp() proxy.AppConns {
	app := &testApp{}
	cc := proxy.NewLocalClientCreator(app)
	return proxy.NewAppConns(cc)
}

func makeAndCommitGoodBlock(
	state sm.State,
	height int64,
	lastCommit *types.Commit,
	proposerAddr []byte,
	blockExec *sm.BlockExecutor,
	privVals map[string]types.PrivValidator,
	evidence []types.Evidence) (sm.State, types.BlockID, *types.Commit, error) {
	// A good block passes
	state, blockID, err := makeAndApplyGoodBlock(state,
		privVals[state.Validators.SelectProposer(state.LastProofHash, height, 0).Address.String()],
		height, lastCommit, proposerAddr, blockExec, evidence)
	if err != nil {
		return state, types.BlockID{}, nil, err
	}

	// Simulate a lastCommit for this block from all validators for the next height
	commit, err := makeValidCommit(height, blockID, state.Validators, privVals)
	if err != nil {
		return state, types.BlockID{}, nil, err
	}
	return state, blockID, commit, nil
}

func makeAndApplyGoodBlock(state sm.State, privVal types.PrivValidator, height int64, lastCommit *types.Commit,
	proposerAddr []byte, blockExec *sm.BlockExecutor, evidence []types.Evidence) (sm.State, types.BlockID, error) {
	message := state.MakeHashMessage(0)
	proof, _ := privVal.GenerateVRFProof(message)
	block, _ := state.MakeBlock(height, makeTxs(height), lastCommit, evidence, proposerAddr, 0, proof)
	if err := blockExec.ValidateBlock(state, 0, block); err != nil {
		return state, types.BlockID{}, err
	}
	blockID := types.BlockID{Hash: block.Hash(),
		PartSetHeader: types.PartSetHeader{Total: 3, Hash: tmrand.Bytes(32)}}
	state, _, err := blockExec.ApplyBlock(state, blockID, block, nil)
	if err != nil {
		return state, types.BlockID{}, err
	}
	return state, blockID, nil
}

func makeValidCommit(
	height int64,
	blockID types.BlockID,
	vals *types.ValidatorSet,
	privVals map[string]types.PrivValidator,
) (*types.Commit, error) {
	sigs := make([]types.CommitSig, 0)
	for i := 0; i < vals.Size(); i++ {
		_, val := vals.GetByIndex(int32(i))
		vote, err := types.MakeVote(height, blockID, vals, privVals[val.Address.String()], chainID, time.Now())
		if err != nil {
			return nil, err
		}
		sigs = append(sigs, vote.CommitSig())
	}
	return types.NewCommit(height, 0, blockID, sigs), nil
}

// make some bogus txs
func makeTxs(height int64) (txs []types.Tx) {
	for i := 0; i < nTxsPerBlock; i++ {
		txs = append(txs, types.Tx([]byte{byte(height), byte(i)}))
	}
	return txs
}

func makePrivVal() types.PrivValidator {
	pk := ed25519.GenPrivKeyFromSecret([]byte("test private validator"))
	return types.NewMockPVWithParams(pk, false, false)
}

func makeState(nVals, height int) (sm.State, dbm.DB, map[string]types.PrivValidator) {
	vals := make([]types.GenesisValidator, nVals)
	privVals := make(map[string]types.PrivValidator, nVals)
	for i := 0; i < nVals; i++ {
		secret := []byte(fmt.Sprintf("test%d", i))
		pk := ed25519.GenPrivKeyFromSecret(secret)
		valAddr := pk.PubKey().Address()
		vals[i] = types.GenesisValidator{
			Address: valAddr,
			PubKey:  pk.PubKey(),
			Power:   1000,
			Name:    fmt.Sprintf("test%d", i),
		}
		privVals[valAddr.String()] = types.NewMockPVWithParams(pk, false, false)
	}
	s, _ := sm.MakeGenesisState(&types.GenesisDoc{
		ChainID:    chainID,
		Validators: vals,
		AppHash:    nil,
	})

	stateDB := dbm.NewMemDB()
	stateStore := sm.NewStore(stateDB)
	if err := stateStore.Save(s); err != nil {
		panic(err)
	}

	for i := 1; i < height; i++ {
		s.LastBlockHeight++
		s.LastValidators = s.Validators.Copy()
		if err := stateStore.Save(s); err != nil {
			panic(err)
		}
	}

	return s, stateDB, privVals
}

func makeBlock(state sm.State, height int64) *types.Block {
	return makeBlockWithPrivVal(state, makePrivVal(), height)
}

func makeBlockWithPrivVal(state sm.State, privVal types.PrivValidator, height int64) *types.Block {
	message := state.MakeHashMessage(0)
	proof, _ := privVal.GenerateVRFProof(message)
	pubKey, _ := privVal.GetPubKey()
	block, _ := state.MakeBlock(
		height,
		makeTxs(state.LastBlockHeight),
		new(types.Commit),
		nil,
		pubKey.Address(),
		0,
		proof,
	)
	return block
}

func genValSet(size int) *types.ValidatorSet {
	vals := make([]*types.Validator, size)
	for i := 0; i < size; i++ {
		vals[i] = types.NewValidator(ed25519.GenPrivKey().PubKey(), 10)
	}
	return types.NewValidatorSet(vals)
}

func makeHeaderPartsResponsesValPubKeyChange(
	state sm.State,
	pubkey crypto.PubKey,
) (types.Header, types.Entropy, types.BlockID, *tmstate.ABCIResponses) {

	block := makeBlock(state, state.LastBlockHeight+1)
	abciResponses := &tmstate.ABCIResponses{
		BeginBlock: &abci.ResponseBeginBlock{},
		EndBlock:   &abci.ResponseEndBlock{ValidatorUpdates: nil},
	}
	// If the pubkey is new, remove the old and add the new.
	_, val := state.NextValidators.GetByIndex(0)
	if !bytes.Equal(pubkey.Bytes(), val.PubKey.Bytes()) {
		abciResponses.EndBlock = &abci.ResponseEndBlock{
			ValidatorUpdates: []abci.ValidatorUpdate{
				types.OC2PB.NewValidatorUpdate(val.PubKey, 0),
				types.OC2PB.NewValidatorUpdate(pubkey, 10),
			},
		}
	}

	return block.Header, block.Entropy, types.BlockID{Hash: block.Hash(), PartSetHeader: types.PartSetHeader{}}, abciResponses
}

func makeHeaderPartsResponsesValPowerChange(
	state sm.State,
	power int64,
) (types.Header, types.Entropy, types.BlockID, *tmstate.ABCIResponses) {

	block := makeBlock(state, state.LastBlockHeight+1)
	abciResponses := &tmstate.ABCIResponses{
		BeginBlock: &abci.ResponseBeginBlock{},
		EndBlock:   &abci.ResponseEndBlock{ValidatorUpdates: nil},
	}

	// If the pubkey is new, remove the old and add the new.
	_, val := state.NextValidators.GetByIndex(0)
	if val.VotingPower != power {
		abciResponses.EndBlock = &abci.ResponseEndBlock{
			ValidatorUpdates: []abci.ValidatorUpdate{
				types.OC2PB.NewValidatorUpdate(val.PubKey, power),
			},
		}
	}

	return block.Header, block.Entropy, types.BlockID{Hash: block.Hash(), PartSetHeader: types.PartSetHeader{}}, abciResponses
}

func makeHeaderPartsResponsesParams(
	state sm.State,
	params tmproto.ConsensusParams,
) (types.Header, types.Entropy, types.BlockID, *tmstate.ABCIResponses) {

	block := makeBlock(state, state.LastBlockHeight+1)
	abciResponses := &tmstate.ABCIResponses{
		BeginBlock: &abci.ResponseBeginBlock{},
		EndBlock:   &abci.ResponseEndBlock{ConsensusParamUpdates: types.OC2PB.ConsensusParams(&params)},
	}
	return block.Header, block.Entropy, types.BlockID{Hash: block.Hash(), PartSetHeader: types.PartSetHeader{}}, abciResponses
}

func randomGenesisDoc() *types.GenesisDoc {
	pubkey := ed25519.GenPrivKey().PubKey()
	return &types.GenesisDoc{
		GenesisTime: tmtime.Now(),
		ChainID:     "abc",
		Validators: []types.GenesisValidator{
			{
				Address: pubkey.Address(),
				PubKey:  pubkey,
				Power:   10,
				Name:    "myval",
			},
		},
		ConsensusParams: types.DefaultConsensusParams(),
	}
}

//----------------------------------------------------------------------------

type testApp struct {
	tmabci.BaseApplication

	CommitVotes         []abci.VoteInfo
	ByzantineValidators []abci.Evidence
	ValidatorUpdates    []abci.ValidatorUpdate
}

var _ tmabci.Application = (*testApp)(nil)
var TestAppVersion uint64 = 66

func (app *testApp) Info(req abci.RequestInfo) (resInfo abci.ResponseInfo) {
	return abci.ResponseInfo{}
}

func (app *testApp) BeginBlock(req tmabci.RequestBeginBlock) abci.ResponseBeginBlock {
	app.CommitVotes = req.LastCommitInfo.Votes
	app.ByzantineValidators = req.ByzantineValidators
	return abci.ResponseBeginBlock{}
}

func (app *testApp) EndBlock(req abci.RequestEndBlock) abci.ResponseEndBlock {
	return abci.ResponseEndBlock{
		ValidatorUpdates: app.ValidatorUpdates,
		ConsensusParamUpdates: &abci.ConsensusParams{
			Version: &tmproto.VersionParams{
				AppVersion: TestAppVersion}}}
}

func (app *testApp) DeliverTx(req abci.RequestDeliverTx) abci.ResponseDeliverTx {
	return abci.ResponseDeliverTx{Events: []abci.Event{}}
}

func (app *testApp) CheckTxSync(req abci.RequestCheckTx) tmabci.ResponseCheckTx {
	return tmabci.ResponseCheckTx{}
}

func (app *testApp) CheckTxAsync(req abci.RequestCheckTx, callback tmabci.CheckTxCallback) {
	callback(tmabci.ResponseCheckTx{})
}

func (app *testApp) Commit() abci.ResponseCommit {
	return abci.ResponseCommit{RetainHeight: 1}
}

func (app *testApp) Query(reqQuery abci.RequestQuery) (resQuery abci.ResponseQuery) {
	return
}