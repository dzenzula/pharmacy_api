// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ShipmentTrackerShipment is an auto generated low-level Go binding around an user-defined struct.
type ShipmentTrackerShipment struct {
	BatchId   string
	FromCity  string
	ToCity    string
	Status    string
	Timestamp string
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fromCity\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toCity\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"}],\"name\":\"ShipmentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"}],\"name\":\"ShipmentUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fromCity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toCity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"}],\"name\":\"addShipment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchId\",\"type\":\"string\"}],\"name\":\"getShipmentHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"batchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fromCity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toCity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"}],\"internalType\":\"structShipmentTracker.Shipment[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shipments\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"batchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fromCity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toCity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"}],\"name\":\"updateShipment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetShipmentHistory is a free data retrieval call binding the contract method 0x03e39f0a.
//
// Solidity: function getShipmentHistory(string batchId) view returns((string,string,string,string,string)[])
func (_Contracts *ContractsCaller) GetShipmentHistory(opts *bind.CallOpts, batchId string) ([]ShipmentTrackerShipment, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getShipmentHistory", batchId)

	if err != nil {
		return *new([]ShipmentTrackerShipment), err
	}

	out0 := *abi.ConvertType(out[0], new([]ShipmentTrackerShipment)).(*[]ShipmentTrackerShipment)

	return out0, err

}

// GetShipmentHistory is a free data retrieval call binding the contract method 0x03e39f0a.
//
// Solidity: function getShipmentHistory(string batchId) view returns((string,string,string,string,string)[])
func (_Contracts *ContractsSession) GetShipmentHistory(batchId string) ([]ShipmentTrackerShipment, error) {
	return _Contracts.Contract.GetShipmentHistory(&_Contracts.CallOpts, batchId)
}

// GetShipmentHistory is a free data retrieval call binding the contract method 0x03e39f0a.
//
// Solidity: function getShipmentHistory(string batchId) view returns((string,string,string,string,string)[])
func (_Contracts *ContractsCallerSession) GetShipmentHistory(batchId string) ([]ShipmentTrackerShipment, error) {
	return _Contracts.Contract.GetShipmentHistory(&_Contracts.CallOpts, batchId)
}

