// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// AutonomousDatabaseDataguardAssociation The properties that define dataguard association between two different Autonomous Databases.
// Note that Autonomous Databases inherit DataGuard association from parent Autonomous Container Database.
// No actions can be taken on AutonomousDatabaseDataguardAssociation, usage is strictly informational.
type AutonomousDatabaseDataguardAssociation struct {

	// The OCID of the Autonomous Dataguard created for Autonomous Container Database where given Autonomous Database resides in.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database that has a relationship with the peer Autonomous Database.
	AutonomousDatabaseId *string `mandatory:"true" json:"autonomousDatabaseId"`

	// The Data Guard role of the Autonomous Container Database, if Autonomous Data Guard is enabled.
	Role AutonomousDatabaseDataguardAssociationRoleEnum `mandatory:"true" json:"role"`

	// The current state of Autonomous Data Guard.
	LifecycleState AutonomousDatabaseDataguardAssociationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Data Guard role of the Autonomous Container Database, if Autonomous Data Guard is enabled.
	PeerRole AutonomousDatabaseDataguardAssociationPeerRoleEnum `mandatory:"true" json:"peerRole"`

	// Additional information about the current lifecycleState, if available.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer Autonomous Database.
	PeerAutonomousDatabaseId *string `mandatory:"false" json:"peerAutonomousDatabaseId"`

	// The current state of Autonomous Data Guard.
	PeerAutonomousDatabaseLifeCycleState AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum `mandatory:"false" json:"peerAutonomousDatabaseLifeCycleState,omitempty"`

	// The protection mode of this Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode AutonomousDatabaseDataguardAssociationProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The lag time between updates to the primary database and application of the redo data on the standby database,
	// as computed by the reporting database.
	// Example: `9 seconds`
	ApplyLag *string `mandatory:"false" json:"applyLag"`

	// The rate at which redo logs are synced between the associated databases.
	// Example: `180 Mb per second`
	ApplyRate *string `mandatory:"false" json:"applyRate"`

	// The approximate number of seconds of redo data not yet available on the standby Autonomous Container Database,
	// as computed by the reporting database.
	// Example: `7 seconds`
	TransportLag *string `mandatory:"false" json:"transportLag"`

	// The date and time of the last update to the apply lag, apply rate, and transport lag values.
	TimeLastSynced *common.SDKTime `mandatory:"false" json:"timeLastSynced"`

	// The date and time the Data Guard association was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when the last role change action happened.
	TimeLastRoleChanged *common.SDKTime `mandatory:"false" json:"timeLastRoleChanged"`
}

func (m AutonomousDatabaseDataguardAssociation) String() string {
	return common.PointerString(m)
}

// AutonomousDatabaseDataguardAssociationRoleEnum Enum with underlying type: string
type AutonomousDatabaseDataguardAssociationRoleEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataguardAssociationRoleEnum
const (
	AutonomousDatabaseDataguardAssociationRolePrimary         AutonomousDatabaseDataguardAssociationRoleEnum = "PRIMARY"
	AutonomousDatabaseDataguardAssociationRoleStandby         AutonomousDatabaseDataguardAssociationRoleEnum = "STANDBY"
	AutonomousDatabaseDataguardAssociationRoleDisabledStandby AutonomousDatabaseDataguardAssociationRoleEnum = "DISABLED_STANDBY"
)

var mappingAutonomousDatabaseDataguardAssociationRole = map[string]AutonomousDatabaseDataguardAssociationRoleEnum{
	"PRIMARY":          AutonomousDatabaseDataguardAssociationRolePrimary,
	"STANDBY":          AutonomousDatabaseDataguardAssociationRoleStandby,
	"DISABLED_STANDBY": AutonomousDatabaseDataguardAssociationRoleDisabledStandby,
}

