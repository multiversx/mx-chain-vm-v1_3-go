package vmpart

import (
	"errors"

	"github.com/multiversx/mx-chain-core-go/data/esdt"
	"github.com/multiversx/mx-chain-vm-common-go"
	"github.com/multiversx/mx-chain-vm-v1_3-go/ipc/common"
)

var _ vmcommon.LegacyBlockchainHook = (*BlockchainHookGateway)(nil)

// BlockchainHookGateway forwards requests to the actual hook
type BlockchainHookGateway struct {
	messenger *VMMessenger
}

// NewBlockchainHookGateway creates a new gateway
func NewBlockchainHookGateway(messenger *VMMessenger) *BlockchainHookGateway {
	return &BlockchainHookGateway{messenger: messenger}
}

// NewAddress forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) NewAddress(creatorAddress []byte, creatorNonce uint64, vmType []byte) ([]byte, error) {

	request := common.NewMessageBlockchainNewAddressRequest(creatorAddress, creatorNonce, vmType)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil, err
	}

	if rawResponse.GetKind() != common.BlockchainNewAddressResponse {
		return nil, common.ErrBadHookResponseFromNode
	}

	response := rawResponse.(*common.MessageBlockchainNewAddressResponse)
	return response.Result, response.GetError()
}

// GetStorageData forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) GetStorageData(accountAddress []byte, index []byte) ([]byte, uint32, error) {

	request := common.NewMessageBlockchainGetStorageDataRequest(accountAddress, index)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil, 0, err
	}

	if rawResponse.GetKind() != common.BlockchainGetStorageDataResponse {
		return nil, 0, common.ErrBadHookResponseFromNode
	}

	response := rawResponse.(*common.MessageBlockchainGetStorageDataResponse)
	return response.Data, 0, response.GetError()
}

// GetBlockhash forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) GetBlockhash(nonce uint64) ([]byte, error) {

	request := common.NewMessageBlockchainGetBlockhashRequest(nonce)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil, err
	}

	if rawResponse.GetKind() != common.BlockchainGetBlockhashResponse {
		return nil, common.ErrBadHookResponseFromNode
	}

	response := rawResponse.(*common.MessageBlockchainGetBlockhashResponse)
	return response.Result, response.GetError()
}

// LastNonce forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) LastNonce() uint64 {

	request := common.NewMessageBlockchainLastNonceRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return 0
	}

	if rawResponse.GetKind() != common.BlockchainLastNonceResponse {
		return 0
	}

	response := rawResponse.(*common.MessageBlockchainLastNonceResponse)
	return response.Result
}

// LastRound forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) LastRound() uint64 {

	request := common.NewMessageBlockchainLastRoundRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return 0
	}

	if rawResponse.GetKind() != common.BlockchainLastRoundResponse {
		return 0
	}

	response := rawResponse.(*common.MessageBlockchainLastRoundResponse)
	return response.Result
}

// LastTimeStamp forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) LastTimeStamp() uint64 {

	request := common.NewMessageBlockchainLastTimeStampRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return 0
	}

	if rawResponse.GetKind() != common.BlockchainLastTimeStampResponse {
		return 0
	}

	response := rawResponse.(*common.MessageBlockchainLastTimeStampResponse)
	return response.Result
}

// LastRandomSeed forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) LastRandomSeed() []byte {

	request := common.NewMessageBlockchainLastRandomSeedRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil
	}

	if rawResponse.GetKind() != common.BlockchainLastRandomSeedResponse {
		return nil
	}

	response := rawResponse.(*common.MessageBlockchainLastRandomSeedResponse)
	return response.Result
}

// LastEpoch forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) LastEpoch() uint32 {

	request := common.NewMessageBlockchainLastEpochRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return 0
	}

	if rawResponse.GetKind() != common.BlockchainLastEpochResponse {
		return 0
	}

	response := rawResponse.(*common.MessageBlockchainLastEpochResponse)
	return response.Result
}

// GetStateRootHash forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) GetStateRootHash() []byte {

	request := common.NewMessageBlockchainGetStateRootHashRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil
	}

	if rawResponse.GetKind() != common.BlockchainGetStateRootHashResponse {
		return nil
	}

	response := rawResponse.(*common.MessageBlockchainGetStateRootHashResponse)
	return response.Result
}

