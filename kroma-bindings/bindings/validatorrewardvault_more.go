// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const ValidatorRewardVaultStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/ValidatorRewardVault.sol:ValidatorRewardVault\",\"label\":\"totalProcessed\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint256\"},{\"astId\":1001,\"contract\":\"contracts/L2/ValidatorRewardVault.sol:ValidatorRewardVault\",\"label\":\"rewards\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1002,\"contract\":\"contracts/L2/ValidatorRewardVault.sol:ValidatorRewardVault\",\"label\":\"isPaid\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_uint256,t_bool)\"},{\"astId\":1003,\"contract\":\"contracts/L2/ValidatorRewardVault.sol:ValidatorRewardVault\",\"label\":\"totalReserved\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var ValidatorRewardVaultStorageLayout = new(solc.StorageLayout)

var ValidatorRewardVaultDeployedBin = "0x6080604052600436106100c05760003560e01c80636ed39f6211610074578063b98debbf1161004e578063b98debbf14610267578063c71b0e1c1461029b578063d3e5792b146102b157600080fd5b80636ed39f62146101f957806370a082311461020e57806384411d651461025157600080fd5b80633ccfd60b116100a55780633ccfd60b1461014c57806354fd4d501461016157806362aba76b146101b757600080fd5b80630d9019e1146100cc57806321670f221461012a57600080fd5b366100c757005b600080fd5b3480156100d857600080fd5b506101007f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561013657600080fd5b5061014a610145366004610956565b6102e5565b005b34801561015857600080fd5b5061014a610624565b34801561016d57600080fd5b506101aa6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161012191906109eb565b3480156101c357600080fd5b506101eb7f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610121565b34801561020557600080fd5b5061014a6106ca565b34801561021a57600080fd5b506101eb610229366004610a05565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b34801561025d57600080fd5b506101eb60005481565b34801561027357600080fd5b506101007f000000000000000000000000000000000000000000000000000000000000000081565b3480156102a757600080fd5b506101eb60035481565b3480156102bd57600080fd5b506101eb7f000000000000000000000000000000000000000000000000000000000000000081565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330173ffffffffffffffffffffffffffffffffffffffff161461040d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f56616c696461746f725265776172645661756c743a2066756e6374696f6e206360448201527f616e206f6e6c792062652063616c6c65642066726f6d207468652056616c696460648201527f61746f72506f6f6c000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166104b0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f56616c696461746f725265776172645661756c743a2076616c696461746f722060448201527f616464726573732063616e6e6f742062652030000000000000000000000000006064820152608401610404565b60008181526002602052604090205460ff1615610575576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604e60248201527f56616c696461746f725265776172645661756c743a207468652072657761726460448201527f2068617320616c7265616479206265656e207061696420666f7220746865204c60648201527f3220626c6f636b206e756d626572000000000000000000000000000000000000608482015260a401610404565b600061057f610786565b600380548201905573ffffffffffffffffffffffffffffffffffffffff84166000818152600160208181526040808420805487019055878452600282529283902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909217909155905183815292935084927f291e8ba3c0f4b0bd86e6e2346fcee1e7ca0975b1cc1886bfbc722d34f3f24d91910160405180910390a3505050565b600061062e6107c6565b604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081529192507342000000000000000000000000000000000000099163e11013dd9184916106959133916188b89190600401610a20565b6000604051808303818588803b1580156106ae57600080fd5b505af11580156106c2573d6000803e3d6000fd5b505050505050565b60006106d46107c6565b905060006106f3335a8460405180602001604052806000815250610913565b905080610782576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f56616c696461746f725265776172645661756c743a20455448207472616e736660448201527f6572206661696c656400000000000000000000000000000000000000000000006064820152608401610404565b5050565b60007f0000000000000000000000000000000000000000000000000000000000000000600354476107b79190610a64565b6107c19190610aa2565b905090565b336000908152600160205260408120547f00000000000000000000000000000000000000000000000000000000000000008110156108ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605660248201527f56616c696461746f725265776172645661756c743a207769746864726177616c60448201527f20616d6f756e74206d7573742062652067726561746572207468616e206d696e60648201527f696d756d207769746864726177616c20616d6f756e7400000000000000000000608482015260a401610404565b33600081815260016020908152604080832083905560038054869003905582548501909255815184815290810183905280820192909252517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1919050565b600080600080845160208601878a8af19695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461095157600080fd5b919050565b6000806040838503121561096957600080fd5b6109728361092d565b946020939093013593505050565b6000815180845260005b818110156109a65760208185018101518683018201520161098a565b818111156109b8576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006109fe6020830184610980565b9392505050565b600060208284031215610a1757600080fd5b6109fe8261092d565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff83166020820152606060408201526000610a5b6060830184610980565b95945050505050565b600082821015610a9d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500390565b600082610ad8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(ValidatorRewardVaultStorageLayoutJSON), ValidatorRewardVaultStorageLayout); err != nil {
		panic(err)
	}

	layouts["ValidatorRewardVault"] = ValidatorRewardVaultStorageLayout
	deployedBytecodes["ValidatorRewardVault"] = ValidatorRewardVaultDeployedBin
	immutableReferences["ValidatorRewardVault"] = true
}
