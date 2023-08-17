package scenarioexec

import (
	"fmt"

	"github.com/multiversx/mx-chain-core-go/core"
	logger "github.com/multiversx/mx-chain-logger-go"
	vmi "github.com/multiversx/mx-chain-vm-common-go"
	"github.com/multiversx/mx-chain-vm-v1_3-go/config"
	"github.com/multiversx/mx-chain-vm-v1_3-go/mock"
	worldhook "github.com/multiversx/mx-chain-vm-v1_3-go/mock/world"
	gasSchedules "github.com/multiversx/mx-chain-vm-v1_3-go/scenarioexec/gasSchedules"
	mc "github.com/multiversx/mx-chain-vm-v1_3-go/scenarios/controller"
	er "github.com/multiversx/mx-chain-vm-v1_3-go/scenarios/expression/reconstructor"
	fr "github.com/multiversx/mx-chain-vm-v1_3-go/scenarios/fileresolver"
	mj "github.com/multiversx/mx-chain-vm-v1_3-go/scenarios/json/model"
	"github.com/multiversx/mx-chain-vm-v1_3-go/vmhost"
	"github.com/multiversx/mx-chain-vm-v1_3-go/vmhost/hostCore"
)

var log = logger.GetOrCreate("vm/scenarios")

// TestVMType is the VM type argument we use in tests.
var TestVMType = []byte{0, 0}

// VMTestExecutor parses, interprets and executes both .test.json tests and .scen.json scenarios with VM.
type VMTestExecutor struct {
	World                 *worldhook.MockWorld
	vm                    vmi.VMExecutionHandler
	checkGas              bool
	scenGasScheduleLoaded bool
	fileResolver          fr.FileResolver
	exprReconstructor     er.ExprReconstructor
}

var _ mc.TestExecutor = (*VMTestExecutor)(nil)
var _ mc.ScenarioExecutor = (*VMTestExecutor)(nil)

// NewVMTestExecutor prepares a new VMTestExecutor instance.
func NewVMTestExecutor() (*VMTestExecutor, error) {
	world := worldhook.NewMockWorld()

	gasScheduleMap := config.MakeGasMapForTests()
	err := world.InitBuiltinFunctions(gasScheduleMap)
	if err != nil {
		return nil, err
	}

	blockGasLimit := uint64(10000000)
	vm, err := hostCore.NewVMHost(world, &vmhost.VMHostParameters{
		VMType:               TestVMType,
		BlockGasLimit:        blockGasLimit,
		GasSchedule:          gasScheduleMap,
		BuiltInFuncContainer: world.BuiltinFuncs.Container,
		ProtectedKeyPrefix:   []byte(core.ProtectedKeyPrefix),
		EnableEpochsHandler: &mock.EnableEpochsHandlerStub{
			IsFlagEnabledInCurrentEpochCalled: func(flag core.EnableEpochFlag) bool {
				return flag == core.SCDeployFlag || flag == core.AheadOfTimeGasUsageFlag || flag == core.RepairCallbackFlag || flag == core.BuiltInFunctionsFlag
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &VMTestExecutor{
		World:                 world,
		vm:                    vm,
		checkGas:              true,
		scenGasScheduleLoaded: false,
		fileResolver:          nil,
		exprReconstructor:     er.ExprReconstructor{},
	}, nil
}

// GetVM yields a reference to the VMExecutionHandler used.
func (ae *VMTestExecutor) GetVM() vmi.VMExecutionHandler {
	return ae.vm
}

func (ae *VMTestExecutor) gasScheduleMapFromScenarios(scenGasSchedule mj.GasSchedule) (config.GasScheduleMap, error) {
	switch scenGasSchedule {
	case mj.GasScheduleDefault:
		return gasSchedules.LoadGasScheduleConfig(gasSchedules.GetV3())
	case mj.GasScheduleDummy:
		return config.MakeGasMapForTests(), nil
	case mj.GasScheduleV1:
		return gasSchedules.LoadGasScheduleConfig(gasSchedules.GetV1())
	case mj.GasScheduleV2:
		return gasSchedules.LoadGasScheduleConfig(gasSchedules.GetV2())
	case mj.GasScheduleV3:
		return gasSchedules.LoadGasScheduleConfig(gasSchedules.GetV3())
	default:
		return nil, fmt.Errorf("unknown scenario GasSchedule: %d", scenGasSchedule)
	}
}

// SetScenariosGasSchedule updates the gas costs based on the scenario config
// only changes the gas schedule once,
// this prevents subsequent gasSchedule declarations in externalSteps to overwrite
func (ae *VMTestExecutor) SetScenariosGasSchedule(newGasSchedule mj.GasSchedule) error {
	if ae.scenGasScheduleLoaded {
		return nil
	}
	gasSchedule, err := ae.gasScheduleMapFromScenarios(newGasSchedule)
	if err != nil {
		return err
	}
	ae.scenGasScheduleLoaded = true
	ae.vm.GasScheduleChange(gasSchedule)
	return nil
}