// GetAutonomousDatabaseDataguardAssociationRoleEnumValues Enumerates the set of values for AutonomousDatabaseDataguardAssociationRoleEnum
func GetAutonomousDatabaseDataguardAssociationRoleEnumValues() []AutonomousDatabaseDataguardAssociationRoleEnum {
	values := make([]AutonomousDatabaseDataguardAssociationRoleEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataguardAssociationRole {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseDataguardAssociationLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseDataguardAssociationLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataguardAssociationLifecycleStateEnum
const (
	AutonomousDatabaseDataguardAssociationLifecycleStateProvisioning         AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseDataguardAssociationLifecycleStateAvailable            AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseDataguardAssociationLifecycleStateRoleChangeInProgress AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseDataguardAssociationLifecycleStateTerminating          AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseDataguardAssociationLifecycleStateTerminated           AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseDataguardAssociationLifecycleStateFailed               AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "FAILED"
	AutonomousDatabaseDataguardAssociationLifecycleStateUnavailable          AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "UNAVAILABLE"
)

var mappingAutonomousDatabaseDataguardAssociationLifecycleState = map[string]AutonomousDatabaseDataguardAssociationLifecycleStateEnum{
	"PROVISIONING":            AutonomousDatabaseDataguardAssociationLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousDatabaseDataguardAssociationLifecycleStateAvailable,
	"ROLE_CHANGE_IN_PROGRESS": AutonomousDatabaseDataguardAssociationLifecycleStateRoleChangeInProgress,
	"TERMINATING":             AutonomousDatabaseDataguardAssociationLifecycleStateTerminating,
	"TERMINATED":              AutonomousDatabaseDataguardAssociationLifecycleStateTerminated,
	"FAILED":                  AutonomousDatabaseDataguardAssociationLifecycleStateFailed,
	"UNAVAILABLE":             AutonomousDatabaseDataguardAssociationLifecycleStateUnavailable,
}

// GetAutonomousDatabaseDataguardAssociationLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseDataguardAssociationLifecycleStateEnum
func GetAutonomousDatabaseDataguardAssociationLifecycleStateEnumValues() []AutonomousDatabaseDataguardAssociationLifecycleStateEnum {
	values := make([]AutonomousDatabaseDataguardAssociationLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataguardAssociationLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseDataguardAssociationPeerRoleEnum Enum with underlying type: string
type AutonomousDatabaseDataguardAssociationPeerRoleEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataguardAssociationPeerRoleEnum
const (
	AutonomousDatabaseDataguardAssociationPeerRolePrimary         AutonomousDatabaseDataguardAssociationPeerRoleEnum = "PRIMARY"
	AutonomousDatabaseDataguardAssociationPeerRoleStandby         AutonomousDatabaseDataguardAssociationPeerRoleEnum = "STANDBY"
	AutonomousDatabaseDataguardAssociationPeerRoleDisabledStandby AutonomousDatabaseDataguardAssociationPeerRoleEnum = "DISABLED_STANDBY"
)

var mappingAutonomousDatabaseDataguardAssociationPeerRole = map[string]AutonomousDatabaseDataguardAssociationPeerRoleEnum{
	"PRIMARY":          AutonomousDatabaseDataguardAssociationPeerRolePrimary,
	"STANDBY":          AutonomousDatabaseDataguardAssociationPeerRoleStandby,
	"DISABLED_STANDBY": AutonomousDatabaseDataguardAssociationPeerRoleDisabledStandby,
}

// GetAutonomousDatabaseDataguardAssociationPeerRoleEnumValues Enumerates the set of values for AutonomousDatabaseDataguardAssociationPeerRoleEnum
func GetAutonomousDatabaseDataguardAssociationPeerRoleEnumValues() []AutonomousDatabaseDataguardAssociationPeerRoleEnum {
	values := make([]AutonomousDatabaseDataguardAssociationPeerRoleEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataguardAssociationPeerRole {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum Enum with underlying type: string
type AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum
const (
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateProvisioning         AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "PROVISIONING"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateAvailable            AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "AVAILABLE"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRoleChangeInProgress AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateTerminating          AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "TERMINATING"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateTerminated           AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "TERMINATED"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateFailed               AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "FAILED"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateUnavailable          AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "UNAVAILABLE"
)

var mappingAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleState = map[string]AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum{
	"PROVISIONING":            AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateProvisioning,
	"AVAILABLE":               AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateAvailable,
	"ROLE_CHANGE_IN_PROGRESS": AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRoleChangeInProgress,
	"TERMINATING":             AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateTerminating,
	"TERMINATED":              AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateTerminated,
	"FAILED":                  AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateFailed,
	"UNAVAILABLE":             AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateUnavailable,
}

// GetAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnumValues Enumerates the set of values for AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum
func GetAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnumValues() []AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum {
	values := make([]AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseDataguardAssociationProtectionModeEnum Enum with underlying type: string
type AutonomousDatabaseDataguardAssociationProtectionModeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataguardAssociationProtectionModeEnum
const (
	AutonomousDatabaseDataguardAssociationProtectionModeAvailability AutonomousDatabaseDataguardAssociationProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	AutonomousDatabaseDataguardAssociationProtectionModePerformance  AutonomousDatabaseDataguardAssociationProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingAutonomousDatabaseDataguardAssociationProtectionMode = map[string]AutonomousDatabaseDataguardAssociationProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": AutonomousDatabaseDataguardAssociationProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  AutonomousDatabaseDataguardAssociationProtectionModePerformance,
}

// GetAutonomousDatabaseDataguardAssociationProtectionModeEnumValues Enumerates the set of values for AutonomousDatabaseDataguardAssociationProtectionModeEnum
func GetAutonomousDatabaseDataguardAssociationProtectionModeEnumValues() []AutonomousDatabaseDataguardAssociationProtectionModeEnum {
	values := make([]AutonomousDatabaseDataguardAssociationProtectionModeEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataguardAssociationProtectionMode {
		values = append(values, v)
	}
	return values
}
