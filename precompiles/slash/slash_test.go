package slash_test

import (
	"math/big"

	sdkmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v14/x/evm/statedb"
	evmtypes "github.com/evmos/evmos/v14/x/evm/types"
	"github.com/exocore/app"
	"github.com/exocore/precompiles/slash"
	"github.com/exocore/x/deposit/keeper"
	depositParams "github.com/exocore/x/deposit/types"
	"github.com/exocore/x/restaking_assets_manage/types"
	slashParams "github.com/exocore/x/slash/types"
)

func (s *PrecompileTestSuite) TestIsTransaction() {
	testCases := []struct {
		name   string
		method string
		isTx   bool
	}{
		{
			slash.MethodSlash,
			s.precompile.Methods[slash.MethodSlash].Name,
			true,
		},
		{
			"invalid",
			"invalid",
			false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Require().Equal(s.precompile.IsTransaction(tc.method), tc.isTx)
		})
	}
}

func paddingClientChainAddress(input []byte, outputLength int) []byte {
	if len(input) < outputLength {
		padding := make([]byte, outputLength-len(input))
		return append(input, padding...)
	}
	return input
}

// TestRun tests the precompiles Run method submitSlash.
func (s *PrecompileTestSuite) TestRunSlash() {
	// deposit params for test
	exoCoreLzAppEventTopic := "0xc6a377bfc4eb120024a8ac08eef205be16b817020812c73223e81d1bdb9708ec"
	usdtAddress := common.FromHex("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	clientChainLzId := 101
	slashAmount := big.NewInt(10)
	depositAmount := big.NewInt(100)
	assetAddr := paddingClientChainAddress(usdtAddress, types.GeneralClientChainAddrLength)
	depositAsset := func(staker []byte, depositAmount sdkmath.Int) {
		// deposit asset for slash test
		params := &keeper.DepositParams{
			ClientChainLzId: 101,
			Action:          types.Deposit,
			StakerAddress:   staker,
			AssetsAddress:   usdtAddress,
			OpAmount:        depositAmount,
		}
		err := s.app.DepositKeeper.Deposit(s.ctx, params)
		s.Require().NoError(err)
	}

	commonMalleate := func() (common.Address, []byte) {
		// Prepare the call input for slash test
		input, err := s.precompile.Pack(
			slash.MethodSlash,
			uint16(clientChainLzId),
			assetAddr,
			paddingClientChainAddress(s.address.Bytes(), types.GeneralClientChainAddrLength),
			slashAmount,
			common.FromHex("0x2E756b8faBeA234b9900767b69D6387400CDC396"),
			common.FromHex("0xceb69f6342ece283b2f5c9088ff249b5d0ae66ea"),
			"5",
			"slash",
		)
		s.Require().NoError(err, "failed to pack input")
		return s.address, input
	}
	successRet, err := s.precompile.Methods[slash.MethodSlash].Outputs.Pack(true)
	s.Require().NoError(err)
	testcases := []struct {
		name        string
		malleate    func() (common.Address, []byte)
		readOnly    bool
		expPass     bool
		errContains string
		returnBytes []byte
	}{
		{
			name: "pass - slash via pre-compiles",
			malleate: func() (common.Address, []byte) {
				depositModuleParam := &depositParams.Params{
					ExoCoreLzAppAddress:    s.address.String(),
					ExoCoreLzAppEventTopic: exoCoreLzAppEventTopic,
				}
				err := s.app.DepositKeeper.SetParams(s.ctx, depositModuleParam)
				s.Require().NoError(err)
				depositAsset(s.address.Bytes(), sdkmath.NewIntFromBigInt(depositAmount))
				slashModuleParam := &slashParams.Params{
					ExoCoreLzAppAddress:    s.address.String(),
					ExoCoreLzAppEventTopic: exoCoreLzAppEventTopic,
				}
				err = s.app.ExoSlashKeeper.SetParams(s.ctx, slashModuleParam)
				s.Require().NoError(err)
				return commonMalleate()
			},
			returnBytes: successRet,
			readOnly:    false,
			expPass:     true,
		},
	}
	for _, tc := range testcases {
		tc := tc
		s.Run(tc.name, func() {
			// setup basic test suite
			s.SetupTest()

			baseFee := s.app.FeeMarketKeeper.GetBaseFee(s.ctx)

			// malleate testcase
			caller, input := tc.malleate()

			contract := vm.NewPrecompile(vm.AccountRef(caller), s.precompile, big.NewInt(0), uint64(1e6))
			contract.Input = input

			contractAddr := contract.Address()
			// Build and sign Ethereum transaction
			txArgs := evmtypes.EvmTxArgs{
				ChainID:   s.app.EvmKeeper.ChainID(),
				Nonce:     0,
				To:        &contractAddr,
				Amount:    nil,
				GasLimit:  100000,
				GasPrice:  app.MainnetMinGasPrices.BigInt(),
				GasFeeCap: baseFee,
				GasTipCap: big.NewInt(1),
				Accesses:  &ethtypes.AccessList{},
			}
			msgEthereumTx := evmtypes.NewTx(&txArgs)

			msgEthereumTx.From = s.address.String()
			err := msgEthereumTx.Sign(s.ethSigner, s.signer)
			s.Require().NoError(err, "failed to sign Ethereum message")

			// Instantiate config
			proposerAddress := s.ctx.BlockHeader().ProposerAddress
			cfg, err := s.app.EvmKeeper.EVMConfig(s.ctx, proposerAddress, s.app.EvmKeeper.ChainID())
			s.Require().NoError(err, "failed to instantiate EVM config")

			msg, err := msgEthereumTx.AsMessage(s.ethSigner, baseFee)
			s.Require().NoError(err, "failed to instantiate Ethereum message")

			// Create StateDB
			s.stateDB = statedb.New(s.ctx, s.app.EvmKeeper, statedb.NewEmptyTxConfig(common.BytesToHash(s.ctx.HeaderHash().Bytes())))
			// Instantiate EVM
			evm := s.app.EvmKeeper.NewEVM(
				s.ctx, msg, cfg, nil, s.stateDB,
			)
			params := s.app.EvmKeeper.GetParams(s.ctx)
			activePrecompiles := params.GetActivePrecompilesAddrs()
			precompileMap := s.app.EvmKeeper.Precompiles(activePrecompiles...)
			err = vm.ValidatePrecompiles(precompileMap, activePrecompiles)
			s.Require().NoError(err, "invalid precompiles", activePrecompiles)
			evm.WithPrecompiles(precompileMap, activePrecompiles)

			// Run precompiled contract
			bz, err := s.precompile.Run(evm, contract, tc.readOnly)

			// Check results
			if tc.expPass {
				s.Require().NoError(err, "expected no error when running the precompile")
				s.Require().Equal(tc.returnBytes, bz, "the return doesn't match the expected result")
			} else {
				s.Require().Error(err, "expected error to be returned when running the precompile")
				s.Require().Nil(bz, "expected returned bytes to be nil")
				s.Require().ErrorContains(err, tc.errContains)
			}
		})
	}
}
