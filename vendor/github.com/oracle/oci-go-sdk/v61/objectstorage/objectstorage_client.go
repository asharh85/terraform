// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"github.com/oracle/oci-go-sdk/v61/common/auth"
	"net/http"
)

//ObjectStorageClient a client for ObjectStorage
type ObjectStorageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewObjectStorageClientWithConfigurationProvider Creates a new default ObjectStorage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewObjectStorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ObjectStorageClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newObjectStorageClientFromBaseClient(baseClient, provider)
}

// NewObjectStorageClientWithOboToken Creates a new default ObjectStorage client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewObjectStorageClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ObjectStorageClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newObjectStorageClientFromBaseClient(baseClient, configProvider)
}

func newObjectStorageClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ObjectStorageClient, err error) {
	// ObjectStorage service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ObjectStorageClient{BaseClient: baseClient}
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ObjectStorageClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("objectstorage", "https://objectstorage.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ObjectStorageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ObjectStorageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AbortMultipartUpload Aborts an in-progress multipart upload and deletes all parts that have been uploaded.
// A default retry strategy applies to this operation AbortMultipartUpload()
func (client ObjectStorageClient) AbortMultipartUpload(ctx context.Context, request AbortMultipartUploadRequest) (response AbortMultipartUploadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.abortMultipartUpload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AbortMultipartUploadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AbortMultipartUploadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AbortMultipartUploadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AbortMultipartUploadResponse")
	}
	return
}

// abortMultipartUpload implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) abortMultipartUpload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AbortMultipartUploadResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkCopyObjects Creates a request to copy a batch of objects within a region or to another region.
// See Object Names (https://docs.cloud.oracle.com/Content/Object/Tasks/managingobjects.htm#namerequirements)
// for object naming requirements.
// A default retry strategy applies to this operation BulkCopyObjects()
func (client ObjectStorageClient) BulkCopyObjects(ctx context.Context, request BulkCopyObjectsRequest) (response BulkCopyObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkCopyObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkCopyObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkCopyObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkCopyObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkCopyObjectsResponse")
	}
	return
}

// bulkCopyObjects implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) bulkCopyObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/bulkCopyObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkCopyObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelWorkRequest Cancels a work request.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client ObjectStorageClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CheckObject Retrieve an object's stored and calculated digests, specifically:
//   - The content MD5;
//   - The SHA-256 of each chunk on each storage server.
// If the MD5 digest calculated from the decrypted object data
// does not match the MD5 digest stored in the object's metadata,
// or any chunk's' SHA-256 does not match the one stored for it,
// this API will signal an error.  The stored and calculated
// digests are always included in the response.
// This internal API allows Object Storage tooling to verify the
// correctness of stored data.  Any tenancy's objects can be
// verified with this API, so it must only protected by BOAT AAA.
// A default retry strategy applies to this operation CheckObject()
func (client ObjectStorageClient) CheckObject(ctx context.Context, request CheckObjectRequest) (response CheckObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.checkObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CheckObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CheckObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CheckObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CheckObjectResponse")
	}
	return
}

// checkObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) checkObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/actions/checkObject/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CheckObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CommitMultipartUpload Commits a multipart upload, which involves checking part numbers and entity tags (ETags) of the parts, to create an aggregate object.
// A default retry strategy applies to this operation CommitMultipartUpload()
func (client ObjectStorageClient) CommitMultipartUpload(ctx context.Context, request CommitMultipartUploadRequest) (response CommitMultipartUploadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.commitMultipartUpload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CommitMultipartUploadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CommitMultipartUploadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CommitMultipartUploadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CommitMultipartUploadResponse")
	}
	return
}

// commitMultipartUpload implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) commitMultipartUpload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CommitMultipartUploadResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CopyObject Creates a request to copy an object within a region or to another region.
// See Object Names (https://docs.cloud.oracle.com/Content/Object/Tasks/managingobjects.htm#namerequirements)
// for object naming requirements.
// A default retry strategy applies to this operation CopyObject()
func (client ObjectStorageClient) CopyObject(ctx context.Context, request CopyObjectRequest) (response CopyObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.copyObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CopyObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CopyObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CopyObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CopyObjectResponse")
	}
	return
}

// copyObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) copyObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/copyObject", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CopyObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CopyPart Copy part of an existing object to a destination object
// A default retry strategy applies to this operation CopyPart()
func (client ObjectStorageClient) CopyPart(ctx context.Context, request CopyPartRequest) (response CopyPartResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.copyPart, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CopyPartResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CopyPartResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CopyPartResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CopyPartResponse")
	}
	return
}

// copyPart implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) copyPart(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/cp/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CopyPartResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBucket Creates a bucket in the given namespace with a bucket name and optional user-defined metadata. Avoid entering
// confidential information in bucket names.
// A default retry strategy applies to this operation CreateBucket()
func (client ObjectStorageClient) CreateBucket(ctx context.Context, request CreateBucketRequest) (response CreateBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBucketResponse")
	}
	return
}

// createBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMultipartUpload Starts a new multipart upload to a specific object in the given bucket in the given namespace.
// See Object Names (https://docs.cloud.oracle.com/Content/Object/Tasks/managingobjects.htm#namerequirements)
// for object naming requirements.
// A default retry strategy applies to this operation CreateMultipartUpload()
func (client ObjectStorageClient) CreateMultipartUpload(ctx context.Context, request CreateMultipartUploadRequest) (response CreateMultipartUploadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createMultipartUpload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMultipartUploadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMultipartUploadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMultipartUploadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMultipartUploadResponse")
	}
	return
}

// createMultipartUpload implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createMultipartUpload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMultipartUploadResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePreauthenticatedRequest Creates a pre-authenticated request specific to the bucket.
// A default retry strategy applies to this operation CreatePreauthenticatedRequest()
func (client ObjectStorageClient) CreatePreauthenticatedRequest(ctx context.Context, request CreatePreauthenticatedRequestRequest) (response CreatePreauthenticatedRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createPreauthenticatedRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePreauthenticatedRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePreauthenticatedRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePreauthenticatedRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePreauthenticatedRequestResponse")
	}
	return
}

// createPreauthenticatedRequest implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createPreauthenticatedRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/p", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePreauthenticatedRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateReplicationPolicy Creates a replication policy for the specified bucket.
// A default retry strategy applies to this operation CreateReplicationPolicy()
func (client ObjectStorageClient) CreateReplicationPolicy(ctx context.Context, request CreateReplicationPolicyRequest) (response CreateReplicationPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createReplicationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateReplicationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateReplicationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateReplicationPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateReplicationPolicyResponse")
	}
	return
}

// createReplicationPolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createReplicationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/replicationPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateReplicationPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRetentionRule Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds.
// Note that a maximum of 100 rules are supported on a bucket.
// A default retry strategy applies to this operation CreateRetentionRule()
func (client ObjectStorageClient) CreateRetentionRule(ctx context.Context, request CreateRetentionRuleRequest) (response CreateRetentionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createRetentionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRetentionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRetentionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRetentionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRetentionRuleResponse")
	}
	return
}

// createRetentionRule implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createRetentionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/retentionRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRetentionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBucket Deletes a bucket if the bucket is already empty. If the bucket is not empty, use
// DeleteObject first. In addition,
// you cannot delete a bucket that has a multipart upload in progress or a pre-authenticated
// request associated with that bucket.
// A default retry strategy applies to this operation DeleteBucket()
func (client ObjectStorageClient) DeleteBucket(ctx context.Context, request DeleteBucketRequest) (response DeleteBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBucketResponse")
	}
	return
}

// deleteBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteObject Deletes an object.
// A default retry strategy applies to this operation DeleteObject()
func (client ObjectStorageClient) DeleteObject(ctx context.Context, request DeleteObjectRequest) (response DeleteObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteObjectResponse")
	}
	return
}

// deleteObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteObjectLifecyclePolicy Deletes the object lifecycle policy for the bucket.
// A default retry strategy applies to this operation DeleteObjectLifecyclePolicy()
func (client ObjectStorageClient) DeleteObjectLifecyclePolicy(ctx context.Context, request DeleteObjectLifecyclePolicyRequest) (response DeleteObjectLifecyclePolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteObjectLifecyclePolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteObjectLifecyclePolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteObjectLifecyclePolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteObjectLifecyclePolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteObjectLifecyclePolicyResponse")
	}
	return
}

// deleteObjectLifecyclePolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteObjectLifecyclePolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/l", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteObjectLifecyclePolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePreauthenticatedRequest Deletes the pre-authenticated request for the bucket.
// A default retry strategy applies to this operation DeletePreauthenticatedRequest()
func (client ObjectStorageClient) DeletePreauthenticatedRequest(ctx context.Context, request DeletePreauthenticatedRequestRequest) (response DeletePreauthenticatedRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePreauthenticatedRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePreauthenticatedRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePreauthenticatedRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePreauthenticatedRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePreauthenticatedRequestResponse")
	}
	return
}

// deletePreauthenticatedRequest implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deletePreauthenticatedRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/p/{parId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePreauthenticatedRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteReplicationPolicy Deletes the replication policy associated with the source bucket.
// A default retry strategy applies to this operation DeleteReplicationPolicy()
func (client ObjectStorageClient) DeleteReplicationPolicy(ctx context.Context, request DeleteReplicationPolicyRequest) (response DeleteReplicationPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteReplicationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteReplicationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteReplicationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteReplicationPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteReplicationPolicyResponse")
	}
	return
}

// deleteReplicationPolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteReplicationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/replicationPolicies/{replicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteReplicationPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRetentionRule Deletes the specified rule. The deletion takes effect typically within 30 seconds.
// A default retry strategy applies to this operation DeleteRetentionRule()
func (client ObjectStorageClient) DeleteRetentionRule(ctx context.Context, request DeleteRetentionRuleRequest) (response DeleteRetentionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRetentionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRetentionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRetentionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRetentionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRetentionRuleResponse")
	}
	return
}

// deleteRetentionRule implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteRetentionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/retentionRules/{retentionRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRetentionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBucket Gets the current representation of the given bucket in the given Object Storage namespace.
// A default retry strategy applies to this operation GetBucket()
func (client ObjectStorageClient) GetBucket(ctx context.Context, request GetBucketRequest) (response GetBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBucketResponse")
	}
	return
}

// getBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBucketOptions Lists the various options associated with the bucket. The options are returned as a JSON-formatted object.
// This API is for internal-use only.
// A default retry strategy applies to this operation GetBucketOptions()
func (client ObjectStorageClient) GetBucketOptions(ctx context.Context, request GetBucketOptionsRequest) (response GetBucketOptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBucketOptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBucketOptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBucketOptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBucketOptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBucketOptionsResponse")
	}
	return
}

// getBucketOptions implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getBucketOptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/actions/options", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBucketOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNamespace Each Oracle Cloud Infrastructure tenant is assigned one unique and uneditable Object Storage namespace. The namespace
// is a system-generated string assigned during account creation. For some older tenancies, the namespace string may be
// the tenancy name in all lower-case letters. You cannot edit a namespace.
// GetNamespace returns the name of the Object Storage namespace for the user making the request.
// If an optional compartmentId query parameter is provided, GetNamespace returns the namespace name of the corresponding
// tenancy, provided the user has access to it.
// A default retry strategy applies to this operation GetNamespace()
func (client ObjectStorageClient) GetNamespace(ctx context.Context, request GetNamespaceRequest) (response GetNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetNamespaceResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNamespaceResponse")
	}
	return
}