// CurrentNonce forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) CurrentNonce() uint64 {

	request := common.NewMessageBlockchainCurrentNonceRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return 0
	}

	if rawResponse.GetKind() != common.BlockchainCurrentNonceResponse {
		return 0
	}

	response := rawResponse.(*common.MessageBlockchainCurrentNonceResponse)
	return response.Result
}

// CurrentRound forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) CurrentRound() uint64 {

	request := common.NewMessageBlockchainCurrentRoundRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return 0
	}

	if rawResponse.GetKind() != common.BlockchainCurrentRoundResponse {
		return 0
	}

	response := rawResponse.(*common.MessageBlockchainCurrentRoundResponse)
	return response.Result
}

// CurrentTimeStamp forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) CurrentTimeStamp() uint64 {

	request := common.NewMessageBlockchainCurrentTimeStampRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return 0
	}

	if rawResponse.GetKind() != common.BlockchainCurrentTimeStampResponse {
		return 0
	}

	response := rawResponse.(*common.MessageBlockchainCurrentTimeStampResponse)
	return response.Result
}

// CurrentRandomSeed forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) CurrentRandomSeed() []byte {

	request := common.NewMessageBlockchainCurrentRandomSeedRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil
	}

	if rawResponse.GetKind() != common.BlockchainCurrentRandomSeedResponse {
		return nil
	}

	response := rawResponse.(*common.MessageBlockchainCurrentRandomSeedResponse)
	return response.Result
}

// CurrentEpoch forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) CurrentEpoch() uint32 {

	request := common.NewMessageBlockchainCurrentEpochRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return 0
	}

	if rawResponse.GetKind() != common.BlockchainCurrentEpochResponse {
		return 0
	}

	response := rawResponse.(*common.MessageBlockchainCurrentEpochResponse)
	return response.Result
}

// ProcessBuiltInFunction forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) ProcessBuiltInFunction(input *vmcommon.ContractCallInput) (*vmcommon.VMOutput, error) {

	request := common.NewMessageBlockchainProcessBuiltInFunctionRequest(input)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil, err
	}

	if rawResponse.GetKind() != common.BlockchainProcessBuiltInFunctionResponse {
		return nil, common.ErrBadHookResponseFromNode
	}

	response := rawResponse.(*common.MessageBlockchainProcessBuiltInFunctionResponse)
	return response.VmOutput.ConvertToVMOutput(), response.GetError()
}

// GetBuiltinFunctionNames forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) GetBuiltinFunctionNames() vmcommon.FunctionNames {

	request := common.NewMessageBlockchainGetBuiltinFunctionNamesRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil
	}

	if rawResponse.GetKind() != common.BlockchainGetBuiltinFunctionNamesResponse {
		return nil
	}

	response := rawResponse.(*common.MessageBlockchainGetBuiltinFunctionNamesResponse)
	return response.Result
}

// GetAllState forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) GetAllState(address []byte) (map[string][]byte, error) {

	request := common.NewMessageBlockchainGetAllStateRequest(address)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil, err
	}

	if rawResponse.GetKind() != common.BlockchainGetAllStateResponse {
		return nil, common.ErrBadHookResponseFromNode
	}

	response := rawResponse.(*common.MessageBlockchainGetAllStateResponse)
	return response.Result.ConvertToMap(), response.GetError()
}

// GetUserAccount forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) GetUserAccount(address []byte) (vmcommon.UserAccountHandler, error) {

	request := common.NewMessageBlockchainGetUserAccountRequest(address)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil, err
	}

	if rawResponse.GetKind() != common.BlockchainGetUserAccountResponse {
		return nil, common.ErrBadHookResponseFromNode
	}

	response := rawResponse.(*common.MessageBlockchainGetUserAccountResponse)
	return response.Result, response.GetError()
}

// GetCode forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) GetCode(account vmcommon.UserAccountHandler) []byte {

	requestAccount := &common.Account{
		Nonce:           account.GetNonce(),
		Balance:         account.GetBalance(),
		CodeHash:        account.GetCodeHash(),
		RootHash:        account.GetRootHash(),
		Address:         account.AddressBytes(),
		DeveloperReward: account.GetDeveloperReward(),
		OwnerAddress:    account.GetOwnerAddress(),
		UserName:        account.GetUserName(),
		CodeMetadata:    account.GetCodeMetadata(),
	}
	request := common.NewMessageBlockchainGetCodeRequest(requestAccount)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil
	}

	if rawResponse.GetKind() != common.BlockchainGetCodeResponse {
		return nil
	}

	response := rawResponse.(*common.MessageBlockchainGetCodeResponse)
	return response.Code
}

