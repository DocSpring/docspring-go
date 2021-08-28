# \PDFApi

All URIs are relative to *https://api.docspring.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFieldsToTemplate**](PDFApi.md#AddFieldsToTemplate) | **Put** /templates/{template_id}/add_fields | Add new fields to a Template
[**BatchGeneratePdfV1**](PDFApi.md#BatchGeneratePdfV1) | **Post** /templates/{template_id}/submissions/batch | Generates multiple PDFs
[**BatchGeneratePdfs**](PDFApi.md#BatchGeneratePdfs) | **Post** /submissions/batches | Generates multiple PDFs
[**CombinePdfs**](PDFApi.md#CombinePdfs) | **Post** /combined_submissions?v&#x3D;2 | Merge submission PDFs, template PDFs, or custom files
[**CombineSubmissions**](PDFApi.md#CombineSubmissions) | **Post** /combined_submissions | Merge generated PDFs together
[**CopyTemplate**](PDFApi.md#CopyTemplate) | **Post** /templates/{template_id}/copy | Copy a Template
[**CreateCustomFileFromUpload**](PDFApi.md#CreateCustomFileFromUpload) | **Post** /custom_files | Create a new custom file from a cached presign upload
[**CreateDataRequestToken**](PDFApi.md#CreateDataRequestToken) | **Post** /data_requests/{data_request_id}/tokens | Creates a new data request token for form authentication
[**CreateFolder**](PDFApi.md#CreateFolder) | **Post** /folders/ | Create a folder
[**CreateHTMLTemplate**](PDFApi.md#CreateHTMLTemplate) | **Post** /templates?desc&#x3D;html | Create a new HTML template
[**CreatePDFTemplate**](PDFApi.md#CreatePDFTemplate) | **Post** /templates | Create a new PDF template with a form POST file upload
[**CreatePDFTemplateFromUpload**](PDFApi.md#CreatePDFTemplateFromUpload) | **Post** /templates?desc&#x3D;cached_upload | Create a new PDF template from a cached presign upload
[**DeleteFolder**](PDFApi.md#DeleteFolder) | **Delete** /folders/{folder_id} | Delete a folder
[**ExpireCombinedSubmission**](PDFApi.md#ExpireCombinedSubmission) | **Delete** /combined_submissions/{combined_submission_id} | Expire a combined submission
[**ExpireSubmission**](PDFApi.md#ExpireSubmission) | **Delete** /submissions/{submission_id} | Expire a PDF submission
[**GeneratePDF**](PDFApi.md#GeneratePDF) | **Post** /templates/{template_id}/submissions | Generates a new PDF
[**GetCombinedSubmission**](PDFApi.md#GetCombinedSubmission) | **Get** /combined_submissions/{combined_submission_id} | Check the status of a combined submission (merged PDFs)
[**GetDataRequest**](PDFApi.md#GetDataRequest) | **Get** /data_requests/{data_request_id} | Look up a submission data request
[**GetPresignUrl**](PDFApi.md#GetPresignUrl) | **Get** /uploads/presign | Get a presigned URL so that you can upload a file to our AWS S3 bucket
[**GetSubmission**](PDFApi.md#GetSubmission) | **Get** /submissions/{submission_id} | Check the status of a PDF
[**GetSubmissionBatch**](PDFApi.md#GetSubmissionBatch) | **Get** /submissions/batches/{submission_batch_id} | Check the status of a submission batch job
[**GetTemplate**](PDFApi.md#GetTemplate) | **Get** /templates/{template_id} | Get a single template
[**GetTemplateSchema**](PDFApi.md#GetTemplateSchema) | **Get** /templates/{template_id}/schema | Fetch the JSON schema for a template
[**ListFolders**](PDFApi.md#ListFolders) | **Get** /folders/ | Get a list of all folders
[**ListTemplates**](PDFApi.md#ListTemplates) | **Get** /templates | Get a list of all templates
[**MoveFolderToFolder**](PDFApi.md#MoveFolderToFolder) | **Post** /folders/{folder_id}/move | Move a folder
[**MoveTemplateToFolder**](PDFApi.md#MoveTemplateToFolder) | **Post** /templates/{template_id}/move | Move Template to folder
[**RenameFolder**](PDFApi.md#RenameFolder) | **Post** /folders/{folder_id}/rename | Rename a folder
[**TestAuthentication**](PDFApi.md#TestAuthentication) | **Get** /authentication | Test Authentication
[**UpdateDataRequest**](PDFApi.md#UpdateDataRequest) | **Put** /data_requests/{data_request_id} | Update a submission data request
[**UpdateTemplate**](PDFApi.md#UpdateTemplate) | **Put** /templates/{template_id} | Update a Template


# **AddFieldsToTemplate**
> AddFieldsTemplateResponse AddFieldsToTemplate(ctx, templateId, addFieldsData)
Add new fields to a Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**|  | 
  **addFieldsData** | [**AddFieldsData**](AddFieldsData.md)|  | 

### Return type

[**AddFieldsTemplateResponse**](add_fields_template_response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchGeneratePdfV1**
> []CreateSubmissionResponse BatchGeneratePdfV1(ctx, templateId, mapStringinterface)
Generates multiple PDFs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**|  | 
  **mapStringinterface** | [**[]map[string]interface{}**](array.md)|  | 

### Return type

[**[]CreateSubmissionResponse**](create_submission_response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchGeneratePdfs**
> CreateSubmissionBatchResponse BatchGeneratePdfs(ctx, submissionBatchData)
Generates multiple PDFs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **submissionBatchData** | [**SubmissionBatchData**](SubmissionBatchData.md)|  | 

### Return type

[**CreateSubmissionBatchResponse**](create_submission_batch_response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CombinePdfs**
> CreateCombinedSubmissionResponse CombinePdfs(ctx, combinePdfsData)
Merge submission PDFs, template PDFs, or custom files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **combinePdfsData** | [**CombinePdfsData**](CombinePdfsData.md)|  | 

### Return type

[**CreateCombinedSubmissionResponse**](create_combined_submission_response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CombineSubmissions**
> CreateCombinedSubmissionResponse CombineSubmissions(ctx, combinedSubmissionData)
Merge generated PDFs together

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **combinedSubmissionData** | [**CombinedSubmissionData**](CombinedSubmissionData.md)|  | 

### Return type

[**CreateCombinedSubmissionResponse**](create_combined_submission_response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyTemplate**
> Template CopyTemplate(ctx, templateId, copyTemplateData)
Copy a Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**|  | 
  **copyTemplateData** | [**CopyTemplateData**](CopyTemplateData.md)|  | 

### Return type

[**Template**](template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomFileFromUpload**
> CreateCustomFileResponse CreateCustomFileFromUpload(ctx, createCustomFileData)
Create a new custom file from a cached presign upload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createCustomFileData** | [**CreateCustomFileData**](CreateCustomFileData.md)|  | 

### Return type

[**CreateCustomFileResponse**](create_custom_file_response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDataRequestToken**
> CreateSubmissionDataRequestTokenResponse CreateDataRequestToken(ctx, dataRequestId)
Creates a new data request token for form authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataRequestId** | **string**|  | 

### Return type

[**CreateSubmissionDataRequestTokenResponse**](create_submission_data_request_token_response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFolder**
> Folder CreateFolder(ctx, createFolderData)
Create a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createFolderData** | [**CreateFolderData**](CreateFolderData.md)|  | 

### Return type

[**Folder**](folder.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHTMLTemplate**
> PendingTemplate CreateHTMLTemplate(ctx, createHtmlTemplateData)
Create a new HTML template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createHtmlTemplateData** | [**CreateHtmlTemplateData**](CreateHtmlTemplateData.md)|  | 

### Return type

[**PendingTemplate**](pending_template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePDFTemplate**
> PendingTemplate CreatePDFTemplate(ctx, templateDocument, templateName, optional)
Create a new PDF template with a form POST file upload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateDocument** | ***os.File*****os.File**|  | 
  **templateName** | **string**|  | 
 **optional** | ***CreatePDFTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreatePDFTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **templateParentFolderId** | **optional.String**|  | 

### Return type

[**PendingTemplate**](pending_template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePDFTemplateFromUpload**
> PendingTemplate CreatePDFTemplateFromUpload(ctx, createTemplateFromUploadData)
Create a new PDF template from a cached presign upload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createTemplateFromUploadData** | [**CreateTemplateFromUploadData**](CreateTemplateFromUploadData.md)|  | 

### Return type

[**PendingTemplate**](pending_template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFolder**
> Folder DeleteFolder(ctx, folderId)
Delete a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**|  | 

### Return type

[**Folder**](folder.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpireCombinedSubmission**
> CombinedSubmission ExpireCombinedSubmission(ctx, combinedSubmissionId)
Expire a combined submission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **combinedSubmissionId** | **string**|  | 

### Return type

[**CombinedSubmission**](combined_submission.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpireSubmission**
> Submission ExpireSubmission(ctx, submissionId)
Expire a PDF submission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **submissionId** | **string**|  | 

### Return type

[**Submission**](submission.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GeneratePDF**
> CreateSubmissionResponse GeneratePDF(ctx, templateId, submissionData)
Generates a new PDF

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**|  | 
  **submissionData** | [**SubmissionData**](SubmissionData.md)|  | 

### Return type

[**CreateSubmissionResponse**](create_submission_response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCombinedSubmission**
> CombinedSubmission GetCombinedSubmission(ctx, combinedSubmissionId)
Check the status of a combined submission (merged PDFs)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **combinedSubmissionId** | **string**|  | 

### Return type

[**CombinedSubmission**](combined_submission.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDataRequest**
> SubmissionDataRequest GetDataRequest(ctx, dataRequestId)
Look up a submission data request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataRequestId** | **string**|  | 

### Return type

[**SubmissionDataRequest**](submission_data_request.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPresignUrl**
> map[string]map[string]interface{} GetPresignUrl(ctx, )
Get a presigned URL so that you can upload a file to our AWS S3 bucket

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubmission**
> Submission GetSubmission(ctx, submissionId, optional)
Check the status of a PDF

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **submissionId** | **string**|  | 
 **optional** | ***GetSubmissionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSubmissionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeData** | **optional.Bool**|  | 

### Return type

[**Submission**](submission.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubmissionBatch**
> SubmissionBatch GetSubmissionBatch(ctx, submissionBatchId, optional)
Check the status of a submission batch job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **submissionBatchId** | **string**|  | 
 **optional** | ***GetSubmissionBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSubmissionBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSubmissions** | **optional.Bool**|  | 

### Return type

[**SubmissionBatch**](submission_batch.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplate**
> Template GetTemplate(ctx, templateId)
Get a single template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**|  | 

### Return type

[**Template**](template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplateSchema**
> map[string]map[string]interface{} GetTemplateSchema(ctx, templateId)
Fetch the JSON schema for a template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**|  | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFolders**
> []Folder ListFolders(ctx, optional)
Get a list of all folders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFoldersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentFolderId** | **optional.String**| Filter By Folder Id | 

### Return type

[**[]Folder**](folder.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTemplates**
> []Template ListTemplates(ctx, optional)
Get a list of all templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| Search By Name | 
 **parentFolderId** | **optional.String**| Filter By Folder Id | 
 **page** | **optional.Int32**| Default: 1 | 
 **perPage** | **optional.Int32**| Default: 50 | 

### Return type

[**[]Template**](template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveFolderToFolder**
> Folder MoveFolderToFolder(ctx, folderId, moveFolderData)
Move a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**|  | 
  **moveFolderData** | [**MoveFolderData**](MoveFolderData.md)|  | 

### Return type

[**Folder**](folder.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveTemplateToFolder**
> Template MoveTemplateToFolder(ctx, templateId, moveTemplateData)
Move Template to folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**|  | 
  **moveTemplateData** | [**MoveTemplateData**](MoveTemplateData.md)|  | 

### Return type

[**Template**](template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenameFolder**
> RenameFolder(ctx, folderId, renameFolderData)
Rename a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**|  | 
  **renameFolderData** | [**RenameFolderData**](RenameFolderData.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestAuthentication**
> AuthenticationSuccessResponse TestAuthentication(ctx, )
Test Authentication

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthenticationSuccessResponse**](authentication_success_response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDataRequest**
> UpdateDataRequestResponse UpdateDataRequest(ctx, dataRequestId, updateSubmissionDataRequestData)
Update a submission data request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataRequestId** | **string**|  | 
  **updateSubmissionDataRequestData** | [**UpdateSubmissionDataRequestData**](UpdateSubmissionDataRequestData.md)|  | 

### Return type

[**UpdateDataRequestResponse**](update_data_request_response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTemplate**
> UpdateTemplateResponse UpdateTemplate(ctx, templateId, updateTemplateData)
Update a Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**|  | 
  **updateTemplateData** | [**UpdateTemplateData**](UpdateTemplateData.md)|  | 

### Return type

[**UpdateTemplateResponse**](update_template_response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

