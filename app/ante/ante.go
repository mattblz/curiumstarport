package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(
	ak auth.AccountKeeper, bankKeeper types.BankKeeper,
	sigGasConsumer auth.SignatureVerificationGasConsumer,
	signModeHandler signing.SignModeHandler,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		auth.NewRejectExtensionOptionsDecorator(),
		auth.NewMempoolFeeDecorator(),
		auth.NewValidateBasicDecorator(),
		auth.TxTimeoutHeightDecorator{},
		auth.NewValidateMemoDecorator(ak),
		auth.NewConsumeGasForTxSizeDecorator(ak),
		auth.NewRejectFeeGranterDecorator(),
		auth.NewSetPubKeyDecorator(ak), // SetPubKeyDecorator must be called before all signature verification decorators
		auth.NewValidateSigCountDecorator(ak),
		NewDeductFeeDecorator(ak, bankKeeper),
		auth.NewSigGasConsumeDecorator(ak, sigGasConsumer),
		auth.NewSigVerificationDecorator(ak, signModeHandler),
		auth.NewIncrementSequenceDecorator(ak),
	)
}
