// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AmdMilanBmPlatformConfig The platform configuration of a bare metal instance that uses an E4 shape.
// (the AMD Milan platform).
type AmdMilanBmPlatformConfig struct {

	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `mandatory:"false" json:"isSecureBootEnabled"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `mandatory:"false" json:"isTrustedPlatformModuleEnabled"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `mandatory:"false" json:"isMeasuredBootEnabled"`

	// The number of NUMA nodes per socket (NPS).
	NumaNodesPerSocket AmdMilanBmPlatformConfigNumaNodesPerSocketEnum `mandatory:"false" json:"numaNodesPerSocket,omitempty"`
}

//GetIsSecureBootEnabled returns IsSecureBootEnabled
func (m AmdMilanBmPlatformConfig) GetIsSecureBootEnabled() *bool {
	return m.IsSecureBootEnabled
}

//GetIsTrustedPlatformModuleEnabled returns IsTrustedPlatformModuleEnabled
func (m AmdMilanBmPlatformConfig) GetIsTrustedPlatformModuleEnabled() *bool {
	return m.IsTrustedPlatformModuleEnabled
}

//GetIsMeasuredBootEnabled returns IsMeasuredBootEnabled
func (m AmdMilanBmPlatformConfig) GetIsMeasuredBootEnabled() *bool {
	return m.IsMeasuredBootEnabled
}

func (m AmdMilanBmPlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AmdMilanBmPlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAmdMilanBmPlatformConfigNumaNodesPerSocketEnum(string(m.NumaNodesPerSocket)); !ok && m.NumaNodesPerSocket != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaNodesPerSocket: %s. Supported values are: %s.", m.NumaNodesPerSocket, strings.Join(GetAmdMilanBmPlatformConfigNumaNodesPerSocketEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AmdMilanBmPlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAmdMilanBmPlatformConfig AmdMilanBmPlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAmdMilanBmPlatformConfig
	}{
		"AMD_MILAN_BM",
		(MarshalTypeAmdMilanBmPlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// AmdMilanBmPlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type AmdMilanBmPlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for AmdMilanBmPlatformConfigNumaNodesPerSocketEnum
const (
	AmdMilanBmPlatformConfigNumaNodesPerSocketNps0 AmdMilanBmPlatformConfigNumaNodesPerSocketEnum = "NPS0"
	AmdMilanBmPlatformConfigNumaNodesPerSocketNps1 AmdMilanBmPlatformConfigNumaNodesPerSocketEnum = "NPS1"
	AmdMilanBmPlatformConfigNumaNodesPerSocketNps2 AmdMilanBmPlatformConfigNumaNodesPerSocketEnum = "NPS2"
	AmdMilanBmPlatformConfigNumaNodesPerSocketNps4 AmdMilanBmPlatformConfigNumaNodesPerSocketEnum = "NPS4"
)

var mappingAmdMilanBmPlatformConfigNumaNodesPerSocketEnum = map[string]AmdMilanBmPlatformConfigNumaNodesPerSocketEnum{
	"NPS0": AmdMilanBmPlatformConfigNumaNodesPerSocketNps0,
	"NPS1": AmdMilanBmPlatformConfigNumaNodesPerSocketNps1,
	"NPS2": AmdMilanBmPlatformConfigNumaNodesPerSocketNps2,
	"NPS4": AmdMilanBmPlatformConfigNumaNodesPerSocketNps4,
}

var mappingAmdMilanBmPlatformConfigNumaNodesPerSocketEnumLowerCase = map[string]AmdMilanBmPlatformConfigNumaNodesPerSocketEnum{
	"nps0": AmdMilanBmPlatformConfigNumaNodesPerSocketNps0,
	"nps1": AmdMilanBmPlatformConfigNumaNodesPerSocketNps1,
	"nps2": AmdMilanBmPlatformConfigNumaNodesPerSocketNps2,
	"nps4": AmdMilanBmPlatformConfigNumaNodesPerSocketNps4,
}

// GetAmdMilanBmPlatformConfigNumaNodesPerSocketEnumValues Enumerates the set of values for AmdMilanBmPlatformConfigNumaNodesPerSocketEnum
func GetAmdMilanBmPlatformConfigNumaNodesPerSocketEnumValues() []AmdMilanBmPlatformConfigNumaNodesPerSocketEnum {
	values := make([]AmdMilanBmPlatformConfigNumaNodesPerSocketEnum, 0)
	for _, v := range mappingAmdMilanBmPlatformConfigNumaNodesPerSocketEnum {
		values = append(values, v)
	}
	return values
}

// GetAmdMilanBmPlatformConfigNumaNodesPerSocketEnumStringValues Enumerates the set of values in String for AmdMilanBmPlatformConfigNumaNodesPerSocketEnum
func GetAmdMilanBmPlatformConfigNumaNodesPerSocketEnumStringValues() []string {
	return []string{
		"NPS0",
		"NPS1",
		"NPS2",
		"NPS4",
	}
}

// GetMappingAmdMilanBmPlatformConfigNumaNodesPerSocketEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAmdMilanBmPlatformConfigNumaNodesPerSocketEnum(val string) (AmdMilanBmPlatformConfigNumaNodesPerSocketEnum, bool) {
	enum, ok := mappingAmdMilanBmPlatformConfigNumaNodesPerSocketEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
