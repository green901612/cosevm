// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
//
// This config provides a convenient way to modify x/evm params and values.
// Its primary purpose is to be used during application initialization.

package types

import (
	"errors"
	"math/big"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	geth "github.com/ethereum/go-ethereum/params"
	"github.com/green901612/cosevm/utils"
)

// chainConfig is the chain configuration used in the EVM to defined which
// opcodes are active based on Ethereum upgrades.
var chainConfig *ChainConfig

// EthereumConfig returns an Ethereum ChainConfig for EVM state transitions.
// All the negative or nil values are converted to nil
func (cc ChainConfig) EthereumConfig(chainID *big.Int) *geth.ChainConfig {
	cID := new(big.Int).SetUint64(cc.ChainId)
	if chainID != nil {
		cID = chainID
	}
	return &geth.ChainConfig{
		ChainID:                 cID,
		HomesteadBlock:          getBlockValue(cc.HomesteadBlock),
		DAOForkBlock:            getBlockValue(cc.DAOForkBlock),
		DAOForkSupport:          cc.DAOForkSupport,
		EIP150Block:             getBlockValue(cc.EIP150Block),
		EIP150Hash:              common.HexToHash(cc.EIP150Hash),
		EIP155Block:             getBlockValue(cc.EIP155Block),
		EIP158Block:             getBlockValue(cc.EIP158Block),
		ByzantiumBlock:          getBlockValue(cc.ByzantiumBlock),
		ConstantinopleBlock:     getBlockValue(cc.ConstantinopleBlock),
		PetersburgBlock:         getBlockValue(cc.PetersburgBlock),
		IstanbulBlock:           getBlockValue(cc.IstanbulBlock),
		MuirGlacierBlock:        getBlockValue(cc.MuirGlacierBlock),
		BerlinBlock:             getBlockValue(cc.BerlinBlock),
		LondonBlock:             getBlockValue(cc.LondonBlock),
		ArrowGlacierBlock:       getBlockValue(cc.ArrowGlacierBlock),
		GrayGlacierBlock:        getBlockValue(cc.GrayGlacierBlock),
		MergeNetsplitBlock:      getBlockValue(cc.MergeNetsplitBlock),
		ShanghaiBlock:           getBlockValue(cc.ShanghaiBlock),
		CancunBlock:             getBlockValue(cc.CancunBlock),
		TerminalTotalDifficulty: nil,
		Ethash:                  nil,
		Clique:                  nil,
	}
}

func DefaultChainConfig(chainID string) *ChainConfig {
	if chainID == "" {
		chainID = utils.DevnetChainID
	}

	eip155ChainID, err := utils.ParseChainID(chainID)
	if err != nil {
		panic(err)
	}

	homesteadBlock := sdkmath.ZeroInt()
	daoForkBlock := sdkmath.ZeroInt()
	eip150Block := sdkmath.ZeroInt()
	eip155Block := sdkmath.ZeroInt()
	eip158Block := sdkmath.ZeroInt()
	byzantiumBlock := sdkmath.ZeroInt()
	constantinopleBlock := sdkmath.ZeroInt()
	petersburgBlock := sdkmath.ZeroInt()
	istanbulBlock := sdkmath.ZeroInt()
	muirGlacierBlock := sdkmath.ZeroInt()
	berlinBlock := sdkmath.ZeroInt()
	londonBlock := sdkmath.ZeroInt()
	arrowGlacierBlock := sdkmath.ZeroInt()
	grayGlacierBlock := sdkmath.ZeroInt()
	mergeNetsplitBlock := sdkmath.ZeroInt()
	shanghaiBlock := sdkmath.ZeroInt()
	cancunBlock := sdkmath.ZeroInt()
	cfg := &ChainConfig{
		ChainId:             eip155ChainID.Uint64(),
		HomesteadBlock:      &homesteadBlock,
		DAOForkBlock:        &daoForkBlock,
		DAOForkSupport:      true,
		EIP150Block:         &eip150Block,
		EIP150Hash:          common.Hash{}.String(),
		EIP155Block:         &eip155Block,
		EIP158Block:         &eip158Block,
		ByzantiumBlock:      &byzantiumBlock,
		ConstantinopleBlock: &constantinopleBlock,
		PetersburgBlock:     &petersburgBlock,
		IstanbulBlock:       &istanbulBlock,
		MuirGlacierBlock:    &muirGlacierBlock,
		BerlinBlock:         &berlinBlock,
		LondonBlock:         &londonBlock,
		ArrowGlacierBlock:   &arrowGlacierBlock,
		GrayGlacierBlock:    &grayGlacierBlock,
		MergeNetsplitBlock:  &mergeNetsplitBlock,
		ShanghaiBlock:       &shanghaiBlock,
		CancunBlock:         &cancunBlock,
	}
	return cfg
}

