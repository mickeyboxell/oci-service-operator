// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// AnalyticsClusterMemoryEstimate DEPRECATED -- please use HeatWave API instead.
// Analytics Cluster memory estimate
// that can be used to determine a suitable Analytics Cluster size. For each MySQL user table the estimated memory
// footprint when the table is loaded to the Analytics Cluster memory is returned.
type AnalyticsClusterMemoryEstimate struct {

	// The OCID of the DB System the Analytics Cluster memory estimate is associated with.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// Current status of the Work Request generating the Analytics Cluster memory estimate.
	Status AnalyticsClusterMemoryEstimateStatusEnum `mandatory:"true" json:"status"`

	// The date and time that the Work Request to generate the Analytics Cluster memory estimate was issued, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc333).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time that the Analytics Cluster memory estimate was generated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc333).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Collection of schemas with estimated memory footprints for MySQL user tables of each schema
	// when loaded to Analytics Cluster memory.
	TableSchemas []AnalyticsClusterSchemaMemoryEstimate `mandatory:"true" json:"tableSchemas"`
}

func (m AnalyticsClusterMemoryEstimate) String() string {
	return common.PointerString(m)
}
