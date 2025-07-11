package types_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/green901612/cosevm/x/evm/types"
)

func (suite *TxDataTestSuite) TestNewLegacyTx() {
	testCases := []struct {
		name string
		tx   *ethtypes.Transaction
	}{
		{
			"non-empty Transaction",
			ethtypes.NewTx(&ethtypes.AccessListTx{
				Nonce:      1,
				Data:       []byte("data"),
				Gas:        100,
				Value:      big.NewInt(1),
				AccessList: ethtypes.AccessList{},
				To:         &suite.addr,
				V:          big.NewInt(1),
				R:          big.NewInt(1),
				S:          big.NewInt(1),
			}),
		},
	}

	for _, tc := range testCases {
		tx, err := types.NewLegacyTx(tc.tx)
		suite.Require().NoError(err)

		suite.Require().NotEmpty(tc.tx)
		suite.Require().Equal(uint8(0), tx.TxType())
	}
}

func (suite *TxDataTestSuite) TestLegacyTxTxType() {
	tx := types.LegacyTx{}
	actual := tx.TxType()

	suite.Require().Equal(uint8(0), actual)
}

func (suite *TxDataTestSuite) TestLegacyTxCopy() {
	tx := &types.LegacyTx{}
	txData := tx.Copy()

	suite.Require().Equal(&types.LegacyTx{}, txData)
	// TODO: Test for different pointers
}

func (suite *TxDataTestSuite) TestLegacyTxGetChainID() {
	tx := types.LegacyTx{}
	actual := tx.GetChainID()

	suite.Require().Nil(actual)
}

func (suite *TxDataTestSuite) TestLegacyTxGetAccessList() {
	tx := types.LegacyTx{}
	actual := tx.GetAccessList()

	suite.Require().Nil(actual)
}