// Shipments is a free data retrieval call binding the contract method 0xab265fd2.
//
// Solidity: function shipments(string , uint256 ) view returns(string batchId, string fromCity, string toCity, string status, string timestamp)
func (_Contracts *ContractsCaller) Shipments(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	BatchId   string
	FromCity  string
	ToCity    string
	Status    string
	Timestamp string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "shipments", arg0, arg1)

	outstruct := new(struct {
		BatchId   string
		FromCity  string
		ToCity    string
		Status    string
		Timestamp string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.FromCity = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ToCity = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// Shipments is a free data retrieval call binding the contract method 0xab265fd2.
//
// Solidity: function shipments(string , uint256 ) view returns(string batchId, string fromCity, string toCity, string status, string timestamp)
func (_Contracts *ContractsSession) Shipments(arg0 string, arg1 *big.Int) (struct {
	BatchId   string
	FromCity  string
	ToCity    string
	Status    string
	Timestamp string
}, error) {
	return _Contracts.Contract.Shipments(&_Contracts.CallOpts, arg0, arg1)
}

// Shipments is a free data retrieval call binding the contract method 0xab265fd2.
//
// Solidity: function shipments(string , uint256 ) view returns(string batchId, string fromCity, string toCity, string status, string timestamp)
func (_Contracts *ContractsCallerSession) Shipments(arg0 string, arg1 *big.Int) (struct {
	BatchId   string
	FromCity  string
	ToCity    string
	Status    string
	Timestamp string
}, error) {
	return _Contracts.Contract.Shipments(&_Contracts.CallOpts, arg0, arg1)
}

// AddShipment is a paid mutator transaction binding the contract method 0x50b09f33.
//
// Solidity: function addShipment(string batchId, string fromCity, string toCity, string status, string timestamp) returns()
func (_Contracts *ContractsTransactor) AddShipment(opts *bind.TransactOpts, batchId string, fromCity string, toCity string, status string, timestamp string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addShipment", batchId, fromCity, toCity, status, timestamp)
}

// AddShipment is a paid mutator transaction binding the contract method 0x50b09f33.
//
// Solidity: function addShipment(string batchId, string fromCity, string toCity, string status, string timestamp) returns()
func (_Contracts *ContractsSession) AddShipment(batchId string, fromCity string, toCity string, status string, timestamp string) (*types.Transaction, error) {
	return _Contracts.Contract.AddShipment(&_Contracts.TransactOpts, batchId, fromCity, toCity, status, timestamp)
}

// AddShipment is a paid mutator transaction binding the contract method 0x50b09f33.
//
// Solidity: function addShipment(string batchId, string fromCity, string toCity, string status, string timestamp) returns()
func (_Contracts *ContractsTransactorSession) AddShipment(batchId string, fromCity string, toCity string, status string, timestamp string) (*types.Transaction, error) {
	return _Contracts.Contract.AddShipment(&_Contracts.TransactOpts, batchId, fromCity, toCity, status, timestamp)
}

// UpdateShipment is a paid mutator transaction binding the contract method 0x10bdab13.
//
// Solidity: function updateShipment(string batchId, string status, string timestamp) returns()
func (_Contracts *ContractsTransactor) UpdateShipment(opts *bind.TransactOpts, batchId string, status string, timestamp string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateShipment", batchId, status, timestamp)
}

// UpdateShipment is a paid mutator transaction binding the contract method 0x10bdab13.
//
// Solidity: function updateShipment(string batchId, string status, string timestamp) returns()
func (_Contracts *ContractsSession) UpdateShipment(batchId string, status string, timestamp string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateShipment(&_Contracts.TransactOpts, batchId, status, timestamp)
}

// UpdateShipment is a paid mutator transaction binding the contract method 0x10bdab13.
//
// Solidity: function updateShipment(string batchId, string status, string timestamp) returns()
func (_Contracts *ContractsTransactorSession) UpdateShipment(batchId string, status string, timestamp string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateShipment(&_Contracts.TransactOpts, batchId, status, timestamp)
}

// ContractsShipmentAddedIterator is returned from FilterShipmentAdded and is used to iterate over the raw logs and unpacked data for ShipmentAdded events raised by the Contracts contract.
type ContractsShipmentAddedIterator struct {
	Event *ContractsShipmentAdded // Event containing the contract specifics and raw log

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
func (it *ContractsShipmentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsShipmentAdded)
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
		it.Event = new(ContractsShipmentAdded)
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
func (it *ContractsShipmentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsShipmentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsShipmentAdded represents a ShipmentAdded event raised by the Contracts contract.
type ContractsShipmentAdded struct {
	BatchId   string
	FromCity  string
	ToCity    string
	Status    string
	Timestamp string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterShipmentAdded is a free log retrieval operation binding the contract event 0xc745a2b15644e3f90e12f77cc9de662d7076e1ef22416041ecc558d37f59bcfa.
//
// Solidity: event ShipmentAdded(string batchId, string fromCity, string toCity, string status, string timestamp)
func (_Contracts *ContractsFilterer) FilterShipmentAdded(opts *bind.FilterOpts) (*ContractsShipmentAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ShipmentAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsShipmentAddedIterator{contract: _Contracts.contract, event: "ShipmentAdded", logs: logs, sub: sub}, nil
}

// WatchShipmentAdded is a free log subscription operation binding the contract event 0xc745a2b15644e3f90e12f77cc9de662d7076e1ef22416041ecc558d37f59bcfa.
//
// Solidity: event ShipmentAdded(string batchId, string fromCity, string toCity, string status, string timestamp)
func (_Contracts *ContractsFilterer) WatchShipmentAdded(opts *bind.WatchOpts, sink chan<- *ContractsShipmentAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ShipmentAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsShipmentAdded)
				if err := _Contracts.contract.UnpackLog(event, "ShipmentAdded", log); err != nil {
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

// ParseShipmentAdded is a log parse operation binding the contract event 0xc745a2b15644e3f90e12f77cc9de662d7076e1ef22416041ecc558d37f59bcfa.
//
// Solidity: event ShipmentAdded(string batchId, string fromCity, string toCity, string status, string timestamp)
func (_Contracts *ContractsFilterer) ParseShipmentAdded(log types.Log) (*ContractsShipmentAdded, error) {
	event := new(ContractsShipmentAdded)
	if err := _Contracts.contract.UnpackLog(event, "ShipmentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsShipmentUpdatedIterator is returned from FilterShipmentUpdated and is used to iterate over the raw logs and unpacked data for ShipmentUpdated events raised by the Contracts contract.
type ContractsShipmentUpdatedIterator struct {
	Event *ContractsShipmentUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsShipmentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsShipmentUpdated)
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
		it.Event = new(ContractsShipmentUpdated)
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
func (it *ContractsShipmentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsShipmentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsShipmentUpdated represents a ShipmentUpdated event raised by the Contracts contract.
type ContractsShipmentUpdated struct {
	BatchId   string
	Status    string
	Timestamp string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterShipmentUpdated is a free log retrieval operation binding the contract event 0x1a685dec5d8e809f1274a44af8e424d4950e7217303920f760c964949c7c0858.
//
// Solidity: event ShipmentUpdated(string batchId, string status, string timestamp)
func (_Contracts *ContractsFilterer) FilterShipmentUpdated(opts *bind.FilterOpts) (*ContractsShipmentUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ShipmentUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsShipmentUpdatedIterator{contract: _Contracts.contract, event: "ShipmentUpdated", logs: logs, sub: sub}, nil
}

// WatchShipmentUpdated is a free log subscription operation binding the contract event 0x1a685dec5d8e809f1274a44af8e424d4950e7217303920f760c964949c7c0858.
//
// Solidity: event ShipmentUpdated(string batchId, string status, string timestamp)
func (_Contracts *ContractsFilterer) WatchShipmentUpdated(opts *bind.WatchOpts, sink chan<- *ContractsShipmentUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ShipmentUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsShipmentUpdated)
				if err := _Contracts.contract.UnpackLog(event, "ShipmentUpdated", log); err != nil {
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

// ParseShipmentUpdated is a log parse operation binding the contract event 0x1a685dec5d8e809f1274a44af8e424d4950e7217303920f760c964949c7c0858.
//
// Solidity: event ShipmentUpdated(string batchId, string status, string timestamp)
func (_Contracts *ContractsFilterer) ParseShipmentUpdated(log types.Log) (*ContractsShipmentUpdated, error) {
	event := new(ContractsShipmentUpdated)
	if err := _Contracts.contract.UnpackLog(event, "ShipmentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
