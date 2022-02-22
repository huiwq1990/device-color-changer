package driver

import (
	"fmt"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	//"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
)

const (
	//defaultSize 해놔야함******************
	//defaultArrayValueSize = 5
)
const (
	ValueTypeCarID		= "CarID"
	ValueTypeVelocity	= "Velocity"
	ValueTypeLength		= "Length"
	ValueTypeGPS		= "GPS"
)

type vehicleDevice struct {
	resourceCarid       *resourceCarid
	resourceGPS			*resourceGPS
	resourceLength		*resourceLength
	resourceVelocity	*resourceVelocity
}



func (d *vehicleDevice) read(deviceName string, deviceResourceName string, typeName string) (*models.CommandValue, error) {
	result := &models.CommandValue{}
	switch deviceResourceName {
	case ValueTypeCarID:
		return d.resourceCarid.value(deviceName, deviceResourceName)
	/*case common.ValueTypeVelocity:
		return d.resourceVelocity.value(parm)
	case common.ValueTypeLength:
		return d.resourceLength.value(parm)
	case common.ValueTypeGPS:
		return d.resourceGPS.value(parm)
	*/
	default:
		return result, fmt.Errorf("make changge: %s", deviceResourceName)
	}
}


func (d *vehicleDevice) write(param *models.CommandValue, deviceName string) error {
	switch param.DeviceResourceName {
	case ValueTypeCarID:
		return d.resourceCarid.write(param, deviceName)

	default:
		return fmt.Errorf("VehicleDriver.HandleWriteCommands: there is no matched device resource for %s", param.String())
	}
}

func newVehicleDevice() *vehicleDevice {
	return &vehicleDevice{
		resourceCarid:      &resourceCarid{},
		resourceGPS:  		&resourceGPS{},
		resourceLength:     &resourceLength{},
		resourceVelocity:   &resourceVelocity{},
	}
}