// getNamespace implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNamespaceMetadata Gets the metadata for the Object Storage namespace, which contains defaultS3CompartmentId and
// defaultSwiftCompartmentId.
// Any user with the OBJECTSTORAGE_NAMESPACE_READ permission will be able to see the current metadata. If you are
// not authorized, talk to an administrator. If you are an administrator who needs to write policies
// to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// A default retry strategy applies to this operation GetNamespaceMetadata()
func (client ObjectStorageClient) GetNamespaceMetadata(ctx context.Context, request GetNamespaceMetadataRequest) (response GetNamespaceMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNamespaceMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNamespaceMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNamespaceMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNamespaceMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNamespaceMetadataResponse")
	}
	return
}

// getNamespaceMetadata implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getNamespaceMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNamespaceMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetObject Gets the metadata and body of an object.
// A default retry strategy applies to this operation GetObject()
func (client ObjectStorageClient) GetObject(ctx context.Context, request GetObjectRequest) (response GetObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetObjectResponse")
	}
	return
}

// getObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetObjectLifecyclePolicy Gets the object lifecycle policy for the bucket.
// A default retry strategy applies to this operation GetObjectLifecyclePolicy()
func (client ObjectStorageClient) GetObjectLifecyclePolicy(ctx context.Context, request GetObjectLifecyclePolicyRequest) (response GetObjectLifecyclePolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getObjectLifecyclePolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetObjectLifecyclePolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetObjectLifecyclePolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetObjectLifecyclePolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetObjectLifecyclePolicyResponse")
	}
	return
}

// getObjectLifecyclePolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getObjectLifecyclePolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/l", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetObjectLifecyclePolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPreauthenticatedRequest Gets the pre-authenticated request for the bucket.
// A default retry strategy applies to this operation GetPreauthenticatedRequest()
func (client ObjectStorageClient) GetPreauthenticatedRequest(ctx context.Context, request GetPreauthenticatedRequestRequest) (response GetPreauthenticatedRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPreauthenticatedRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPreauthenticatedRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPreauthenticatedRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPreauthenticatedRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPreauthenticatedRequestResponse")
	}
	return
}

// getPreauthenticatedRequest implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getPreauthenticatedRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p/{parId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPreauthenticatedRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetReplicationPolicy Get the replication policy.
// A default retry strategy applies to this operation GetReplicationPolicy()
func (client ObjectStorageClient) GetReplicationPolicy(ctx context.Context, request GetReplicationPolicyRequest) (response GetReplicationPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReplicationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReplicationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReplicationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReplicationPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReplicationPolicyResponse")
	}
	return
}

// getReplicationPolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getReplicationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/replicationPolicies/{replicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetReplicationPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRetentionRule Get the specified retention rule.
// A default retry strategy applies to this operation GetRetentionRule()
func (client ObjectStorageClient) GetRetentionRule(ctx context.Context, request GetRetentionRuleRequest) (response GetRetentionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRetentionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRetentionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRetentionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRetentionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRetentionRuleResponse")
	}
	return
}

// getRetentionRule implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getRetentionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/retentionRules/{retentionRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRetentionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request for the given ID.
// A default retry strategy applies to this operation GetWorkRequest()
func (client ObjectStorageClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// HeadBucket Efficiently checks to see if a bucket exists and gets the current entity tag (ETag) for the bucket.
// A default retry strategy applies to this operation HeadBucket()
func (client ObjectStorageClient) HeadBucket(ctx context.Context, request HeadBucketRequest) (response HeadBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.headBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = HeadBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = HeadBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(HeadBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into HeadBucketResponse")
	}
	return
}

// headBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) headBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodHead, "/n/{namespaceName}/b/{bucketName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response HeadBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// HeadObject Gets the user-defined metadata and entity tag (ETag) for an object.
// A default retry strategy applies to this operation HeadObject()
func (client ObjectStorageClient) HeadObject(ctx context.Context, request HeadObjectRequest) (response HeadObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.headObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = HeadObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = HeadObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(HeadObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into HeadObjectResponse")
	}
	return
}

// headObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) headObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodHead, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response HeadObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBuckets Gets a list of all BucketSummary items in a compartment. A BucketSummary contains only summary fields for the bucket
// and does not contain fields like the user-defined metadata.
// ListBuckets returns a BucketSummary containing at most 1000 buckets. To paginate through more buckets, use the returned
// `opc-next-page` value with the `page` request parameter.
// To use this and other API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// A default retry strategy applies to this operation ListBuckets()
func (client ObjectStorageClient) ListBuckets(ctx context.Context, request ListBucketsRequest) (response ListBucketsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBuckets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBucketsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBucketsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBucketsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBucketsResponse")
	}
	return
}

