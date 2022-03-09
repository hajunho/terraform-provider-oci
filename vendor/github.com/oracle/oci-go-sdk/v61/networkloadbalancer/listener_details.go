// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// ListenerDetails The listener's configuration.
// For more information about backend set configuration, see
// Managing Load Balancer Listeners (https://docs.cloud.oracle.com/Content/Balance/Tasks/managinglisteners.htm).
type ListenerDetails struct {

	// A friendly name for the listener. It must be unique and it cannot be changed.
	// Example: `example_listener`
	Name *string `mandatory:"true" json:"name"`

	// The name of the associated backend set.
	// Example: `example_backend_set`
	DefaultBackendSetName *string `mandatory:"true" json:"defaultBackendSetName"`

	// The communication port for the listener.
	// Example: `80`
	Port *int `mandatory:"true" json:"port"`

	// The protocol on which the listener accepts connection requests.
	// For public network load balancers, ANY protocol refers to TCP/UDP with the wildcard port.
	// For private network load balancers, ANY protocol refers to TCP/UDP/ICMP (note that ICMP requires isPreserveSourceDestination to be set to true).
	// "ListNetworkLoadBalancersProtocols" API is deprecated and it will not return the updated values. Use the allowed values for the protocol instead.
	// Example: `TCP`
	Protocol ListenerProtocolsEnum `mandatory:"true" json:"protocol"`

	// IP version associated with the listener.
	IpVersion IpVersionEnum `mandatory:"false" json:"ipVersion,omitempty"`

	// Property to enable/disable PPv2 feature for this listener.
	IsPpv2Enabled *bool `mandatory:"false" json:"isPpv2Enabled"`

	// An array that represents the PPV2 Options that can be enabled on TCP Listeners.
	// Example: ["VCN_ID"]
	InternalProxyProtocolOptions []MetadataOptionsEnum `mandatory:"false" json:"internalProxyProtocolOptions"`

	// Override to use 0xE1 custom TLV for encoding Class E IP Address in IP Options. (Default is 0xE2)
	IsSgwNatIpTlvTypeOverrideEnabled *bool `mandatory:"false" json:"isSgwNatIpTlvTypeOverrideEnabled"`
}

func (m ListenerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListenerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListenerProtocolsEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetListenerProtocolsEnumStringValues(), ",")))
	}

	if _, ok := GetMappingIpVersionEnum(string(m.IpVersion)); !ok && m.IpVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IpVersion: %s. Supported values are: %s.", m.IpVersion, strings.Join(GetIpVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
