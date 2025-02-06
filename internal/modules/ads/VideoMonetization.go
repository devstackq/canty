// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ads

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AdsMetaData contains all meta data concerning the Ads contract.
var AdsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"adText\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"adImage\",\"type\":\"string\"}],\"name\":\"AdPlaced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"adText\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"adImage\",\"type\":\"string\"}],\"name\":\"placeAd\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"updateAdPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052670de0b6b3a7640000600155348015601a575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506106d5806100675f395ff3fe60806040526004361061003e575f3560e01c80634fb00867146100425780638da5cb5b1461005e578063b52f824514610088578063d5e10d35146100b2575b5f5ffd5b61005c600480360381019061005791906103e6565b6100da565b005b348015610069575f5ffd5b506100726101d7565b60405161007f919061049b565b60405180910390f35b348015610093575f5ffd5b5061009c6101fb565b6040516100a991906104cc565b60405180910390f35b3480156100bd575f5ffd5b506100d860048036038101906100d3919061050f565b610201565b005b60015434101561011f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161011690610594565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167ff444b15df736f5cd73459aeaf1d94bd9ede5e1699f0758989638b875347b1da38383604051610167929190610602565b60405180910390a25f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3490811502906040515f60405180830381858888f193505050501580156101d2573d5f5f3e3d5ffd5b505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461028f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028690610681565b60405180910390fd5b8060018190555050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6102f8826102b2565b810181811067ffffffffffffffff82111715610317576103166102c2565b5b80604052505050565b5f610329610299565b905061033582826102ef565b919050565b5f67ffffffffffffffff821115610354576103536102c2565b5b61035d826102b2565b9050602081019050919050565b828183375f83830152505050565b5f61038a6103858461033a565b610320565b9050828152602081018484840111156103a6576103a56102ae565b5b6103b184828561036a565b509392505050565b5f82601f8301126103cd576103cc6102aa565b5b81356103dd848260208601610378565b91505092915050565b5f5f604083850312156103fc576103fb6102a2565b5b5f83013567ffffffffffffffff811115610419576104186102a6565b5b610425858286016103b9565b925050602083013567ffffffffffffffff811115610446576104456102a6565b5b610452858286016103b9565b9150509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104858261045c565b9050919050565b6104958161047b565b82525050565b5f6020820190506104ae5f83018461048c565b92915050565b5f819050919050565b6104c6816104b4565b82525050565b5f6020820190506104df5f8301846104bd565b92915050565b6104ee816104b4565b81146104f8575f5ffd5b50565b5f81359050610509816104e5565b92915050565b5f60208284031215610524576105236102a2565b5b5f610531848285016104fb565b91505092915050565b5f82825260208201905092915050565b7f496e73756666696369656e74207061796d656e740000000000000000000000005f82015250565b5f61057e60148361053a565b91506105898261054a565b602082019050919050565b5f6020820190508181035f8301526105ab81610572565b9050919050565b5f81519050919050565b8281835e5f83830152505050565b5f6105d4826105b2565b6105de818561053a565b93506105ee8185602086016105bc565b6105f7816102b2565b840191505092915050565b5f6040820190508181035f83015261061a81856105ca565b9050818103602083015261062e81846105ca565b90509392505050565b7f4f6e6c79206f776e65722063616e2075706461746520616420707269636500005f82015250565b5f61066b601e8361053a565b915061067682610637565b602082019050919050565b5f6020820190508181035f8301526106988161065f565b905091905056fea26469706673582212206f04b92beb77de756ddfcaef2918fd702cec25b91cc5124e12706e2588645cfc64736f6c634300081c0033",
}

// AdsABI is the input ABI used to generate the binding from.
// Deprecated: Use AdsMetaData.ABI instead.
var AdsABI = AdsMetaData.ABI

// AdsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AdsMetaData.Bin instead.
var AdsBin = AdsMetaData.Bin

