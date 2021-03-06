package keeper

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/bluzelle/curium/x/curium"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gogo/protobuf/sortkeys"
	"github.com/tendermint/tendermint/privval"
	"strings"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/bluzelle/curium/x/voting/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

var voteQueue = map[int64][]types.MsgVote{}
var proofQueue = make([]types.MsgVoteProof, 0)
var voteHandlers = map[string]types.VoteHandler{}

type (
	Keeper struct {
		cdc           codec.Marshaler
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		homeDir       string
		stakingKeeper stakingkeeper.Keeper
		accKeeper     authkeeper.AccountKeeper
		// this line is used by starport scaffolding # ibc/keeper/attribute
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	homeDir string,
	stakingKeeper stakingkeeper.Keeper,
	accKeeper authkeeper.AccountKeeper,
// this line is used by starport scaffolding # ibc/keeper/parameter
) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		homeDir:       homeDir,
		stakingKeeper: stakingKeeper,
		accKeeper:     accKeeper,
		// this line is used by starport scaffolding # ibc/keeper/return
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) RegisterVoteHandler(handler types.VoteHandler) {
	voteHandlers[handler.VoteType()] = handler
}

// Vote this is the function that other modules will call to send a vote
func (k Keeper) Vote(ctx sdk.Context, votingReq types.VotingRequest) {
	valcons := k.GetValconsAddress()
	proof := types.MsgVoteProof{
		Creator:   votingReq.Creator,
		Valcons:   valcons,
		Signature: k.SignProofSig(votingReq.Value),
		VoteType:  votingReq.VoteType,
		Id:        votingReq.Id,
		From:      votingReq.From,
		Batch:     GenerateBatchNow(),
	}

	proofQueue = append(proofQueue, proof)

	vote := types.MsgVote{
		Creator:  proof.Creator,
		Valcons:  proof.Valcons,
		VoteType: proof.VoteType,
		Id:       proof.Id,
		From:     proof.From,
		Batch:    proof.Batch,
		Value:    votingReq.Value,
		Block:    ctx.BlockHeight() + 3,
	}

	voteQueue[ctx.BlockHeight()+3] = append(voteQueue[ctx.BlockHeight() + 3], vote)
}

func MakeProofStoreKey(valcons string, voteType string, voteId uint64) []byte{
	return append([]byte(valcons + voteType), uint64ToByteArray(voteId)...)
}

func uint64ToByteArray(n uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(n))
	return b
}

func MakeVoteStoreKey(block int64, voteType string, voteId uint64, valcons string) []byte {
	bytes := uint64ToByteArray(uint64(block))
	bytes = append(bytes, []byte(voteType)...)
	bytes = append(bytes, uint64ToByteArray(voteId)...)
	bytes = append(bytes, []byte(valcons)...)
	return bytes
}

func (k Keeper) GetProofStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofPrefix))
}

func (k Keeper) GetVoteStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VotePrefix))
}


func (k Keeper) TransmitProofQueue(ctx sdk.Context) {
	if len(proofQueue) > 0 {
		var msgs []sdk.Msg
		for i := 0; i < len(proofQueue); i++{
			msgs = append(msgs, &proofQueue[i])
		}
		result, err := curium.BroadcastMessages(ctx, msgs, &k.accKeeper, proofQueue[0].From, k.homeDir)
		if err != nil {
			k.Logger(ctx).Error("Error broadcasting proofs", "err", err)
			return
		}
		k.Logger(ctx).Info("Broadcast proofs successful", "result", result)
		proofQueue = make([]types.MsgVoteProof, 0)
	}
}

func (k Keeper) TransmitVoteQueue(ctx sdk.Context) {
	if voteQueue[ctx.BlockHeight()] != nil {
		var msgs []sdk.Msg
		for i := 0; i < len(voteQueue[ctx.BlockHeight()]); i++ {
			msgs = append(msgs, &voteQueue[ctx.BlockHeight()][i])
		}
		curium.BroadcastMessages(ctx, msgs, &k.accKeeper, voteQueue[ctx.BlockHeight()][0].From, k.homeDir)
		delete(voteQueue, ctx.BlockHeight())
	}
}

func GenerateBatchNow() string {
	return GenerateBatch(time.Now())
}