// GetShardOfAddress forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) GetShardOfAddress(address []byte) uint32 {

	request := common.NewMessageBlockchainGetShardOfAddressRequest(address)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return 0
	}

	if rawResponse.GetKind() != common.BlockchainGetShardOfAddressResponse {
		return 0
	}

	response := rawResponse.(*common.MessageBlockchainGetShardOfAddressResponse)
	return response.Result
}

// IsSmartContract forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) IsSmartContract(address []byte) bool {

	request := common.NewMessageBlockchainIsSmartContractRequest(address)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return false
	}

	if rawResponse.GetKind() != common.BlockchainIsSmartContractResponse {
		return false
	}

	response := rawResponse.(*common.MessageBlockchainIsSmartContractResponse)
	return response.Result
}

// IsPayable forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) IsPayable(_, address []byte) (bool, error) {

	request := common.NewMessageBlockchainIsPayableRequest(address)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return false, err
	}

	if rawResponse.GetKind() != common.BlockchainIsPayableResponse {
		return false, common.ErrBadHookResponseFromNode
	}

	response := rawResponse.(*common.MessageBlockchainIsPayableResponse)
	return response.Result, response.GetError()
}

// SaveCompiledCode forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) SaveCompiledCode(codeHash []byte, code []byte) {

	request := common.NewMessageBlockchainSaveCompiledCodeRequest(codeHash, code)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return
	}

	if rawResponse.GetKind() != common.BlockchainSaveCompiledCodeResponse {
		return
	}

	return
}

// GetCompiledCode forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) GetCompiledCode(codeHash []byte) (bool, []byte) {

	request := common.NewMessageBlockchainGetCompiledCodeRequest(codeHash)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return false, nil
	}

	if rawResponse.GetKind() != common.BlockchainGetCompiledCodeResponse {
		return false, nil
	}

	response := rawResponse.(*common.MessageBlockchainGetCompiledCodeResponse)
	return response.Found, response.Code
}

// ClearCompiledCodes forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) ClearCompiledCodes() {

	request := common.NewMessageBlockchainClearCompiledCodesRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return
	}

	if rawResponse.GetKind() != common.BlockchainClearCompiledCodesResponse {
		return
	}

	return
}

// GetESDTToken forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) GetESDTToken(address []byte, tokenID []byte, nonce uint64) (*esdt.ESDigitalToken, error) {

	request := common.NewMessageBlockchainGetESDTTokenRequest(address, tokenID, nonce)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return nil, err
	}

	if rawResponse.GetKind() != common.BlockchainGetESDTTokenResponse {
		return nil, common.ErrBadHookResponseFromNode
	}

	response := rawResponse.(*common.MessageBlockchainGetESDTTokenResponse)
	return response.Result, response.GetError()
}

// IsPaused not used in v1.3
func (blockchain *BlockchainHookGateway) IsPaused(_ []byte) bool {
	return false
}

// IsLimitedTransfer not used in v1.3
func (blockchain *BlockchainHookGateway) IsLimitedTransfer(_ []byte) bool {
	return false
}

// IsInterfaceNil forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) IsInterfaceNil() bool {
	return blockchain == nil
}

// GetSnapshot forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) GetSnapshot() int {

	request := common.NewMessageBlockchainGetSnapshotRequest()
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return 0
	}

	if rawResponse.GetKind() != common.BlockchainGetSnapshotResponse {
		return 0
	}

	response := rawResponse.(*common.MessageBlockchainGetSnapshotResponse)
	return response.Result
}

// RevertToSnapshot forwards a message to the actual hook
func (blockchain *BlockchainHookGateway) RevertToSnapshot(snapshot int) error {

	request := common.NewMessageBlockchainRevertToSnapshotRequest(snapshot)
	rawResponse, err := blockchain.messenger.SendHookCallRequest(request)
	if err != nil {
		return err
	}

	if rawResponse.GetKind() != common.BlockchainRevertToSnapshotResponse {
		return common.ErrBadHookResponseFromNode
	}

	return err
}

// ExecuteSmartContractCallOnOtherVM -
func (blockchain *BlockchainHookGateway) ExecuteSmartContractCallOnOtherVM(_ *vmcommon.ContractCallInput) (*vmcommon.VMOutput, error) {
	return nil, errors.New("not implemented")
}
