// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// DatasetCollection Results of a dataset list operation. Contains DatasetSummary items.
type DatasetCollection struct {

	// List of datasets.
	Items []DatasetSummary `mandatory:"true" json:"items"`
}

func (m DatasetCollection) String() string {
	return common.PointerString(m)
}