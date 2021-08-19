// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
)

// CreateNewDatabaseDetails Details for creating a new database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateNewDatabaseDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"true" json:"dbHomeId"`

	Database *CreateDatabaseDetails `mandatory:"true" json:"database"`

	// A valid Oracle Database version. To get a list of supported versions, use the ListDbVersions operation.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`
}

//GetDbHomeId returns DbHomeId
func (m CreateNewDatabaseDetails) GetDbHomeId() *string {
	return m.DbHomeId
}

//GetDbVersion returns DbVersion
func (m CreateNewDatabaseDetails) GetDbVersion() *string {
	return m.DbVersion
}

//GetKmsKeyId returns KmsKeyId
func (m CreateNewDatabaseDetails) GetKmsKeyId() *string {
	return m.KmsKeyId
}

//GetKmsKeyVersionId returns KmsKeyVersionId
func (m CreateNewDatabaseDetails) GetKmsKeyVersionId() *string {
	return m.KmsKeyVersionId
}

func (m CreateNewDatabaseDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateNewDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateNewDatabaseDetails CreateNewDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateNewDatabaseDetails
	}{
		"NONE",
		(MarshalTypeCreateNewDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
