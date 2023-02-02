package vmhooks

// // Declare the function signatures (see [cgo](https://golang.org/cmd/cgo/)).
//
// #include <stdlib.h>
// typedef unsigned char uint8_t;
// typedef int int32_t;
//
// extern void			v1_3_getSCAddress(void *context, int32_t resultOffset);
// extern void			v1_3_getOwnerAddress(void *context, int32_t resultOffset);
// extern int32_t		v1_3_getShardOfAddress(void *context, int32_t addressOffset);
// extern int32_t		v1_3_isSmartContract(void *context, int32_t addressOffset);
// extern void			v1_3_getExternalBalance(void *context, int32_t addressOffset, int32_t resultOffset);
// extern int32_t		v1_3_blockHash(void *context, long long nonce, int32_t resultOffset);
// extern int32_t		v1_3_transferValue(void *context, int32_t dstOffset, int32_t valueOffset, int32_t dataOffset, int32_t length);
// extern int32_t		v1_3_transferESDT(void *context, int32_t dstOffset, int32_t tokenIDOffset, int32_t tokenIdLen, int32_t valueOffset, long long gasLimit, int32_t dataOffset, int32_t length);
// extern int32_t		v1_3_transferESDTExecute(void *context, int32_t dstOffset, int32_t tokenIDOffset, int32_t tokenIdLen, int32_t valueOffset, long long gasLimit, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t		v1_3_transferESDTNFTExecute(void *context, int32_t dstOffset, int32_t tokenIDOffset, int32_t tokenIdLen, int32_t valueOffset, long long nonce, long long gasLimit, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t		v1_3_transferValueExecute(void *context, int32_t dstOffset, int32_t valueOffset, long long gasLimit, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t		v1_3_getArgumentLength(void *context, int32_t id);
// extern int32_t		v1_3_getArgument(void *context, int32_t id, int32_t argOffset);
// extern int32_t		v1_3_getFunction(void *context, int32_t functionOffset);
// extern int32_t		v1_3_getNumArguments(void *context);
// extern int32_t		v1_3_storageStore(void *context, int32_t keyOffset, int32_t keyLength , int32_t dataOffset, int32_t dataLength);
// extern int32_t		v1_3_storageLoadLength(void *context, int32_t keyOffset, int32_t keyLength );
// extern int32_t		v1_3_storageLoad(void *context, int32_t keyOffset, int32_t keyLength , int32_t dataOffset);
// extern int32_t		v1_3_storageLoadFromAddress(void *context, int32_t addressOffset, int32_t keyOffset, int32_t keyLength , int32_t dataOffset);
// extern void			v1_3_getCaller(void *context, int32_t resultOffset);
// extern void			v1_3_checkNoPayment(void *context);
// extern int32_t		v1_3_callValue(void *context, int32_t resultOffset);
// extern int32_t		v1_3_getESDTValue(void *context, int32_t resultOffset);
// extern int32_t		v1_3_getESDTTokenName(void *context, int32_t resultOffset);
// extern long long	v1_3_getESDTTokenNonce(void *context);
// extern int32_t		v1_3_getESDTTokenType(void *context);
// extern long long v1_3_getCurrentESDTNFTNonce(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen);
// extern int32_t		v1_3_getCallValueTokenName(void *context, int32_t callValueOffset, int32_t tokenNameOffset);
// extern void			v1_3_writeLog(void *context, int32_t pointer, int32_t length, int32_t topicPtr, int32_t numTopics);
// extern void			v1_3_writeEventLog(void *context, int32_t numTopics, int32_t topicLengthsOffset, int32_t topicOffset, int32_t dataOffset, int32_t dataLength);
// extern void			v1_3_returnData(void* context, int32_t dataOffset, int32_t length);
// extern void			v1_3_signalError(void* context, int32_t messageOffset, int32_t messageLength);
// extern long long v1_3_getGasLeft(void *context);
// extern int32_t		v1_3_getESDTBalance(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce, int32_t resultOffset);
// extern int32_t		v1_3_getESDTNFTNameLength(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce);
// extern int32_t		v1_3_getESDTNFTAttributeLength(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce);
// extern int32_t		v1_3_getESDTNFTURILength(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce);
// extern int32_t		v1_3_getESDTTokenData(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce, int32_t valueOffset, int32_t propertiesOffset, int32_t hashOffset, int32_t nameOffset, int32_t attributesOffset, int32_t creatorOffset, int32_t royaltiesOffset, int32_t urisOffset);
//
// extern int32_t		v1_3_executeOnDestContext(void *context, long long gas, int32_t addressOffset, int32_t valueOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t		v1_3_executeOnDestContextByCaller(void *context, long long gas, int32_t addressOffset, int32_t valueOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t		v1_3_executeOnSameContext(void *context, long long gas, int32_t addressOffset, int32_t valueOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t		v1_3_delegateExecution(void *context, long long gas, int32_t addressOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t		v1_3_executeReadOnly(void *context, long long gas, int32_t addressOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t		v1_3_createContract(void *context, long long gas, int32_t valueOffset, int32_t codeOffset, int32_t codeMetadataOffset, int32_t length, int32_t resultOffset, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t		v1_3_deployFromSourceContract(void *context, long long gas, int32_t valueOffset, int32_t addressOffset, int32_t codeMetadataOffset, int32_t resultOffset, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern void			v1_3_upgradeContract(void *context, int32_t dstOffset, long long gas, int32_t valueOffset, int32_t codeOffset, int32_t codeMetadataOffset, int32_t length, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern void			v1_3_upgradeFromSourceContract(void *context, int32_t dstOffset, long long gas, int32_t valueOffset, int32_t addressOffset, int32_t codeMetadataOffset, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern void			v1_3_asyncCall(void *context, int32_t dstOffset, int32_t valueOffset, int32_t dataOffset, int32_t length);
// extern void			v1_3_createAsyncCall(void *context, int32_t identifierOffset, int32_t identifierLength, int32_t dstOffset, int32_t valueOffset, int32_t dataOffset, int32_t length, int32_t successCallback, int32_t successLength, int32_t errorCallback, int32_t errorLength, long long gas);
// extern int32_t		v1_3_setAsyncContextCallback(void *context, int32_t identifierOffset, int32_t identifierLength, int32_t callback, int32_t callbackLength);
//
// extern int32_t		v1_3_getNumReturnData(void *context);
// extern int32_t		v1_3_getReturnDataSize(void *context, int32_t resultID);
// extern int32_t		v1_3_getReturnData(void *context, int32_t resultID, int32_t dataOffset);
//
// extern int32_t		v1_3_setStorageLock(void *context, int32_t keyOffset, int32_t keyLength, long long lockTimestamp);
// extern long long v1_3_getStorageLock(void *context, int32_t keyOffset, int32_t keyLength);
// extern int32_t		v1_3_isStorageLocked(void *context, int32_t keyOffset, int32_t keyLength);
// extern int32_t		v1_3_clearStorageLock(void *context, int32_t keyOffset, int32_t keyLength);
//
// extern long long v1_3_getBlockTimestamp(void *context);
// extern long long v1_3_getBlockNonce(void *context);
// extern long long v1_3_getBlockRound(void *context);
// extern long long v1_3_getBlockEpoch(void *context);
// extern void			v1_3_getBlockRandomSeed(void *context, int32_t resultOffset);
// extern void			v1_3_getStateRootHash(void *context, int32_t resultOffset);
//
// extern long long v1_3_getPrevBlockTimestamp(void *context);
// extern long long v1_3_getPrevBlockNonce(void *context);
// extern long long v1_3_getPrevBlockRound(void *context);
// extern long long v1_3_getPrevBlockEpoch(void *context);
// extern void			v1_3_getPrevBlockRandomSeed(void *context, int32_t resultOffset);
// extern void			v1_3_getOriginalTxHash(void *context, int32_t resultOffset);
import "C"

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"unsafe"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/esdt"
	"github.com/multiversx/mx-chain-core-go/data/vm"
	logger "github.com/multiversx/mx-chain-logger-go"
	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
	"github.com/multiversx/mx-chain-vm-common-go/parsers"
	"github.com/multiversx/mx-chain-vm-v1_3-go/math"
	"github.com/multiversx/mx-chain-vm-v1_3-go/vmhost"
	"github.com/multiversx/mx-chain-vm-v1_3-go/wasmer"
)

var logEEI = logger.GetOrCreate("vm/eei")

func getFirstESDTTransferIfExist(vmInput *vmcommon.VMInput) *vmcommon.ESDTTransfer {
	esdtTransfers := vmInput.ESDTTransfers
	if len(esdtTransfers) > 0 {
		return esdtTransfers[0]
	}
	return &vmcommon.ESDTTransfer{
		ESDTValue: big.NewInt(0),
	}
}

