// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v48/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v48/goldengate"
)

func init() {
	RegisterResource("oci_golden_gate_deployment_backup", GoldenGateDeploymentBackupResource())
}

func GoldenGateDeploymentBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createGoldenGateDeploymentBackup,
		Read:     readGoldenGateDeploymentBackup,
		Update:   updateGoldenGateDeploymentBackup,
		Delete:   deleteGoldenGateDeploymentBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deployment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"object": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"backup_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_automatic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ogg_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_backup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createGoldenGateDeploymentBackup(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).goldenGateClient()

	return CreateResource(d, sync)
}

func readGoldenGateDeploymentBackup(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).goldenGateClient()

	return ReadResource(sync)
}

func updateGoldenGateDeploymentBackup(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).goldenGateClient()

	return UpdateResource(d, sync)
}

func deleteGoldenGateDeploymentBackup(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).goldenGateClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type GoldenGateDeploymentBackupResourceCrud struct {
	BaseCrud
	Client                 *oci_golden_gate.GoldenGateClient
	Res                    *oci_golden_gate.DeploymentBackup
	DisableNotFoundRetries bool
}

func (s *GoldenGateDeploymentBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GoldenGateDeploymentBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateCreating),
		string(oci_golden_gate.LifecycleStateInProgress),
	}
}

func (s *GoldenGateDeploymentBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateActive),
		string(oci_golden_gate.LifecycleStateNeedsAttention),
		string(oci_golden_gate.LifecycleStateSucceeded),
	}
}

func (s *GoldenGateDeploymentBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateDeleting),
	}
}

func (s *GoldenGateDeploymentBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_golden_gate.LifecycleStateDeleted),
	}
}

func (s *GoldenGateDeploymentBackupResourceCrud) Create() error {
	request := oci_golden_gate.CreateDeploymentBackupRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.CreateDeploymentBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeploymentBackupFromWorkRequest(workId, GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GoldenGateDeploymentBackupResourceCrud) getDeploymentBackupFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_golden_gate.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	deploymentId, err := goldenGateDeploymentBackupWaitForWorkRequest(workId, "deploymentbackup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] operation failed: %v for identifier: %v\n", workId, deploymentId)
		return err
	}
	s.D.SetId(*deploymentId)

	return s.Get()
}

func deploymentBackupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "golden_gate", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_golden_gate.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func goldenGateDeploymentBackupWaitForWorkRequest(wId *string, entityType string, action oci_golden_gate.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_golden_gate.GoldenGateClient) (*string, error) {
	retryPolicy := GetRetryPolicy(disableFoundRetries, "golden_gate")
	retryPolicy.ShouldRetryOperation = deploymentBackupWorkRequestShouldRetryFunc(timeout)

	response := oci_golden_gate.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_golden_gate.OperationStatusInProgress),
			string(oci_golden_gate.OperationStatusAccepted),
		},
		Target: []string{
			string(oci_golden_gate.OperationStatusSucceeded),
			string(oci_golden_gate.OperationStatusFailed),
			string(oci_golden_gate.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_golden_gate.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_golden_gate.OperationStatusFailed || response.Status == oci_golden_gate.OperationStatusCanceled {
		return nil, getErrorFromGoldenGateDeploymentBackupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGoldenGateDeploymentBackupWorkRequest(client *oci_golden_gate.GoldenGateClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_golden_gate.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_golden_gate.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *GoldenGateDeploymentBackupResourceCrud) Get() error {
	request := oci_golden_gate.GetDeploymentBackupRequest{}

	tmp := s.D.Id()
	request.DeploymentBackupId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.GetDeploymentBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DeploymentBackup
	return nil
}

func (s *GoldenGateDeploymentBackupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_golden_gate.UpdateDeploymentBackupRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.DeploymentBackupId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.UpdateDeploymentBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DeploymentBackup
	return nil
}

func (s *GoldenGateDeploymentBackupResourceCrud) Delete() error {
	request := oci_golden_gate.DeleteDeploymentBackupRequest{}

	tmp := s.D.Id()
	request.DeploymentBackupId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.DeleteDeploymentBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := goldenGateDeploymentBackupWaitForWorkRequest(workId, "deploymentbackup",
		oci_golden_gate.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GoldenGateDeploymentBackupResourceCrud) SetData() error {
	s.D.Set("backup_type", s.Res.BackupType)

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeploymentId != nil {
		s.D.Set("deployment_id", *s.Res.DeploymentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutomatic != nil {
		s.D.Set("is_automatic", *s.Res.IsAutomatic)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NamespaceName != nil {
		s.D.Set("namespace", *s.Res.NamespaceName)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	if s.Res.OggVersion != nil {
		s.D.Set("ogg_version", *s.Res.OggVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfBackup != nil {
		s.D.Set("time_of_backup", s.Res.TimeOfBackup.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func DeploymentBackupSummaryToMap(obj oci_golden_gate.DeploymentBackupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["backup_type"] = string(obj.BackupType)

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	if obj.DeploymentId != nil {
		result["deployment_id"] = string(*obj.DeploymentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsAutomatic != nil {
		result["is_automatic"] = bool(*obj.IsAutomatic)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	if obj.OggVersion != nil {
		result["ogg_version"] = string(*obj.OggVersion)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = systemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeOfBackup != nil {
		result["time_of_backup"] = obj.TimeOfBackup.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *GoldenGateDeploymentBackupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_golden_gate.ChangeDeploymentBackupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DeploymentBackupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	_, err := s.Client.ChangeDeploymentBackupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := waitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
