// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2OutputOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"startingBlockNumber\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"startingTimestamp\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1004,\"contract\":\"contracts/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"l2Outputs\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_struct(CheckpointOutput)1005_storage)dyn_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_struct(CheckpointOutput)1005_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct Types.CheckpointOutput[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(CheckpointOutput)1005_storage\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(CheckpointOutput)1005_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Types.CheckpointOutput\",\"numberOfBytes\":\"96\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2OutputOracleStorageLayout = new(solc.StorageLayout)

var L2OutputOracleDeployedBin = "0x6080604052600436106101745760003560e01c80639e45e8f4116100cb578063cf8e5cf01161007f578063e4a3011611610059578063e4a30116146104bb578063e6646723146104db578063f4daa291146104fb57600080fd5b8063cf8e5cf014610466578063d1de856c14610486578063dcec3348146104a657600080fd5b8063a48ea6de116100b0578063a48ea6de146103f2578063b0ea09a814610412578063b98debbf1461043257600080fd5b80639e45e8f41461031a578063a25ae5571461037357600080fd5b80635a045f781161012d57806370872aa51161010757806370872aa5146102ce5780637f006420146102e4578063887862721461030457600080fd5b80635a045f781461028f57806369f16eec146102a45780636abcf563146102b957600080fd5b80634599c7881161015e5780634599c788146101f0578063529933df1461020557806354fd4d501461023957600080fd5b80622134cc1461017957806333727c4d146101c0575b600080fd5b34801561018557600080fd5b506101ad7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b3480156101cc57600080fd5b506101e06101db3660046115f1565b61052f565b60405190151581526020016101b7565b3480156101fc57600080fd5b506101ad61059d565b34801561021157600080fd5b506101ad7f000000000000000000000000000000000000000000000000000000000000000081565b34801561024557600080fd5b506102826040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101b7919061160a565b6102a261029d36600461167d565b610610565b005b3480156102b057600080fd5b506101ad610c17565b3480156102c557600080fd5b506003546101ad565b3480156102da57600080fd5b506101ad60015481565b3480156102f057600080fd5b506101ad6102ff3660046115f1565b610c29565b34801561031057600080fd5b506101ad60025481565b34801561032657600080fd5b5061034e7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b7565b34801561037f57600080fd5b5061039361038e3660046115f1565b610e09565b60408051825173ffffffffffffffffffffffffffffffffffffffff16815260208084015190820152828201516fffffffffffffffffffffffffffffffff90811692820192909252606092830151909116918101919091526080016101b7565b3480156103fe57600080fd5b506101ad61040d3660046115f1565b610ec6565b34801561041e57600080fd5b5061034e61042d3660046115f1565b610f32565b34801561043e57600080fd5b5061034e7f000000000000000000000000000000000000000000000000000000000000000081565b34801561047257600080fd5b506103936104813660046115f1565b610f74565b34801561049257600080fd5b506101ad6104a13660046115f1565b610fb3565b3480156104b257600080fd5b506101ad610ffb565b3480156104c757600080fd5b506102a26104d63660046116af565b611041565b3480156104e757600080fd5b506102a26104f63660046116f6565b61125a565b34801561050757600080fd5b506101ad7f000000000000000000000000000000000000000000000000000000000000000081565b6000427f0000000000000000000000000000000000000000000000000000000000000000600384815481106105665761056661172f565b600091825260209091206002600390920201015461059691906fffffffffffffffffffffffffffffffff1661178d565b1092915050565b6003546000901561060757600380546105b8906001906117a5565b815481106105c8576105c861172f565b600091825260209091206003909102016002015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16919050565b6001545b905090565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633a5490466040518163ffffffff1660e01b8152600401602060405180830381865afa15801561067d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a191906117bc565b905073ffffffffffffffffffffffffffffffffffffffff80821614610776573373ffffffffffffffffffffffffffffffffffffffff8216146107765760405162461bcd60e51b815260206004820152604260248201527f4c324f75747075744f7261636c653a206f6e6c7920746865206e65787420736560448201527f6c65637465642076616c696461746f722063616e207375626d6974206f75747060648201527f7574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b61077e610ffb565b84146108185760405162461bcd60e51b815260206004820152604860248201527f4c324f75747075744f7261636c653a20626c6f636b206e756d626572206d757360448201527f7420626520657175616c20746f206e65787420657870656374656420626c6f6360648201527f6b206e756d626572000000000000000000000000000000000000000000000000608482015260a40161076d565b4261082285610fb3565b106108955760405162461bcd60e51b815260206004820152603560248201527f4c324f75747075744f7261636c653a2063616e6e6f74207375626d6974204c3260448201527f206f757470757420696e20746865206675747572650000000000000000000000606482015260840161076d565b846109085760405162461bcd60e51b815260206004820152603c60248201527f4c324f75747075744f7261636c653a204c3220636865636b706f696e74206f7560448201527f747075742063616e6e6f7420626520746865207a65726f206861736800000000606482015260840161076d565b82158015906109175750814015155b156109b857828240146109b85760405162461bcd60e51b815260206004820152604960248201527f4c324f75747075744f7261636c653a20626c6f636b206861736820646f65732060448201527f6e6f74206d61746368207468652068617368206174207468652065787065637460648201527f6564206865696768740000000000000000000000000000000000000000000000608482015260a40161076d565b60006109c360035490565b60408051608081018252338152602081018981526fffffffffffffffffffffffffffffffff428181168486019081528b831660608601908152600380546001810182556000829052965196027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b8101805473ffffffffffffffffffffffffffffffffffffffff989098167fffffffffffffffffffffffff00000000000000000000000000000000000000009098169790971790965593517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c86015551925182167001000000000000000000000000000000000292909116919091177fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d9092019190915590519192508691839189917f457b4388026260019ae0b0b4f16c98235d74fe7359be469bdcba16e6d0d4968991610b209190815260200190565b60405180910390a473ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663d38dc7ee82610b8f7f00000000000000000000000000000000000000000000000000000000000000004261178d565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815260048101929092526fffffffffffffffffffffffffffffffff166024820152604401600060405180830381600087803b158015610bf757600080fd5b505af1158015610c0b573d6000803e3d6000fd5b50505050505050505050565b60035460009061060b906001906117a5565b6000610c3361059d565b821115610cce5760405162461bcd60e51b815260206004820152604960248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f7420666f72206120626c6f636b207468617420686173206e6f74206265656e2060648201527f7375626d69747465640000000000000000000000000000000000000000000000608482015260a40161076d565b600354610d695760405162461bcd60e51b815260206004820152604760248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f74206173206e6f206f7574707574732068617665206265656e207375626d697460648201527f7465642079657400000000000000000000000000000000000000000000000000608482015260a40161076d565b6003546000905b80821015610e025760006002610d86838561178d565b610d9091906117e0565b90508460038281548110610da657610da661172f565b600091825260209091206003909102016002015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff161015610df857610df181600161178d565b9250610dfc565b8091505b50610d70565b5092915050565b60408051608081018252600080825260208201819052918101829052606081019190915260038281548110610e4057610e4061172f565b6000918252602091829020604080516080810182526003909302909101805473ffffffffffffffffffffffffffffffffffffffff1683526001810154938301939093526002909201546fffffffffffffffffffffffffffffffff808216938301939093527001000000000000000000000000000000009004909116606082015292915050565b60007f000000000000000000000000000000000000000000000000000000000000000060038381548110610efc57610efc61172f565b6000918252602090912060026003909202010154610f2c91906fffffffffffffffffffffffffffffffff1661178d565b92915050565b600060038281548110610f4757610f4761172f565b600091825260209091206003909102015473ffffffffffffffffffffffffffffffffffffffff1692915050565b6040805160808101825260008082526020820181905291810182905260608101919091526003610fa383610c29565b81548110610e4057610e4061172f565b60007f000000000000000000000000000000000000000000000000000000000000000060015483610fe491906117a5565b610fee919061181b565b600254610f2c919061178d565b60035460009015611039577f000000000000000000000000000000000000000000000000000000000000000061102f61059d565b61060b919061178d565b61060b61059d565b600054610100900460ff16158080156110615750600054600160ff909116105b8061107b5750303b15801561107b575060005460ff166001145b6110ed5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161076d565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561114b57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b428211156111e85760405162461bcd60e51b8152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201527f74696d6500000000000000000000000000000000000000000000000000000000608482015260a40161076d565b60028290556001839055801561125557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461132b5760405162461bcd60e51b815260206004820152604160248201527f4c324f75747075744f7261636c653a206f6e6c792074686520636f6c6f73736560448201527f756d20636f6e74726163742063616e207265706c61636520616e206f7574707560648201527f7400000000000000000000000000000000000000000000000000000000000000608482015260a40161076d565b73ffffffffffffffffffffffffffffffffffffffff81166113b45760405162461bcd60e51b815260206004820152603060248201527f4c324f75747075744f7261636c653a207375626d69747465722061646472657360448201527f732063616e6e6f74206265207a65726f00000000000000000000000000000000606482015260840161076d565b60035483106114515760405162461bcd60e51b815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f74207265706c616365206160448201527f6e206f757470757420616674657220746865206c6174657374206f757470757460648201527f20696e6465780000000000000000000000000000000000000000000000000000608482015260a40161076d565b6000600384815481106114665761146661172f565b6000918252602090912060039091020160028101549091507f0000000000000000000000000000000000000000000000000000000000000000906114bc906fffffffffffffffffffffffffffffffff16426117a5565b106115555760405162461bcd60e51b815260206004820152604860248201527f4c324f75747075744f7261636c653a2063616e6e6f74207265706c616365206160448201527f6e206f757470757420746861742068617320616c7265616479206265656e206660648201527f696e616c697a6564000000000000000000000000000000000000000000000000608482015260a40161076d565b6001810183905580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831617815560405183815284907fa1b831bb8b6b242db6d0988a6d21f869c610de9f703a5e45e1b7d3dc3137b9069060200160405180910390a250505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60006020828403121561160357600080fd5b5035919050565b600060208083528351808285015260005b818110156116375785810183015185820160400152820161161b565b81811115611649576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b6000806000806080858703121561169357600080fd5b5050823594602084013594506040840135936060013592509050565b600080604083850312156116c257600080fd5b50508035926020909101359150565b73ffffffffffffffffffffffffffffffffffffffff811681146116f357600080fd5b50565b60008060006060848603121561170b57600080fd5b83359250602084013591506040840135611724816116d1565b809150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156117a0576117a061175e565b500190565b6000828210156117b7576117b761175e565b500390565b6000602082840312156117ce57600080fd5b81516117d9816116d1565b9392505050565b600082611816577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156118535761185361175e565b50029056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L2OutputOracleStorageLayoutJSON), L2OutputOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2OutputOracle"] = L2OutputOracleStorageLayout
	deployedBytecodes["L2OutputOracle"] = L2OutputOracleDeployedBin
	immutableReferences["L2OutputOracle"] = true
}