// BaseOpsAPIImports creates a new wasmer.Imports populated with the BaseOpsAPI API methods
func BaseOpsAPIImports() (*wasmer.Imports, error) {
	imports := wasmer.NewImports()
	imports = imports.Namespace("env")

	imports, err := imports.Append("getSCAddress", v1_3_getSCAddress, C.v1_3_getSCAddress)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getOwnerAddress", v1_3_getOwnerAddress, C.v1_3_getOwnerAddress)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getShardOfAddress", v1_3_getShardOfAddress, C.v1_3_getShardOfAddress)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("isSmartContract", v1_3_isSmartContract, C.v1_3_isSmartContract)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getExternalBalance", v1_3_getExternalBalance, C.v1_3_getExternalBalance)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockHash", v1_3_blockHash, C.v1_3_blockHash)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("transferValue", v1_3_transferValue, C.v1_3_transferValue)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("transferESDTExecute", v1_3_transferESDTExecute, C.v1_3_transferESDTExecute)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("transferESDTNFTExecute", v1_3_transferESDTNFTExecute, C.v1_3_transferESDTNFTExecute)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("transferESDT", v1_3_transferESDT, C.v1_3_transferESDT)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("transferValueExecute", v1_3_transferValueExecute, C.v1_3_transferValueExecute)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("asyncCall", v1_3_asyncCall, C.v1_3_asyncCall)
	if err != nil {
		return nil, err
	}

	// imports, err = imports.Append("createAsyncCall", createAsyncCall, C.createAsyncCall)
	// if err != nil {
	// 	return nil, err
	// }

	// imports, err = imports.Append("setAsyncContextCallback", setAsyncContextCallback, C.setAsyncContextCallback)
	// if err != nil {
	// 	return nil, err
	// }

	imports, err = imports.Append("getArgumentLength", v1_3_getArgumentLength, C.v1_3_getArgumentLength)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getArgument", v1_3_getArgument, C.v1_3_getArgument)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getFunction", v1_3_getFunction, C.v1_3_getFunction)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getNumArguments", v1_3_getNumArguments, C.v1_3_getNumArguments)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("storageStore", v1_3_storageStore, C.v1_3_storageStore)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("storageLoadLength", v1_3_storageLoadLength, C.v1_3_storageLoadLength)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("storageLoad", v1_3_storageLoad, C.v1_3_storageLoad)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("storageLoadFromAddress", v1_3_storageLoadFromAddress, C.v1_3_storageLoadFromAddress)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getStorageLock", v1_3_getStorageLock, C.v1_3_getStorageLock)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("setStorageLock", v1_3_setStorageLock, C.v1_3_setStorageLock)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("isStorageLocked", v1_3_isStorageLocked, C.v1_3_isStorageLocked)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("clearStorageLock", v1_3_clearStorageLock, C.v1_3_clearStorageLock)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCaller", v1_3_getCaller, C.v1_3_getCaller)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("checkNoPayment", v1_3_checkNoPayment, C.v1_3_checkNoPayment)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCallValue", v1_3_callValue, C.v1_3_callValue)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTValue", v1_3_getESDTValue, C.v1_3_getESDTValue)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTTokenName", v1_3_getESDTTokenName, C.v1_3_getESDTTokenName)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTTokenType", v1_3_getESDTTokenType, C.v1_3_getESDTTokenType)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTTokenNonce", v1_3_getESDTTokenNonce, C.v1_3_getESDTTokenNonce)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCurrentESDTNFTNonce", v1_3_getCurrentESDTNFTNonce, C.v1_3_getCurrentESDTNFTNonce)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCallValueTokenName", v1_3_getCallValueTokenName, C.v1_3_getCallValueTokenName)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("writeLog", v1_3_writeLog, C.v1_3_writeLog)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("writeEventLog", v1_3_writeEventLog, C.v1_3_writeEventLog)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("finish", v1_3_returnData, C.v1_3_returnData)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("signalError", v1_3_signalError, C.v1_3_signalError)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockTimestamp", v1_3_getBlockTimestamp, C.v1_3_getBlockTimestamp)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockNonce", v1_3_getBlockNonce, C.v1_3_getBlockNonce)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockRound", v1_3_getBlockRound, C.v1_3_getBlockRound)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockEpoch", v1_3_getBlockEpoch, C.v1_3_getBlockEpoch)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockRandomSeed", v1_3_getBlockRandomSeed, C.v1_3_getBlockRandomSeed)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getStateRootHash", v1_3_getStateRootHash, C.v1_3_getStateRootHash)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getPrevBlockTimestamp", v1_3_getPrevBlockTimestamp, C.v1_3_getPrevBlockTimestamp)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getPrevBlockNonce", v1_3_getPrevBlockNonce, C.v1_3_getPrevBlockNonce)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getPrevBlockRound", v1_3_getPrevBlockRound, C.v1_3_getPrevBlockRound)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getPrevBlockEpoch", v1_3_getPrevBlockEpoch, C.v1_3_getPrevBlockEpoch)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getPrevBlockRandomSeed", v1_3_getPrevBlockRandomSeed, C.v1_3_getPrevBlockRandomSeed)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getOriginalTxHash", v1_3_getOriginalTxHash, C.v1_3_getOriginalTxHash)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getGasLeft", v1_3_getGasLeft, C.v1_3_getGasLeft)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("executeOnDestContext", v1_3_executeOnDestContext, C.v1_3_executeOnDestContext)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("executeOnDestContextByCaller", v1_3_executeOnDestContextByCaller, C.v1_3_executeOnDestContextByCaller)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("executeOnSameContext", v1_3_executeOnSameContext, C.v1_3_executeOnSameContext)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("delegateExecution", v1_3_delegateExecution, C.v1_3_delegateExecution)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("createContract", v1_3_createContract, C.v1_3_createContract)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("deployFromSourceContract", v1_3_deployFromSourceContract, C.v1_3_deployFromSourceContract)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("upgradeContract", v1_3_upgradeContract, C.v1_3_upgradeContract)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("upgradeFromSourceContract", v1_3_upgradeFromSourceContract, C.v1_3_upgradeFromSourceContract)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("executeReadOnly", v1_3_executeReadOnly, C.v1_3_executeReadOnly)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getNumReturnData", v1_3_getNumReturnData, C.v1_3_getNumReturnData)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getReturnDataSize", v1_3_getReturnDataSize, C.v1_3_getReturnDataSize)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getReturnData", v1_3_getReturnData, C.v1_3_getReturnData)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTBalance", v1_3_getESDTBalance, C.v1_3_getESDTBalance)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTTokenData", v1_3_getESDTTokenData, C.v1_3_getESDTTokenData)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTNFTNameLength", v1_3_getESDTNFTNameLength, C.v1_3_getESDTNFTNameLength)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTNFTAttributeLength", v1_3_getESDTNFTAttributeLength, C.v1_3_getESDTNFTAttributeLength)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTNFTURILength", v1_3_getESDTNFTURILength, C.v1_3_getESDTNFTURILength)
	if err != nil {
		return nil, err
	}

	return imports, nil
}

//export v1_3_getGasLeft
func v1_3_getGasLeft(context unsafe.Pointer) int64 {
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetGasLeft
	metering.UseGas(gasToUse)

	return int64(metering.GasLeft())
}

//export v1_3_getSCAddress
func v1_3_getSCAddress(context unsafe.Pointer, resultOffset int32) {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetSCAddress
	metering.UseGas(gasToUse)

	owner := runtime.GetSCAddress()
	err := runtime.MemStore(resultOffset, owner)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}
}

//export v1_3_getOwnerAddress
func v1_3_getOwnerAddress(context unsafe.Pointer, resultOffset int32) {
	blockchain := vmhost.GetBlockchainContext(context)
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetOwnerAddress
	metering.UseGas(gasToUse)

	owner, err := blockchain.GetOwnerAddress()
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	err = runtime.MemStore(resultOffset, owner)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}
}

//export v1_3_getShardOfAddress
func v1_3_getShardOfAddress(context unsafe.Pointer, addressOffset int32) int32 {
	blockchain := vmhost.GetBlockchainContext(context)
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetShardOfAddress
	metering.UseGas(gasToUse)

	address, err := runtime.MemLoad(addressOffset, vmhost.AddressLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 0
	}

	return int32(blockchain.GetShardOfAddress(address))
}

