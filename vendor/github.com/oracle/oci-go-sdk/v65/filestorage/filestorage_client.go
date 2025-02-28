// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//FileStorageClient a client for FileStorage
type FileStorageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFileStorageClientWithConfigurationProvider Creates a new default FileStorage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFileStorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FileStorageClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newFileStorageClientFromBaseClient(baseClient, provider)
}

// NewFileStorageClientWithOboToken Creates a new default FileStorage client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewFileStorageClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FileStorageClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFileStorageClientFromBaseClient(baseClient, configProvider)
}

func newFileStorageClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FileStorageClient, err error) {
	// FileStorage service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FileStorage"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FileStorageClient{BaseClient: baseClient}
	client.BasePath = "20171215"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FileStorageClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("filestorage", "https://filestorage.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FileStorageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *FileStorageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeFileSystemCompartment Moves a file system and its associated snapshots into a different compartment within the same tenancy. For information about moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes)
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ChangeFileSystemCompartment.go.html to see an example of how to use ChangeFileSystemCompartment API.
func (client FileStorageClient) ChangeFileSystemCompartment(ctx context.Context, request ChangeFileSystemCompartmentRequest) (response ChangeFileSystemCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeFileSystemCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFileSystemCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFileSystemCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFileSystemCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFileSystemCompartmentResponse")
	}
	return
}

// changeFileSystemCompartment implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) changeFileSystemCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fileSystems/{fileSystemId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFileSystemCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/ChangeFileSystemCompartment"
		err = common.PostProcessServiceError(err, "FileStorage", "ChangeFileSystemCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMountTargetCompartment Moves a mount target and its associated export set into a different compartment within the same tenancy. For information about moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes)
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ChangeMountTargetCompartment.go.html to see an example of how to use ChangeMountTargetCompartment API.
func (client FileStorageClient) ChangeMountTargetCompartment(ctx context.Context, request ChangeMountTargetCompartmentRequest) (response ChangeMountTargetCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeMountTargetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMountTargetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMountTargetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMountTargetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMountTargetCompartmentResponse")
	}
	return
}

// changeMountTargetCompartment implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) changeMountTargetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mountTargets/{mountTargetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMountTargetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTarget/ChangeMountTargetCompartment"
		err = common.PostProcessServiceError(err, "FileStorage", "ChangeMountTargetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExport Creates a new export in the specified export set, path, and
// file system.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/CreateExport.go.html to see an example of how to use CreateExport API.
func (client FileStorageClient) CreateExport(ctx context.Context, request CreateExportRequest) (response CreateExportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createExport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExportResponse")
	}
	return
}

// createExport implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) createExport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Export/CreateExport"
		err = common.PostProcessServiceError(err, "FileStorage", "CreateExport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFileSystem Creates a new file system in the specified compartment and
// availability domain. Instances can mount file systems in
// another availability domain, but doing so might increase
// latency when compared to mounting instances in the same
// availability domain.
// After you create a file system, you can associate it with a mount
// target. Instances can then mount the file system by connecting to the
// mount target's IP address. You can associate a file system with
// more than one mount target at a time.
// For information about access control and compartments, see
// Overview of the IAM Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
// For information about Network Security Groups access control, see
// Network Security Groups (https://docs.cloud.oracle.com/Content/Network/Concepts/networksecuritygroups.htm).
// For information about availability domains, see Regions and
// Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the
// `ListAvailabilityDomains` operation in the Identity and Access
// Management Service API.
// All Oracle Cloud Infrastructure resources, including
// file systems, get an Oracle-assigned, unique ID called an Oracle
// Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)).
// When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource
// type or by viewing the resource in the Console.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/CreateFileSystem.go.html to see an example of how to use CreateFileSystem API.
func (client FileStorageClient) CreateFileSystem(ctx context.Context, request CreateFileSystemRequest) (response CreateFileSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFileSystemResponse")
	}
	return
}

// createFileSystem implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) createFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fileSystems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/CreateFileSystem"
		err = common.PostProcessServiceError(err, "FileStorage", "CreateFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMountTarget Creates a new mount target in the specified compartment and
// subnet. You can associate a file system with a mount
// target only when they exist in the same availability domain. Instances
// can connect to mount targets in another availablity domain, but
// you might see higher latency than with instances in the same
// availability domain as the mount target.
// Mount targets have one or more private IP addresses that you can
// provide as the host portion of remote target parameters in
// client mount commands. These private IP addresses are listed
// in the privateIpIds property of the mount target and are highly available. Mount
// targets also consume additional IP addresses in their subnet.
// Do not use /30 or smaller subnets for mount target creation because they
// do not have sufficient available IP addresses.
// Allow at least three IP addresses for each mount target.
// For information about access control and compartments, see
// Overview of the IAM
// Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
// For information about availability domains, see Regions and
// Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the
// `ListAvailabilityDomains` operation in the Identity and Access
// Management Service API.
// All Oracle Cloud Infrastructure Services resources, including
// mount targets, get an Oracle-assigned, unique ID called an
// Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)).
// When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource
// type, or by viewing the resource in the Console.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/CreateMountTarget.go.html to see an example of how to use CreateMountTarget API.
func (client FileStorageClient) CreateMountTarget(ctx context.Context, request CreateMountTargetRequest) (response CreateMountTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createMountTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMountTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMountTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMountTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMountTargetResponse")
	}
	return
}

// createMountTarget implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) createMountTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mountTargets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMountTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTarget/CreateMountTarget"
		err = common.PostProcessServiceError(err, "FileStorage", "CreateMountTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSnapshot Creates a new snapshot of the specified file system. You
// can access the snapshot at `.snapshot/<name>`.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/CreateSnapshot.go.html to see an example of how to use CreateSnapshot API.
func (client FileStorageClient) CreateSnapshot(ctx context.Context, request CreateSnapshotRequest) (response CreateSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSnapshotResponse")
	}
	return
}

// createSnapshot implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) createSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/snapshots", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Snapshot/CreateSnapshot"
		err = common.PostProcessServiceError(err, "FileStorage", "CreateSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExport Deletes the specified export.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteExport.go.html to see an example of how to use DeleteExport API.
func (client FileStorageClient) DeleteExport(ctx context.Context, request DeleteExportRequest) (response DeleteExportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExportResponse")
	}
	return
}

// deleteExport implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteExport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/exports/{exportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Export/DeleteExport"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteExport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFileSystem Deletes the specified file system. Before you delete the file system,
// verify that no remaining export resources still reference it. Deleting a
// file system also deletes all of its snapshots.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteFileSystem.go.html to see an example of how to use DeleteFileSystem API.
func (client FileStorageClient) DeleteFileSystem(ctx context.Context, request DeleteFileSystemRequest) (response DeleteFileSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFileSystemResponse")
	}
	return
}

