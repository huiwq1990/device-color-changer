//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package driver

import (
	"fmt"
	dsModels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	sdk "github.com/edgexfoundry/device-sdk-go/v2/pkg/service"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"reflect"
	"sync"
)

type VehicleDriver struct {
	logger		logger.LoggingClient
	asyncCh		chan<- *dsModels.AsyncValues
	vehicleDevices	sync.Map
}

var once sync.Once
var driver *VehicleDriver
var sdkService sdk.DeviceService


func NewVehicleDeviceDriver() *VehicleDriver {
	once.Do(func() {
		driver = new(VehicleDriver)
	})

	return driver
}

func (driver *VehicleDriver) retrieveVehicleDevice(deviceName string) (vdv *vehicleDevice, err error) {
	//return vdv, err

	vd, ok := driver.vehicleDevices.LoadOrStore(deviceName, newVehicleDevice())
	if vdv, ok = vd.(*vehicleDevice); !ok {
		driver.logger.Errorf("retrieve vehicleDevice by name: %s, the returned value has to be a reference of "+
			"vehicleDevice struct, but got: %s", deviceName, reflect.TypeOf(vd))
	}
	return vdv, err
}

// Initialize performs protocol-specific initialization for the device
// service.
func (driver *VehicleDriver) Initialize(logger logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues, deviceCh chan<- []dsModels.DiscoveredDevice) error {
	driver.logger = logger
	driver.asyncCh = asyncCh

	//service := sdk.RunningService()
	//devices := service.Devices()
	//handler := NewRestHandler(sdk.RunningService(), logger, asyncCh)
	//return handler.Start()
	return nil
}

// ====================================================================================================
// ====================================================================================================
// HandleReadCommands triggers a protocol Read operation for the specified device.
func (driver *VehicleDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	driver.logger.Debugf("HandleReadCommands, deviceName: %s, protocols: %v, reqs: %v",deviceName,protocols,reqs)

	//param print
	//readMap,_ := json.MarshalIndent(protocols, "", "  ")
	//fmt.Print(string(readMap))


	vd, err := driver.retrieveVehicleDevice(deviceName)
	if err != nil {
		driver.logger.Errorf("retrieveVehicleDevice fail, deviceName: %s, err: %v",deviceName,err)
		return nil, err
	}else{
		driver.logger.Infof("retrieveVehicleDevice success, deviceName: %s",deviceName)
	}

	res = make([]*dsModels.CommandValue, len(reqs))

	for i, req := range reqs {
		if dr, ok := sdkService.DeviceResource(deviceName, req.DeviceResourceName); ok {
			if v, err := vd.read(deviceName, req.DeviceResourceName, dr.Properties.ValueType); err != nil {
				driver.logger.Errorf("HandleReadCommands vehicle device read fail, deviceName: %s, err: %v",deviceName,err)
				return nil, err
			} else {
				res[i] = v
			}
		} else {
			driver.logger.Errorf("retrieveVehicleDevice fail, deviceName: %s, err: %v",deviceName,err)
			return nil, fmt.Errorf("cannot find device resource %s from device %s in cache", req.DeviceResourceName, deviceName)
		}
	}

	return res, nil
}

// HandleWriteCommands passes a slice of CommandRequest struct each representing
// a ResourceOperation for a specific device resource.
// Since the commands are actuation commands, params provide parameters for the individual
// command.
func (driver *VehicleDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest, params []*dsModels.CommandValue) error {
	driver.logger.Debugf("HandleWriteCommands, deviceName: %s, protocols: %v, reqs: %v, params: %v",deviceName,protocols,reqs,params)

	vd, err := driver.retrieveVehicleDevice(deviceName)
	if err != nil {
		driver.logger.Errorf("retrieveVehicleDevice shipe")
		return err
	}else{
		driver.logger.Infof("retrieveVehicleDevice succ, deviceName: %s",deviceName)
	}

	for _, param := range params {
		if err := vd.write(param, deviceName); err != nil {
			driver.logger.Errorf("HandleWriteCommands write params fail, deviceName: %s, err: %v",deviceName,err)
			return err
		}
	}
	return nil
}
// ====================================================================================================
// ====================================================================================================
func (driver *VehicleDriver) Stop(force bool) error {
	driver.logger.Debugf("VehicleDriver.go -> stop device-vehicle driver is stopping....")
	return nil
}

func (driver *VehicleDriver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) (err error) {
	driver.logger.Debugf("a new Device is added: %s", deviceName)
	return err
}

func (driver *VehicleDriver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState)(err error) {
	driver.logger.Debugf("Device %s is updated", deviceName)
	return nil
}

func (driver *VehicleDriver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) (err error) {
	driver.logger.Debugf("Device %s is removed", deviceName)
	return err
}
// ====================================================================================================
// ====================================================================================================
