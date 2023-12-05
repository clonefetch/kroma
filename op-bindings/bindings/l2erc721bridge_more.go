// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2ERC721BridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2ERC721Bridge.sol:L2ERC721Bridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_array(t_uint256)1001_storage\"}],\"types\":{\"t_array(t_uint256)1001_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2ERC721BridgeStorageLayout = new(solc.StorageLayout)

var L2ERC721BridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100725760003560e01c80637f46ddb2116100505780637f46ddb2146100bd578063927ede2d14610109578063aa5574521461013057600080fd5b80633687011a1461007757806354fd4d501461008c578063761f4493146100aa575b600080fd5b61008a6100853660046111d1565b610143565b005b6100946101ef565b6040516100a191906112ce565b60405180910390f35b61008a6100b83660046112e1565b610292565b6100e47f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100a1565b6100e47f000000000000000000000000000000000000000000000000000000000000000081565b61008a61013e366004611379565b6107f9565b333b156101d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b6101e786863333888888886108b5565b505050505050565b606061021a7f0000000000000000000000000000000000000000000000000000000000000000610e53565b6102437f0000000000000000000000000000000000000000000000000000000000000000610e53565b61026c7f0000000000000000000000000000000000000000000000000000000000000000610e53565b60405160200161027e939291906113f0565b604051602081830303815290604052905090565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156103b057507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610374573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103989190611466565b73ffffffffffffffffffffffffffffffffffffffff16145b61043c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f74686572206272696467650060648201526084016101ce565b3073ffffffffffffffffffffffffffffffffffffffff8816036104e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c660000000000000000000000000000000000000000000060648201526084016101ce565b61050b877f74259ebf00000000000000000000000000000000000000000000000000000000610f11565b610597576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e20696e746560448201527f7266616365206973206e6f7420636f6d706c69616e740000000000000000000060648201526084016101ce565b8673ffffffffffffffffffffffffffffffffffffffff1663033964be6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106069190611466565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16146106e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f4c324552433732314272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204b726f6d61204d696e7461626c6520455243373231206c6f6360648201527f616c20746f6b656e000000000000000000000000000000000000000000000000608482015260a4016101ce565b6040517fa144819400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905288169063a144819490604401600060405180830381600087803b15801561075657600080fd5b505af115801561076a573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac878787876040516107e894939291906114cc565b60405180910390a450505050505050565b73ffffffffffffffffffffffffffffffffffffffff851661089c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f7420626520616464726573732830290000000000000000000000000000000060648201526084016101ce565b6108ac87873388888888886108b5565b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff8716610958576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f74206265206164647265737328302900000000000000000000000000000060648201526084016101ce565b6040517f6352211e0000000000000000000000000000000000000000000000000000000081526004810185905273ffffffffffffffffffffffffffffffffffffffff891690636352211e90602401602060405180830381865afa1580156109c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e79190611466565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610aa1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4c324552433732314272696467653a205769746864726177616c206973206e6f60448201527f74206265696e6720696e69746961746564206279204e4654206f776e6572000060648201526084016101ce565b60008873ffffffffffffffffffffffffffffffffffffffff1663033964be6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610aee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b129190611466565b90508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610bcf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e20646f6560448201527f73206e6f74206d6174636820676976656e2076616c756500000000000000000060648201526084016101ce565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152602482018790528a1690639dc29fac90604401600060405180830381600087803b158015610c3f57600080fd5b505af1158015610c53573d6000803e3d6000fd5b50505050600063761f449360e01b828b8a8a8a8989604051602401610c7e979695949392919061150c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290517f3dbb202b00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690633dbb202b90610d93907f00000000000000000000000000000000000000000000000000000000000000009085908a90600401611569565b600060405180830381600087803b158015610dad57600080fd5b505af1158015610dc1573d6000803e3d6000fd5b505050508773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a58a8a8989604051610e3f94939291906114cc565b60405180910390a450505050505050505050565b60606000610e6083610f34565b600101905060008167ffffffffffffffff811115610e8057610e806115ae565b6040519080825280601f01601f191660200182016040528015610eaa576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084610eb457509392505050565b6000610f1c83611017565b8015610f2d5750610f2d838361107b565b9392505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310610f7d577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310610fa9576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310610fc757662386f26fc10000830492506010015b6305f5e1008310610fdf576305f5e100830492506008015b6127108310610ff357612710830492506004015b60648310611005576064830492506002015b600a8310611011576001015b92915050565b6000611043827f01ffc9a70000000000000000000000000000000000000000000000000000000061107b565b80156110115750611074827fffffffff0000000000000000000000000000000000000000000000000000000061107b565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015611133575060208210155b801561113f5750600081115b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461116c57600080fd5b50565b803563ffffffff8116811461118357600080fd5b919050565b60008083601f84011261119a57600080fd5b50813567ffffffffffffffff8111156111b257600080fd5b6020830191508360208285010111156111ca57600080fd5b9250929050565b60008060008060008060a087890312156111ea57600080fd5b86356111f58161114a565b955060208701356112058161114a565b94506040870135935061121a6060880161116f565b9250608087013567ffffffffffffffff81111561123657600080fd5b61124289828a01611188565b979a9699509497509295939492505050565b60005b8381101561126f578181015183820152602001611257565b8381111561127e576000848401525b50505050565b6000815180845261129c816020860160208601611254565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610f2d6020830184611284565b600080600080600080600060c0888a0312156112fc57600080fd5b87356113078161114a565b965060208801356113178161114a565b955060408801356113278161114a565b945060608801356113378161114a565b93506080880135925060a088013567ffffffffffffffff81111561135a57600080fd5b6113668a828b01611188565b989b979a50959850939692959293505050565b600080600080600080600060c0888a03121561139457600080fd5b873561139f8161114a565b965060208801356113af8161114a565b955060408801356113bf8161114a565b9450606088013593506113d46080890161116f565b925060a088013567ffffffffffffffff81111561135a57600080fd5b60008451611402818460208901611254565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161143e816001850160208a01611254565b60019201918201528351611459816002840160208801611254565b0160020195945050505050565b60006020828403121561147857600080fd5b8151610f2d8161114a565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152606060408201526000611502606083018486611483565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a083015261155c60c083018486611483565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006115986060830185611284565b905063ffffffff83166040830152949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2ERC721BridgeStorageLayoutJSON), L2ERC721BridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ERC721Bridge"] = L2ERC721BridgeStorageLayout
	deployedBytecodes["L2ERC721Bridge"] = L2ERC721BridgeDeployedBin
}