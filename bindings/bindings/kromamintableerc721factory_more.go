// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/kroma-network/kroma/bindings/solc"
)

const KromaMintableERC721FactoryStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/universal/KromaMintableERC721Factory.sol:KromaMintableERC721Factory\",\"label\":\"isKromaMintableERC721\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"}}}"

var KromaMintableERC721FactoryStorageLayout = new(solc.StorageLayout)

var KromaMintableERC721FactoryDeployedBin = "0x60806040523480156200001157600080fd5b50600436106200006f5760003560e01c8063c8ddda4c1162000056578063c8ddda4c14620000cd578063de679a4e1462000104578063ee9a31a2146200014157600080fd5b806354fd4d5014620000745780637d1d0c5b1462000096575b600080fd5b6200007e62000169565b6040516200008d9190620005db565b60405180910390f35b620000be7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016200008d565b620000f3620000de36600462000621565b60006020819052908152604090205460ff1681565b60405190151581526020016200008d565b6200011b6200011536600462000721565b62000214565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016200008d565b6200011b7f000000000000000000000000000000000000000000000000000000000000000081565b6060620001967f0000000000000000000000000000000000000000000000000000000000000000620003f9565b620001c17f0000000000000000000000000000000000000000000000000000000000000000620003f9565b620001ec7f0000000000000000000000000000000000000000000000000000000000000000620003f9565b60405160200162000200939291906200079e565b604051602081830303815290604052905090565b600073ffffffffffffffffffffffffffffffffffffffff8416620002e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f4b726f6d614d696e7461626c65455243373231466163746f72793a204c31207460448201527f6f6b656e20616464726573732063616e6e6f742062652061646472657373283060648201527f2900000000000000000000000000000000000000000000000000000000000000608482015260a40160405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000086868660405162000339906200054e565b620003499594939291906200081a565b604051809103906000f08015801562000366573d6000803e3d6000fd5b5073ffffffffffffffffffffffffffffffffffffffff8181166000818152602081815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590513381529394509188169290917f0b5bde7eee01f5b9701af85ba2f7e4ddcfe77445964a66b30a8d5d8659667f9f910160405180910390a3949350505050565b6060816000036200043d57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156200046d57806200045481620008aa565b9150620004659050600a8362000914565b915062000441565b60008167ffffffffffffffff8111156200048b576200048b6200063f565b6040519080825280601f01601f191660200182016040528015620004b6576020820181803683370190505b5090505b84156200054657620004ce6001836200092b565b9150620004dd600a8662000945565b620004ea9060306200095c565b60f81b81838151811062000502576200050262000977565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506200053e600a8662000914565b9450620004ba565b949350505050565b612e8280620009a783390190565b60005b83811015620005795781810151838201526020016200055f565b8381111562000589576000848401525b50505050565b60008151808452620005a98160208601602086016200055c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000620005f060208301846200058f565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146200061c57600080fd5b919050565b6000602082840312156200063457600080fd5b620005f082620005f7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126200068057600080fd5b813567ffffffffffffffff808211156200069e576200069e6200063f565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715620006e757620006e76200063f565b816040528381528660208588010111156200070157600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156200073757600080fd5b6200074284620005f7565b9250602084013567ffffffffffffffff808211156200076057600080fd5b6200076e878388016200066e565b935060408601359150808211156200078557600080fd5b5062000794868287016200066e565b9150509250925092565b60008451620007b28184602089016200055c565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551620007f0816001850160208a016200055c565b600192019182015283516200080d8160028401602088016200055c565b0160020195945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835286602084015280861660408401525060a060608301526200085b60a08301856200058f565b82810360808401526200086f81856200058f565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620008de57620008de6200087b565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082620009265762000926620008e5565b500490565b6000828210156200094057620009406200087b565b500390565b600082620009575762000957620008e5565b500690565b600082198211156200097257620009726200087b565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe6101406040523480156200001257600080fd5b5060405162002e8238038062002e82833981016040819052620000359162000633565b60006001818484826200004983826200075c565b5060016200005882826200075c565b50505060809290925260a05260c0526001600160a01b038516620000dc5760405162461bcd60e51b815260206004820152603060248201527f4b726f6d614d696e7461626c654552433732313a206272696467652063616e6e60448201526f6f74206265206164647265737328302960801b60648201526084015b60405180910390fd5b83600003620001545760405162461bcd60e51b815260206004820152603360248201527f4b726f6d614d696e7461626c654552433732313a2072656d6f7465206368616960448201527f6e2069642063616e6e6f74206265207a65726f000000000000000000000000006064820152608401620000d3565b6001600160a01b038316620001d25760405162461bcd60e51b815260206004820152603660248201527f4b726f6d614d696e7461626c654552433732313a2072656d6f746520746f6b6560448201527f6e2063616e6e6f742062652061646472657373283029000000000000000000006064820152608401620000d3565b60e08490526001600160a01b0383811661010081905290861661012052620002089060146200025c602090811b62000df417901c565b6200021e856200041c60201b6200101d1760201c565b6040516020016200023192919062000828565b604051602081830303815290604052600a90816200025091906200075c565b50505050505062000999565b606060006200026d836002620008b2565b6200027a906002620008d4565b6001600160401b0381111562000294576200029462000559565b6040519080825280601f01601f191660200182016040528015620002bf576020820181803683370190505b509050600360fc1b81600081518110620002dd57620002dd620008ef565b60200101906001600160f81b031916908160001a905350600f60fb1b816001815181106200030f576200030f620008ef565b60200101906001600160f81b031916908160001a905350600062000335846002620008b2565b62000342906001620008d4565b90505b6001811115620003c4576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106200037a576200037a620008ef565b1a60f81b828281518110620003935762000393620008ef565b60200101906001600160f81b031916908160001a90535060049490941c93620003bc8162000905565b905062000345565b508315620004155760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401620000d3565b9392505050565b606081600003620004445750506040805180820190915260018152600360fc1b602082015290565b8160005b81156200047457806200045b816200091f565b91506200046c9050600a8362000951565b915062000448565b6000816001600160401b0381111562000491576200049162000559565b6040519080825280601f01601f191660200182016040528015620004bc576020820181803683370190505b5090505b84156200053457620004d460018362000968565b9150620004e3600a8662000982565b620004f0906030620008d4565b60f81b818381518110620005085762000508620008ef565b60200101906001600160f81b031916908160001a9053506200052c600a8662000951565b9450620004c0565b949350505050565b80516001600160a01b03811681146200055457600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b838110156200058c57818101518382015260200162000572565b838111156200059c576000848401525b50505050565b600082601f830112620005b457600080fd5b81516001600160401b0380821115620005d157620005d162000559565b604051601f8301601f19908116603f01168101908282118183101715620005fc57620005fc62000559565b816040528381528660208588010111156200061657600080fd5b620006298460208301602089016200056f565b9695505050505050565b600080600080600060a086880312156200064c57600080fd5b62000657866200053c565b9450602086015193506200066e604087016200053c565b60608701519093506001600160401b03808211156200068c57600080fd5b6200069a89838a01620005a2565b93506080880151915080821115620006b157600080fd5b50620006c088828901620005a2565b9150509295509295909350565b600181811c90821680620006e257607f821691505b6020821081036200070357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200075757600081815260208120601f850160051c81016020861015620007325750805b601f850160051c820191505b8181101562000753578281556001016200073e565b5050505b505050565b81516001600160401b0381111562000778576200077862000559565b6200079081620007898454620006cd565b8462000709565b602080601f831160018114620007c85760008415620007af5750858301515b600019600386901b1c1916600185901b17855562000753565b600085815260208120601f198616915b82811015620007f957888601518255948401946001909101908401620007d8565b5085821015620008185787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6832ba3432b932bab69d60b91b8152600083516200084e8160098501602088016200056f565b600160fe1b60099184019182015283516200087181600a8401602088016200056f565b712f746f6b656e5552493f75696e743235363d60701b600a9290910191820152601c01949350505050565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615620008cf57620008cf6200089c565b500290565b60008219821115620008ea57620008ea6200089c565b500190565b634e487b7160e01b600052603260045260246000fd5b6000816200091757620009176200089c565b506000190190565b6000600182016200093457620009346200089c565b5060010190565b634e487b7160e01b600052601260045260246000fd5b6000826200096357620009636200093b565b500490565b6000828210156200097d576200097d6200089c565b500390565b6000826200099457620009946200093b565b500690565b60805160a05160c05160e051610100516101205161247662000a0c600039600081816103a10152818161042c01528181610a9c0152610b8a0152600081816101e0015261037b0152600081816102e801526103c7015260006109330152600061090a015260006108e101526124766000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c80637d1d0c5b116100ee578063c87b56dd11610097578063e78cea9211610071578063e78cea921461039f578063e9518196146103c5578063e985e9c5146103eb578063ee9a31a21461042757600080fd5b8063c87b56dd1461035e578063d547cfb714610371578063d6c0b2c41461037957600080fd5b8063a1448194116100c8578063a144819414610325578063a22cb46514610338578063b88d4fde1461034b57600080fd5b80637d1d0c5b146102e357806395d89b411461030a5780639dc29fac1461031257600080fd5b806323b872dd1161015b5780634f6ccce7116101355780634f6ccce7146102a257806354fd4d50146102b55780636352211e146102bd57806370a08231146102d057600080fd5b806323b872dd146102695780632f745c591461027c57806342842e0e1461028f57600080fd5b8063081812fc1161018c578063081812fc1461022f578063095ea7b31461024257806318160ddd1461025757600080fd5b806301ffc9a7146101b3578063033964be146101db57806306fdde031461021a575b600080fd5b6101c66101c1366004611e67565b61044e565b60405190151581526020015b60405180910390f35b6102027f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101d2565b6102226104fd565b6040516101d29190611efa565b61020261023d366004611f0d565b61058f565b610255610250366004611f42565b6105b6565b005b6008545b6040519081526020016101d2565b610255610277366004611f6c565b6106ec565b61025b61028a366004611f42565b610773565b61025561029d366004611f6c565b61081b565b61025b6102b0366004611f0d565b610836565b6102226108da565b6102026102cb366004611f0d565b61097d565b61025b6102de366004611fa8565b6109e8565b61025b7f000000000000000000000000000000000000000000000000000000000000000081565b610222610a82565b610255610320366004611f42565b610a91565b610255610333366004611f42565b610b7f565b610255610346366004611fc3565b610c62565b61025561035936600461202e565b610c71565b61022261036c366004611f0d565b610cff565b610222610d66565b7f0000000000000000000000000000000000000000000000000000000000000000610202565b7f0000000000000000000000000000000000000000000000000000000000000000610202565b7f000000000000000000000000000000000000000000000000000000000000000061025b565b6101c66103f9366004612128565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6102027f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f74259ebf000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000084168214806104e657507fffffffff00000000000000000000000000000000000000000000000000000000848116908216145b806104f557506104f584611152565b949350505050565b60606000805461050c9061215b565b80601f01602080910402602001604051908101604052809291908181526020018280546105389061215b565b80156105855780601f1061055a57610100808354040283529160200191610585565b820191906000526020600020905b81548152906001019060200180831161056857829003601f168201915b5050505050905090565b600061059a826111a8565b506000908152600460205260409020546001600160a01b031690565b60006105c18261097d565b9050806001600160a01b0316836001600160a01b03160361064f5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560448201527f720000000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b336001600160a01b038216148061066b575061066b81336103f9565b6106dd5760405162461bcd60e51b815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c00006064820152608401610646565b6106e7838361120f565b505050565b6106f63382611295565b6107685760405162461bcd60e51b815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f7665640000000000000000000000000000000000006064820152608401610646565b6106e7838383611313565b600061077e836109e8565b82106107f25760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201527f74206f6620626f756e64730000000000000000000000000000000000000000006064820152608401610646565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b6106e783838360405180602001604052806000815250610c71565b600061084160085490565b82106108b55760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201527f7574206f6620626f756e647300000000000000000000000000000000000000006064820152608401610646565b600882815481106108c8576108c86121ae565b90600052602060002001549050919050565b60606109057f000000000000000000000000000000000000000000000000000000000000000061101d565b61092e7f000000000000000000000000000000000000000000000000000000000000000061101d565b6109577f000000000000000000000000000000000000000000000000000000000000000061101d565b604051602001610969939291906121dd565b604051602081830303815290604052905090565b6000818152600260205260408120546001600160a01b0316806109e25760405162461bcd60e51b815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152606401610646565b92915050565b60006001600160a01b038216610a665760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f74206120766160448201527f6c6964206f776e657200000000000000000000000000000000000000000000006064820152608401610646565b506001600160a01b031660009081526003602052604090205490565b60606001805461050c9061215b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610b2f5760405162461bcd60e51b815260206004820152603760248201527f4b726f6d614d696e7461626c654552433732313a206f6e6c792062726964676560448201527f2063616e2063616c6c20746869732066756e6374696f6e0000000000000000006064820152608401610646565b610b3881611503565b816001600160a01b03167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca582604051610b7391815260200190565b60405180910390a25050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c1d5760405162461bcd60e51b815260206004820152603760248201527f4b726f6d614d696e7461626c654552433732313a206f6e6c792062726964676560448201527f2063616e2063616c6c20746869732066756e6374696f6e0000000000000000006064820152608401610646565b610c2782826115c2565b816001600160a01b03167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688582604051610b7391815260200190565b610c6d3383836115dc565b5050565b610c7b3383611295565b610ced5760405162461bcd60e51b815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f7665640000000000000000000000000000000000006064820152608401610646565b610cf9848484846116c8565b50505050565b6060610d0a826111a8565b6000610d14611751565b90506000815111610d345760405180602001604052806000815250610d5f565b80610d3e8461101d565b604051602001610d4f929190612253565b6040516020818303038152906040525b9392505050565b600a8054610d739061215b565b80601f0160208091040260200160405190810160405280929190818152602001828054610d9f9061215b565b8015610dec5780601f10610dc157610100808354040283529160200191610dec565b820191906000526020600020905b815481529060010190602001808311610dcf57829003601f168201915b505050505081565b60606000610e038360026122b1565b610e0e9060026122ee565b67ffffffffffffffff811115610e2657610e26611fff565b6040519080825280601f01601f191660200182016040528015610e50576020820181803683370190505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110610e8757610e876121ae565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610eea57610eea6121ae565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000610f268460026122b1565b610f319060016122ee565b90505b6001811115610fce577f303132333435363738396162636465660000000000000000000000000000000085600f1660108110610f7257610f726121ae565b1a60f81b828281518110610f8857610f886121ae565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c93610fc781612306565b9050610f34565b508315610d5f5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610646565b60608160000361106057505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561108a57806110748161233b565b91506110839050600a836123a2565b9150611064565b60008167ffffffffffffffff8111156110a5576110a5611fff565b6040519080825280601f01601f1916602001820160405280156110cf576020820181803683370190505b5090505b84156104f5576110e46001836123b6565b91506110f1600a866123cd565b6110fc9060306122ee565b60f81b818381518110611111576111116121ae565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061114b600a866123a2565b94506110d3565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f780e9d630000000000000000000000000000000000000000000000000000000014806109e257506109e282611760565b6000818152600260205260409020546001600160a01b031661120c5760405162461bcd60e51b815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152606401610646565b50565b600081815260046020526040902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038416908117909155819061125c8261097d565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000806112a18361097d565b9050806001600160a01b0316846001600160a01b031614806112e857506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b806104f55750836001600160a01b03166113018461058f565b6001600160a01b031614949350505050565b826001600160a01b03166113268261097d565b6001600160a01b0316146113a25760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201527f6f776e65720000000000000000000000000000000000000000000000000000006064820152608401610646565b6001600160a01b03821661141d5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610646565b611428838383611843565b61143360008261120f565b6001600160a01b038316600090815260036020526040812080546001929061145c9084906123b6565b90915550506001600160a01b038216600090815260036020526040812080546001929061148a9084906122ee565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b600061150e8261097d565b905061151c81600084611843565b61152760008361120f565b6001600160a01b03811660009081526003602052604081208054600192906115509084906123b6565b909155505060008281526002602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055518391906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b610c6d8282604051806020016040528060008152506118fb565b816001600160a01b0316836001600160a01b03160361163d5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610646565b6001600160a01b0383811660008181526005602090815260408083209487168084529482529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6116d3848484611313565b6116df84848484611984565b610cf95760405162461bcd60e51b815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608401610646565b6060600a805461050c9061215b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f80ac58cd0000000000000000000000000000000000000000000000000000000014806117f357507fffffffff0000000000000000000000000000000000000000000000000000000082167f5b5e139f00000000000000000000000000000000000000000000000000000000145b806109e257507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146109e2565b6001600160a01b03831661189e5761189981600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b6118c1565b816001600160a01b0316836001600160a01b0316146118c1576118c18382611b43565b6001600160a01b0382166118d8576106e781611be0565b826001600160a01b0316826001600160a01b0316146106e7576106e78282611c8f565b6119058383611cd3565b6119126000848484611984565b6106e75760405162461bcd60e51b815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608401610646565b60006001600160a01b0384163b15611b38576040517f150b7a020000000000000000000000000000000000000000000000000000000081526001600160a01b0385169063150b7a02906119e19033908990889088906004016123e1565b6020604051808303816000875af1925050508015611a3a575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611a379181019061241d565b60015b611aed573d808015611a68576040519150601f19603f3d011682016040523d82523d6000602084013e611a6d565b606091505b508051600003611ae55760405162461bcd60e51b815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608401610646565b805181602001fd5b7fffffffff00000000000000000000000000000000000000000000000000000000167f150b7a02000000000000000000000000000000000000000000000000000000001490506104f5565b506001949350505050565b60006001611b50846109e8565b611b5a91906123b6565b600083815260076020526040902054909150808214611bad576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b600854600090611bf2906001906123b6565b60008381526009602052604081205460088054939450909284908110611c1a57611c1a6121ae565b906000526020600020015490508060088381548110611c3b57611c3b6121ae565b6000918252602080832090910192909255828152600990915260408082208490558582528120556008805480611c7357611c7361243a565b6001900381819060005260206000200160009055905550505050565b6000611c9a836109e8565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b6001600160a01b038216611d295760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610646565b6000818152600260205260409020546001600160a01b031615611d8e5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610646565b611d9a60008383611843565b6001600160a01b0382166000908152600360205260408120805460019290611dc39084906122ee565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461120c57600080fd5b600060208284031215611e7957600080fd5b8135610d5f81611e39565b60005b83811015611e9f578181015183820152602001611e87565b83811115610cf95750506000910152565b60008151808452611ec8816020860160208601611e84565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610d5f6020830184611eb0565b600060208284031215611f1f57600080fd5b5035919050565b80356001600160a01b0381168114611f3d57600080fd5b919050565b60008060408385031215611f5557600080fd5b611f5e83611f26565b946020939093013593505050565b600080600060608486031215611f8157600080fd5b611f8a84611f26565b9250611f9860208501611f26565b9150604084013590509250925092565b600060208284031215611fba57600080fd5b610d5f82611f26565b60008060408385031215611fd657600080fd5b611fdf83611f26565b915060208301358015158114611ff457600080fd5b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000806080858703121561204457600080fd5b61204d85611f26565b935061205b60208601611f26565b925060408501359150606085013567ffffffffffffffff8082111561207f57600080fd5b818701915087601f83011261209357600080fd5b8135818111156120a5576120a5611fff565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156120eb576120eb611fff565b816040528281528a602084870101111561210457600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b6000806040838503121561213b57600080fd5b61214483611f26565b915061215260208401611f26565b90509250929050565b600181811c9082168061216f57607f821691505b6020821081036121a8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600084516121ef818460208901611e84565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161222b816001850160208a01611e84565b60019201918201528351612246816002840160208801611e84565b0160020195945050505050565b60008351612265818460208801611e84565b835190830190612279818360208801611e84565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156122e9576122e9612282565b500290565b6000821982111561230157612301612282565b500190565b60008161231557612315612282565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361236c5761236c612282565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826123b1576123b1612373565b500490565b6000828210156123c8576123c8612282565b500390565b6000826123dc576123dc612373565b500690565b60006001600160a01b038087168352808616602084015250836040830152608060608301526124136080830184611eb0565b9695505050505050565b60006020828403121561242f57600080fd5b8151610d5f81611e39565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c634300080f000aa164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(KromaMintableERC721FactoryStorageLayoutJSON), KromaMintableERC721FactoryStorageLayout); err != nil {
		panic(err)
	}

	layouts["KromaMintableERC721Factory"] = KromaMintableERC721FactoryStorageLayout
	deployedBytecodes["KromaMintableERC721Factory"] = KromaMintableERC721FactoryDeployedBin
}