// deleteFileSystem implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fileSystems/{fileSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/DeleteFileSystem"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMountTarget Deletes the specified mount target. This operation also deletes the
// mount target's VNICs.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteMountTarget.go.html to see an example of how to use DeleteMountTarget API.
func (client FileStorageClient) DeleteMountTarget(ctx context.Context, request DeleteMountTargetRequest) (response DeleteMountTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMountTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMountTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMountTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMountTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMountTargetResponse")
	}
	return
}

// deleteMountTarget implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteMountTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/mountTargets/{mountTargetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMountTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTarget/DeleteMountTarget"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteMountTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSnapshot Deletes the specified snapshot.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteSnapshot.go.html to see an example of how to use DeleteSnapshot API.
func (client FileStorageClient) DeleteSnapshot(ctx context.Context, request DeleteSnapshotRequest) (response DeleteSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSnapshotResponse")
	}
	return
}

// deleteSnapshot implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) deleteSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/snapshots/{snapshotId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Snapshot/DeleteSnapshot"
		err = common.PostProcessServiceError(err, "FileStorage", "DeleteSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExport Gets the specified export's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetExport.go.html to see an example of how to use GetExport API.
func (client FileStorageClient) GetExport(ctx context.Context, request GetExportRequest) (response GetExportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExportResponse")
	}
	return
}

// getExport implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getExport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exports/{exportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Export/GetExport"
		err = common.PostProcessServiceError(err, "FileStorage", "GetExport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExportSet Gets the specified export set's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetExportSet.go.html to see an example of how to use GetExportSet API.
func (client FileStorageClient) GetExportSet(ctx context.Context, request GetExportSetRequest) (response GetExportSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExportSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExportSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExportSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExportSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExportSetResponse")
	}
	return
}

// getExportSet implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getExportSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exportSets/{exportSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExportSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ExportSet/GetExportSet"
		err = common.PostProcessServiceError(err, "FileStorage", "GetExportSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFileSystem Gets the specified file system's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetFileSystem.go.html to see an example of how to use GetFileSystem API.
func (client FileStorageClient) GetFileSystem(ctx context.Context, request GetFileSystemRequest) (response GetFileSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFileSystemResponse")
	}
	return
}

// getFileSystem implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fileSystems/{fileSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/GetFileSystem"
		err = common.PostProcessServiceError(err, "FileStorage", "GetFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMountTarget Gets the specified mount target's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetMountTarget.go.html to see an example of how to use GetMountTarget API.
func (client FileStorageClient) GetMountTarget(ctx context.Context, request GetMountTargetRequest) (response GetMountTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMountTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMountTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMountTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMountTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMountTargetResponse")
	}
	return
}

// getMountTarget implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getMountTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mountTargets/{mountTargetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMountTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTarget/GetMountTarget"
		err = common.PostProcessServiceError(err, "FileStorage", "GetMountTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSnapshot Gets the specified snapshot's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/GetSnapshot.go.html to see an example of how to use GetSnapshot API.
func (client FileStorageClient) GetSnapshot(ctx context.Context, request GetSnapshotRequest) (response GetSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSnapshotResponse")
	}
	return
}

