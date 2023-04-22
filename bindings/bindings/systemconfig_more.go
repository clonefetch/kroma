// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/kroma-network/kroma/bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1010_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1009_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b506004361061011b5760003560e01c80638f974d7f116100b2578063e81b2c6d11610081578063f45e65d811610066578063f45e65d814610286578063f68016b71461028f578063ffa1ad74146102a357600080fd5b8063e81b2c6d1461026a578063f2fde38b1461027357600080fd5b80638f974d7f1461021e578063935f029e14610231578063b40a817c14610244578063c9b26f611461025757600080fd5b80634f16540b116100ee5780634f16540b146101bc57806354fd4d50146101e3578063715018a6146101f85780638da5cb5b1461020057600080fd5b80630c18c1621461012057806318d139181461013c5780631fd19ee11461015157806329477e8614610199575b600080fd5b61012960655481565b6040519081526020015b60405180910390f35b61014f61014a366004610cb6565b6102ab565b005b7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08545b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610133565b6101a3627a120081565b60405167ffffffffffffffff9091168152602001610133565b6101297f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b6101eb61036f565b6040516101339190610d52565b61014f610412565b60335473ffffffffffffffffffffffffffffffffffffffff16610174565b61014f61022c366004610d7d565b610426565b61014f61023f366004610ddc565b6106aa565b61014f610252366004610dfe565b610743565b61014f610265366004610e19565b61081b565b61012960675481565b61014f610281366004610cb6565b61084b565b61012960665481565b6068546101a39067ffffffffffffffff1681565b610129600081565b6102b361091e565b6102db817f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0855565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516103639190610d52565b60405180910390a35050565b606061039a7f000000000000000000000000000000000000000000000000000000000000000061099f565b6103c37f000000000000000000000000000000000000000000000000000000000000000061099f565b6103ec7f000000000000000000000000000000000000000000000000000000000000000061099f565b6040516020016103fe93929190610e32565b604051602081830303815290604052905090565b61041a61091e565b6104246000610adc565b565b600054610100900460ff16158080156104465750600054600160ff909116105b806104605750303b158015610460575060005460ff166001145b6104f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561054f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b627a120067ffffffffffffffff841610156105c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016104e8565b6105ce610b53565b6105d78761084b565b606586905560668590556067849055606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff85161790557f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0882905580156106a157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b6106b261091e565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516107369190610d52565b60405180910390a3505050565b61074b61091e565b627a120067ffffffffffffffff821610156107c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016104e8565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040805160208082019390935281518082039093018352810190526002610332565b61082361091e565b6067819055604080516020808201849052825180830390910181529082019091526000610332565b61085361091e565b73ffffffffffffffffffffffffffffffffffffffff81166108f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104e8565b6108ff81610adc565b50565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60335473ffffffffffffffffffffffffffffffffffffffff163314610424576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104e8565b6060816000036109e257505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610a0c57806109f681610ed7565b9150610a059050600a83610f3e565b91506109e6565b60008167ffffffffffffffff811115610a2757610a27610f52565b6040519080825280601f01601f191660200182016040528015610a51576020820181803683370190505b5090505b8415610ad457610a66600183610f81565b9150610a73600a86610f98565b610a7e906030610fac565b60f81b818381518110610a9357610a93610fc4565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610acd600a86610f3e565b9450610a55565b949350505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610bea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104e8565b610424600054610100900460ff16610c84576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104e8565b61042433610adc565b803573ffffffffffffffffffffffffffffffffffffffff81168114610cb157600080fd5b919050565b600060208284031215610cc857600080fd5b610cd182610c8d565b9392505050565b60005b83811015610cf3578181015183820152602001610cdb565b83811115610d02576000848401525b50505050565b60008151808452610d20816020860160208601610cd8565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610cd16020830184610d08565b803567ffffffffffffffff81168114610cb157600080fd5b60008060008060008060c08789031215610d9657600080fd5b610d9f87610c8d565b9550602087013594506040870135935060608701359250610dc260808801610d65565b9150610dd060a08801610c8d565b90509295509295509295565b60008060408385031215610def57600080fd5b50508035926020909101359150565b600060208284031215610e1057600080fd5b610cd182610d65565b600060208284031215610e2b57600080fd5b5035919050565b60008451610e44818460208901610cd8565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610e80816001850160208a01610cd8565b60019201918201528351610e9b816002840160208801610cd8565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f0857610f08610ea8565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610f4d57610f4d610f0f565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015610f9357610f93610ea8565b500390565b600082610fa757610fa7610f0f565b500690565b60008219821115610fbf57610fbf610ea8565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}