func (suite *TxDataTestSuite) TestLegacyTxGetData() {
	testCases := []struct {
		name string
		tx   types.LegacyTx
	}{
		{
			"non-empty transaction",
			types.LegacyTx{
				Data: nil,
			},
		},
	}

	for _, tc := range testCases {
		actual := tc.tx.GetData()

		suite.Require().Equal(tc.tx.Data, actual, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxGetGas() {
	testCases := []struct {
		name string
		tx   types.LegacyTx
		exp  uint64
	}{
		{
			"non-empty gas",
			types.LegacyTx{
				GasLimit: suite.uint64,
			},
			suite.uint64,
		},
	}

	for _, tc := range testCases {
		actual := tc.tx.GetGas()

		suite.Require().Equal(tc.exp, actual, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxGetGasPrice() {
	testCases := []struct {
		name string
		tx   types.LegacyTx
		exp  *big.Int
	}{
		{
			"empty gasPrice",
			types.LegacyTx{
				GasPrice: nil,
			},
			nil,
		},
		{
			"non-empty gasPrice",
			types.LegacyTx{
				GasPrice: &suite.sdkInt,
			},
			(&suite.sdkInt).BigInt(),
		},
	}

	for _, tc := range testCases {
		actual := tc.tx.GetGasFeeCap()

		suite.Require().Equal(tc.exp, actual, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxGetGasTipCap() {
	testCases := []struct {
		name string
		tx   types.LegacyTx
		exp  *big.Int
	}{
		{
			"non-empty gasPrice",
			types.LegacyTx{
				GasPrice: &suite.sdkInt,
			},
			(&suite.sdkInt).BigInt(),
		},
	}

	for _, tc := range testCases {
		actual := tc.tx.GetGasTipCap()

		suite.Require().Equal(tc.exp, actual, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxGetGasFeeCap() {
	testCases := []struct {
		name string
		tx   types.LegacyTx
		exp  *big.Int
	}{
		{
			"non-empty gasPrice",
			types.LegacyTx{
				GasPrice: &suite.sdkInt,
			},
			(&suite.sdkInt).BigInt(),
		},
	}

	for _, tc := range testCases {
		actual := tc.tx.GetGasFeeCap()

		suite.Require().Equal(tc.exp, actual, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxGetValue() {
	testCases := []struct {
		name string
		tx   types.LegacyTx
		exp  *big.Int
	}{
		{
			"empty amount",
			types.LegacyTx{
				Amount: nil,
			},
			nil,
		},
		{
			"non-empty amount",
			types.LegacyTx{
				Amount: &suite.sdkInt,
			},
			(&suite.sdkInt).BigInt(),
		},
	}

	for _, tc := range testCases {
		actual := tc.tx.GetValue()

		suite.Require().Equal(tc.exp, actual, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxGetNonce() {
	testCases := []struct {
		name string
		tx   types.LegacyTx
		exp  uint64
	}{
		{
			"none-empty nonce",
			types.LegacyTx{
				Nonce: suite.uint64,
			},
			suite.uint64,
		},
	}
	for _, tc := range testCases {
		actual := tc.tx.GetNonce()

		suite.Require().Equal(tc.exp, actual)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxGetTo() {
	testCases := []struct {
		name string
		tx   types.LegacyTx
		exp  *common.Address
	}{
		{
			"empty address",
			types.LegacyTx{
				To: "",
			},
			nil,
		},
		{
			"non-empty address",
			types.LegacyTx{
				To: suite.hexAddr,
			},
			&suite.addr,
		},
	}

	for _, tc := range testCases {
		actual := tc.tx.GetTo()

		suite.Require().Equal(tc.exp, actual, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxAsEthereumData() {
	tx := &types.LegacyTx{}
	txData := tx.AsEthereumData()

	suite.Require().Equal(&ethtypes.LegacyTx{}, txData)
}

func (suite *TxDataTestSuite) TestLegacyTxSetSignatureValues() {
	testCases := []struct {
		name string
		v    *big.Int
		r    *big.Int
		s    *big.Int
	}{
		{
			"non-empty values",
			suite.bigInt,
			suite.bigInt,
			suite.bigInt,
		},
	}
	for _, tc := range testCases {
		tx := &types.LegacyTx{}
		tx.SetSignatureValues(nil, tc.v, tc.r, tc.s)

		v, r, s := tx.GetRawSignatureValues()

		suite.Require().Equal(tc.v, v, tc.name)
		suite.Require().Equal(tc.r, r, tc.name)
		suite.Require().Equal(tc.s, s, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxValidate() {
	testCases := []struct {
		name     string
		tx       types.LegacyTx
		expError bool
	}{
		{
			"empty",
			types.LegacyTx{},
			true,
		},
		{
			"gas price is nil",
			types.LegacyTx{
				GasPrice: nil,
			},
			true,
		},
		{
			"gas price is negative",
			types.LegacyTx{
				GasPrice: &suite.sdkMinusOneInt,
			},
			true,
		},
		{
			"amount is negative",
			types.LegacyTx{
				GasPrice: &suite.sdkInt,
				Amount:   &suite.sdkMinusOneInt,
			},
			true,
		},
		{
			"to address is invalid",
			types.LegacyTx{
				GasPrice: &suite.sdkInt,
				Amount:   &suite.sdkInt,
				To:       suite.invalidAddr,
			},
			true,
		},
	}

	for _, tc := range testCases {
		err := tc.tx.Validate()

		if tc.expError {
			suite.Require().Error(err, tc.name)
			continue
		}

		suite.Require().NoError(err, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxEffectiveGasPrice() {
	testCases := []struct {
		name    string
		tx      types.LegacyTx
		baseFee *big.Int
		exp     *big.Int
	}{
		{
			"non-empty legacy tx",
			types.LegacyTx{
				GasPrice: &suite.sdkInt,
			},
			(&suite.sdkInt).BigInt(),
			(&suite.sdkInt).BigInt(),
		},
	}

	for _, tc := range testCases {
		actual := tc.tx.EffectiveGasPrice(tc.baseFee)

		suite.Require().Equal(tc.exp, actual, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxEffectiveFee() {
	testCases := []struct {
		name    string
		tx      types.LegacyTx
		baseFee *big.Int
		exp     *big.Int
	}{
		{
			"non-empty legacy tx",
			types.LegacyTx{
				GasPrice: &suite.sdkInt,
				GasLimit: uint64(1),
			},
			(&suite.sdkInt).BigInt(),
			(&suite.sdkInt).BigInt(),
		},
	}

	for _, tc := range testCases {
		actual := tc.tx.EffectiveFee(tc.baseFee)

		suite.Require().Equal(tc.exp, actual, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxEffectiveCost() {
	testCases := []struct {
		name    string
		tx      types.LegacyTx
		baseFee *big.Int
		exp     *big.Int
	}{
		{
			"non-empty legacy tx",
			types.LegacyTx{
				GasPrice: &suite.sdkInt,
				GasLimit: uint64(1),
				Amount:   &suite.sdkZeroInt,
			},
			(&suite.sdkInt).BigInt(),
			(&suite.sdkInt).BigInt(),
		},
	}

	for _, tc := range testCases {
		actual := tc.tx.EffectiveCost(tc.baseFee)

		suite.Require().Equal(tc.exp, actual, tc.name)
	}
}

func (suite *TxDataTestSuite) TestLegacyTxFeeCost() {
	tx := &types.LegacyTx{}

	suite.Require().Panics(func() { tx.Fee() }, "should panic")
	suite.Require().Panics(func() { tx.Cost() }, "should panic")
}
