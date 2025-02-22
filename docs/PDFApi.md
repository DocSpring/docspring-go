# \PDFAPI

All URIs are relative to *https://sync.api.docspring.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFieldsToTemplate**](PDFAPI.md#AddFieldsToTemplate) | **Put** /templates/{template_id}/add_fields | Add new fields to a Template
[**BatchGeneratePdfs**](PDFAPI.md#BatchGeneratePdfs) | **Post** /submissions/batches | Generates multiple PDFs
[**CombinePdfs**](PDFAPI.md#CombinePdfs) | **Post** /combined_submissions?v&#x3D;2 | Merge submission PDFs, template PDFs, or custom files
[**CombineSubmissions**](PDFAPI.md#CombineSubmissions) | **Post** /combined_submissions | Merge generated PDFs together
[**CopyTemplate**](PDFAPI.md#CopyTemplate) | **Post** /templates/{template_id}/copy | Copy a Template
[**CreateCustomFileFromUpload**](PDFAPI.md#CreateCustomFileFromUpload) | **Post** /custom_files | Create a new custom file from a cached presign upload
[**CreateDataRequestEvent**](PDFAPI.md#CreateDataRequestEvent) | **Post** /data_requests/{data_request_id}/events | Creates a new event for emailing a signee a request for signature
[**CreateDataRequestToken**](PDFAPI.md#CreateDataRequestToken) | **Post** /data_requests/{data_request_id}/tokens | Creates a new data request token for form authentication
[**CreateFolder**](PDFAPI.md#CreateFolder) | **Post** /folders/ | Create a folder
[**CreateHTMLTemplate**](PDFAPI.md#CreateHTMLTemplate) | **Post** /templates?endpoint_description&#x3D;html | Create a new HTML template
[**CreatePDFTemplate**](PDFAPI.md#CreatePDFTemplate) | **Post** /templates | Create a new PDF template with a form POST file upload
[**CreatePDFTemplateFromUpload**](PDFAPI.md#CreatePDFTemplateFromUpload) | **Post** /templates?endpoint_description&#x3D;cached_upload | Create a new PDF template from a cached presign upload
[**DeleteFolder**](PDFAPI.md#DeleteFolder) | **Delete** /folders/{folder_id} | Delete a folder
[**DeleteTemplate**](PDFAPI.md#DeleteTemplate) | **Delete** /templates/{template_id} | Delete a template
[**ExpireCombinedSubmission**](PDFAPI.md#ExpireCombinedSubmission) | **Delete** /combined_submissions/{combined_submission_id} | Expire a combined submission
[**ExpireSubmission**](PDFAPI.md#ExpireSubmission) | **Delete** /submissions/{submission_id} | Expire a PDF submission
[**GeneratePdf**](PDFAPI.md#GeneratePdf) | **Post** /templates/{template_id}/submissions | Generates a new PDF
[**GeneratePdfForHtmlTemplate**](PDFAPI.md#GeneratePdfForHtmlTemplate) | **Post** /templates/{template_id}/submissions?endpoint_description&#x3D;html_templates | Generates a new PDF for an HTML template
[**GeneratePreview**](PDFAPI.md#GeneratePreview) | **Post** /submissions/{submission_id}/generate_preview | Generated a preview PDF for partially completed data requests
[**GetCombinedSubmission**](PDFAPI.md#GetCombinedSubmission) | **Get** /combined_submissions/{combined_submission_id} | Check the status of a combined submission (merged PDFs)
[**GetDataRequest**](PDFAPI.md#GetDataRequest) | **Get** /data_requests/{data_request_id} | Look up a submission data request
[**GetFullTemplate**](PDFAPI.md#GetFullTemplate) | **Get** /templates/{template_id}?full&#x3D;true | Fetch the full template attributes
[**GetPresignUrl**](PDFAPI.md#GetPresignUrl) | **Get** /uploads/presign | Get a presigned URL so that you can upload a file to our AWS S3 bucket
[**GetSubmission**](PDFAPI.md#GetSubmission) | **Get** /submissions/{submission_id} | Check the status of a PDF
[**GetSubmissionBatch**](PDFAPI.md#GetSubmissionBatch) | **Get** /submissions/batches/{submission_batch_id} | Check the status of a submission batch job
[**GetTemplate**](PDFAPI.md#GetTemplate) | **Get** /templates/{template_id} | Check the status of an uploaded template
[**GetTemplateSchema**](PDFAPI.md#GetTemplateSchema) | **Get** /templates/{template_id}/schema | Fetch the JSON schema for a template
[**ListCombinedSubmissions**](PDFAPI.md#ListCombinedSubmissions) | **Get** /combined_submissions | Get a list of all combined submissions
[**ListFolders**](PDFAPI.md#ListFolders) | **Get** /folders/ | Get a list of all folders
[**ListSubmissions**](PDFAPI.md#ListSubmissions) | **Get** /submissions | List all submissions
[**ListTemplateSubmissions**](PDFAPI.md#ListTemplateSubmissions) | **Get** /templates/{template_id}/submissions | List all submissions for a given template
[**ListTemplates**](PDFAPI.md#ListTemplates) | **Get** /templates | Get a list of all templates
[**MoveFolderToFolder**](PDFAPI.md#MoveFolderToFolder) | **Post** /folders/{folder_id}/move | Move a folder
[**MoveTemplateToFolder**](PDFAPI.md#MoveTemplateToFolder) | **Post** /templates/{template_id}/move | Move Template to folder
[**RenameFolder**](PDFAPI.md#RenameFolder) | **Post** /folders/{folder_id}/rename | Rename a folder
[**TestAuthentication**](PDFAPI.md#TestAuthentication) | **Get** /authentication | Test Authentication
[**UpdateDataRequest**](PDFAPI.md#UpdateDataRequest) | **Put** /data_requests/{data_request_id} | Update a submission data request
[**UpdateTemplate**](PDFAPI.md#UpdateTemplate) | **Put** /templates/{template_id} | Update a Template



## AddFieldsToTemplate

> TemplateAddFieldsResponse AddFieldsToTemplate(ctx, templateId).Data(data).Execute()

Add new fields to a Template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateId := "tpl_1234567890abcdef02" // string | 
	data := *openapiclient.NewAddFieldsData([]map[string]interface{}{map[string]interface{}(123)}) // AddFieldsData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.AddFieldsToTemplate(context.Background(), templateId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.AddFieldsToTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFieldsToTemplate`: TemplateAddFieldsResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.AddFieldsToTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFieldsToTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**AddFieldsData**](AddFieldsData.md) |  | 

### Return type

[**TemplateAddFieldsResponse**](TemplateAddFieldsResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchGeneratePdfs

> BatchGeneratePdfs201Response BatchGeneratePdfs(ctx).Data(data).Wait(wait).Execute()

Generates multiple PDFs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	data := *openapiclient.NewSubmissionBatchData([]map[string]interface{}{map[string]interface{}(123)}) // SubmissionBatchData | 
	wait := true // bool | Wait for submission batch to be processed before returning. Set to false to return immediately. Default: true (on sync.* subdomain) (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.BatchGeneratePdfs(context.Background()).Data(data).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.BatchGeneratePdfs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchGeneratePdfs`: BatchGeneratePdfs201Response
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.BatchGeneratePdfs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchGeneratePdfsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**SubmissionBatchData**](SubmissionBatchData.md) |  | 
 **wait** | **bool** | Wait for submission batch to be processed before returning. Set to false to return immediately. Default: true (on sync.* subdomain) | [default to true]

### Return type

[**BatchGeneratePdfs201Response**](BatchGeneratePdfs201Response.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CombinePdfs

> CreateCombinedSubmissionResponse CombinePdfs(ctx).Data(data).Execute()

Merge submission PDFs, template PDFs, or custom files

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	data := *openapiclient.NewCombinePdfsData([]map[string]interface{}{map[string]interface{}(123)}) // CombinePdfsData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.CombinePdfs(context.Background()).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.CombinePdfs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CombinePdfs`: CreateCombinedSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.CombinePdfs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCombinePdfsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**CombinePdfsData**](CombinePdfsData.md) |  | 

### Return type

[**CreateCombinedSubmissionResponse**](CreateCombinedSubmissionResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CombineSubmissions

> CreateCombinedSubmissionResponse CombineSubmissions(ctx).Data(data).Wait(wait).Execute()

Merge generated PDFs together

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	data := *openapiclient.NewCombinedSubmissionData([]string{"SubmissionIds_example"}) // CombinedSubmissionData | 
	wait := true // bool | Wait for combined submission to be processed before returning. Set to false to return immediately. Default: true (on sync.* subdomain) (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.CombineSubmissions(context.Background()).Data(data).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.CombineSubmissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CombineSubmissions`: CreateCombinedSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.CombineSubmissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCombineSubmissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**CombinedSubmissionData**](CombinedSubmissionData.md) |  | 
 **wait** | **bool** | Wait for combined submission to be processed before returning. Set to false to return immediately. Default: true (on sync.* subdomain) | [default to true]

### Return type

[**CreateCombinedSubmissionResponse**](CreateCombinedSubmissionResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyTemplate

> TemplatePreview CopyTemplate(ctx, templateId).Options(options).Execute()

Copy a Template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | 
	options := *openapiclient.NewCopyTemplateOptions("ParentFolderId_example") // CopyTemplateOptions |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.CopyTemplate(context.Background(), templateId).Options(options).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.CopyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyTemplate`: TemplatePreview
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.CopyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **options** | [**CopyTemplateOptions**](CopyTemplateOptions.md) |  | 

### Return type

[**TemplatePreview**](TemplatePreview.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomFileFromUpload

> CreateCustomFileResponse CreateCustomFileFromUpload(ctx).Data(data).Execute()

Create a new custom file from a cached presign upload

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	data := *openapiclient.NewCreateCustomFileData("CacheId_example") // CreateCustomFileData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.CreateCustomFileFromUpload(context.Background()).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.CreateCustomFileFromUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomFileFromUpload`: CreateCustomFileResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.CreateCustomFileFromUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFileFromUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**CreateCustomFileData**](CreateCustomFileData.md) |  | 

### Return type

[**CreateCustomFileResponse**](CreateCustomFileResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataRequestEvent

> CreateSubmissionDataRequestEventResponse CreateDataRequestEvent(ctx, dataRequestId).Event(event).Execute()

Creates a new event for emailing a signee a request for signature

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dataRequestId := "drq_1234567890abcdef01" // string | 
	event := *openapiclient.NewCreateSubmissionDataRequestEventRequest("EventType_example") // CreateSubmissionDataRequestEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.CreateDataRequestEvent(context.Background(), dataRequestId).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.CreateDataRequestEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataRequestEvent`: CreateSubmissionDataRequestEventResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.CreateDataRequestEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataRequestEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **event** | [**CreateSubmissionDataRequestEventRequest**](CreateSubmissionDataRequestEventRequest.md) |  | 

### Return type

[**CreateSubmissionDataRequestEventResponse**](CreateSubmissionDataRequestEventResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataRequestToken

> CreateSubmissionDataRequestTokenResponse CreateDataRequestToken(ctx, dataRequestId).Type_(type_).Execute()

Creates a new data request token for form authentication

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dataRequestId := "drq_1234567890abcdef01" // string | 
	type_ := "api" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.CreateDataRequestToken(context.Background(), dataRequestId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.CreateDataRequestToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataRequestToken`: CreateSubmissionDataRequestTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.CreateDataRequestToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataRequestTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** |  | 

### Return type

[**CreateSubmissionDataRequestTokenResponse**](CreateSubmissionDataRequestTokenResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFolder

> Folder CreateFolder(ctx).Data(data).Execute()

Create a folder

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	data := *openapiclient.NewCreateFolderData(map[string]interface{}(123)) // CreateFolderData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.CreateFolder(context.Background()).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.CreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolder`: Folder
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.CreateFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**CreateFolderData**](CreateFolderData.md) |  | 

### Return type

[**Folder**](Folder.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHTMLTemplate

> TemplatePreview CreateHTMLTemplate(ctx).Data(data).Execute()

Create a new HTML template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	data := *openapiclient.NewCreateHtmlTemplate(map[string]interface{}(123)) // CreateHtmlTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.CreateHTMLTemplate(context.Background()).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.CreateHTMLTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHTMLTemplate`: TemplatePreview
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.CreateHTMLTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHTMLTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**CreateHtmlTemplate**](CreateHtmlTemplate.md) |  | 

### Return type

[**TemplatePreview**](TemplatePreview.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePDFTemplate

> TemplatePreview CreatePDFTemplate(ctx).TemplateDocument(templateDocument).TemplateName(templateName).Wait(wait).TemplateDescription(templateDescription).TemplateParentFolderId(templateParentFolderId).Execute()

Create a new PDF template with a form POST file upload

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateDocument := os.NewFile(1234, "some_file") // *os.File | 
	templateName := "templateName_example" // string | 
	wait := true // bool | Wait for template document to be processed before returning. Set to false to return immediately. Default: true (on sync.* subdomain) (optional) (default to true)
	templateDescription := "templateDescription_example" // string |  (optional)
	templateParentFolderId := "templateParentFolderId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.CreatePDFTemplate(context.Background()).TemplateDocument(templateDocument).TemplateName(templateName).Wait(wait).TemplateDescription(templateDescription).TemplateParentFolderId(templateParentFolderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.CreatePDFTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePDFTemplate`: TemplatePreview
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.CreatePDFTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePDFTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateDocument** | ***os.File** |  | 
 **templateName** | **string** |  | 
 **wait** | **bool** | Wait for template document to be processed before returning. Set to false to return immediately. Default: true (on sync.* subdomain) | [default to true]
 **templateDescription** | **string** |  | 
 **templateParentFolderId** | **string** |  | 

### Return type

[**TemplatePreview**](TemplatePreview.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePDFTemplateFromUpload

> TemplatePreview CreatePDFTemplateFromUpload(ctx).Data(data).Execute()

Create a new PDF template from a cached presign upload

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	data := *openapiclient.NewCreatePdfTemplate(map[string]interface{}(123)) // CreatePdfTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.CreatePDFTemplateFromUpload(context.Background()).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.CreatePDFTemplateFromUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePDFTemplateFromUpload`: TemplatePreview
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.CreatePDFTemplateFromUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePDFTemplateFromUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**CreatePdfTemplate**](CreatePdfTemplate.md) |  | 

### Return type

[**TemplatePreview**](TemplatePreview.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFolder

> Folder DeleteFolder(ctx, folderId).Execute()

Delete a folder

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	folderId := "fld_1234567890abcdef01" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.DeleteFolder(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.DeleteFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFolder`: Folder
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.DeleteFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Folder**](Folder.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplate

> SuccessMultipleErrorsResponse DeleteTemplate(ctx, templateId).Execute()

Delete a template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.DeleteTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.DeleteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTemplate`: SuccessMultipleErrorsResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.DeleteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessMultipleErrorsResponse**](SuccessMultipleErrorsResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireCombinedSubmission

> CombinedSubmission ExpireCombinedSubmission(ctx, combinedSubmissionId).Execute()

Expire a combined submission

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	combinedSubmissionId := "com_1234567890abcdef01" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.ExpireCombinedSubmission(context.Background(), combinedSubmissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.ExpireCombinedSubmission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpireCombinedSubmission`: CombinedSubmission
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.ExpireCombinedSubmission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**combinedSubmissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireCombinedSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CombinedSubmission**](CombinedSubmission.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireSubmission

> SubmissionPreview ExpireSubmission(ctx, submissionId).Execute()

Expire a PDF submission

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	submissionId := "sub_1234567890abcdef01" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.ExpireSubmission(context.Background(), submissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.ExpireSubmission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpireSubmission`: SubmissionPreview
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.ExpireSubmission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmissionPreview**](SubmissionPreview.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneratePdf

> CreateSubmissionResponse GeneratePdf(ctx, templateId).Submission(submission).Wait(wait).Execute()

Generates a new PDF

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | 
	submission := *openapiclient.NewCreatePdfSubmissionData(map[string]interface{}(123)) // CreatePdfSubmissionData | 
	wait := true // bool | Wait for submission to be processed before returning. Set to false to return immediately. Default: true (on sync.* subdomain) (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.GeneratePdf(context.Background(), templateId).Submission(submission).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.GeneratePdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GeneratePdf`: CreateSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.GeneratePdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submission** | [**CreatePdfSubmissionData**](CreatePdfSubmissionData.md) |  | 
 **wait** | **bool** | Wait for submission to be processed before returning. Set to false to return immediately. Default: true (on sync.* subdomain) | [default to true]

### Return type

[**CreateSubmissionResponse**](CreateSubmissionResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneratePdfForHtmlTemplate

> CreateSubmissionResponse GeneratePdfForHtmlTemplate(ctx, templateId).Submission(submission).Wait(wait).Execute()

Generates a new PDF for an HTML template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateId := "tpl_HTML567890abcdef01" // string | 
	submission := *openapiclient.NewCreateHtmlSubmissionData() // CreateHtmlSubmissionData | 
	wait := true // bool | Wait for submission to be processed before returning. Set to false to return immediately. Default: true (on sync.* subdomain) (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.GeneratePdfForHtmlTemplate(context.Background(), templateId).Submission(submission).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.GeneratePdfForHtmlTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GeneratePdfForHtmlTemplate`: CreateSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.GeneratePdfForHtmlTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePdfForHtmlTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submission** | [**CreateHtmlSubmissionData**](CreateHtmlSubmissionData.md) |  | 
 **wait** | **bool** | Wait for submission to be processed before returning. Set to false to return immediately. Default: true (on sync.* subdomain) | [default to true]

### Return type

[**CreateSubmissionResponse**](CreateSubmissionResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneratePreview

> SuccessErrorResponse GeneratePreview(ctx, submissionId).Execute()

Generated a preview PDF for partially completed data requests

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	submissionId := "sub_1234567890abcdef01" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.GeneratePreview(context.Background(), submissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.GeneratePreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GeneratePreview`: SuccessErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.GeneratePreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessErrorResponse**](SuccessErrorResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCombinedSubmission

> CombinedSubmission GetCombinedSubmission(ctx, combinedSubmissionId).Execute()

Check the status of a combined submission (merged PDFs)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	combinedSubmissionId := "com_1234567890abcdef01" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.GetCombinedSubmission(context.Background(), combinedSubmissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.GetCombinedSubmission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCombinedSubmission`: CombinedSubmission
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.GetCombinedSubmission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**combinedSubmissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCombinedSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CombinedSubmission**](CombinedSubmission.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataRequest

> SubmissionDataRequestShow GetDataRequest(ctx, dataRequestId).Execute()

Look up a submission data request

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dataRequestId := "drq_1234567890abcdef01" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.GetDataRequest(context.Background(), dataRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.GetDataRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataRequest`: SubmissionDataRequestShow
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.GetDataRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmissionDataRequestShow**](SubmissionDataRequestShow.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFullTemplate

> Template GetFullTemplate(ctx, templateId).Execute()

Fetch the full template attributes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.GetFullTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.GetFullTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFullTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.GetFullTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFullTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Template**](Template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPresignUrl

> UploadPresignResponse GetPresignUrl(ctx).Execute()

Get a presigned URL so that you can upload a file to our AWS S3 bucket

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.GetPresignUrl(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.GetPresignUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPresignUrl`: UploadPresignResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.GetPresignUrl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPresignUrlRequest struct via the builder pattern


### Return type

[**UploadPresignResponse**](UploadPresignResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubmission

> Submission GetSubmission(ctx, submissionId).IncludeData(includeData).Execute()

Check the status of a PDF

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	submissionId := "sub_1234567890abcdef01" // string | 
	includeData := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.GetSubmission(context.Background(), submissionId).IncludeData(includeData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.GetSubmission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubmission`: Submission
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.GetSubmission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeData** | **bool** |  | 

### Return type

[**Submission**](Submission.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubmissionBatch

> SubmissionBatchWithSubmissions GetSubmissionBatch(ctx, submissionBatchId).IncludeSubmissions(includeSubmissions).Execute()

Check the status of a submission batch job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	submissionBatchId := "sbb_1234567890abcdef01" // string | 
	includeSubmissions := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.GetSubmissionBatch(context.Background(), submissionBatchId).IncludeSubmissions(includeSubmissions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.GetSubmissionBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubmissionBatch`: SubmissionBatchWithSubmissions
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.GetSubmissionBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionBatchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubmissionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSubmissions** | **bool** |  | 

### Return type

[**SubmissionBatchWithSubmissions**](SubmissionBatchWithSubmissions.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplate

> TemplatePreview GetTemplate(ctx, templateId).Execute()

Check the status of an uploaded template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.GetTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.GetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplate`: TemplatePreview
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.GetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplatePreview**](TemplatePreview.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateSchema

> JsonSchema GetTemplateSchema(ctx, templateId).Execute()

Fetch the JSON schema for a template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.GetTemplateSchema(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.GetTemplateSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateSchema`: JsonSchema
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.GetTemplateSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JsonSchema**](JsonSchema.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCombinedSubmissions

> []CombinedSubmission ListCombinedSubmissions(ctx).Page(page).PerPage(perPage).Execute()

Get a list of all combined submissions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(2) // int32 | Default: 1 (optional)
	perPage := int32(1) // int32 | Default: 50 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.ListCombinedSubmissions(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.ListCombinedSubmissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCombinedSubmissions`: []CombinedSubmission
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.ListCombinedSubmissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCombinedSubmissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Default: 1 | 
 **perPage** | **int32** | Default: 50 | 

### Return type

[**[]CombinedSubmission**](CombinedSubmission.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFolders

> []Folder ListFolders(ctx).ParentFolderId(parentFolderId).Execute()

Get a list of all folders

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	parentFolderId := "fld_1234567890abcdef02" // string | Filter By Folder Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.ListFolders(context.Background()).ParentFolderId(parentFolderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.ListFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFolders`: []Folder
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.ListFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentFolderId** | **string** | Filter By Folder Id | 

### Return type

[**[]Folder**](Folder.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubmissions

> ListSubmissionsResponse ListSubmissions(ctx).Cursor(cursor).Limit(limit).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Type_(type_).IncludeData(includeData).Execute()

List all submissions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cursor := "sub_1234567890abcdef12" // string |  (optional)
	limit := float32(3) // float32 |  (optional)
	createdAfter := "2019-01-01T09:00:00-05:00" // string |  (optional)
	createdBefore := "2020-01-01T09:00:00.000+0200" // string |  (optional)
	type_ := "test" // string |  (optional)
	includeData := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.ListSubmissions(context.Background()).Cursor(cursor).Limit(limit).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Type_(type_).IncludeData(includeData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.ListSubmissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubmissions`: ListSubmissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.ListSubmissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubmissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **limit** | **float32** |  | 
 **createdAfter** | **string** |  | 
 **createdBefore** | **string** |  | 
 **type_** | **string** |  | 
 **includeData** | **bool** |  | 

### Return type

[**ListSubmissionsResponse**](ListSubmissionsResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplateSubmissions

> ListSubmissionsResponse ListTemplateSubmissions(ctx, templateId).Cursor(cursor).Limit(limit).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Type_(type_).IncludeData(includeData).Execute()

List all submissions for a given template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateId := "tpl_1234567890abcdef02" // string | 
	cursor := "cursor_example" // string |  (optional)
	limit := float32(8.14) // float32 |  (optional)
	createdAfter := "createdAfter_example" // string |  (optional)
	createdBefore := "createdBefore_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	includeData := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.ListTemplateSubmissions(context.Background(), templateId).Cursor(cursor).Limit(limit).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Type_(type_).IncludeData(includeData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.ListTemplateSubmissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTemplateSubmissions`: ListSubmissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.ListTemplateSubmissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplateSubmissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** |  | 
 **limit** | **float32** |  | 
 **createdAfter** | **string** |  | 
 **createdBefore** | **string** |  | 
 **type_** | **string** |  | 
 **includeData** | **bool** |  | 

### Return type

[**ListSubmissionsResponse**](ListSubmissionsResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplates

> []TemplatePreview ListTemplates(ctx).Query(query).ParentFolderId(parentFolderId).Page(page).PerPage(perPage).Execute()

Get a list of all templates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	query := "2" // string | Search By Name (optional)
	parentFolderId := "fld_1234567890abcdef01" // string | Filter By Folder Id (optional)
	page := int32(2) // int32 | Default: 1 (optional)
	perPage := int32(1) // int32 | Default: 50 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.ListTemplates(context.Background()).Query(query).ParentFolderId(parentFolderId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.ListTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTemplates`: []TemplatePreview
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.ListTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Search By Name | 
 **parentFolderId** | **string** | Filter By Folder Id | 
 **page** | **int32** | Default: 1 | 
 **perPage** | **int32** | Default: 50 | 

### Return type

[**[]TemplatePreview**](TemplatePreview.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveFolderToFolder

> Folder MoveFolderToFolder(ctx, folderId).Data(data).Execute()

Move a folder

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	folderId := "fld_1234567890abcdef01" // string | 
	data := *openapiclient.NewMoveFolderData() // MoveFolderData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.MoveFolderToFolder(context.Background(), folderId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.MoveFolderToFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveFolderToFolder`: Folder
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.MoveFolderToFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveFolderToFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**MoveFolderData**](MoveFolderData.md) |  | 

### Return type

[**Folder**](Folder.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveTemplateToFolder

> TemplatePreview MoveTemplateToFolder(ctx, templateId).Data(data).Execute()

Move Template to folder

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | 
	data := *openapiclient.NewMoveTemplateData("ParentFolderId_example") // MoveTemplateData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.MoveTemplateToFolder(context.Background(), templateId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.MoveTemplateToFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveTemplateToFolder`: TemplatePreview
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.MoveTemplateToFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveTemplateToFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**MoveTemplateData**](MoveTemplateData.md) |  | 

### Return type

[**TemplatePreview**](TemplatePreview.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameFolder

> Folder RenameFolder(ctx, folderId).Data(data).Execute()

Rename a folder

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	folderId := "fld_1234567890abcdef01" // string | 
	data := *openapiclient.NewRenameFolderData("Name_example") // RenameFolderData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.RenameFolder(context.Background(), folderId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.RenameFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameFolder`: Folder
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.RenameFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**RenameFolderData**](RenameFolderData.md) |  | 

### Return type

[**Folder**](Folder.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestAuthentication

> SuccessErrorResponse TestAuthentication(ctx).Execute()

Test Authentication

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.TestAuthentication(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.TestAuthentication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestAuthentication`: SuccessErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.TestAuthentication`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestAuthenticationRequest struct via the builder pattern


### Return type

[**SuccessErrorResponse**](SuccessErrorResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataRequest

> CreateSubmissionDataRequestResponse UpdateDataRequest(ctx, dataRequestId).Data(data).Execute()

Update a submission data request

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dataRequestId := "drq_1234567890abcdef01" // string | 
	data := *openapiclient.NewUpdateSubmissionDataRequestData() // UpdateSubmissionDataRequestData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.UpdateDataRequest(context.Background(), dataRequestId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.UpdateDataRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataRequest`: CreateSubmissionDataRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.UpdateDataRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**UpdateSubmissionDataRequestData**](UpdateSubmissionDataRequestData.md) |  | 

### Return type

[**CreateSubmissionDataRequestResponse**](CreateSubmissionDataRequestResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplate

> SuccessMultipleErrorsResponse UpdateTemplate(ctx, templateId).Data(data).Execute()

Update a Template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	templateId := "tpl_1234567890abcdef03" // string | 
	data := *openapiclient.NewUpdateHtmlTemplate(map[string]interface{}(123)) // UpdateHtmlTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFAPI.UpdateTemplate(context.Background(), templateId).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFAPI.UpdateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTemplate`: SuccessMultipleErrorsResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFAPI.UpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**UpdateHtmlTemplate**](UpdateHtmlTemplate.md) |  | 

### Return type

[**SuccessMultipleErrorsResponse**](SuccessMultipleErrorsResponse.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

