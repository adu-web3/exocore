package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/ExocoreNetwork/exocore/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

var s *AvsTaskTestSuite

type AvsTaskTestSuite struct {
	testutil.BaseTestSuite

	// needed by test
	operatorAddr          sdk.AccAddress
	avsAddr               string
	assetID               string
	stakerID              string
	assetAddr             common.Address
	assetDecimal          uint32
	clientChainLzID       uint64
	depositAmount         sdkmath.Int
	delegationAmount      sdkmath.Int
	updatedAmountForOptIn sdkmath.Int
}

func TestAvsTaskTestSuite(t *testing.T) {
	s = new(AvsTaskTestSuite)
	suite.Run(t, s)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "avs-task module Suite")
}

func (suite *AvsTaskTestSuite) SetupTest() {
	suite.DoSetupTest()
}