//export v1_3_isSmartContract
func v1_3_isSmartContract(context unsafe.Pointer, addressOffset int32) int32 {
	blockchain := vmhost.GetBlockchainContext(context)
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.IsSmartContract
	metering.UseGas(gasToUse)

	address, err := runtime.MemLoad(addressOffset, vmhost.AddressLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 0
	}

	isSmartContract := blockchain.IsSmartContract(address)

	return int32(vmhost.BooleanToInt(isSmartContract))
}

//export v1_3_signalError
func v1_3_signalError(context unsafe.Pointer, messageOffset int32, messageLength int32) {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.SignalError
	metering.UseGas(gasToUse)

	message, err := runtime.MemLoad(messageOffset, messageLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}
	runtime.SignalUserError(string(message))
}

//export v1_3_getExternalBalance
func v1_3_getExternalBalance(context unsafe.Pointer, addressOffset int32, resultOffset int32) {
	blockchain := vmhost.GetBlockchainContext(context)
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetExternalBalance
	metering.UseGas(gasToUse)

	address, err := runtime.MemLoad(addressOffset, vmhost.AddressLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	balance := blockchain.GetBalance(address)

	err = runtime.MemStore(resultOffset, balance)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}
}

//export v1_3_blockHash
func v1_3_blockHash(context unsafe.Pointer, nonce int64, resultOffset int32) int32 {
	blockchain := vmhost.GetBlockchainContext(context)
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockHash
	metering.UseGas(gasToUse)

	hash := blockchain.BlockHash(nonce)
	err := runtime.MemStore(resultOffset, hash)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

func isBuiltInCall(data string, host vmhost.VMHost) bool {
	argParser := parsers.NewCallArgsParser()
	functionName, _, _ := argParser.ParseData(data)
	return host.IsBuiltinFunctionName(functionName)
}

func getESDTDataFromBlockchainHook(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
) (*esdt.ESDigitalToken, error) {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)
	blockchain := vmhost.GetBlockchainContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetExternalBalance
	metering.UseGas(gasToUse)

	address, err := runtime.MemLoad(addressOffset, vmhost.AddressLen)
	if err != nil {
		return nil, err
	}

	tokenID, err := runtime.MemLoad(tokenIDOffset, tokenIDLen)
	if err != nil {
		return nil, err
	}

	esdtToken, err := blockchain.GetESDTToken(address, tokenID, uint64(nonce))
	if err != nil {
		return nil, err
	}

	return esdtToken, nil
}

//export v1_3_getESDTBalance
func v1_3_getESDTBalance(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
	resultOffset int32,
) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	esdtData, err := getESDTDataFromBlockchainHook(context, addressOffset, tokenIDOffset, tokenIDLen, nonce)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 0
	}
	err = runtime.MemStore(resultOffset, esdtData.Value.Bytes())
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 0
	}

	return int32(len(esdtData.Value.Bytes()))
}

//export v1_3_getESDTNFTNameLength
func v1_3_getESDTNFTNameLength(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	esdtData, err := getESDTDataFromBlockchainHook(context, addressOffset, tokenIDOffset, tokenIDLen, nonce)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 0
	}
	if esdtData == nil || esdtData.TokenMetaData == nil {
		return 0
	}

	return int32(len(esdtData.TokenMetaData.Name))
}

//export v1_3_getESDTNFTAttributeLength
func v1_3_getESDTNFTAttributeLength(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	esdtData, err := getESDTDataFromBlockchainHook(context, addressOffset, tokenIDOffset, tokenIDLen, nonce)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 0
	}
	if esdtData == nil || esdtData.TokenMetaData == nil {
		return 0
	}

	return int32(len(esdtData.TokenMetaData.Attributes))
}

//export v1_3_getESDTNFTURILength
func v1_3_getESDTNFTURILength(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	esdtData, err := getESDTDataFromBlockchainHook(context, addressOffset, tokenIDOffset, tokenIDLen, nonce)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 0
	}
	if esdtData == nil || esdtData.TokenMetaData == nil {
		return 0
	}
	if len(esdtData.TokenMetaData.URIs) == 0 {
		return 0
	}

	return int32(len(esdtData.TokenMetaData.URIs[0]))
}

//export v1_3_getESDTTokenData
func v1_3_getESDTTokenData(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
	valueHandle int32,
	propertiesOffset int32,
	hashOffset int32,
	nameOffset int32,
	attributesOffset int32,
	creatorOffset int32,
	royaltiesHandle int32,
	urisOffset int32,
) int32 {
	bigInt := vmhost.GetBigIntContext(context)
	runtime := vmhost.GetRuntimeContext(context)
	esdtData, err := getESDTDataFromBlockchainHook(context, addressOffset, tokenIDOffset, tokenIDLen, nonce)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 0
	}

	value := bigInt.GetOne(valueHandle)
	value.Set(esdtData.Value)

	err = runtime.MemStore(propertiesOffset, esdtData.Properties)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 0
	}

	if esdtData.TokenMetaData != nil {
		err = runtime.MemStore(hashOffset, esdtData.TokenMetaData.Hash)
		if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
			return 0
		}
		err = runtime.MemStore(nameOffset, esdtData.TokenMetaData.Name)
		if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
			return 0
		}
		err = runtime.MemStore(attributesOffset, esdtData.TokenMetaData.Attributes)
		if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
			return 0
		}
		err = runtime.MemStore(creatorOffset, esdtData.TokenMetaData.Creator)
		if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
			return 0
		}

		royalties := bigInt.GetOne(royaltiesHandle)
		royalties.SetUint64(uint64(esdtData.TokenMetaData.Royalties))

		if len(esdtData.TokenMetaData.URIs) > 0 {
			err = runtime.MemStore(urisOffset, esdtData.TokenMetaData.URIs[0])
			if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
				return 0
			}
		}
	}
	return int32(len(esdtData.Value.Bytes()))
}

