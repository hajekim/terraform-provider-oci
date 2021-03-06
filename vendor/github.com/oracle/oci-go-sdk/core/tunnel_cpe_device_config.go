// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TunnelCpeDeviceConfig The set of CPE configuration answers for the tunnel, which the customer provides in
// UpdateTunnelCpeDeviceConfig.
// The answers correlate to the questions that are specific to the CPE device type (see the
// `parameters` attribute of CpeDeviceShapeDetail).
// See these related operations:
//   * GetTunnelCpeDeviceConfig
//   * GetTunnelCpeDeviceConfigContent
//   * GetIpsecCpeDeviceConfigContent
//   * GetCpeDeviceConfigContent
type TunnelCpeDeviceConfig struct {
	TunnelCpeDeviceConfigParameter []CpeDeviceConfigAnswer `mandatory:"false" json:"tunnelCpeDeviceConfigParameter"`
}

func (m TunnelCpeDeviceConfig) String() string {
	return common.PointerString(m)
}
