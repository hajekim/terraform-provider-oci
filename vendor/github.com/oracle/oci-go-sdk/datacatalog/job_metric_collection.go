// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// JobMetricCollection Results of a job metrics listing. Job metrics are datum about a job execution in key value pairs.
type JobMetricCollection struct {

	// Collection of job metrics.
	Items []JobMetricSummary `mandatory:"true" json:"items"`
}

func (m JobMetricCollection) String() string {
	return common.PointerString(m)
}
