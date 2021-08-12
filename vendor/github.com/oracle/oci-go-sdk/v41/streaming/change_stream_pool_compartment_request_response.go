// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package streaming

import (
	"github.com/oracle/oci-go-sdk/v41/common"
	"net/http"
)

// ChangeStreamPoolCompartmentRequest wrapper for the ChangeStreamPoolCompartment operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ChangeStreamPoolCompartment.go.html to see an example of how to use ChangeStreamPoolCompartmentRequest.
type ChangeStreamPoolCompartmentRequest struct {

	// The OCID of the stream pool.
	StreamPoolId *string `mandatory:"true" contributesTo:"path" name:"streamPoolId"`

	// The stream pool will be moved into the compartment specified within this entity.
	ChangeStreamPoolCompartmentDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the if-match parameter to the value of the etag from a previous GET or POST response for that resource. The resource will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ChangeStreamPoolCompartmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ChangeStreamPoolCompartmentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ChangeStreamPoolCompartmentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ChangeStreamPoolCompartmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ChangeStreamPoolCompartmentResponse wrapper for the ChangeStreamPoolCompartment operation
type ChangeStreamPoolCompartmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ChangeStreamPoolCompartmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ChangeStreamPoolCompartmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
