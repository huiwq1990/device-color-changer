// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	device_virtual "github.com/edgexfoundry/device-color-changer"
	"github.com/edgexfoundry/device-color-changer/internal/driver"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	serviceName string = "color-changer"
)

func main() {
	d := driver.NewColorChangerDeviceDriver()
	startup.Bootstrap(serviceName, device_virtual.Version, d)
}