// getSnapshot implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) getSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/snapshots/{snapshotId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Snapshot/GetSnapshot"
		err = common.PostProcessServiceError(err, "FileStorage", "GetSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExportSets Lists the export set resources in the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListExportSets.go.html to see an example of how to use ListExportSets API.
func (client FileStorageClient) ListExportSets(ctx context.Context, request ListExportSetsRequest) (response ListExportSetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExportSets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExportSetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExportSetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExportSetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExportSetsResponse")
	}
	return
}

// listExportSets implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listExportSets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exportSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExportSetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ExportSetSummary/ListExportSets"
		err = common.PostProcessServiceError(err, "FileStorage", "ListExportSets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExports Lists export resources by compartment, file system, or export
// set. You must specify an export set ID, a file system ID, and
// / or a compartment ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListExports.go.html to see an example of how to use ListExports API.
func (client FileStorageClient) ListExports(ctx context.Context, request ListExportsRequest) (response ListExportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExportsResponse")
	}
	return
}

// listExports implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listExports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExportsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ExportSummary/ListExports"
		err = common.PostProcessServiceError(err, "FileStorage", "ListExports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFileSystems Lists the file system resources in the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListFileSystems.go.html to see an example of how to use ListFileSystems API.
func (client FileStorageClient) ListFileSystems(ctx context.Context, request ListFileSystemsRequest) (response ListFileSystemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFileSystems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFileSystemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFileSystemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFileSystemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFileSystemsResponse")
	}
	return
}

// listFileSystems implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listFileSystems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fileSystems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFileSystemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystemSummary/ListFileSystems"
		err = common.PostProcessServiceError(err, "FileStorage", "ListFileSystems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMountTargets Lists the mount target resources in the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListMountTargets.go.html to see an example of how to use ListMountTargets API.
func (client FileStorageClient) ListMountTargets(ctx context.Context, request ListMountTargetsRequest) (response ListMountTargetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMountTargets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMountTargetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMountTargetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMountTargetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMountTargetsResponse")
	}
	return
}

// listMountTargets implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listMountTargets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mountTargets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMountTargetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTargetSummary/ListMountTargets"
		err = common.PostProcessServiceError(err, "FileStorage", "ListMountTargets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSnapshots Lists snapshots of the specified file system.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListSnapshots.go.html to see an example of how to use ListSnapshots API.
func (client FileStorageClient) ListSnapshots(ctx context.Context, request ListSnapshotsRequest) (response ListSnapshotsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSnapshots, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSnapshotsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSnapshotsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSnapshotsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSnapshotsResponse")
	}
	return
}

// listSnapshots implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) listSnapshots(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/snapshots", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSnapshotsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/SnapshotSummary/ListSnapshots"
		err = common.PostProcessServiceError(err, "FileStorage", "ListSnapshots", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExport Updates the specified export's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateExport.go.html to see an example of how to use UpdateExport API.
func (client FileStorageClient) UpdateExport(ctx context.Context, request UpdateExportRequest) (response UpdateExportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExportResponse")
	}
	return
}

// updateExport implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateExport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/exports/{exportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Export/UpdateExport"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateExport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExportSet Updates the specified export set's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateExportSet.go.html to see an example of how to use UpdateExportSet API.
func (client FileStorageClient) UpdateExportSet(ctx context.Context, request UpdateExportSetRequest) (response UpdateExportSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExportSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExportSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExportSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExportSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExportSetResponse")
	}
	return
}

// updateExportSet implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateExportSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/exportSets/{exportSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExportSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/ExportSet/UpdateExportSet"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateExportSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFileSystem Updates the specified file system's information.
// You can use this operation to rename a file system.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateFileSystem.go.html to see an example of how to use UpdateFileSystem API.
func (client FileStorageClient) UpdateFileSystem(ctx context.Context, request UpdateFileSystemRequest) (response UpdateFileSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFileSystemResponse")
	}
	return
}

// updateFileSystem implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fileSystems/{fileSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/FileSystem/UpdateFileSystem"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMountTarget Updates the specified mount target's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateMountTarget.go.html to see an example of how to use UpdateMountTarget API.
func (client FileStorageClient) UpdateMountTarget(ctx context.Context, request UpdateMountTargetRequest) (response UpdateMountTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMountTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMountTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMountTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMountTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMountTargetResponse")
	}
	return
}

// updateMountTarget implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateMountTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mountTargets/{mountTargetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMountTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/MountTarget/UpdateMountTarget"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateMountTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSnapshot Updates the specified snapshot's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/UpdateSnapshot.go.html to see an example of how to use UpdateSnapshot API.
func (client FileStorageClient) UpdateSnapshot(ctx context.Context, request UpdateSnapshotRequest) (response UpdateSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSnapshotResponse")
	}
	return
}

// updateSnapshot implements the OCIOperation interface (enables retrying operations)
func (client FileStorageClient) updateSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/snapshots/{snapshotId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/filestorage/20171215/Snapshot/UpdateSnapshot"
		err = common.PostProcessServiceError(err, "FileStorage", "UpdateSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
