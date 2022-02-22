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
	"os"
	"sync"
)

type Dht11Driver struct {
	logger		logger.LoggingClient
	asyncCh		chan<- *dsModels.AsyncValues
	vehicleDevices	sync.Map
}

var once sync.Once
var driver *Dht11Driver
var sdkService sdk.DeviceService

var raspberrypiServer = "raspberrypi-mocker:5000"

func NewDeviceDriver() *Dht11Driver {
	once.Do(func() {
		driver = new(Dht11Driver)
	})

	if os.Getenv("RASPBERRY_SERVER") != "" {
		raspberrypiServer = os.Getenv("RASPBERRY_SERVER")
	}

	fmt.Printf("raspberrypi server, env val: %s\n",raspberrypiServer)

	return driver
}

// Initialize performs protocol-specific initialization for the device
// service.
func (driver *Dht11Driver) Initialize(logger logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues, deviceCh chan<- []dsModels.DiscoveredDevice) error {
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
func (driver *Dht11Driver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	driver.logger.Debugf("HandleReadCommands, deviceName: %s, protocols: %v, reqs: %v",deviceName,protocols,reqs)

	deviceInfo,err := sdkService.GetDeviceByName(deviceName)
	if err != nil {
		driver.logger.Errorf("HandleReadCommands, getDevice fail, name: %s, err : %v",deviceName,err)
		return nil, err
	}else{
		driver.logger.Debugf("HandleReadCommands, device info, deviceName: %s, serviceName: %s, profileName: %s",deviceName,deviceInfo.ServiceName,deviceInfo.ProfileName)
	}

	res = make([]*dsModels.CommandValue, len(reqs))

	for i, req := range reqs {
		if dr, ok := sdkService.DeviceResource(deviceName, req.DeviceResourceName); ok {
			switch req.DeviceResourceName {
			case "Color":
				if color,err := callColorServer(raspberrypiServer,deviceName);err != nil {
					return nil, err
				}else {
					if cv,err := dsModels.NewCommandValue(req.DeviceResourceName,dr.Properties.ValueType,color);err != nil {
						return nil,err
					}else{
						res[i] = cv
					}
				}

			default:
				return nil, fmt.Errorf("deviceResource: %s not support",req.DeviceResourceName)
			}
		} else {
			driver.logger.Errorf("retrieveVehicleDevice fail, deviceName: %s, err: %v",deviceName,err)
			return nil, fmt.Errorf("cannot find device resource %s from device %s in cache", req.DeviceResourceName, deviceName)
		}
	}
	//
	return res, nil
}

// HandleWriteCommands passes a slice of CommandRequest struct each representing
// a ResourceOperation for a specific device resource.
// Since the commands are actuation commands, params provide parameters for the individual
// command.
func (driver *Dht11Driver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest, params []*dsModels.CommandValue) error {
	driver.logger.Debugf("HandleWriteCommands, deviceName: %s, protocols: %v, reqs: %v, params: %v",deviceName,protocols,reqs,params)

	deviceInfo,err := sdkService.GetDeviceByName(deviceName)
	if err != nil {
		driver.logger.Errorf("HandleWriteCommands, getDevice fail, name: %s, err : %v",deviceName,err)
		return err
	}else{
		driver.logger.Debugf("HandleWriteCommands, device info, deviceName: %s, serviceName: %s, profileName: %s",deviceName,deviceInfo.ServiceName,deviceInfo.ProfileName)
	}

	for i, param := range params {
		if _, ok := sdkService.DeviceResource(deviceName, param.DeviceResourceName); ok {
			switch param.DeviceResourceName {
			case "Color":
				if color,err := setDeviceColor(raspberrypiServer,deviceName,params[i].Value.(string));err != nil {
					return err
				}else{
					driver.logger.Debugf("HandleWriteCommands, deviceName: %s, response: %s",deviceName,color)
				}
			default:
				return fmt.Errorf("deviceResource: %s not support",param.DeviceResourceName)
			}
		} else {
			return fmt.Errorf("cannot find device resource %s from device %s in cache", param.DeviceResourceName, deviceName)
		}
	}

	return nil
}
// ====================================================================================================
// ====================================================================================================
func (driver *Dht11Driver) Stop(force bool) error {
	driver.logger.Debugf("VehicleDriver.go -> stop device-vehicle driver is stopping....")
	return nil
}

func (driver *Dht11Driver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) (err error) {
	driver.logger.Debugf("a new Device is added: %s", deviceName)
	return err
}

func (driver *Dht11Driver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState)(err error) {
	driver.logger.Debugf("Device %s is updated", deviceName)
	return nil
}

func (driver *Dht11Driver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) (err error) {
	driver.logger.Debugf("Device %s is removed", deviceName)
	return err
}
// ====================================================================================================
// ====================================================================================================
