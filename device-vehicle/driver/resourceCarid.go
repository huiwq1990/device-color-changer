package driver

import (
	"fmt"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
)

type resourceCarid struct{
	Id string
}


//vehicleDevice.read에서 호출 devicedriver->device->device resource
func (rc *resourceCarid) value(deviceName string, deviceResourceName string) (*models.CommandValue, error) {
	result := &models.CommandValue{}

	//var newValueInt int64
	//
	//newValueInt = 0

	result, _ = models.NewCommandValue(deviceResourceName, common.ValueTypeString, rc.Id)

	return result, nil

}

func (rc *resourceCarid) write(param *models.CommandValue, deviceName string) error {
	switch param.Type {
	case common.ValueTypeString:
		rc.Id = param.Value.(string)
		return nil
	default:
		return fmt.Errorf("resourceCarID.write: unknown device resource: %s", param.DeviceResourceName)
	}

}