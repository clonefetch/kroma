// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/kroma-network/kroma/bindings/solc"
)

const KromaPortalStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"params\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_struct(ResourceParams)1010_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)1008_storage\"},{\"astId\":1004,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"l2Sender\",\"offset\":0,\"slot\":\"50\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"finalizedWithdrawals\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1006,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"provenWithdrawals\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_mapping(t_bytes32,t_struct(ProvenWithdrawal)1009_storage)\"},{\"astId\":1007,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"paused\",\"offset\":0,\"slot\":\"53\",\"type\":\"t_bool\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1008_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[48]\",\"numberOfBytes\":\"1536\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_mapping(t_bytes32,t_struct(ProvenWithdrawal)1009_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e struct KromaPortal.ProvenWithdrawal)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_struct(ProvenWithdrawal)1009_storage\"},\"t_struct(ProvenWithdrawal)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"struct KromaPortal.ProvenWithdrawal\",\"numberOfBytes\":\"64\"},\"t_struct(ResourceParams)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceParams\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var KromaPortalStorageLayout = new(solc.StorageLayout)

var KromaPortalDeployedBin = "0x6080604052600436106101835760003560e01c8063724c184c116100d6578063ca3e99ba1161007f578063d53a822f11610059578063d53a822f14610523578063e965084c14610543578063e9e05c42146105cf57600080fd5b8063ca3e99ba14610458578063cd7c97891461046d578063cff0ab961461048257600080fd5b80638c3152e9116100b05780638c3152e9146103db5780639bf62d82146103fb578063a14238e71461042857600080fd5b8063724c184c1461037a5780638456cb59146103ae578063867ead13146103c357600080fd5b80635865b6071161013857806364b792081161011257806364b792081461032d5780636bb0291e146103455780636dbffb781461035a57600080fd5b80635865b6071461029d5780635c1f2827146102d15780635c975abb1461030357600080fd5b806313620abd1161016957806313620abd1461022d5780633f4ba83a1461026657806354fd4d501461027b57600080fd5b80621c2ff6146101af5780630757b2441461020d57600080fd5b366101aa576101a83334620186a06000604051806020016040528060008152506105dd565b005b600080fd5b3480156101bb57600080fd5b506101e37f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561021957600080fd5b506101a8610228366004612a2b565b6107c3565b34801561023957600080fd5b50610245633b9aca0081565b6040516fffffffffffffffffffffffffffffffff9091168152602001610204565b34801561027257600080fd5b506101a8610dc1565b34801561028757600080fd5b50610290610eca565b6040516102049190612b82565b3480156102a957600080fd5b506101e37f000000000000000000000000000000000000000000000000000000000000000081565b3480156102dd57600080fd5b506102f56fffffffffffffffffffffffffffffffff81565b604051908152602001610204565b34801561030f57600080fd5b5060355461031d9060ff1681565b6040519015158152602001610204565b34801561033957600080fd5b506102f56301312d0081565b34801561035157600080fd5b506102f5600a81565b34801561036657600080fd5b5061031d610375366004612b95565b610f6d565b34801561038657600080fd5b506101e37f000000000000000000000000000000000000000000000000000000000000000081565b3480156103ba57600080fd5b506101a8611044565b3480156103cf57600080fd5b506102f5633b9aca0081565b3480156103e757600080fd5b506101a86103f6366004612bae565b611149565b34801561040757600080fd5b506032546101e39073ffffffffffffffffffffffffffffffffffffffff1681565b34801561043457600080fd5b5061031d610443366004612b95565b60336020526000908152604090205460ff1681565b34801561046457600080fd5b506102f5611914565b34801561047957600080fd5b506102f5600881565b34801561048e57600080fd5b506001546104ea906fffffffffffffffffffffffffffffffff81169067ffffffffffffffff7001000000000000000000000000000000008204811691780100000000000000000000000000000000000000000000000090041683565b604080516fffffffffffffffffffffffffffffffff909416845267ffffffffffffffff9283166020850152911690820152606001610204565b34801561052f57600080fd5b506101a861053e366004612bf4565b611926565b34801561054f57600080fd5b506105a161055e366004612b95565b603460205260009081526040902080546001909101546fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041683565b604080519384526fffffffffffffffffffffffffffffffff9283166020850152911690820152606001610204565b6101a86105dd366004612c11565b8260005a905083156106785773ffffffffffffffffffffffffffffffffffffffff8716156106785760405162461bcd60e51b815260206004820152603d60248201527f4b726f6d61506f7274616c3a206d7573742073656e6420746f2061646472657360448201527f73283029207768656e206372656174696e67206120636f6e747261637400000060648201526084015b60405180910390fd5b6152088567ffffffffffffffff1610156106fa5760405162461bcd60e51b815260206004820152603560248201527f4b726f6d61506f7274616c3a20676173206c696d6974206d75737420636f766560448201527f7220696e737472696e7369632067617320636f73740000000000000000000000606482015260840161066f565b3332811461071b575033731111000000000000000000000000000000001111015b60003488888888604051602001610736959493929190612c99565b604051602081830303815290604052905060008973ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32846040516107a69190612b82565b60405180910390a450506107ba8282611b15565b50505050505050565b60355460ff16156108165760405162461bcd60e51b815260206004820152601360248201527f4b726f6d61506f7274616c3a2070617573656400000000000000000000000000604482015260640161066f565b3073ffffffffffffffffffffffffffffffffffffffff16856040015173ffffffffffffffffffffffffffffffffffffffff16036108bb5760405162461bcd60e51b815260206004820152603c60248201527f4b726f6d61506f7274616c3a20796f752063616e6e6f742073656e64206d657360448201527f736167657320746f2074686520706f7274616c20636f6e747261637400000000606482015260840161066f565b6040517fa25ae557000000000000000000000000000000000000000000000000000000008152600481018590526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae55790602401606060405180830381865afa158015610949573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096d9190612d1e565b51905061098761098236869003860186612d83565b611df4565b81146109fb5760405162461bcd60e51b815260206004820152602660248201527f4b726f6d61506f7274616c3a20696e76616c6964206f757470757420726f6f7460448201527f2070726f6f660000000000000000000000000000000000000000000000000000606482015260840161066f565b6000610a0687611eaa565b6000818152603460209081526040918290208251606081018452815481526001909101546fffffffffffffffffffffffffffffffff8082169383018490527001000000000000000000000000000000009091041692810192909252919250901580610b385750805160408083015190517fa25ae5570000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff90911660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae55790602401606060405180830381865afa158015610b10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b349190612d1e565b5114155b610baa5760405162461bcd60e51b815260206004820152603460248201527f4b726f6d61506f7274616c3a207769746864726177616c20686173682068617360448201527f20616c7265616479206265656e2070726f76656e000000000000000000000000606482015260840161066f565b60408051602080820185905260008284015282518083038401815260608301808552815191909201207f12e64a7200000000000000000000000000000000000000000000000000000000909152917f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16916312e64a7291610c4c9185918b918b918e013590606401612e3c565b602060405180830381865afa158015610c69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8d9190612f3d565b610cff5760405162461bcd60e51b815260206004820152602f60248201527f4b726f6d61506f7274616c3a20696e76616c6964207769746864726177616c2060448201527f696e636c7573696f6e2070726f6f660000000000000000000000000000000000606482015260840161066f565b604080516060810182528581526fffffffffffffffffffffffffffffffff42811660208084019182528c831684860190815260008981526034835286812095518655925190518416700100000000000000000000000000000000029316929092176001909301929092558b830151908c0151925173ffffffffffffffffffffffffffffffffffffffff918216939091169186917f67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f629190a4505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610e6c5760405162461bcd60e51b815260206004820152602660248201527f4b726f6d61506f7274616c3a206f6e6c7920677561726469616e2063616e207560448201527f6e70617573650000000000000000000000000000000000000000000000000000606482015260840161066f565b603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040513381527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa906020015b60405180910390a1565b6060610ef57f0000000000000000000000000000000000000000000000000000000000000000611ef7565b610f1e7f0000000000000000000000000000000000000000000000000000000000000000611ef7565b610f477f0000000000000000000000000000000000000000000000000000000000000000611ef7565b604051602001610f5993929190612f5a565b604051602081830303815290604052905090565b6040517fa25ae5570000000000000000000000000000000000000000000000000000000081526004810182905260009061103e9073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063a25ae55790602401606060405180830381865afa158015610fff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110239190612d1e565b602001516fffffffffffffffffffffffffffffffff16612034565b92915050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146110ee5760405162461bcd60e51b8152602060048201526024808201527f4b726f6d61506f7274616c3a206f6e6c7920677561726469616e2063616e207060448201527f6175736500000000000000000000000000000000000000000000000000000000606482015260840161066f565b603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040513381527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25890602001610ec0565b60355460ff161561119c5760405162461bcd60e51b815260206004820152601360248201527f4b726f6d61506f7274616c3a2070617573656400000000000000000000000000604482015260640161066f565b60325473ffffffffffffffffffffffffffffffffffffffff1661dead1461122b5760405162461bcd60e51b815260206004820152603c60248201527f4b726f6d61506f7274616c3a2063616e206f6e6c792074726967676572206f6e60448201527f65207769746864726177616c20706572207472616e73616374696f6e00000000606482015260840161066f565b600061123682611eaa565b60008181526034602090815260408083208151606081018352815481526001909101546fffffffffffffffffffffffffffffffff808216948301859052700100000000000000000000000000000000909104169181019190915292935090036113075760405162461bcd60e51b815260206004820152602f60248201527f4b726f6d61506f7274616c3a207769746864726177616c20686173206e6f742060448201527f6265656e2070726f76656e207965740000000000000000000000000000000000606482015260840161066f565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663887862726040518163ffffffff1660e01b8152600401602060405180830381865afa158015611372573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113969190612fd0565b81602001516fffffffffffffffffffffffffffffffff1610156114475760405162461bcd60e51b815260206004820152604860248201527f4b726f6d61506f7274616c3a207769746864726177616c2074696d657374616d60448201527f70206c657373207468616e204c32204f7261636c65207374617274696e67207460648201527f696d657374616d70000000000000000000000000000000000000000000000000608482015260a40161066f565b61146681602001516fffffffffffffffffffffffffffffffff16612034565b6114fe5760405162461bcd60e51b815260206004820152604260248201527f4b726f6d61506f7274616c3a2070726f76656e207769746864726177616c206660448201527f696e616c697a6174696f6e20706572696f6420686173206e6f7420656c61707360648201527f6564000000000000000000000000000000000000000000000000000000000000608482015260a40161066f565b60408181015190517fa25ae5570000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff90911660048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae55790602401606060405180830381865afa1580156115a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115c79190612d1e565b82518151919250146116675760405162461bcd60e51b815260206004820152604660248201527f4b726f6d61506f7274616c3a206f757470757420726f6f742070726f76656e2060448201527f6973206e6f74207468652073616d652061732063757272656e74206f7574707560648201527f7420726f6f740000000000000000000000000000000000000000000000000000608482015260a40161066f565b61168681602001516fffffffffffffffffffffffffffffffff16612034565b61171e5760405162461bcd60e51b815260206004820152604260248201527f4b726f6d61506f7274616c3a20636865636b706f696e74206f7574707574206660448201527f696e616c697a6174696f6e20706572696f6420686173206e6f7420656c61707360648201527f6564000000000000000000000000000000000000000000000000000000000000608482015260a40161066f565b60008381526033602052604090205460ff16156117a35760405162461bcd60e51b815260206004820152603260248201527f4b726f6d61506f7274616c3a207769746864726177616c2068617320616c726560448201527f616479206265656e2066696e616c697a65640000000000000000000000000000606482015260840161066f565b600083815260336020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055908601516032805473ffffffffffffffffffffffffffffffffffffffff9092167fffffffffffffffffffffffff00000000000000000000000000000000000000009092169190911790558501516080860151606087015160a0880151611845939291906120d7565b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905560405190915084907fdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b906118aa90841515815260200190565b60405180910390a2801580156118c05750326001145b1561190d5760405162461bcd60e51b815260206004820152601e60248201527f4b726f6d61506f7274616c3a207769746864726177616c206661696c65640000604482015260640161066f565b5050505050565b611923600a6301312d00613047565b81565b600054610100900460ff16158080156119465750600054600160ff909116105b806119605750303b158015611960575060005460ff166001145b6119d25760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161066f565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611a3057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055603580548315157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909116179055611a92612131565b8015611af557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b600154600090611b4b907801000000000000000000000000000000000000000000000000900467ffffffffffffffff16436130af565b90508015611c83576000611b64600a6301312d00613047565b600154611b8f9190700100000000000000000000000000000000900467ffffffffffffffff166130c6565b905060006008611ba4600a6301312d00613047565b611bae919061313a565b600154611bce9084906fffffffffffffffffffffffffffffffff1661313a565b611bd89190613047565b600154909150600090611c1d90611c029084906fffffffffffffffffffffffffffffffff166131f6565b633b9aca006fffffffffffffffffffffffffffffffff6121fa565b90506001841115611c4457611c41611c02826008611c3c6001896130af565b61220f565b90505b6fffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff4316021760015550505b60018054849190601090611cb6908490700100000000000000000000000000000000900467ffffffffffffffff1661326a565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506301312d00600160000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff161315611d795760405162461bcd60e51b815260206004820152603e60248201527f5265736f757263654d65746572696e673a2063616e6e6f7420627579206d6f7260448201527f6520676173207468616e20617661696c61626c6520676173206c696d69740000606482015260840161066f565b600154600090611da5906fffffffffffffffffffffffffffffffff1667ffffffffffffffff8616613296565b90506000611db748633b9aca00612264565b611dc190836132d3565b905060005a611dd090866130af565b905080821115611dec57611dec611de782846130af565b61227d565b505050505050565b8051600090611e065761103e826122ab565b81517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611e375761103e826122ea565b60405162461bcd60e51b815260206004820152602a60248201527f48617368696e673a20756e6b6e6f776e206f757470757420726f6f742070726f60448201527f6f662076657273696f6e00000000000000000000000000000000000000000000606482015260840161066f565b919050565b80516020808301516040808501516060860151608087015160a08801519351600097611eda9790969591016132e7565b604051602081830303815290604052805190602001209050919050565b606081600003611f3a57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611f645780611f4e8161333e565b9150611f5d9050600a836132d3565b9150611f3e565b60008167ffffffffffffffff811115611f7f57611f7f612888565b6040519080825280601f01601f191660200182016040528015611fa9576020820181803683370190505b5090505b841561202c57611fbe6001836130af565b9150611fcb600a86613376565b611fd690603061338a565b60f81b818381518110611feb57611feb6133a2565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350612025600a866132d3565b9450611fad565b949350505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f4daa2916040518163ffffffff1660e01b8152600401602060405180830381865afa1580156120a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120c59190612fd0565b6120cf908361338a565b421192915050565b600080603f60c88601604002045a101561211a576308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b600080845160208601878a5af19695505050505050565b600054610100900460ff166121ae5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161066f565b60408051606081018252633b9aca00808252600060208301524367ffffffffffffffff169190920181905278010000000000000000000000000000000000000000000000000217600155565b600061202c612209858561232d565b8361233d565b6000670de0b6b3a76400006122506122278583613047565b61223990670de0b6b3a76400006130c6565b61224b85670de0b6b3a764000061313a565b61234c565b61225a908661313a565b61202c9190613047565b6000818310156122745781612276565b825b9392505050565b6000805a90505b825a61229090836130af565b10156122a65761229f8261333e565b9150612284565b505050565b60008160000151826020015183604001518460600151604051602001611eda949392919093845260208401929092526040830152606082015260800190565b80516020808301516040808501516060808701516080808901518551978801989098529386019490945284015282015260a081019190915260009060c001611eda565b6000818312156122745781612276565b60008183126122745781612276565b6000612276670de0b6b3a7640000836123648661237d565b61236e919061313a565b6123789190613047565b6125a7565b60008082136123ce5760405162461bcd60e51b815260206004820152600960248201527f554e444546494e45440000000000000000000000000000000000000000000000604482015260640161066f565b600060606123db846127cc565b03609f8181039490941b90931c6c465772b2bbbb5f824b15207a3081018102606090811d6d0388eaa27412d5aca026815d636e018202811d6d0df99ac502031bf953eff472fdcc018202811d6d13cdffb29d51d99322bdff5f2211018202811d6d0a0f742023def783a307a986912e018202811d6d01920d8043ca89b5239253284e42018202811d6c0b7a86d7375468fac667a0a527016c29508e458543d8aa4df2abee7883018302821d6d0139601a2efabe717e604cbb4894018302821d6d02247f7a7b6594320649aa03aba1018302821d7fffffffffffffffffffffffffffffffffffffff73c0c716a594e00d54e3c4cbc9018302821d7ffffffffffffffffffffffffffffffffffffffdc7b88c420e53a9890533129f6f01830290911d7fffffffffffffffffffffffffffffffffffffff465fda27eb4d63ded474e5f832019091027ffffffffffffffff5f6af8f7b3396644f18e157960000000000000000000000000105711340daa0d5f769dba1915cef59f0815a5506027d0267a36c0c95b3975ab3ee5b203a7614a3f75373f047d803ae7b6687f2b393909302929092017d57115e47018c7177eebf7cd370a3356a1b7863008a5ae8028c72b88642840160ae1d92915050565b60007ffffffffffffffffffffffffffffffffffffffffffffffffdb731c958f34d94c182136125d857506000919050565b680755bf798b4a1bf1e582126126305760405162461bcd60e51b815260206004820152600c60248201527f4558505f4f564552464c4f570000000000000000000000000000000000000000604482015260640161066f565b6503782dace9d9604e83901b059150600060606bb17217f7d1cf79abc9e3b39884821b056b80000000000000000000000001901d6bb17217f7d1cf79abc9e3b39881029093037fffffffffffffffffffffffffffffffffffffffdbf3ccf1604d263450f02a550481018102606090811d6d0277594991cfc85f6e2461837cd9018202811d7fffffffffffffffffffffffffffffffffffffe5adedaa1cb095af9e4da10e363c018202811d6db1bbb201f443cf962f1a1d3db4a5018202811d7ffffffffffffffffffffffffffffffffffffd38dc772608b0ae56cce01296c0eb018202811d6e05180bb14799ab47a8a8cb2a527d57016d02d16720577bd19bf614176fe9ea6c10fe68e7fd37d0007b713f765084018402831d9081019084017ffffffffffffffffffffffffffffffffffffffe2c69812cf03b0763fd454a8f7e010290911d6e0587f503bb6ea29d25fcb7401964500190910279d835ebba824c98fb31b83b2ca45c000000000000000000000000010574029d9dc38563c32e5c2f6dc192ee70ef65f9978af30260c3939093039290921c92915050565b600080821161281d5760405162461bcd60e51b815260206004820152600960248201527f554e444546494e45440000000000000000000000000000000000000000000000604482015260640161066f565b5060016fffffffffffffffffffffffffffffffff821160071b82811c67ffffffffffffffff1060061b1782811c63ffffffff1060051b1782811c61ffff1060041b1782811c60ff10600390811b90911783811c600f1060021b1783811c909110821b1791821c111790565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b803573ffffffffffffffffffffffffffffffffffffffff81168114611ea557600080fd5b600082601f8301126128ec57600080fd5b813567ffffffffffffffff8082111561290757612907612888565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561294d5761294d612888565b8160405283815286602085880101111561296657600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060c0828403121561299857600080fd5b60405160c0810167ffffffffffffffff82821081831117156129bc576129bc612888565b81604052829350843583526129d3602086016128b7565b60208401526129e4604086016128b7565b6040840152606085013560608401526080850135608084015260a0850135915080821115612a1157600080fd5b50612a1e858286016128db565b60a0830152505092915050565b6000806000806000858703610100811215612a4557600080fd5b863567ffffffffffffffff80821115612a5d57600080fd5b612a698a838b01612986565b97506020890135965060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc084011215612aa257600080fd5b60408901955060e0890135925080831115612abc57600080fd5b828901925089601f840112612ad057600080fd5b8235915080821115612ae157600080fd5b508860208260051b8401011115612af757600080fd5b959894975092955050506020019190565b60005b83811015612b23578181015183820152602001612b0b565b83811115612b32576000848401525b50505050565b60008151808452612b50816020860160208601612b08565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006122766020830184612b38565b600060208284031215612ba757600080fd5b5035919050565b600060208284031215612bc057600080fd5b813567ffffffffffffffff811115612bd757600080fd5b61202c84828501612986565b8015158114612bf157600080fd5b50565b600060208284031215612c0657600080fd5b813561227681612be3565b600080600080600060a08688031215612c2957600080fd5b612c32866128b7565b945060208601359350604086013567ffffffffffffffff8082168214612c5757600080fd5b909350606087013590612c6982612be3565b90925060808701359080821115612c7f57600080fd5b50612c8c888289016128db565b9150509295509295909350565b8581528460208201527fffffffffffffffff0000000000000000000000000000000000000000000000008460c01b16604082015282151560f81b604882015260008251612ced816049850160208701612b08565b919091016049019695505050505050565b80516fffffffffffffffffffffffffffffffff81168114611ea557600080fd5b600060608284031215612d3057600080fd5b6040516060810181811067ffffffffffffffff82111715612d5357612d53612888565b60405282518152612d6660208401612cfe565b6020820152612d7760408401612cfe565b60408201529392505050565b600060a08284031215612d9557600080fd5b60405160a0810181811067ffffffffffffffff82111715612db857612db8612888565b806040525082358152602083013560208201526040830135604082015260608301356060820152608083013560808201528091505092915050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8481526000602060808184015280608084015260018060a085015260c0840160c060408601528087825260e08601905060e08860051b87010191508860005b89811015612f25577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2088850301835281357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18c3603018112612edc57600080fd5b8b01868101903567ffffffffffffffff811115612ef857600080fd5b803603821315612f0757600080fd5b612f12868284612df3565b9550505091850191908501908401612e7b565b50505080935050505082606083015295945050505050565b600060208284031215612f4f57600080fd5b815161227681612be3565b60008451612f6c818460208901612b08565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551612fa8816001850160208a01612b08565b60019201918201528351612fc3816002840160208801612b08565b0160020195945050505050565b600060208284031215612fe257600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008261305657613056612fe9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83147f8000000000000000000000000000000000000000000000000000000000000000831416156130aa576130aa613018565b500590565b6000828210156130c1576130c1613018565b500390565b6000808312837f80000000000000000000000000000000000000000000000000000000000000000183128115161561310057613100613018565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01831381161561313457613134613018565b50500390565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60008413600084138583048511828216161561317b5761317b613018565b7f800000000000000000000000000000000000000000000000000000000000000060008712868205881281841616156131b6576131b6613018565b600087129250878205871284841616156131d2576131d2613018565b878505871281841616156131e8576131e8613018565b505050929093029392505050565b6000808212827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0384138115161561323057613230613018565b827f800000000000000000000000000000000000000000000000000000000000000003841281161561326457613264613018565b50500190565b600067ffffffffffffffff80831681851680830382111561328d5761328d613018565b01949350505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156132ce576132ce613018565b500290565b6000826132e2576132e2612fe9565b500490565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a083015261333260c0830184612b38565b98975050505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361336f5761336f613018565b5060010190565b60008261338557613385612fe9565b500690565b6000821982111561339d5761339d613018565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(KromaPortalStorageLayoutJSON), KromaPortalStorageLayout); err != nil {
		panic(err)
	}

	layouts["KromaPortal"] = KromaPortalStorageLayout
	deployedBytecodes["KromaPortal"] = KromaPortalDeployedBin
}