//export v1_3_transferValue
func v1_3_transferValue(context unsafe.Pointer, destOffset int32, valueOffset int32, dataOffset int32, length int32) int32 {
	host := vmhost.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()
	output := host.Output()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.TransferValue
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	dest, err := runtime.MemLoad(destOffset, vmhost.AddressLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	valueBytes, err := runtime.MemLoad(valueOffset, vmhost.BalanceLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.PersistPerByte, uint64(length))
	metering.UseGas(gasToUse)

	data, err := runtime.MemLoad(dataOffset, length)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	if isBuiltInCall(string(data), host) {
		return 1
	}

	err = output.Transfer(dest, sender, 0, 0, big.NewInt(0).SetBytes(valueBytes), data, vm.DirectCall)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

type indirectContractCallArguments struct {
	dest      []byte
	value     *big.Int
	function  []byte
	args      [][]byte
	actualLen int32
}

func extractIndirectContractCallArgumentsWithValue(
	host vmhost.VMHost,
	destOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) (*indirectContractCallArguments, error) {
	return extractIndirectContractCallArguments(
		host,
		destOffset,
		valueOffset,
		true,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

func extractIndirectContractCallArgumentsWithoutValue(
	host vmhost.VMHost,
	destOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) (*indirectContractCallArguments, error) {
	return extractIndirectContractCallArguments(
		host,
		destOffset,
		0,
		false,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

func extractIndirectContractCallArguments(
	host vmhost.VMHost,
	destOffset int32,
	valueOffset int32,
	hasValueOffset bool,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) (*indirectContractCallArguments, error) {
	runtime := host.Runtime()
	metering := host.Metering()

	dest, err := runtime.MemLoad(destOffset, vmhost.AddressLen)
	if err != nil {
		return nil, err
	}

	var value *big.Int

	if hasValueOffset {
		valueBytes, err := runtime.MemLoad(valueOffset, vmhost.BalanceLen)
		if err != nil {
			return nil, err
		}
		value = big.NewInt(0).SetBytes(valueBytes)
	}

	function, err := runtime.MemLoad(functionOffset, functionLength)
	if err != nil {
		return nil, err
	}

	args, actualLen, err := getArgumentsFromMemory(
		host,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)

	gasToUse := math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(actualLen))
	metering.UseGas(gasToUse)

	return &indirectContractCallArguments{
		dest:      dest,
		value:     value,
		function:  function,
		args:      args,
		actualLen: actualLen,
	}, nil
}

//export v1_3_transferValueExecute
func v1_3_transferValueExecute(
	context unsafe.Pointer,
	destOffset int32,
	valueOffset int32,
	gasLimit int64,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := vmhost.GetVMHost(context)
	return TransferValueExecuteWithHost(
		host,
		destOffset,
		valueOffset,
		gasLimit,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// TransferValueExecuteWithHost - transferValueExecute with host instead of pointer context
func TransferValueExecuteWithHost(
	host vmhost.VMHost,
	destOffset int32,
	valueOffset int32,
	gasLimit int64,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.TransferValue
	metering.UseGas(gasToUse)

	callArgs, err := extractIndirectContractCallArgumentsWithValue(
		host, destOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return TransferValueExecuteWithTypedArgs(
		host,
		callArgs.dest,
		callArgs.value,
		gasLimit,
		callArgs.function,
		callArgs.args,
	)
}

// TransferValueExecuteWithTypedArgs - transferValueExecute with args already read from memory
func TransferValueExecuteWithTypedArgs(
	host vmhost.VMHost,
	dest []byte,
	value *big.Int,
	gasLimit int64,
	function []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()
	output := host.Output()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.TransferValue
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()

	var err error
	var contractCallInput *vmcommon.ContractCallInput

	if function != nil {
		contractCallInput, err = prepareIndirectContractCallInput(
			host,
			sender,
			value,
			gasLimit,
			dest,
			function,
			args,
			gasToUse,
			false,
		)
		if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
			return 1
		}
	}

	if contractCallInput != nil {
		if host.IsBuiltinFunctionName(contractCallInput.Function) {
			return 1
		}
	}

	if host.AreInSameShard(sender, dest) && contractCallInput != nil && host.Blockchain().IsSmartContract(dest) {
		logEEI.Trace("eGLD pre-transfer execution begin")
		_, _, err = host.ExecuteOnDestContext(contractCallInput)
		if err != nil {
			logEEI.Trace("eGLD pre-transfer execution failed", "error", err)
			return 1
		}

		return 0
	}

	data := makeCrossShardCallFromInput(contractCallInput)
	err = output.Transfer(dest, sender, uint64(gasLimit), 0, value, []byte(data), vm.DirectCall)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

func makeCrossShardCallFromInput(vmInput *vmcommon.ContractCallInput) string {
	if vmInput == nil {
		return ""
	}

	txData := vmInput.Function
	for _, arg := range vmInput.Arguments {
		txData += "@" + hex.EncodeToString(arg)
	}

	return txData
}

//export v1_3_transferESDT
func v1_3_transferESDT(
	context unsafe.Pointer,
	destOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	valueOffset int32,
	gasLimit int64,
	dataOffset int32,
	length int32,
) int32 {
	host := vmhost.GetVMHost(context)
	metering := host.Metering()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.TransferValue
	metering.UseGas(gasToUse)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.PersistPerByte, uint64(length))
	metering.UseGas(gasToUse)
	logEEI.Warn("transferESDT() is deprecated")
	// this is only for backward compatibility - function deprecated
	return 1
}

//export v1_3_transferESDTExecute
func v1_3_transferESDTExecute(
	context unsafe.Pointer,
	destOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	valueOffset int32,
	gasLimit int64,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	return v1_3_transferESDTNFTExecute(context, destOffset, tokenIDOffset, tokenIDLen, valueOffset, 0,
		gasLimit, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
}

//export v1_3_transferESDTNFTExecute
func v1_3_transferESDTNFTExecute(
	context unsafe.Pointer,
	destOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	valueOffset int32,
	nonce int64,
	gasLimit int64,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := vmhost.GetVMHost(context)
	return TransferESDTNFTExecuteWithHost(
		host,
		destOffset,
		tokenIDOffset,
		tokenIDLen,
		valueOffset,
		nonce,
		gasLimit,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset)
}

// TransferESDTNFTExecuteWithHost contains only memory reading of arguments
func TransferESDTNFTExecuteWithHost(
	host vmhost.VMHost,
	destOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	valueOffset int32,
	nonce int64,
	gasLimit int64,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	tokenIdentifier, executeErr := runtime.MemLoad(tokenIDOffset, tokenIDLen)
	if vmhost.WithFaultAndHost(host, executeErr, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	callArgs, err := extractIndirectContractCallArgumentsWithValue(
		host, destOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	gasToUse := math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(callArgs.actualLen))
	metering.UseGas(gasToUse)

	return TransferESDTNFTExecuteWithTypedArgs(
		host,
		callArgs.value,
		tokenIdentifier,
		callArgs.dest,
		nonce,
		gasLimit,
		callArgs.function,
		callArgs.args,
	)
}

// TransferESDTNFTExecuteWithTypedArgs defines the actual transfer ESDT execute logic
func TransferESDTNFTExecuteWithTypedArgs(
	host vmhost.VMHost,
	esdtValue *big.Int,
	esdtTokenName []byte,
	dest []byte,
	nonce int64,
	gasLimit int64,
	function []byte,
	data [][]byte,
) int32 {

	var executeErr error

	runtime := host.Runtime()
	metering := host.Metering()

	output := host.Output()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.TransferValue
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()

	var contractCallInput *vmcommon.ContractCallInput
	if function != nil {
		contractCallInput, executeErr = prepareIndirectContractCallInput(
			host,
			sender,
			big.NewInt(0),
			gasLimit,
			dest,
			function,
			data,
			gasToUse,
			false,
		)
		if vmhost.WithFaultAndHost(host, executeErr, runtime.SyncExecAPIErrorShouldFailExecution()) {
			return 1
		}

		esdtTokenType := core.Fungible
		if nonce > 0 {
			esdtTokenType = core.NonFungible
		}
		contractCallInput.ESDTTransfers = make([]*vmcommon.ESDTTransfer, 1)
		contractCallInput.ESDTTransfers[0] = &vmcommon.ESDTTransfer{
			ESDTValue:      esdtValue,
			ESDTTokenName:  esdtTokenName,
			ESDTTokenType:  uint32(esdtTokenType),
			ESDTTokenNonce: uint64(nonce),
		}
	}

	snapshotBeforeTransfer := host.Blockchain().GetSnapshot()

	gasLimitForExec, executeErr := output.TransferESDT(dest, sender, esdtTokenName, uint64(nonce), esdtValue, contractCallInput)
	if vmhost.WithFaultAndHost(host, executeErr, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	if host.AreInSameShard(sender, dest) && contractCallInput != nil && host.Blockchain().IsSmartContract(dest) {
		contractCallInput.GasProvided = gasLimitForExec
		logEEI.Trace("ESDT post-transfer execution begin")
		_, _, executeErr = host.ExecuteOnDestContext(contractCallInput)
		if executeErr != nil {
			logEEI.Trace("ESDT post-transfer execution failed", "error", executeErr)
			host.Blockchain().RevertToSnapshot(snapshotBeforeTransfer)
			return 1
		}

		return 0
	}

	return 0
}

//export v1_3_createAsyncCall
func v1_3_createAsyncCall(context unsafe.Pointer,
	asyncContextIdentifier int32,
	identifierLength int32,
	destOffset int32,
	valueOffset int32,
	dataOffset int32,
	length int32,
	successOffset int32,
	successLength int32,
	errorOffset int32,
	errorLength int32,
	gas int64,
) {
	host := vmhost.GetVMHost(context)
	runtime := host.Runtime()

	// TODO consume gas

	acIdentifier, err := runtime.MemLoad(asyncContextIdentifier, identifierLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	calledSCAddress, err := runtime.MemLoad(destOffset, vmhost.AddressLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	value, err := runtime.MemLoad(valueOffset, vmhost.BalanceLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	data, err := runtime.MemLoad(dataOffset, length)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	successFunc, err := runtime.MemLoad(successOffset, successLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	errorFunc, err := runtime.MemLoad(errorOffset, errorLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	err = runtime.AddAsyncContextCall(acIdentifier, &vmhost.AsyncGeneratedCall{
		Destination:     calledSCAddress,
		Data:            data,
		ValueBytes:      value,
		SuccessCallback: string(successFunc),
		ErrorCallback:   string(errorFunc),
		ProvidedGas:     uint64(gas),
	})
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}
}

//export v1_3_setAsyncContextCallback
func v1_3_setAsyncContextCallback(context unsafe.Pointer,
	asyncContextIdentifier int32,
	identifierLength int32,
	callback int32,
	callbackLength int32,
) int32 {
	host := vmhost.GetVMHost(context)
	runtime := host.Runtime()

	// TODO consume gas

	acIdentifier, err := runtime.MemLoad(asyncContextIdentifier, identifierLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	asyncContext, err := runtime.GetAsyncContext(acIdentifier)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	callbackFunc, err := runtime.MemLoad(callback, callbackLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	asyncContext.Callback = string(callbackFunc)

	return 0
}

//export v1_3_upgradeContract
func v1_3_upgradeContract(
	context unsafe.Pointer,
	destOffset int32,
	gasLimit int64,
	valueOffset int32,
	codeOffset int32,
	codeMetadataOffset int32,
	length int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) {
	host := vmhost.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.CreateContract
	metering.UseGas(gasToUse)

	value, err := runtime.MemLoad(valueOffset, vmhost.BalanceLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	code, err := runtime.MemLoad(codeOffset, length)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	codeMetadata, err := runtime.MemLoad(codeMetadataOffset, vmhost.CodeMetadataLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	data, actualLen, err := getArgumentsFromMemory(
		host,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(actualLen))
	metering.UseGas(gasToUse)

	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	calledSCAddress, err := runtime.MemLoad(destOffset, vmhost.AddressLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	gasSchedule := metering.GasSchedule()
	gasToUse = math.MulUint64(gasSchedule.BaseOperationCost.DataCopyPerByte, uint64(length))
	metering.UseGas(gasToUse)

	upgradeContract(host, calledSCAddress, code, codeMetadata, value, data, gasLimit)
}

//export v1_3_upgradeFromSourceContract
func v1_3_upgradeFromSourceContract(
	context unsafe.Pointer,
	destOffset int32,
	gasLimit int64,
	valueOffset int32,
	sourceContractAddressOffset int32,
	codeMetadataOffset int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) {
	host := vmhost.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.CreateContract
	metering.UseGas(gasToUse)

	value, err := runtime.MemLoad(valueOffset, vmhost.BalanceLen)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	sourceContractAddress, err := runtime.MemLoad(sourceContractAddressOffset, vmhost.AddressLen)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	codeMetadata, err := runtime.MemLoad(codeMetadataOffset, vmhost.CodeMetadataLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	data, actualLen, err := getArgumentsFromMemory(
		host,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(actualLen))
	metering.UseGas(gasToUse)

	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	calledSCAddress, err := runtime.MemLoad(destOffset, vmhost.AddressLen)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	UpgradeFromSourceContractWithTypedArgs(
		host,
		sourceContractAddress,
		calledSCAddress,
		value,
		data,
		gasLimit,
		codeMetadata,
	)
}

// UpgradeFromSourceContractWithTypedArgs - upgradeFromSourceContract with args already read from memory
func UpgradeFromSourceContractWithTypedArgs(
	host vmhost.VMHost,
	sourceContractAddress []byte,
	destContractAddress []byte,
	value []byte,
	data [][]byte,
	gasLimit int64,
	codeMetadata []byte,
) {
	runtime := host.Runtime()
	blockchain := host.Blockchain()

	code, err := blockchain.GetCode(sourceContractAddress)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	upgradeContract(host, destContractAddress, code, codeMetadata, value, data, gasLimit)
}

func upgradeContract(
	host vmhost.VMHost,
	destContractAddress []byte,
	code []byte,
	codeMetadata []byte,
	value []byte,
	data [][]byte,
	gasLimit int64,
) {
	runtime := host.Runtime()
	metering := host.Metering()
	gasSchedule := metering.GasSchedule()
	minAsyncCallCost := math.AddUint64(
		math.MulUint64(2, gasSchedule.BaseOpsAPICost.AsyncCallStep),
		gasSchedule.BaseOpsAPICost.AsyncCallbackGasLock)
	if uint64(gasLimit) < minAsyncCallCost {
		runtime.SetRuntimeBreakpointValue(vmhost.BreakpointOutOfGas)
		return
	}

	// Set up the async call as if it is not known whether the called SC
	// is in the same shard with the caller or not. This will be later resolved
	// by runtime.ExecuteAsyncCall().
	callData := vmhost.UpgradeFunctionName + "@" + hex.EncodeToString(code) + "@" + hex.EncodeToString(codeMetadata)
	for _, arg := range data {
		callData += "@" + hex.EncodeToString(arg)
	}

	err := runtime.ExecuteAsyncCall(
		destContractAddress,
		[]byte(callData),
		value,
	)
	logEEI.Trace("upgradeContract", "error", err)
}

//export v1_3_asyncCall
func v1_3_asyncCall(context unsafe.Pointer, destOffset int32, valueOffset int32, dataOffset int32, length int32) {
	host := vmhost.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()

	gasSchedule := metering.GasSchedule()
	gasToUse := gasSchedule.BaseOpsAPICost.AsyncCallStep
	metering.UseGas(gasToUse)

	calledSCAddress, err := runtime.MemLoad(destOffset, vmhost.AddressLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	value, err := runtime.MemLoad(valueOffset, vmhost.BalanceLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	gasToUse = math.MulUint64(gasSchedule.BaseOperationCost.DataCopyPerByte, uint64(length))
	metering.UseGas(gasToUse)

	data, err := runtime.MemLoad(dataOffset, length)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	err = runtime.ExecuteAsyncCall(calledSCAddress, data, value)
	if errors.Is(err, vmhost.ErrNotEnoughGas) {
		runtime.SetRuntimeBreakpointValue(vmhost.BreakpointOutOfGas)
		return
	}
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}
}

//export v1_3_getArgumentLength
func v1_3_getArgumentLength(context unsafe.Pointer, id int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetArgument
	metering.UseGas(gasToUse)

	args := runtime.Arguments()
	if id < 0 || int32(len(args)) <= id {
		return -1
	}

	return int32(len(args[id]))
}

//export v1_3_getArgument
func v1_3_getArgument(context unsafe.Pointer, id int32, argOffset int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetArgument
	metering.UseGas(gasToUse)

	args := runtime.Arguments()
	if id < 0 || int32(len(args)) <= id {
		return -1
	}

	err := runtime.MemStore(argOffset, args[id])
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(args[id]))
}

//export v1_3_getFunction
func v1_3_getFunction(context unsafe.Pointer, functionOffset int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetFunction
	metering.UseGas(gasToUse)

	function := runtime.Function()
	err := runtime.MemStore(functionOffset, []byte(function))
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(function))
}

//export v1_3_getNumArguments
func v1_3_getNumArguments(context unsafe.Pointer) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetNumArguments
	metering.UseGas(gasToUse)

	args := runtime.Arguments()
	return int32(len(args))
}

//export v1_3_storageStore
func v1_3_storageStore(context unsafe.Pointer, keyOffset int32, keyLength int32, dataOffset int32, dataLength int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	storage := vmhost.GetStorageContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.StorageStore
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	data, err := runtime.MemLoad(dataOffset, dataLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	storageStatus, err := storage.SetStorage(key, data)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	return int32(storageStatus)
}

//export v1_3_storageLoadLength
func v1_3_storageLoadLength(context unsafe.Pointer, keyOffset int32, keyLength int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	storage := vmhost.GetStorageContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.StorageLoad
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	data := storage.GetStorageUnmetered(key)

	return int32(len(data))
}

//export v1_3_storageLoadFromAddress
func v1_3_storageLoadFromAddress(context unsafe.Pointer, addressOffset int32, keyOffset int32, keyLength int32, dataOffset int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	storage := vmhost.GetStorageContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.StorageLoad
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	address, err := runtime.MemLoad(addressOffset, vmhost.AddressLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	data := storage.GetStorageFromAddress(address, key)

	err = runtime.MemStore(dataOffset, data)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(data))
}

//export v1_3_storageLoad
func v1_3_storageLoad(context unsafe.Pointer, keyOffset int32, keyLength int32, dataOffset int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	storage := vmhost.GetStorageContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.StorageLoad
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	data := storage.GetStorage(key)

	err = runtime.MemStore(dataOffset, data)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(data))
}

//export v1_3_setStorageLock
func v1_3_setStorageLock(context unsafe.Pointer, keyOffset int32, keyLength int32, lockTimestamp int64) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	storage := vmhost.GetStorageContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.Int64StorageStore
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	timeLockKey := vmhost.CustomStorageKey(vmhost.TimeLockKeyPrefix, key)
	bigTimestamp := big.NewInt(0).SetInt64(lockTimestamp)
	storageStatus, err := storage.SetProtectedStorage(timeLockKey, bigTimestamp.Bytes())
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}
	return int32(storageStatus)
}

//export v1_3_getStorageLock
func v1_3_getStorageLock(context unsafe.Pointer, keyOffset int32, keyLength int32) int64 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)
	storage := vmhost.GetStorageContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.StorageLoad
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	timeLockKey := vmhost.CustomStorageKey(vmhost.TimeLockKeyPrefix, key)
	data := storage.GetStorage(timeLockKey)
	timeLock := big.NewInt(0).SetBytes(data).Int64()

	return timeLock
}

//export v1_3_isStorageLocked
func v1_3_isStorageLocked(context unsafe.Pointer, keyOffset int32, keyLength int32) int32 {
	timeLock := v1_3_getStorageLock(context, keyOffset, keyLength)
	if timeLock < 0 {
		return -1
	}

	currentTimestamp := v1_3_getBlockTimestamp(context)
	if timeLock <= currentTimestamp {
		return 0
	}

	return 1
}

//export v1_3_clearStorageLock
func v1_3_clearStorageLock(context unsafe.Pointer, keyOffset int32, keyLength int32) int32 {
	return v1_3_setStorageLock(context, keyOffset, keyLength, 0)
}

//export v1_3_getCaller
func v1_3_getCaller(context unsafe.Pointer, resultOffset int32) {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCaller
	metering.UseGas(gasToUse)

	caller := runtime.GetVMInput().CallerAddr

	err := runtime.MemStore(resultOffset, caller)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}
}

//export v1_3_checkNoPayment
func v1_3_checkNoPayment(context unsafe.Pointer) {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCallValue
	metering.UseGas(gasToUse)

	vmInput := runtime.GetVMInput()
	if vmInput.CallValue.Sign() > 0 {
		runtime := vmhost.GetRuntimeContext(context)
		vmhost.WithFault(vmhost.ErrNonPayableFunctionEgld, context, runtime.BaseOpsErrorShouldFailExecution())
		return
	}
	if len(vmInput.ESDTTransfers) > 0 {
		runtime := vmhost.GetRuntimeContext(context)
		vmhost.WithFault(vmhost.ErrNonPayableFunctionEsdt, context, runtime.BaseOpsErrorShouldFailExecution())
		return
	}
}

//export v1_3_callValue
func v1_3_callValue(context unsafe.Pointer, resultOffset int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCallValue
	metering.UseGas(gasToUse)

	value := runtime.GetVMInput().CallValue.Bytes()
	value = vmhost.PadBytesLeft(value, vmhost.BalanceLen)

	err := runtime.MemStore(resultOffset, value)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(value))
}

//export v1_3_getESDTValue
func v1_3_getESDTValue(context unsafe.Pointer, resultOffset int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCallValue
	metering.UseGas(gasToUse)

	var value []byte

	esdtTransfer := getFirstESDTTransferIfExist(runtime.GetVMInput())
	if esdtTransfer.ESDTValue.Cmp(vmhost.Zero) > 0 {
		value = esdtTransfer.ESDTValue.Bytes()
		value = vmhost.PadBytesLeft(value, vmhost.BalanceLen)
	}

	err := runtime.MemStore(resultOffset, value)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(value))
}

//export v1_3_getESDTTokenName
func v1_3_getESDTTokenName(context unsafe.Pointer, resultOffset int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCallValue
	metering.UseGas(gasToUse)

	esdtTransfer := getFirstESDTTransferIfExist(runtime.GetVMInput())
	tokenName := esdtTransfer.ESDTTokenName

	err := runtime.MemStore(resultOffset, tokenName)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(tokenName))
}

//export v1_3_getESDTTokenNonce
func v1_3_getESDTTokenNonce(context unsafe.Pointer) int64 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCallValue
	metering.UseGas(gasToUse)

	esdtTransfer := getFirstESDTTransferIfExist(runtime.GetVMInput())
	return int64(esdtTransfer.ESDTTokenNonce)
}

//export v1_3_getCurrentESDTNFTNonce
func v1_3_getCurrentESDTNFTNonce(context unsafe.Pointer, addressOffset int32, tokenIDOffset int32, tokenIDLen int32) int64 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)
	storage := vmhost.GetStorageContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.StorageLoad
	metering.UseGas(gasToUse)

	destination, err := runtime.MemLoad(addressOffset, vmhost.AddressLen)
	if err != nil {
		return 0
	}

	tokenID, err := runtime.MemLoad(tokenIDOffset, tokenIDLen)
	if err != nil {
		return 0
	}

	key := []byte(core.ProtectedKeyPrefix + core.ESDTNFTLatestNonceIdentifier + string(tokenID))
	data := storage.GetStorageFromAddress(destination, key)

	nonce := big.NewInt(0).SetBytes(data).Uint64()
	return int64(nonce)
}

//export v1_3_getESDTTokenType
func v1_3_getESDTTokenType(context unsafe.Pointer) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCallValue
	metering.UseGas(gasToUse)

	esdtTransfer := getFirstESDTTransferIfExist(runtime.GetVMInput())
	return int32(esdtTransfer.ESDTTokenType)
}

//export v1_3_getCallValueTokenName
func v1_3_getCallValueTokenName(context unsafe.Pointer, callValueOffset int32, tokenNameOffset int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetCallValue
	metering.UseGas(gasToUse)

	callValue := runtime.GetVMInput().CallValue.Bytes()
	tokenName := make([]byte, 0)
	esdtTransfer := getFirstESDTTransferIfExist(runtime.GetVMInput())

	if len(esdtTransfer.ESDTTokenName) > 0 {
		tokenName = make([]byte, 0, len(esdtTransfer.ESDTTokenName))
		copy(tokenName, esdtTransfer.ESDTTokenName)
		callValue = esdtTransfer.ESDTValue.Bytes()
	}
	callValue = vmhost.PadBytesLeft(callValue, vmhost.BalanceLen)

	err := runtime.MemStore(tokenNameOffset, tokenName)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	err = runtime.MemStore(callValueOffset, callValue)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(tokenName))
}

//export v1_3_writeLog
func v1_3_writeLog(context unsafe.Pointer, dataPointer int32, dataLength int32, topicPtr int32, numTopics int32) {
	// note: deprecated
	runtime := vmhost.GetRuntimeContext(context)
	output := vmhost.GetOutputContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.Log
	gas := math.MulUint64(metering.GasSchedule().BaseOperationCost.PersistPerByte, uint64(numTopics*vmhost.HashLen+dataLength))
	gasToUse = math.AddUint64(gasToUse, gas)
	metering.UseGas(gasToUse)

	log, err := runtime.MemLoad(dataPointer, dataLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	topics, err := vmhost.GuardedMakeByteSlice2D(numTopics)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	for i := int32(0); i < numTopics; i++ {
		topics[i], err = runtime.MemLoad(topicPtr+i*vmhost.HashLen, vmhost.HashLen)
		if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
			return
		}
	}

	output.WriteLog(runtime.GetSCAddress(), topics, log)
}

//export v1_3_writeEventLog
func v1_3_writeEventLog(
	context unsafe.Pointer,
	numTopics int32,
	topicLengthsOffset int32,
	topicOffset int32,
	dataOffset int32,
	dataLength int32) {

	host := vmhost.GetVMHost(context)
	runtime := vmhost.GetRuntimeContext(context)
	output := vmhost.GetOutputContext(context)
	metering := vmhost.GetMeteringContext(context)

	topics, topicDataTotalLen, err := getArgumentsFromMemory(
		host,
		numTopics,
		topicLengthsOffset,
		topicOffset,
	)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	data, err := runtime.MemLoad(dataOffset, dataLength)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	gasToUse := metering.GasSchedule().BaseOpsAPICost.Log
	gasForData := math.MulUint64(
		metering.GasSchedule().BaseOperationCost.DataCopyPerByte,
		uint64(topicDataTotalLen+dataLength))
	gasToUse = math.AddUint64(gasToUse, gasForData)
	metering.UseGas(gasToUse)

	output.WriteLog(runtime.GetSCAddress(), topics, data)
}

//export v1_3_getBlockTimestamp
func v1_3_getBlockTimestamp(context unsafe.Pointer) int64 {
	blockchain := vmhost.GetBlockchainContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockTimeStamp
	metering.UseGas(gasToUse)

	return int64(blockchain.CurrentTimeStamp())
}

//export v1_3_getBlockNonce
func v1_3_getBlockNonce(context unsafe.Pointer) int64 {
	blockchain := vmhost.GetBlockchainContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockNonce
	metering.UseGas(gasToUse)

	return int64(blockchain.CurrentNonce())
}

//export v1_3_getBlockRound
func v1_3_getBlockRound(context unsafe.Pointer) int64 {
	blockchain := vmhost.GetBlockchainContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockRound
	metering.UseGas(gasToUse)

	return int64(blockchain.CurrentRound())
}

//export v1_3_getBlockEpoch
func v1_3_getBlockEpoch(context unsafe.Pointer) int64 {
	blockchain := vmhost.GetBlockchainContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockEpoch
	metering.UseGas(gasToUse)

	return int64(blockchain.CurrentEpoch())
}

//export v1_3_getBlockRandomSeed
func v1_3_getBlockRandomSeed(context unsafe.Pointer, pointer int32) {
	runtime := vmhost.GetRuntimeContext(context)
	blockchain := vmhost.GetBlockchainContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockRandomSeed
	metering.UseGas(gasToUse)

	randomSeed := blockchain.CurrentRandomSeed()
	err := runtime.MemStore(pointer, randomSeed)
	vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution())
}

//export v1_3_getStateRootHash
func v1_3_getStateRootHash(context unsafe.Pointer, pointer int32) {
	runtime := vmhost.GetRuntimeContext(context)
	blockchain := vmhost.GetBlockchainContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetStateRootHash
	metering.UseGas(gasToUse)

	stateRootHash := blockchain.GetStateRootHash()
	err := runtime.MemStore(pointer, stateRootHash)
	vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution())
}

//export v1_3_getPrevBlockTimestamp
func v1_3_getPrevBlockTimestamp(context unsafe.Pointer) int64 {
	blockchain := vmhost.GetBlockchainContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockTimeStamp
	metering.UseGas(gasToUse)

	return int64(blockchain.LastTimeStamp())
}

//export v1_3_getPrevBlockNonce
func v1_3_getPrevBlockNonce(context unsafe.Pointer) int64 {
	blockchain := vmhost.GetBlockchainContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockNonce
	metering.UseGas(gasToUse)

	return int64(blockchain.LastNonce())
}

//export v1_3_getPrevBlockRound
func v1_3_getPrevBlockRound(context unsafe.Pointer) int64 {
	blockchain := vmhost.GetBlockchainContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockRound
	metering.UseGas(gasToUse)

	return int64(blockchain.LastRound())
}

//export v1_3_getPrevBlockEpoch
func v1_3_getPrevBlockEpoch(context unsafe.Pointer) int64 {
	blockchain := vmhost.GetBlockchainContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockEpoch
	metering.UseGas(gasToUse)

	return int64(blockchain.LastEpoch())
}

//export v1_3_getPrevBlockRandomSeed
func v1_3_getPrevBlockRandomSeed(context unsafe.Pointer, pointer int32) {
	runtime := vmhost.GetRuntimeContext(context)
	blockchain := vmhost.GetBlockchainContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockRandomSeed
	metering.UseGas(gasToUse)

	randomSeed := blockchain.LastRandomSeed()
	err := runtime.MemStore(pointer, randomSeed)
	vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution())
}

//export v1_3_returnData
func v1_3_returnData(context unsafe.Pointer, pointer int32, length int32) {
	runtime := vmhost.GetRuntimeContext(context)
	output := vmhost.GetOutputContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.Finish
	gas := math.MulUint64(metering.GasSchedule().BaseOperationCost.PersistPerByte, uint64(length))
	gasToUse = math.AddUint64(gasToUse, gas)
	metering.UseGas(gasToUse)

	data, err := runtime.MemLoad(pointer, length)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return
	}

	output.Finish(data)
}

//export v1_3_executeOnSameContext
func v1_3_executeOnSameContext(
	context unsafe.Pointer,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := vmhost.GetVMHost(context)
	return ExecuteOnSameContextWithHost(
		host,
		gasLimit,
		addressOffset,
		valueOffset,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// ExecuteOnSameContextWithHost - executeOnSameContext with host instead of pointer context
func ExecuteOnSameContextWithHost(
	host vmhost.VMHost,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()

	callArgs, err := extractIndirectContractCallArgumentsWithValue(
		host, addressOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return ExecuteOnSameContextWithTypedArgs(
		host,
		gasLimit,
		callArgs.value,
		callArgs.function,
		callArgs.dest,
		callArgs.args,
	)
}

// ExecuteOnSameContextWithTypedArgs - executeOnSameContext with args already read from memory
func ExecuteOnSameContextWithTypedArgs(
	host vmhost.VMHost,
	gasLimit int64,
	value *big.Int,
	function []byte,
	dest []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.ExecuteOnSameContext
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	contractCallInput, err := prepareIndirectContractCallInput(
		host,
		sender,
		value,
		gasLimit,
		dest,
		function,
		args,
		gasToUse,
		true,
	)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	if isBuiltInCall(contractCallInput.Function, host) {
		return 1
	}

	_, err = host.ExecuteOnSameContext(contractCallInput)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_3_executeOnDestContext
func v1_3_executeOnDestContext(
	context unsafe.Pointer,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := vmhost.GetVMHost(context)
	return ExecuteOnDestContextWithHost(
		host,
		gasLimit,
		addressOffset,
		valueOffset,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// ExecuteOnDestContextWithHost - executeOnDestContext with host instead of pointer context
func ExecuteOnDestContextWithHost(
	host vmhost.VMHost,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()

	callArgs, err := extractIndirectContractCallArgumentsWithValue(
		host, addressOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return ExecuteOnDestContextWithTypedArgs(
		host,
		gasLimit,
		callArgs.value,
		callArgs.function,
		callArgs.dest,
		callArgs.args,
	)
}

// ExecuteOnDestContextWithTypedArgs - executeOnDestContext with args already read from memory
func ExecuteOnDestContextWithTypedArgs(
	host vmhost.VMHost,
	gasLimit int64,
	value *big.Int,
	function []byte,
	dest []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.ExecuteOnDestContext
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	contractCallInput, err := prepareIndirectContractCallInput(
		host,
		sender,
		value,
		gasLimit,
		dest,
		function,
		args,
		gasToUse,
		true,
	)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	_, _, err = host.ExecuteOnDestContext(contractCallInput)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_3_executeOnDestContextByCaller
func v1_3_executeOnDestContextByCaller(
	context unsafe.Pointer,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := vmhost.GetVMHost(context)
	return ExecuteOnDestContextByCallerWithHost(
		host,
		gasLimit,
		addressOffset,
		valueOffset,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// ExecuteOnDestContextByCallerWithHost - executeOnDestContextByCaller with host instead of pointer context
func ExecuteOnDestContextByCallerWithHost(
	host vmhost.VMHost,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()

	callArgs, err := extractIndirectContractCallArgumentsWithValue(
		host, addressOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return ExecuteOnDestContextByCallerWithTypedArgs(
		host,
		gasLimit,
		callArgs.value,
		callArgs.function,
		callArgs.dest,
		callArgs.args,
	)
}

// ExecuteOnDestContextByCallerWithTypedArgs - executeOnDestContextByCaller with args already read from memory
func ExecuteOnDestContextByCallerWithTypedArgs(
	host vmhost.VMHost,
	gasLimit int64,
	value *big.Int,
	function []byte,
	dest []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.ExecuteOnDestContext
	metering.UseGas(gasToUse)

	send := runtime.GetVMInput().CallerAddr
	contractCallInput, err := prepareIndirectContractCallInput(
		host,
		send,
		value,
		gasLimit,
		dest,
		function,
		args,
		gasToUse,
		true,
	)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	if isBuiltInCall(contractCallInput.Function, host) {
		return 1
	}

	_, _, err = host.ExecuteOnDestContext(contractCallInput)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_3_delegateExecution
func v1_3_delegateExecution(
	context unsafe.Pointer,
	gasLimit int64,
	addressOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := vmhost.GetVMHost(context)
	return DelegateExecutionWithHost(
		host,
		gasLimit,
		addressOffset,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// DelegateExecutionWithHost - delegateExecution with host instead of pointer context
func DelegateExecutionWithHost(
	host vmhost.VMHost,
	gasLimit int64,
	addressOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()

	callArgs, err := extractIndirectContractCallArgumentsWithoutValue(
		host, addressOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return DelegateExecutionWithTypedArgs(
		host,
		gasLimit,
		callArgs.function,
		callArgs.dest,
		callArgs.args,
	)
}

// DelegateExecutionWithTypedArgs - delegateExecution with args already read from memory
func DelegateExecutionWithTypedArgs(
	host vmhost.VMHost,
	gasLimit int64,
	function []byte,
	dest []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.DelegateExecution
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	value := runtime.GetVMInput().CallValue
	contractCallInput, err := prepareIndirectContractCallInput(
		host,
		sender,
		value,
		gasLimit,
		dest,
		function,
		args,
		gasToUse,
		true,
	)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	if isBuiltInCall(contractCallInput.Function, host) {
		return 1
	}

	_, err = host.ExecuteOnSameContext(contractCallInput)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_3_executeReadOnly
func v1_3_executeReadOnly(
	context unsafe.Pointer,
	gasLimit int64,
	addressOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := vmhost.GetVMHost(context)
	return ExecuteReadOnlyWithHost(
		host,
		gasLimit,
		addressOffset,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// ExecuteReadOnlyWithHost - executeReadOnly with host instead of pointer context
func ExecuteReadOnlyWithHost(
	host vmhost.VMHost,
	gasLimit int64,
	addressOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()

	callArgs, err := extractIndirectContractCallArgumentsWithoutValue(
		host, addressOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return ExecuteReadOnlyWithTypedArguments(
		host,
		gasLimit,
		callArgs.function,
		callArgs.dest,
		callArgs.args,
	)
}

// ExecuteReadOnlyWithTypedArguments - executeReadOnly with args already read from memory
func ExecuteReadOnlyWithTypedArguments(
	host vmhost.VMHost,
	gasLimit int64,
	function []byte,
	dest []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.ExecuteReadOnly
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	value := runtime.GetVMInput().CallValue
	contractCallInput, err := prepareIndirectContractCallInput(
		host,
		sender,
		value,
		gasLimit,
		dest,
		function,
		args,
		gasToUse,
		true,
	)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	if isBuiltInCall(contractCallInput.Function, host) {
		return 1
	}

	runtime.SetReadOnly(true)
	_, err = host.ExecuteOnSameContext(contractCallInput)
	runtime.SetReadOnly(false)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_3_createContract
func v1_3_createContract(
	context unsafe.Pointer,
	gasLimit int64,
	valueOffset int32,
	codeOffset int32,
	codeMetadataOffset int32,
	length int32,
	resultOffset int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := vmhost.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.CreateContract
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	value, err := runtime.MemLoad(valueOffset, vmhost.BalanceLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	code, err := runtime.MemLoad(codeOffset, length)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	codeMetadata, err := runtime.MemLoad(codeMetadataOffset, vmhost.CodeMetadataLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	data, actualLen, err := getArgumentsFromMemory(
		host,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(actualLen))
	metering.UseGas(gasToUse)

	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	valueAsInt := big.NewInt(0).SetBytes(value)
	newAddress, err := createContract(sender, data, valueAsInt, metering, gasLimit, code, codeMetadata, host, runtime)

	if err != nil {
		return 1
	}

	err = runtime.MemStore(resultOffset, newAddress)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_3_deployFromSourceContract
func v1_3_deployFromSourceContract(
	context unsafe.Pointer,
	gasLimit int64,
	valueOffset int32,
	sourceContractAddressOffset int32,
	codeMetadataOffset int32,
	resultAddressOffset int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := vmhost.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().BaseOpsAPICost.CreateContract
	metering.UseGas(gasToUse)

	value, err := runtime.MemLoad(valueOffset, vmhost.BalanceLen)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	sourceContractAddress, err := runtime.MemLoad(sourceContractAddressOffset, vmhost.AddressLen)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	codeMetadata, err := runtime.MemLoad(codeMetadataOffset, vmhost.CodeMetadataLen)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	data, actualLen, err := getArgumentsFromMemory(
		host,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(actualLen))
	metering.UseGas(gasToUse)

	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	newAddress, err := DeployFromSourceContractWithTypedArgs(
		host,
		sourceContractAddress,
		codeMetadata,
		big.NewInt(0).SetBytes(value),
		data,
		gasLimit,
	)

	if err != nil {
		return 1
	}

	err = runtime.MemStore(resultAddressOffset, newAddress)
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

// DeployFromSourceContractWithTypedArgs - deployFromSourceContract with args already read from memory
func DeployFromSourceContractWithTypedArgs(
	host vmhost.VMHost,
	sourceContractAddress []byte,
	codeMetadata []byte,
	value *big.Int,
	data [][]byte,
	gasLimit int64,
) ([]byte, error) {
	runtime := host.Runtime()
	metering := host.Metering()
	sender := runtime.GetSCAddress()

	blockchain := host.Blockchain()
	code, err := blockchain.GetCode(sourceContractAddress)
	if vmhost.WithFaultAndHost(host, err, runtime.BaseOpsErrorShouldFailExecution()) {
		return nil, err
	}

	return createContract(sender, data, value, metering, gasLimit, code, codeMetadata, host, runtime)
}

func createContract(
	sender []byte,
	data [][]byte,
	value *big.Int,
	metering vmhost.MeteringContext,
	gasLimit int64,
	code []byte,
	codeMetadata []byte,
	host vmhost.VMHost,
	runtime vmhost.RuntimeContext,
) ([]byte, error) {
	contractCreate := &vmcommon.ContractCreateInput{
		VMInput: vmcommon.VMInput{
			CallerAddr:  sender,
			Arguments:   data,
			CallValue:   value,
			GasPrice:    0,
			GasProvided: metering.BoundGasLimit(gasLimit),
		},
		ContractCode:         code,
		ContractCodeMetadata: codeMetadata,
	}

	return host.CreateNewContract(contractCreate)
}

//export v1_3_getNumReturnData
func v1_3_getNumReturnData(context unsafe.Pointer) int32 {
	output := vmhost.GetOutputContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetNumReturnData
	metering.UseGas(gasToUse)

	returnData := output.ReturnData()
	return int32(len(returnData))
}

//export v1_3_getReturnDataSize
func v1_3_getReturnDataSize(context unsafe.Pointer, resultID int32) int32 {
	output := vmhost.GetOutputContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetReturnDataSize
	metering.UseGas(gasToUse)

	returnData := output.ReturnData()
	if resultID >= int32(len(returnData)) {
		return 0
	}

	return int32(len(returnData[resultID]))
}

//export v1_3_getReturnData
func v1_3_getReturnData(context unsafe.Pointer, resultID int32, dataOffset int32) int32 {
	runtime := vmhost.GetRuntimeContext(context)
	output := vmhost.GetOutputContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetReturnData
	metering.UseGas(gasToUse)

	returnData := output.ReturnData()
	if resultID >= int32(len(returnData)) {
		return 0
	}

	err := runtime.MemStore(dataOffset, returnData[resultID])
	if vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution()) {
		return 0
	}

	return int32(len(returnData[resultID]))
}

//export v1_3_getOriginalTxHash
func v1_3_getOriginalTxHash(context unsafe.Pointer, dataOffset int32) {
	runtime := vmhost.GetRuntimeContext(context)
	metering := vmhost.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BaseOpsAPICost.GetBlockHash
	metering.UseGas(gasToUse)

	err := runtime.MemStore(dataOffset, runtime.GetOriginalTxHash())
	_ = vmhost.WithFault(err, context, runtime.BaseOpsErrorShouldFailExecution())
}

func prepareIndirectContractCallInput(
	host vmhost.VMHost,
	sender []byte,
	value *big.Int,
	gasLimit int64,
	destination []byte,
	function []byte,
	data [][]byte,
	gasToUse uint64,
	syncExecutionRequired bool,
) (*vmcommon.ContractCallInput, error) {
	runtime := host.Runtime()
	metering := host.Metering()

	if syncExecutionRequired && !host.AreInSameShard(runtime.GetSCAddress(), destination) {
		return nil, vmhost.ErrSyncExecutionNotInSameShard
	}

	contractCallInput := &vmcommon.ContractCallInput{
		VMInput: vmcommon.VMInput{
			CallerAddr:  sender,
			Arguments:   data,
			CallValue:   value,
			GasPrice:    0,
			GasProvided: metering.BoundGasLimit(gasLimit),
		},
		RecipientAddr: destination,
		Function:      string(function),
	}

	return contractCallInput, nil
}

func getArgumentsFromMemory(
	host vmhost.VMHost,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) ([][]byte, int32, error) {
	runtime := host.Runtime()

	if numArguments < 0 {
		return nil, 0, fmt.Errorf("negative numArguments (%d)", numArguments)
	}

	argumentsLengthData, err := runtime.MemLoad(argumentsLengthOffset, numArguments*4)
	if err != nil {
		return nil, 0, err
	}

	argumentLengths := createInt32Array(argumentsLengthData, numArguments)
	data, err := runtime.MemLoadMultiple(dataOffset, argumentLengths)
	if err != nil {
		return nil, 0, err
	}

	totalArgumentBytes := int32(0)
	for _, length := range argumentLengths {
		totalArgumentBytes += length
	}

	return data, totalArgumentBytes, nil
}

func createInt32Array(rawData []byte, numIntegers int32) []int32 {
	integers := make([]int32, numIntegers)
	index := 0
	for cursor := 0; cursor < len(rawData); cursor += 4 {
		rawInt := rawData[cursor : cursor+4]
		actualInt := binary.LittleEndian.Uint32(rawInt)
		integers[index] = int32(actualInt)
		index++
	}
	return integers
}