func GenerateBatch(t time.Time) string {
	out := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		0,
		0,
		time.UTC,
	).String()
	out = strings.Replace(out, ":00 +0000 UTC", "", 1)
	out = strings.Replace(out, " ", "-", -1)
	return out

}

func (k Keeper) StoreVote(ctx sdk.Context, vote *types.Vote) {
	store := k.GetVoteStore(ctx)
	key := MakeVoteStoreKey(ctx.BlockHeight(), vote.VoteType, vote.Id, vote.Valcons)
	store.Set(key, k.cdc.MustMarshalBinaryBare(vote))
}



func (k Keeper) CheckDeliverVotes(ctx sdk.Context) {
	start := make([]byte, 8)
	binary.LittleEndian.PutUint64(start, uint64(ctx.BlockHeight() - 2))
	store := k.GetVoteStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, start)
	var votes = map[string]map[uint64][]*types.Vote{}
	for ; iterator.Valid(); iterator.Next() {
		var vote types.Vote
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &vote)
		if votes[vote.VoteType] == nil {
			votes[vote.VoteType] = map[uint64][]*types.Vote{}
		}
		votes[vote.VoteType][vote.Id] = append(votes[vote.VoteType][vote.Id], &vote)
		store.Delete(iterator.Key())
	}

	for voteType := range votes  {
		if voteHandlers[voteType] == nil {
			k.Logger(ctx).Error("No vote handler registered", "type", voteType)
			continue
		}

		var ids []uint64
		for id := range votes[voteType] {
			ids = append(ids, id)
		}
		sortkeys.Uint64s(ids)
		for _, id := range ids {
			voteHandlers[voteType].VotesReceived(&ctx, id, votes[voteType][id])
		}
	}
}

func (k Keeper) GetPrivateValidator() *privval.FilePV {
	return privval.LoadFilePV(k.homeDir+"/config/priv_validator_key.json", k.homeDir+"/data/priv_validator_state.json")
}

func (k Keeper) GetValconsAddress() string {
	validator := k.GetPrivateValidator()
	address := validator.GetAddress()
	consAddress := (sdk.ConsAddress)(address)
	addressString := consAddress.String()
	return addressString
}

func (k Keeper) SignProofSig(value []byte) string {
	v := k.GetPrivateValidator()
	s, _ := v.Key.PrivKey.Sign([]byte(value))
	return hex.EncodeToString(s)
}

func (k Keeper) GetValidator(ctx sdk.Context, valcons string) (validator stakingtypes.Validator, found bool) {
	consAddr, _ := sdk.ConsAddressFromBech32(valcons)
	return k.stakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
}

func (k Keeper) GetValidatorWeight(ctx sdk.Context, valcons string) int64 {
	validator, validatorFound := k.GetValidator(ctx, valcons)
	if !validatorFound {
		return 0
	}
	return validator.ConsensusPower()
}

func (k Keeper) IsVoteValid(ctx sdk.Context, msg *types.MsgVote) bool {
	validator, validatorFound := k.GetValidator(ctx, msg.Valcons)

	if validator.Jailed {
		k.Logger(ctx).Info("Vote received from jailed validator", "name", msg.Id, "valcons", msg.Valcons)
		return false
	}

	if !validatorFound {
		k.Logger(ctx).Info("Vote received from unknown validator", "name", msg.Id, "valcons", msg.Valcons)
		return false
	}

	proofSignatureString := k.GetVoteProof(ctx, msg).Signature
	proofSignature, _ := hex.DecodeString(proofSignatureString)

	pubKey, _ := validator.ConsPubKey()

	isGood := pubKey.VerifySignature(msg.Value, proofSignature)

	if !isGood {
		k.Logger(ctx).Info("vote/proof mismatch", "name", msg.Id, "valcons", msg.Valcons)
	}

	return isGood
}

func (k Keeper) GetVoteProof(ctx sdk.Context, vote *types.MsgVote) types.MsgVoteProof {
	var msg types.MsgVoteProof
	proofStore := k.GetProofStore(ctx)

	var keys []string
	iterator := proofStore.Iterator(nil, nil)
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, string(iterator.Key()))
	}

	proof := proofStore.Get(MakeProofStoreKey(vote.Valcons, vote.VoteType, vote.Id))
	k.cdc.MustUnmarshalBinaryBare(proof, &msg)
	return msg
}