// listBuckets implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listBuckets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBucketsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMultipartUploadParts Lists the parts of an in-progress multipart upload.
// A default retry strategy applies to this operation ListMultipartUploadParts()
func (client ObjectStorageClient) ListMultipartUploadParts(ctx context.Context, request ListMultipartUploadPartsRequest) (response ListMultipartUploadPartsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMultipartUploadParts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMultipartUploadPartsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMultipartUploadPartsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMultipartUploadPartsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMultipartUploadPartsResponse")
	}
	return
}

// listMultipartUploadParts implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listMultipartUploadParts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMultipartUploadPartsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMultipartUploads Lists all of the in-progress multipart uploads for the given bucket in the given Object Storage namespace.
// A default retry strategy applies to this operation ListMultipartUploads()
func (client ObjectStorageClient) ListMultipartUploads(ctx context.Context, request ListMultipartUploadsRequest) (response ListMultipartUploadsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMultipartUploads, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMultipartUploadsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMultipartUploadsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMultipartUploadsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMultipartUploadsResponse")
	}
	return
}

// listMultipartUploads implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listMultipartUploads(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMultipartUploadsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListObjectVersions Lists the object versions in a bucket.
// ListObjectVersions returns an ObjectVersionCollection containing at most 1000 object versions. To paginate through
// more object versions, use the returned `opc-next-page` value with the `page` request parameter.
// To use this and other API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// A default retry strategy applies to this operation ListObjectVersions()
func (client ObjectStorageClient) ListObjectVersions(ctx context.Context, request ListObjectVersionsRequest) (response ListObjectVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listObjectVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListObjectVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListObjectVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListObjectVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListObjectVersionsResponse")
	}
	return
}

// listObjectVersions implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listObjectVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/objectversions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListObjectVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListObjects Lists the objects in a bucket. By default, ListObjects returns object names only. See the `fields`
// parameter for other fields that you can optionally include in ListObjects response.
// ListObjects returns at most 1000 objects. To paginate through more objects, use the returned 'nextStartWith'
// value with the 'start' parameter. To filter which objects ListObjects returns, use the 'start' and 'end'
// parameters.
// To use this and other API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// A default retry strategy applies to this operation ListObjects()
func (client ObjectStorageClient) ListObjects(ctx context.Context, request ListObjectsRequest) (response ListObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListObjectsResponse")
	}
	return
}

// listObjects implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPreauthenticatedRequests Lists pre-authenticated requests for the bucket.
// A default retry strategy applies to this operation ListPreauthenticatedRequests()
func (client ObjectStorageClient) ListPreauthenticatedRequests(ctx context.Context, request ListPreauthenticatedRequestsRequest) (response ListPreauthenticatedRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPreauthenticatedRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPreauthenticatedRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPreauthenticatedRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPreauthenticatedRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPreauthenticatedRequestsResponse")
	}
	return
}

// listPreauthenticatedRequests implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listPreauthenticatedRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPreauthenticatedRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReplicationPolicies List the replication policies associated with a bucket.
// A default retry strategy applies to this operation ListReplicationPolicies()
func (client ObjectStorageClient) ListReplicationPolicies(ctx context.Context, request ListReplicationPoliciesRequest) (response ListReplicationPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReplicationPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReplicationPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReplicationPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReplicationPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReplicationPoliciesResponse")
	}
	return
}

