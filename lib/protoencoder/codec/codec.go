package codec

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodecsdk "github.com/cosmos/cosmos-sdk/crypto/codec"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	authz "github.com/cosmos/cosmos-sdk/x/authz"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govbeta "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgrade "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	icafeestypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/core/types"
	lightclientssolo "github.com/cosmos/ibc-go/v7/modules/light-clients/06-solomachine"
	lightclientstendermint "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	lightclientslocalhost "github.com/cosmos/ibc-go/v7/modules/light-clients/09-localhost"

	cryptocodec "github.com/evmos/evmos/v18/crypto/codec"
	evmostypes "github.com/evmos/evmos/v18/types"
	erc20 "github.com/evmos/evmos/v18/x/erc20/types"
	evm "github.com/evmos/evmos/v18/x/evm/types"
	feemarket "github.com/evmos/evmos/v18/x/feemarket/types"
	inflation "github.com/evmos/evmos/v18/x/inflation/v1/types"
	revenue "github.com/evmos/evmos/v18/x/revenue/v1/types"
	vesting "github.com/evmos/evmos/v18/x/vesting/types"

	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authTx "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

var (
	ClientCtx client.Context
	Encoder   *codec.ProtoCodec
)

func init() {
	registry := codectypes.NewInterfaceRegistry()
	evmostypes.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry)
	cryptocodecsdk.RegisterInterfaces(registry)
	feegrant.RegisterInterfaces(registry)
	evm.RegisterInterfaces(registry)
	bank.RegisterInterfaces(registry)
	auth.RegisterInterfaces(registry)
	gov.RegisterInterfaces(registry)
	govbeta.RegisterInterfaces(registry)
	authz.RegisterInterfaces(registry)
	slashing.RegisterInterfaces(registry)
	staking.RegisterInterfaces(registry)
	upgrade.RegisterInterfaces(registry)
	evidence.RegisterInterfaces(registry)
	consensus.RegisterInterfaces(registry)
	feemarket.RegisterInterfaces(registry)
	inflation.RegisterInterfaces(registry)
	erc20.RegisterInterfaces(registry)
	vesting.RegisterInterfaces(registry)
	distr.RegisterInterfaces(registry)
	ibctypes.RegisterInterfaces(registry)
	ibctransfer.RegisterInterfaces(registry)
	lightclientssolo.RegisterInterfaces(registry)
	lightclientstendermint.RegisterInterfaces(registry)
	lightclientslocalhost.RegisterInterfaces(registry)
	icatypes.RegisterInterfaces(registry)
	icafeestypes.RegisterInterfaces(registry)
	revenue.RegisterInterfaces(registry)
	Encoder = codec.NewProtoCodec(registry)

	ClientCtx = client.Context{}.
		WithCodec(Encoder).
		WithTxConfig(authTx.NewTxConfig(Encoder, []signing.SignMode{signing.SignMode_SIGN_MODE_DIRECT}))
}
