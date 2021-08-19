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

// DbSystemSnapshot Snapshot of the DbSystem details at the time of the backup
type DbSystemSnapshot struct {

	// The OCID of the DB System.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the DB System. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment the DB System belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the subnet the DB System is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Name of the MySQL Version in use for the DB System.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// Initial size of the data volume in GiBs that will be created and attached.
	DataStorageSizeInGBs *int `mandatory:"true" json:"dataStorageSizeInGBs"`

	Maintenance *MaintenanceDetails `mandatory:"true" json:"maintenance"`

	// User-provided data about the DB System.
	Description *string `mandatory:"false" json:"description"`

	// The Availability Domain where the primary DB System should be located.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The name of the Fault Domain the DB System is located in.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The shape of the primary instances of the DB System. The shape
	// determines resources allocated to a DB System - CPU cores
	// and memory for VM shapes; CPU cores, memory and storage for non-VM
	// (or bare metal) shapes. To get a list of shapes, use (the
	// ListShapes operation.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// The username for the administrative user.
	AdminUsername *string `mandatory:"false" json:"adminUsername"`

	BackupPolicy *BackupPolicy `mandatory:"false" json:"backupPolicy"`

	// The OCID of the Configuration to be used for Instances in this DB System.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// The hostname for the primary endpoint of the DB System. Used for DNS.
	// The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN)
	// (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com").
	// Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// The IP address the DB System is configured to listen on. A private
	// IP address of the primary endpoint of the DB System. Must be an
	// available IP address within the subnet's CIDR. This will be a
	// "dotted-quad" style IPv4 address.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The port for primary endpoint of the DB System to listen on.
	Port *int `mandatory:"false" json:"port"`

	// The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port.
	PortX *int `mandatory:"false" json:"portX"`

	// If the policy is to enable high availability of the instance, by
	// maintaining secondary/failover capacity as necessary.
	IsHighlyAvailable *bool `mandatory:"false" json:"isHighlyAvailable"`

	// The network endpoints available for this DB System.
	Endpoints []DbSystemEndpoint `mandatory:"false" json:"endpoints"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DbSystemSnapshot) String() string {
	return common.PointerString(m)
}