// listReplicationPolicies implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listReplicationPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/replicationPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListReplicationPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReplicationSources List the replication sources of a destination bucket.
// A default retry strategy applies to this operation ListReplicationSources()
func (client ObjectStorageClient) ListReplicationSources(ctx context.Context, request ListReplicationSourcesRequest) (response ListReplicationSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReplicationSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReplicationSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReplicationSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReplicationSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReplicationSourcesResponse")
	}
	return
}

// listReplicationSources implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listReplicationSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/replicationSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListReplicationSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRetentionRules List the retention rules for a bucket. The retention rules are sorted based on creation time,
// with the most recently created retention rule returned first.
// A default retry strategy applies to this operation ListRetentionRules()
func (client ObjectStorageClient) ListRetentionRules(ctx context.Context, request ListRetentionRulesRequest) (response ListRetentionRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRetentionRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRetentionRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRetentionRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRetentionRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRetentionRulesResponse")
	}
	return
}

// listRetentionRules implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listRetentionRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/retentionRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRetentionRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Lists the errors of the work request with the given ID.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client ObjectStorageClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the logs of the work request with the given ID.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client ObjectStorageClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
// A default retry strategy applies to this operation ListWorkRequests()
func (client ObjectStorageClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MakeBucketWritable Stops replication to the destination bucket and removes the replication policy. When the replication
// policy was created, this destination bucket became read-only except for new and changed objects replicated
// automatically from the source bucket. MakeBucketWritable removes the replication policy. This bucket is no
// longer the target for replication and is now writable, allowing users to make changes to bucket contents.
// A default retry strategy applies to this operation MakeBucketWritable()
func (client ObjectStorageClient) MakeBucketWritable(ctx context.Context, request MakeBucketWritableRequest) (response MakeBucketWritableResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.makeBucketWritable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MakeBucketWritableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MakeBucketWritableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MakeBucketWritableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MakeBucketWritableResponse")
	}
	return
}

// makeBucketWritable implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) makeBucketWritable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/makeBucketWritable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MakeBucketWritableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MergeObjectMetadata Merges an Object's user-defined metadata with existing metadata.
// A default retry strategy applies to this operation MergeObjectMetadata()
func (client ObjectStorageClient) MergeObjectMetadata(ctx context.Context, request MergeObjectMetadataRequest) (response MergeObjectMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.mergeObjectMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MergeObjectMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MergeObjectMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MergeObjectMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MergeObjectMetadataResponse")
	}
	return
}

// mergeObjectMetadata implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) mergeObjectMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/mergeObjectMetadata/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MergeObjectMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutObject Creates a new object or overwrites an existing object with the same name. The maximum object size allowed by
// PutObject is 50 GiB.
// See Object Names (https://docs.cloud.oracle.com/Content/Object/Tasks/managingobjects.htm#namerequirements)
// for object naming requirements.
// See Special Instructions for Object Storage PUT (https://docs.cloud.oracle.com/Content/API/Concepts/signingrequests.htm#ObjectStoragePut)
// for request signature requirements.
// A default retry strategy applies to this operation PutObject()
func (client ObjectStorageClient) PutObject(ctx context.Context, request PutObjectRequest) (response PutObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutObjectResponse")
	}
	return
}

// putObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) putObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutObjectResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutObjectLifecyclePolicy Creates or replaces the object lifecycle policy for the bucket.
// A default retry strategy applies to this operation PutObjectLifecyclePolicy()
func (client ObjectStorageClient) PutObjectLifecyclePolicy(ctx context.Context, request PutObjectLifecyclePolicyRequest) (response PutObjectLifecyclePolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putObjectLifecyclePolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutObjectLifecyclePolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutObjectLifecyclePolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutObjectLifecyclePolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutObjectLifecyclePolicyResponse")
	}
	return
}

// putObjectLifecyclePolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) putObjectLifecyclePolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/l", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutObjectLifecyclePolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// QueryObject Returns query result as a byte string in the response body.
// A default retry strategy applies to this operation QueryObject()
func (client ObjectStorageClient) QueryObject(ctx context.Context, request QueryObjectRequest) (response QueryObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.queryObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = QueryObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = QueryObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(QueryObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into QueryObjectResponse")
	}
	return
}

// queryObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) queryObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/query", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response QueryObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReencryptBucket Re-encrypts the unique data encryption key that encrypts each object written to the bucket by using the most recent
// version of the master encryption key assigned to the bucket. (All data encryption keys are encrypted by a master
// encryption key. Master encryption keys are assigned to buckets and managed by Oracle by default, but you can assign
// a key that you created and control through the Oracle Cloud Infrastructure Key Management service.) The kmsKeyId property
// of the bucket determines which master encryption key is assigned to the bucket. If you assigned a different Key Management
// master encryption key to the bucket, you can call this API to re-encrypt all data encryption keys with the newly
// assigned key. Similarly, you might want to re-encrypt all data encryption keys if the assigned key has been rotated to
// a new key version since objects were last added to the bucket. If you call this API and there is no kmsKeyId associated
// with the bucket, the call will fail.
// Calling this API starts a work request task to re-encrypt the data encryption key of all objects in the bucket. Only
// objects created before the time of the API call will be re-encrypted. The call can take a long time, depending on how many
// objects are in the bucket and how big they are. This API returns a work request ID that you can use to retrieve the status
// of the work request task.
// All the versions of objects will be re-encrypted whether versioning is enabled or suspended at the bucket.
// A default retry strategy applies to this operation ReencryptBucket()
func (client ObjectStorageClient) ReencryptBucket(ctx context.Context, request ReencryptBucketRequest) (response ReencryptBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.reencryptBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReencryptBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReencryptBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReencryptBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReencryptBucketResponse")
	}
	return
}

// reencryptBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) reencryptBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/reencrypt", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReencryptBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReencryptObject Re-encrypts the data encryption keys that encrypt the object and its chunks. By default, when you create a bucket, the Object Storage
// service manages the master encryption key used to encrypt each object's data encryption keys. The encryption mechanism that you specify for
// the bucket applies to the objects it contains.
// You can alternatively employ one of these encryption strategies for an object:
// - You can assign a key that you created and control through the Oracle Cloud Infrastructure Vault service.
// - You can encrypt an object using your own encryption key. The key you supply is known as a customer-provided encryption key (SSE-C).
// A default retry strategy applies to this operation ReencryptObject()
func (client ObjectStorageClient) ReencryptObject(ctx context.Context, request ReencryptObjectRequest) (response ReencryptObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.reencryptObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReencryptObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReencryptObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReencryptObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReencryptObjectResponse")
	}
	return
}

// reencryptObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) reencryptObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/reencrypt/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReencryptObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RenameObject Rename an object in the given Object Storage namespace.
// See Object Names (https://docs.cloud.oracle.com/Content/Object/Tasks/managingobjects.htm#namerequirements)
// for object naming requirements.
// A default retry strategy applies to this operation RenameObject()
func (client ObjectStorageClient) RenameObject(ctx context.Context, request RenameObjectRequest) (response RenameObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.renameObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RenameObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RenameObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RenameObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RenameObjectResponse")
	}
	return
}

// renameObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) renameObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/renameObject", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RenameObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReplaceObjectMetadata Overwrites an object's user-defined metadata.
// A default retry strategy applies to this operation ReplaceObjectMetadata()
func (client ObjectStorageClient) ReplaceObjectMetadata(ctx context.Context, request ReplaceObjectMetadataRequest) (response ReplaceObjectMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.replaceObjectMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReplaceObjectMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReplaceObjectMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReplaceObjectMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReplaceObjectMetadataResponse")
	}
	return
}

// replaceObjectMetadata implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) replaceObjectMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/replaceObjectMetadata/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReplaceObjectMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestoreObjects Restores one or more objects specified by the objectName parameter.
// By default objects will be restored for 24 hours. Duration can be configured using the hours parameter.
// A default retry strategy applies to this operation RestoreObjects()
func (client ObjectStorageClient) RestoreObjects(ctx context.Context, request RestoreObjectsRequest) (response RestoreObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.restoreObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestoreObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestoreObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreObjectsResponse")
	}
	return
}

