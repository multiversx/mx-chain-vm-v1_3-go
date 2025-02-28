package nodepart

import (
	"fmt"
	"os"
	"time"

	"github.com/multiversx/mx-chain-vm-common-go"
	"github.com/multiversx/mx-chain-vm-v1_3-go/ipc/common"
	"github.com/multiversx/mx-chain-vm-v1_3-go/ipc/marshaling"
)

// NodePart is the endpoint that implements the message loop on Node's side
type NodePart struct {
	Messenger  *NodeMessenger
	blockchain vmcommon.LegacyBlockchainHook
	Repliers   []common.MessageReplier
	config     Config
}

// NewNodePart creates the Node part
func NewNodePart(
	input *os.File,
	output *os.File,
	blockchain vmcommon.LegacyBlockchainHook,
	config Config,
	marshalizer marshaling.Marshalizer,
) (*NodePart, error) {
	messenger := NewNodeMessenger(input, output, marshalizer)

	part := &NodePart{
		Messenger:  messenger,
		blockchain: blockchain,
		config:     config,
	}

	part.Repliers = common.CreateReplySlots(part.noopReplier)
	part.Repliers[common.BlockchainNewAddressRequest] = part.replyToBlockchainNewAddress
	part.Repliers[common.BlockchainGetStorageDataRequest] = part.replyToBlockchainGetStorageData
	part.Repliers[common.BlockchainGetBlockhashRequest] = part.replyToBlockchainGetBlockhash
	part.Repliers[common.BlockchainLastNonceRequest] = part.replyToBlockchainLastNonce
	part.Repliers[common.BlockchainLastRoundRequest] = part.replyToBlockchainLastRound
	part.Repliers[common.BlockchainLastTimeStampRequest] = part.replyToBlockchainLastTimeStamp
	part.Repliers[common.BlockchainLastRandomSeedRequest] = part.replyToBlockchainLastRandomSeed
	part.Repliers[common.BlockchainLastEpochRequest] = part.replyToBlockchainLastEpoch
	part.Repliers[common.BlockchainGetStateRootHashRequest] = part.replyToBlockchainGetStateRootHash
	part.Repliers[common.BlockchainCurrentNonceRequest] = part.replyToBlockchainCurrentNonce
	part.Repliers[common.BlockchainCurrentRoundRequest] = part.replyToBlockchainCurrentRound
	part.Repliers[common.BlockchainCurrentTimeStampRequest] = part.replyToBlockchainCurrentTimeStamp
	part.Repliers[common.BlockchainCurrentRandomSeedRequest] = part.replyToBlockchainCurrentRandomSeed
	part.Repliers[common.BlockchainCurrentEpochRequest] = part.replyToBlockchainCurrentEpoch
	part.Repliers[common.BlockchainProcessBuiltInFunctionRequest] = part.replyToBlockchainProcessBuiltInFunction
	part.Repliers[common.BlockchainGetBuiltinFunctionNamesRequest] = part.replyToBlockchainGetBuiltinFunctionNames
	part.Repliers[common.BlockchainGetAllStateRequest] = part.replyToBlockchainGetAllState
	part.Repliers[common.BlockchainGetUserAccountRequest] = part.replyToBlockchainGetUserAccount
	part.Repliers[common.BlockchainGetCodeRequest] = part.replyToBlockchainGetCode
	part.Repliers[common.BlockchainGetShardOfAddressRequest] = part.replyToBlockchainGetShardOfAddress
	part.Repliers[common.BlockchainIsSmartContractRequest] = part.replyToBlockchainIsSmartContract
	part.Repliers[common.BlockchainIsPayableRequest] = part.replyToBlockchainIsPayable
	part.Repliers[common.BlockchainSaveCompiledCodeRequest] = part.replyToBlockchainSaveCompiledCode
	part.Repliers[common.BlockchainGetCompiledCodeRequest] = part.replyToBlockchainGetCompiledCode
	part.Repliers[common.BlockchainClearCompiledCodesRequest] = part.replyToBlockchainClearCompiledCodes
	part.Repliers[common.BlockchainGetESDTTokenRequest] = part.replyToBlockchainGetESDTToken
	part.Repliers[common.BlockchainIsInterfaceNilRequest] = part.replyToBlockchainIsInterfaceNil
	part.Repliers[common.BlockchainGetSnapshotRequest] = part.replyToBlockchainGetSnapshot
	part.Repliers[common.BlockchainRevertToSnapshotRequest] = part.replyToBlockchainRevertToSnapshot

	return part, nil
}

func (part *NodePart) noopReplier(_ common.MessageHandler) common.MessageHandler {
	log.Error("noopReplier called")
	return common.CreateMessage(common.UndefinedRequestOrResponse)
}

// StartLoop runs the main loop
func (part *NodePart) StartLoop(request common.MessageHandler) (common.MessageHandler, error) {
	defer part.timeTrack(time.Now(), "[NODE] end of loop")

	err := part.Messenger.SendContractRequest(request)
	if err != nil {
		return nil, err
	}

	response, err := part.doLoop()
	if err != nil {
		log.Warn("[NODE]: end of loop", "err", err)
	}

	part.Messenger.ResetDialogue()
	return response, err
}

// doLoop ends when processing the transaction ends or in the case of a critical failure
// Critical failure = VM timeouts or crashes
// The error result is set only in case of critical failure
func (part *NodePart) doLoop() (common.MessageHandler, error) {
	remainingMilliseconds := part.config.MaxLoopTime

	for {
		message, duration, err := part.Messenger.ReceiveHookCallRequestOrContractResponse(remainingMilliseconds)
		if err != nil {
			return nil, err
		}

		remainingMilliseconds -= duration
		if remainingMilliseconds < 0 {
			return nil, common.ErrVMTimeExpired
		}

		if common.IsHookCall(message) {
			err := part.replyToHookCallRequest(message)
			if err != nil {
				return nil, err
			}

			continue
		}

		if common.IsVersionResponse(message) {
			return message, nil
		}
		if common.IsContractResponse(message) {
			return message, nil
		}
		if common.IsDiagnose(message) {
			return message, nil
		}
		if common.IsGasScheduleChangeResponse(message) {
			return message, nil
		}

		return nil, common.ErrBadMessageFromVM
	}
}

func (part *NodePart) replyToHookCallRequest(request common.MessageHandler) error {
	defer part.timeTrack(time.Now(), fmt.Sprintf("replyToHookCallRequest %s", request.GetKindName()))

	replier := part.Repliers[request.GetKind()]
	hookResponse := replier(request)
	err := part.Messenger.SendHookCallResponse(hookResponse)
	return err
}

// SendStopSignal sends a stop signal to VM
// Should only be used for tests!
func (part *NodePart) SendStopSignal() error {
	request := common.NewMessageStop()

	err := part.Messenger.Send(request)
	if err != nil {
		return err
	}

	log.Warn("Node sent stop signal to VM.")
	return nil
}

func (part *NodePart) timeTrack(start time.Time, message string) {
	elapsed := time.Since(start)
	log.Trace(message, "duration", elapsed)
}
