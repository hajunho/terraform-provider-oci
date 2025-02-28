// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ValuePhoneNumber Phone number field value.
type ValuePhoneNumber struct {

	// Confidence score between 0 to 1.
	Confidence *float32 `mandatory:"true" json:"confidence"`

	BoundingPolygon *BoundingPolygon `mandatory:"true" json:"boundingPolygon"`

	// Indexes of the words in the field value.
	WordIndexes []int `mandatory:"true" json:"wordIndexes"`

	// Phone number field value.
	Value *string `mandatory:"true" json:"value"`

	// Detected text of a field.
	Text *string `mandatory:"false" json:"text"`
}

//GetText returns Text
func (m ValuePhoneNumber) GetText() *string {
	return m.Text
}

//GetConfidence returns Confidence
func (m ValuePhoneNumber) GetConfidence() *float32 {
	return m.Confidence
}

//GetBoundingPolygon returns BoundingPolygon
func (m ValuePhoneNumber) GetBoundingPolygon() *BoundingPolygon {
	return m.BoundingPolygon
}

//GetWordIndexes returns WordIndexes
func (m ValuePhoneNumber) GetWordIndexes() []int {
	return m.WordIndexes
}

func (m ValuePhoneNumber) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValuePhoneNumber) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ValuePhoneNumber) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeValuePhoneNumber ValuePhoneNumber
	s := struct {
		DiscriminatorParam string `json:"valueType"`
		MarshalTypeValuePhoneNumber
	}{
		"PHONE_NUMBER",
		(MarshalTypeValuePhoneNumber)(m),
	}

	return json.Marshal(&s)
}