// DeployAds deploys a new Ethereum contract, binding an instance of Ads to it.
func DeployAds(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ads, error) {
	parsed, err := AdsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ads{AdsCaller: AdsCaller{contract: contract}, AdsTransactor: AdsTransactor{contract: contract}, AdsFilterer: AdsFilterer{contract: contract}}, nil
}

// Ads is an auto generated Go binding around an Ethereum contract.
type Ads struct {
	AdsCaller     // Read-only binding to the contract
	AdsTransactor // Write-only binding to the contract
	AdsFilterer   // Log filterer for contract events
}

// AdsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdsSession struct {
	Contract     *Ads              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdsCallerSession struct {
	Contract *AdsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AdsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdsTransactorSession struct {
	Contract     *AdsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdsRaw struct {
	Contract *Ads // Generic contract binding to access the raw methods on
}

// AdsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdsCallerRaw struct {
	Contract *AdsCaller // Generic read-only contract binding to access the raw methods on
}

// AdsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdsTransactorRaw struct {
	Contract *AdsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAds creates a new instance of Ads, bound to a specific deployed contract.
func NewAds(address common.Address, backend bind.ContractBackend) (*Ads, error) {
	contract, err := bindAds(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ads{AdsCaller: AdsCaller{contract: contract}, AdsTransactor: AdsTransactor{contract: contract}, AdsFilterer: AdsFilterer{contract: contract}}, nil
}

// NewAdsCaller creates a new read-only instance of Ads, bound to a specific deployed contract.
func NewAdsCaller(address common.Address, caller bind.ContractCaller) (*AdsCaller, error) {
	contract, err := bindAds(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdsCaller{contract: contract}, nil
}

// NewAdsTransactor creates a new write-only instance of Ads, bound to a specific deployed contract.
func NewAdsTransactor(address common.Address, transactor bind.ContractTransactor) (*AdsTransactor, error) {
	contract, err := bindAds(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdsTransactor{contract: contract}, nil
}

// NewAdsFilterer creates a new log filterer instance of Ads, bound to a specific deployed contract.
func NewAdsFilterer(address common.Address, filterer bind.ContractFilterer) (*AdsFilterer, error) {
	contract, err := bindAds(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdsFilterer{contract: contract}, nil
}

// bindAds binds a generic wrapper to an already deployed contract.
func bindAds(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AdsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ads *AdsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ads.Contract.AdsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ads *AdsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ads.Contract.AdsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ads *AdsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ads.Contract.AdsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ads *AdsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ads.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ads *AdsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ads.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ads *AdsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ads.Contract.contract.Transact(opts, method, params...)
}

// AdPrice is a free data retrieval call binding the contract method 0xb52f8245.
//
// Solidity: function adPrice() view returns(uint256)
func (_Ads *AdsCaller) AdPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ads.contract.Call(opts, &out, "adPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdPrice is a free data retrieval call binding the contract method 0xb52f8245.
//
// Solidity: function adPrice() view returns(uint256)
func (_Ads *AdsSession) AdPrice() (*big.Int, error) {
	return _Ads.Contract.AdPrice(&_Ads.CallOpts)
}

// AdPrice is a free data retrieval call binding the contract method 0xb52f8245.
//
// Solidity: function adPrice() view returns(uint256)
func (_Ads *AdsCallerSession) AdPrice() (*big.Int, error) {
	return _Ads.Contract.AdPrice(&_Ads.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ads *AdsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ads.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ads *AdsSession) Owner() (common.Address, error) {
	return _Ads.Contract.Owner(&_Ads.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ads *AdsCallerSession) Owner() (common.Address, error) {
	return _Ads.Contract.Owner(&_Ads.CallOpts)
}

// PlaceAd is a paid mutator transaction binding the contract method 0x4fb00867.
//
// Solidity: function placeAd(string adText, string adImage) payable returns()
func (_Ads *AdsTransactor) PlaceAd(opts *bind.TransactOpts, adText string, adImage string) (*types.Transaction, error) {
	return _Ads.contract.Transact(opts, "placeAd", adText, adImage)
}

// PlaceAd is a paid mutator transaction binding the contract method 0x4fb00867.
//
// Solidity: function placeAd(string adText, string adImage) payable returns()
func (_Ads *AdsSession) PlaceAd(adText string, adImage string) (*types.Transaction, error) {
	return _Ads.Contract.PlaceAd(&_Ads.TransactOpts, adText, adImage)
}

// PlaceAd is a paid mutator transaction binding the contract method 0x4fb00867.
//
// Solidity: function placeAd(string adText, string adImage) payable returns()
func (_Ads *AdsTransactorSession) PlaceAd(adText string, adImage string) (*types.Transaction, error) {
	return _Ads.Contract.PlaceAd(&_Ads.TransactOpts, adText, adImage)
}

// UpdateAdPrice is a paid mutator transaction binding the contract method 0xd5e10d35.
//
// Solidity: function updateAdPrice(uint256 newPrice) returns()
func (_Ads *AdsTransactor) UpdateAdPrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _Ads.contract.Transact(opts, "updateAdPrice", newPrice)
}

// UpdateAdPrice is a paid mutator transaction binding the contract method 0xd5e10d35.
//
// Solidity: function updateAdPrice(uint256 newPrice) returns()
func (_Ads *AdsSession) UpdateAdPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Ads.Contract.UpdateAdPrice(&_Ads.TransactOpts, newPrice)
}

// UpdateAdPrice is a paid mutator transaction binding the contract method 0xd5e10d35.
//
// Solidity: function updateAdPrice(uint256 newPrice) returns()
func (_Ads *AdsTransactorSession) UpdateAdPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Ads.Contract.UpdateAdPrice(&_Ads.TransactOpts, newPrice)
}

// AdsAdPlacedIterator is returned from FilterAdPlaced and is used to iterate over the raw logs and unpacked data for AdPlaced events raised by the Ads contract.
type AdsAdPlacedIterator struct {
	Event *AdsAdPlaced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdsAdPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdsAdPlaced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdsAdPlaced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdsAdPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdsAdPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdsAdPlaced represents a AdPlaced event raised by the Ads contract.
type AdsAdPlaced struct {
	From    common.Address
	AdText  string
	AdImage string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdPlaced is a free log retrieval operation binding the contract event 0xf444b15df736f5cd73459aeaf1d94bd9ede5e1699f0758989638b875347b1da3.
//
// Solidity: event AdPlaced(address indexed from, string adText, string adImage)
func (_Ads *AdsFilterer) FilterAdPlaced(opts *bind.FilterOpts, from []common.Address) (*AdsAdPlacedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Ads.contract.FilterLogs(opts, "AdPlaced", fromRule)
	if err != nil {
		return nil, err
	}
	return &AdsAdPlacedIterator{contract: _Ads.contract, event: "AdPlaced", logs: logs, sub: sub}, nil
}

// WatchAdPlaced is a free log subscription operation binding the contract event 0xf444b15df736f5cd73459aeaf1d94bd9ede5e1699f0758989638b875347b1da3.
//
// Solidity: event AdPlaced(address indexed from, string adText, string adImage)
func (_Ads *AdsFilterer) WatchAdPlaced(opts *bind.WatchOpts, sink chan<- *AdsAdPlaced, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Ads.contract.WatchLogs(opts, "AdPlaced", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdsAdPlaced)
				if err := _Ads.contract.UnpackLog(event, "AdPlaced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdPlaced is a log parse operation binding the contract event 0xf444b15df736f5cd73459aeaf1d94bd9ede5e1699f0758989638b875347b1da3.
//
// Solidity: event AdPlaced(address indexed from, string adText, string adImage)
func (_Ads *AdsFilterer) ParseAdPlaced(log types.Log) (*AdsAdPlaced, error) {
	event := new(AdsAdPlaced)
	if err := _Ads.contract.UnpackLog(event, "AdPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
