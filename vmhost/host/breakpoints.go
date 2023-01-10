package host

import (
	"github.com/multiversx/mx-chain-vm-v1_3-go/vmhost"
)

func (host *vmHost) handleBreakpointIfAny(executionErr error) error {
	if executionErr == nil {
		return nil
	}

	runtime := host.Runtime()
	breakpointValue := runtime.GetRuntimeBreakpointValue()
	if breakpointValue != arwen.BreakpointNone {
		err := host.handleBreakpoint(breakpointValue)
		runtime.AddError(err)
		return err
	}

	log.Trace("wasmer execution error", "err", executionErr)
	return arwen.ErrExecutionFailed
}

func (host *vmHost) handleBreakpoint(breakpointValue arwen.BreakpointValue) error {
	if breakpointValue == arwen.BreakpointAsyncCall {
		return host.handleAsyncCallBreakpoint()
	}
	if breakpointValue == arwen.BreakpointExecutionFailed {
		return arwen.ErrExecutionFailed
	}
	if breakpointValue == arwen.BreakpointSignalError {
		return arwen.ErrSignalError
	}
	if breakpointValue == arwen.BreakpointOutOfGas {
		return arwen.ErrNotEnoughGas
	}

	return arwen.ErrUnhandledRuntimeBreakpoint
}
