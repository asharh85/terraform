// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"net/http"
)

// ListReplicationPoliciesRequest wrapper for the ListReplicationPolicies operation
type ListReplicationPoliciesRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket. Avoid entering confidential information.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// The page at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReplicationPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReplicationPoliciesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReplicationPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListReplicationPoliciesResponse wrapper for the ListReplicationPolicies operation
type ListReplicationPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ReplicationPolicySummary instances
	Items []ReplicationPolicySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Paginating a list of replication policies. In the GET request, set the limit to the number of buckets items that you want returned in the response. If the opc-next-page header appears in the response, then this is a partial list and there are additional policies to get. Include the header's value as the `page` parameter in the subsequent GET request to get the next batch of policies. Repeat this process to retrieve the entire list of policies.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListReplicationPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReplicationPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}