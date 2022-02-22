// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	device_color "github.com/edgexfoundry/device-color-changer"
	"github.com/edgexfoundry/device-color-changer/internal/driver"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	serviceName string = "color-changer"
)

func main() {
	fmt.Printf("version %s",device_color.Version)
	d := driver.NewColorChangerDeviceDriver()
	//startup.Bootstrap(serviceName, device_color.Version, d)
	startup.Bootstrap(serviceName, "v0.0.1", d)

}
