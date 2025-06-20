// GetIBCDenomAddress returns the address from the hash of the ICS20's DenomTrace Path.
package utils

import (
	"fmt"
	"strings"

	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	"github.com/ethereum/go-ethereum/common"
)

func GetIBCDenomAddress(denom string) (common.Address, error) {
	if !strings.HasPrefix(denom, "ibc/") {
		return common.Address{}, ibctransfertypes.ErrInvalidDenomForTransfer.Wrapf("coin %s does not have 'ibc/' prefix", denom)
	}

	if len(denom) < 5 || strings.TrimSpace(denom[4:]) == "" {
		return common.Address{}, ibctransfertypes.ErrInvalidDenomForTransfer.Wrapf("coin %s is not a valid IBC voucher hash", denom)
	}

	// Get the address from the hash of the ICS20's DenomTrace Path
	bz, err := ibctransfertypes.ParseHexHash(denom[4:])
	if err != nil {
		return common.Address{}, ibctransfertypes.ErrInvalidDenomForTransfer.Wrap(err.Error())
	}

	return common.BytesToAddress(bz), nil
}

// ComputeIBCDenomTrace compute the ibc voucher denom trace associated with
// the portID, channelID, and the given a token denomination.
func ComputeIBCDenomTrace(
	portID, channelID,
	denom string,
) ibctransfertypes.DenomTrace {
	denomTrace := ibctransfertypes.DenomTrace{
		Path:      fmt.Sprintf("%s/%s", portID, channelID),
		BaseDenom: denom,
	}

	return denomTrace
}

// ComputeIBCDenom compute the ibc voucher denom associated to
// the portID, channelID, and the given a token denomination.
func ComputeIBCDenom(
	portID, channelID,
	denom string,
) string {
	return ComputeIBCDenomTrace(portID, channelID, denom).IBCDenom()
}