// restoreObjects implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) restoreObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/restoreObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestoreObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBucket Performs a partial or full update of a bucket's user-defined metadata.
// Use UpdateBucket to move a bucket from one compartment to another within the same tenancy. Supply the compartmentID
// of the compartment that you want to move the bucket to. For more information about moving resources between compartments,
// see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// A default retry strategy applies to this operation UpdateBucket()
func (client ObjectStorageClient) UpdateBucket(ctx context.Context, request UpdateBucketRequest) (response UpdateBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBucketResponse")
	}
	return
}

// updateBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBucketOptions Updates internal options associated with a bucket. The options to be updated/added/removed are specified in a
// JSON object in the body of the request. This API only affects the bucket options that are specified in the JSON
// body. Other options that have been set (by prior use of this API) are left unchanged.
// Options that have a value set to null are removed.
// All the existing options are removed, if the value associated with "freeformOptions" is null.
// This API is for internal-use only.
// A default retry strategy applies to this operation UpdateBucketOptions()
func (client ObjectStorageClient) UpdateBucketOptions(ctx context.Context, request UpdateBucketOptionsRequest) (response UpdateBucketOptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBucketOptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBucketOptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBucketOptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBucketOptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBucketOptionsResponse")
	}
	return
}

// updateBucketOptions implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateBucketOptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/options", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBucketOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateNamespaceMetadata By default, buckets created using the Amazon S3 Compatibility API or the Swift API are created in the root
// compartment of the Oracle Cloud Infrastructure tenancy.
// You can change the default Swift/Amazon S3 compartmentId designation to a different compartmentId. All
// subsequent bucket creations will use the new default compartment, but no previously created
// buckets will be modified. A user must have OBJECTSTORAGE_NAMESPACE_UPDATE permission to make changes to the default
// compartments for Amazon S3 and Swift.
// A default retry strategy applies to this operation UpdateNamespaceMetadata()
func (client ObjectStorageClient) UpdateNamespaceMetadata(ctx context.Context, request UpdateNamespaceMetadataRequest) (response UpdateNamespaceMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNamespaceMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNamespaceMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNamespaceMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNamespaceMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNamespaceMetadataResponse")
	}
	return
}

// updateNamespaceMetadata implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateNamespaceMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNamespaceMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateObjectStorageTier Changes the storage tier of the object specified by the objectName parameter.
// A default retry strategy applies to this operation UpdateObjectStorageTier()
func (client ObjectStorageClient) UpdateObjectStorageTier(ctx context.Context, request UpdateObjectStorageTierRequest) (response UpdateObjectStorageTierResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateObjectStorageTier, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateObjectStorageTierResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateObjectStorageTierResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateObjectStorageTierResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateObjectStorageTierResponse")
	}
	return
}

// updateObjectStorageTier implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateObjectStorageTier(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/updateObjectStorageTier", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateObjectStorageTierResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRetentionRule Updates the specified retention rule. Rule changes take effect typically within 30 seconds.
// A default retry strategy applies to this operation UpdateRetentionRule()
func (client ObjectStorageClient) UpdateRetentionRule(ctx context.Context, request UpdateRetentionRuleRequest) (response UpdateRetentionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRetentionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRetentionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRetentionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRetentionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRetentionRuleResponse")
	}
	return
}

// updateRetentionRule implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateRetentionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/retentionRules/{retentionRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRetentionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UploadPart Uploads a single part of a multipart upload.
// A default retry strategy applies to this operation UploadPart()
func (client ObjectStorageClient) UploadPart(ctx context.Context, request UploadPartRequest) (response UploadPartResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.uploadPart, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadPartResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadPartResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadPartResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadPartResponse")
	}
	return
}

// uploadPart implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) uploadPart(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadPartResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}