// setChainConfig allows to set the `chainConfig` variable modifying the
// default values. The method is private because it should only be called once
// in the EVMConfigurator.
func setChainConfig(cc *ChainConfig) error {
	if chainConfig != nil {
		return errors.New("chainConfig already set. Cannot set again the chainConfig")
	}
	config := DefaultChainConfig("")
	if cc != nil {
		config = cc
	}
	if err := config.Validate(); err != nil {
		return err
	}
	chainConfig = config

	return nil
}

func getBlockValue(block *sdkmath.Int) *big.Int {
	if block == nil || block.IsNegative() {
		return nil
	}

	return block.BigInt()
}

// Validate performs a basic validation of the ChainConfig params. The function will return an error
// if any of the block values is uninitialized (i.e nil) or if the EIP150Hash is an invalid hash.
func (cc ChainConfig) Validate() error {
	if err := validateBlock(cc.HomesteadBlock); err != nil {
		return errorsmod.Wrap(err, "homesteadBlock")
	}
	if err := validateBlock(cc.DAOForkBlock); err != nil {
		return errorsmod.Wrap(err, "daoForkBlock")
	}
	if err := validateBlock(cc.EIP150Block); err != nil {
		return errorsmod.Wrap(err, "eip150Block")
	}
	if err := validateHash(cc.EIP150Hash); err != nil {
		return err
	}
	if err := validateBlock(cc.EIP155Block); err != nil {
		return errorsmod.Wrap(err, "eip155Block")
	}
	if err := validateBlock(cc.EIP158Block); err != nil {
		return errorsmod.Wrap(err, "eip158Block")
	}
	if err := validateBlock(cc.ByzantiumBlock); err != nil {
		return errorsmod.Wrap(err, "byzantiumBlock")
	}
	if err := validateBlock(cc.ConstantinopleBlock); err != nil {
		return errorsmod.Wrap(err, "constantinopleBlock")
	}
	if err := validateBlock(cc.PetersburgBlock); err != nil {
		return errorsmod.Wrap(err, "petersburgBlock")
	}
	if err := validateBlock(cc.IstanbulBlock); err != nil {
		return errorsmod.Wrap(err, "istanbulBlock")
	}
	if err := validateBlock(cc.MuirGlacierBlock); err != nil {
		return errorsmod.Wrap(err, "muirGlacierBlock")
	}
	if err := validateBlock(cc.BerlinBlock); err != nil {
		return errorsmod.Wrap(err, "berlinBlock")
	}
	if err := validateBlock(cc.LondonBlock); err != nil {
		return errorsmod.Wrap(err, "londonBlock")
	}
	if err := validateBlock(cc.ArrowGlacierBlock); err != nil {
		return errorsmod.Wrap(err, "arrowGlacierBlock")
	}
	if err := validateBlock(cc.GrayGlacierBlock); err != nil {
		return errorsmod.Wrap(err, "GrayGlacierBlock")
	}
	if err := validateBlock(cc.MergeNetsplitBlock); err != nil {
		return errorsmod.Wrap(err, "MergeNetsplitBlock")
	}
	if err := validateBlock(cc.ShanghaiBlock); err != nil {
		return errorsmod.Wrap(err, "ShanghaiBlock")
	}
	if err := validateBlock(cc.CancunBlock); err != nil {
		return errorsmod.Wrap(err, "CancunBlock")
	}
	// NOTE: chain ID is not needed to check config order
	if err := cc.EthereumConfig(nil).CheckConfigForkOrder(); err != nil {
		return errorsmod.Wrap(err, "invalid config fork order")
	}
	return nil
}

func validateHash(hex string) error {
	if hex != "" && strings.TrimSpace(hex) == "" {
		return errorsmod.Wrap(ErrInvalidChainConfig, "hash cannot be blank")
	}

	return nil
}

func validateBlock(block *sdkmath.Int) error {
	// nil value means that the fork has not yet been applied
	if block == nil {
		return nil
	}

	if block.IsNegative() {
		return errorsmod.Wrapf(
			ErrInvalidChainConfig, "block value cannot be negative: %s", block,
		)
	}

	return nil
}
