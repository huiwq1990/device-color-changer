package driver

import ("fmt"
	"github.com/edgexfoundry/device-color-changer/internal/driver/temperaturesensor"
	"sync"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	sdk "github.com/edgexfoundry/device-sdk-go/pkg/service"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/models"

	_ "modernc.org/ql/driver"
)

// Driver 是对device的操作
type VirtualDriver struct {
	lc      logger.LoggingClient
	asyncCh chan<- *dsModels.AsyncValues
	locker         sync.Mutex
	temperatureSensor *temperaturesensor.TemperatureSensor
}


var once sync.Once
var driver *VirtualDriver
var sdkService sdk.DeviceService


func NewColorChangerDeviceDriver() dsModels.ProtocolDriver {
	once.Do(func() {
		driver = new(VirtualDriver)
	})
	return driver
}


func (d *VirtualDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues, deviceCh chan<- []dsModels.DiscoveredDevice) error {
	d.lc = lc
	d.asyncCh = asyncCh


	//service := sdk.RunningService()
	//devices := service.Devices()
	//for _, device := range devices {
	//	err := prepareVirtualResources(d, device.Name)
	//	if err != nil {
	//		return fmt.Errorf("prepare virtual resources failed: %v", err)
	//	}
	//}

	return nil
}

func (d *VirtualDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	d.locker.Lock()
	defer func() {
		d.locker.Unlock()
	}()


	return res, nil
}

func (d *VirtualDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest,
	params []*dsModels.CommandValue) error {

	return nil
}

func (d *VirtualDriver) Stop(force bool) error {
	d.lc.Info("VirtualDriver.Stop: device-virtual driver is stopping...")

	return nil
}

func (d *VirtualDriver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.lc.Debug(fmt.Sprintf("a new Device is added: %s", deviceName))
	return nil
}

func (d *VirtualDriver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.lc.Debug(fmt.Sprintf("Device %s is updated", deviceName))

	return nil
}

func (d *VirtualDriver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	d.lc.Debug(fmt.Sprintf("Device %s is removed", deviceName))
	return nil
}
