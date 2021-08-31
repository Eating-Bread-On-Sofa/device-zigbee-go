// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/edgexfoundry/device-zigbee-go"
	"github.com/edgexfoundry/device-zigbee-go/driver"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	serviceName string = "edgex-device-zigbee"
)

func main() {
	d := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, device_zigbee.Version, d)
}
