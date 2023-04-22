package withdrawals

import (
	"encoding/json"
	"math/big"
	"os"
	"path"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/bindings/predeploys"
	"github.com/kroma-network/kroma/components/node/testutils"
)

func TestParseMessagePassed(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected *bindings.L2ToL1MessagePasserMessagePassed
	}{
		{
			"withdrawal through bridge",
			"bridge-withdrawal.json",
			&bindings.L2ToL1MessagePasserMessagePassed{
				Nonce:    new(big.Int),
				Sender:   common.HexToAddress(predeploys.L2CrossDomainMessenger),
				Target:   common.HexToAddress(predeploys.DevL1CrossDomainMessenger),
				Value:    new(big.Int),
				GasLimit: big.NewInt(203648),
				Data: hexutil.MustDecode(
					"0xd764ad0b00010000000000000000000000000000000000000000000000000000" +
						"00000000000000000000000000000000420000000000000000000000000000000000" +
						"00100000000000000000000000006900000000000000000000000000000000000003" +
						"00000000000000000000000000000000000000000000000000000000000000000000" +
						"00000000000000000000000000000000000000000000000000000000000000000000" +
						"000000000000000000000000000000000000000000000000000000c0000000000000" +
						"00000000000000000000000000000000000000000000000000e40166a07a00000000" +
						"000000000000000089d51be807d98fc974a0f41b2e67a8228d7846ef000000000000" +
						"0000000000007c6b91d9be155a6db01f749217d76ff02a7227f20000000000000000" +
						"00000000c20c5ec92fda6e611a08485123cdc0d5b84bd3a200000000000000000000" +
						"0000c20c5ec92fda6e611a08485123cdc0d5b84bd3a2000000000000000000000000" +
						"00000000000000000000000000000000000001f40000000000000000000000000000" +
						"0000000000000000000000000000000000c000000000000000000000000000000000" +
						"00000000000000000000000000000000000000000000000000000000000000000000" +
						"00000000000000000000",
				),
				WithdrawalHash: common.HexToHash("0x0d827f8148288e3a2466018f71b968ece4ea9f9e2a81c30da9bd46cce2868285"),
				Raw: types.Log{
					Address: common.HexToAddress(predeploys.L2ToL1MessagePasser),
					Topics: []common.Hash{
						common.HexToHash("0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054"),
						common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
						common.HexToHash(predeploys.L2CrossDomainMessenger),
						common.HexToHash(predeploys.DevL1CrossDomainMessenger),
					},
					Data: hexutil.MustDecode(
						"0x00000000000000000000000000000000000000000000000000000000000000" +
							"000000000000000000000000000000000000000000000000000000000000031b80" +
							"00000000000000000000000000000000000000000000000000000000000000800d" +
							"827f8148288e3a2466018f71b968ece4ea9f9e2a81c30da9bd46cce28682850000" +
							"0000000000000000000000000000000000000000000000000000000001e4d764ad" +
							"0b0001000000000000000000000000000000000000000000000000000000000000" +
							"000000000000000000000000420000000000000000000000000000000000001000" +
							"000000000000000000000069000000000000000000000000000000000000030000" +
							"000000000000000000000000000000000000000000000000000000000000000000" +
							"000000000000000000000000000000000000000000000000000000000000000000" +
							"000000000000000000000000000000000000000000000000000000c00000000000" +
							"0000000000000000000000000000000000000000000000000000e40166a07a0000" +
							"0000000000000000000089d51be807d98fc974a0f41b2e67a8228d7846ef000000" +
							"0000000000000000007c6b91d9be155a6db01f749217d76ff02a7227f200000000" +
							"0000000000000000c20c5ec92fda6e611a08485123cdc0d5b84bd3a20000000000" +
							"00000000000000c20c5ec92fda6e611a08485123cdc0d5b84bd3a2000000000000" +
							"00000000000000000000000000000000000000000000000001f400000000000000" +
							"000000000000000000000000000000000000000000000000c00000000000000000" +
							"000000000000000000000000000000000000000000000000000000000000000000" +
							"000000000000000000000000000000000000000000000000000000000000000000" +
							"0000000000000000000000000000",
					),
					BlockNumber: 0x36,
					TxHash:      common.HexToHash("0x9346381068b59d2098495baa72ed2f773c1e09458610a7a208984859dff73add"),
					TxIndex:     1,
					BlockHash:   common.HexToHash("0xfdd4ad8a984b45687aca0463db491cbd0e85273d970019a3f8bf618b614938df"),
					Index:       2,
					Removed:     false,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(path.Join("testdata", test.file))
			require.NoError(t, err)
			dec := json.NewDecoder(f)
			receipt := new(types.Receipt)
			require.NoError(t, dec.Decode(receipt))
			parsed, err := ParseMessagePassed(receipt)
			require.NoError(t, err)

			// Have to do this weird thing to compare zero bigints.
			// When they're deserialized from JSON, the internal byte
			// array is an empty array whereas it is nil in the expectation.
			parsedNonce := parsed.Nonce
			parsedValue := parsed.Value
			expNonce := test.expected.Nonce
			expValue := test.expected.Value
			testutils.RequireBigEqual(t, expNonce, parsedNonce)
			testutils.RequireBigEqual(t, expValue, parsedValue)
			parsed.Nonce = nil
			parsed.Value = nil
			test.expected.Nonce = nil
			test.expected.Value = nil

			require.EqualValues(t, test.expected, parsed)
		})
	}
}
