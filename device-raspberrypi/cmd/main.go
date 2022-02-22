// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/startup"
	device_raspirest "github.com/huiwq1990/device-raspberrypi"
	"github.com/huiwq1990/device-raspberrypi/internal/driver"

)

const (
	serviceName string = "device-raspberrypi"
)

func main() {
	d := driver.NewDeviceDriver()
	startup.Bootstrap(serviceName, device_raspirest.Version, d)
}
