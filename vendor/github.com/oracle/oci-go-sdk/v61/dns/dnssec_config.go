// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// DnssecConfig DNSSEC configuration data.
// A zone may have a maximum total of 10 DnssecKeyVersions, regardless of signing key type.
type DnssecConfig struct {

	// A read only array of KSK DnssecKeyVersions.
	KskDnssecKeyVersions []KskDnssecKeyVersion `mandatory:"false" json:"kskDnssecKeyVersions"`

	// A read only array of ZSK DnssecKeyVersions.
	ZskDnssecKeyVersions []ZskDnssecKeyVersion `mandatory:"false" json:"zskDnssecKeyVersions"`
}

func (m DnssecConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DnssecConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}