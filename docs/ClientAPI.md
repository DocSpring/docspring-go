# \ClientAPI

All URIs are relative to *https://sync.api.docspring.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTemplateFields**](ClientAPI.md#AddTemplateFields) | **Post** /templates/{template_id}/fields | Add fields to a template
[**CopyTemplate**](ClientAPI.md#CopyTemplate) | **Post** /templates/{template_id}/copy | Copy a template
[**CreateCombinedPdf**](ClientAPI.md#CreateCombinedPdf) | **Post** /combined_pdfs | Combine PDFs
[**CreateCustomFileFromUpload**](ClientAPI.md#CreateCustomFileFromUpload) | **Post** /custom_files | Create a custom file from an upload
[**CreateFolder**](ClientAPI.md#CreateFolder) | **Post** /folders | Create a folder
[**CreateSignatureRequestEvent**](ClientAPI.md#CreateSignatureRequestEvent) | **Post** /signature_requests/{signature_request_id}/events | Record a signature request event
[**CreateSignatureRequestToken**](ClientAPI.md#CreateSignatureRequestToken) | **Post** /signature_requests/{signature_request_id}/tokens | Create a signature request token
[**CreateSubmission**](ClientAPI.md#CreateSubmission) | **Post** /templates/{template_id}/submissions | Generate a PDF
[**CreateSubmissionBatch**](ClientAPI.md#CreateSubmissionBatch) | **Post** /submissions/batches | Generate PDFs in a batch
[**CreateTemplate**](ClientAPI.md#CreateTemplate) | **Post** /templates | Create a template
[**CreateWebhook**](ClientAPI.md#CreateWebhook) | **Post** /webhooks | Create a webhook
[**DeleteCustomFile**](ClientAPI.md#DeleteCustomFile) | **Delete** /custom_files/{custom_file_id} | Delete a custom file
[**DeleteFolder**](ClientAPI.md#DeleteFolder) | **Delete** /folders/{folder_id} | Delete a folder
[**DeleteTemplate**](ClientAPI.md#DeleteTemplate) | **Delete** /templates/{template_id} | Delete a template
[**DeleteWebhook**](ClientAPI.md#DeleteWebhook) | **Delete** /webhooks/{webhook_id} | Delete a webhook
[**ExpireCombinedPdf**](ClientAPI.md#ExpireCombinedPdf) | **Delete** /combined_pdfs/{combined_pdf_id} | Expire a combined PDF
[**ExpireSubmission**](ClientAPI.md#ExpireSubmission) | **Delete** /submissions/{submission_id} | Expire a submission
[**GenerateSubmissionFiles**](ClientAPI.md#GenerateSubmissionFiles) | **Post** /submissions/{submission_id}/generate_files | Generate more files
[**GenerateSubmissionPreview**](ClientAPI.md#GenerateSubmissionPreview) | **Post** /submissions/{submission_id}/generate_preview | Generate a preview PDF
[**GetCombinedPdf**](ClientAPI.md#GetCombinedPdf) | **Get** /combined_pdfs/{combined_pdf_id} | Get a combined PDF
[**GetCustomFile**](ClientAPI.md#GetCustomFile) | **Get** /custom_files/{custom_file_id} | Get a custom file
[**GetFolder**](ClientAPI.md#GetFolder) | **Get** /folders/{folder_id} | Get a folder
[**GetPresignUrl**](ClientAPI.md#GetPresignUrl) | **Get** /uploads/presign | Get a presigned URL to upload a file
[**GetSignatureRequest**](ClientAPI.md#GetSignatureRequest) | **Get** /signature_requests/{signature_request_id} | Get a signature request
[**GetSubmission**](ClientAPI.md#GetSubmission) | **Get** /submissions/{submission_id} | Get a submission
[**GetSubmissionBatch**](ClientAPI.md#GetSubmissionBatch) | **Get** /submissions/batches/{submission_batch_id} | Get a batch
[**GetTemplate**](ClientAPI.md#GetTemplate) | **Get** /templates/{template_id} | Get a template
[**GetTemplateDiff**](ClientAPI.md#GetTemplateDiff) | **Get** /templates/{template_id}/diff | Compare a version of a template with the one before it
[**GetTemplateModelCode**](ClientAPI.md#GetTemplateModelCode) | **Get** /templates/{template_id}/model_code | Generate typed code for a template's submission data
[**GetTemplatePageImages**](ClientAPI.md#GetTemplatePageImages) | **Get** /templates/{template_id}/page_images | Get images of a template's pages
[**GetTemplateSchema**](ClientAPI.md#GetTemplateSchema) | **Get** /templates/{template_id}/schema | Get the JSON Schema of a template's submission data
[**GetWebhook**](ClientAPI.md#GetWebhook) | **Get** /webhooks/{webhook_id} | Get a webhook
[**ListCombinedPdfs**](ClientAPI.md#ListCombinedPdfs) | **Get** /combined_pdfs | List combined PDFs
[**ListCustomFiles**](ClientAPI.md#ListCustomFiles) | **Get** /custom_files | List custom files
[**ListFolders**](ClientAPI.md#ListFolders) | **Get** /folders | List folders
[**ListSubmissions**](ClientAPI.md#ListSubmissions) | **Get** /submissions | List submissions
[**ListTemplateFields**](ClientAPI.md#ListTemplateFields) | **Get** /templates/{template_id}/fields | List a template's fields
[**ListTemplateSubmissions**](ClientAPI.md#ListTemplateSubmissions) | **Get** /templates/{template_id}/submissions | List a template's submissions
[**ListTemplates**](ClientAPI.md#ListTemplates) | **Get** /templates | List templates
[**ListWebhooks**](ClientAPI.md#ListWebhooks) | **Get** /webhooks | List webhooks
[**MoveFolder**](ClientAPI.md#MoveFolder) | **Post** /folders/{folder_id}/move | Move a folder
[**MoveTemplate**](ClientAPI.md#MoveTemplate) | **Post** /templates/{template_id}/move | Move a template to a folder
[**PublishTemplateVersion**](ClientAPI.md#PublishTemplateVersion) | **Post** /templates/{template_id}/publish_version | Publish a template version
[**RemoveTemplateFields**](ClientAPI.md#RemoveTemplateFields) | **Post** /templates/{template_id}/fields/remove | Remove fields from a template
[**RenameFolder**](ClientAPI.md#RenameFolder) | **Post** /folders/{folder_id}/rename | Rename a folder
[**ReorderTemplateFields**](ClientAPI.md#ReorderTemplateFields) | **Post** /templates/{template_id}/fields/reorder | Reorder a template's fields
[**RestoreTemplateVersion**](ClientAPI.md#RestoreTemplateVersion) | **Post** /templates/{template_id}/restore_version | Restore a template version
[**TestAuthentication**](ClientAPI.md#TestAuthentication) | **Get** /authentication | Test authentication
[**UpdateSignatureRequest**](ClientAPI.md#UpdateSignatureRequest) | **Put** /signature_requests/{signature_request_id} | Update a signature request
[**UpdateTemplate**](ClientAPI.md#UpdateTemplate) | **Put** /templates/{template_id} | Update a template
[**UpdateTemplateDocument**](ClientAPI.md#UpdateTemplateDocument) | **Put** /templates/{template_id}/document | Replace a template's PDF
[**UpdateTemplateFields**](ClientAPI.md#UpdateTemplateFields) | **Put** /templates/{template_id}/fields | Update a template's fields
[**UpdateWebhook**](ClientAPI.md#UpdateWebhook) | **Put** /webhooks/{webhook_id} | Update a webhook
[**UploadTemplate**](ClientAPI.md#UploadTemplate) | **Post** /templates/upload | Create a PDF template from an uploaded PDF
[**UploadTemplateDocument**](ClientAPI.md#UploadTemplateDocument) | **Put** /templates/{template_id}/document/upload | Replace a template's PDF with an uploaded PDF
[**ValidateSubmission**](ClientAPI.md#ValidateSubmission) | **Post** /templates/{template_id}/submissions/validate | Check a submission without creating it



## AddTemplateFields

> TemplateFields AddTemplateFields(ctx, templateId).Body(body).Execute()

Add fields to a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewAddTemplateFieldsRequest([]openapiclient.FieldInput{*openapiclient.NewFieldInput()}) // AddTemplateFieldsRequest | The fields to add to the template's draft.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.AddTemplateFields(context.Background(), templateId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.AddTemplateFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTemplateFields`: TemplateFields
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.AddTemplateFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTemplateFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AddTemplateFieldsRequest**](AddTemplateFieldsRequest.md) | The fields to add to the template's draft. | 

### Return type

[**TemplateFields**](TemplateFields.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyTemplate

> Template CopyTemplate(ctx, templateId).Body(body).Execute()

Copy a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewCopyTemplateRequest() // CopyTemplateRequest | Settings for the copy. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.CopyTemplate(context.Background(), templateId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CopyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.CopyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CopyTemplateRequest**](CopyTemplateRequest.md) | Settings for the copy. | 

### Return type

[**Template**](Template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCombinedPdf

> CombinedPdf CreateCombinedPdf(ctx).Body(body).Wait(wait).Execute()

Combine PDFs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	body := *openapiclient.NewCreateCombinedPdfRequest([]openapiclient.CombinedPdfSourceInput{openapiclient.CombinedPdfSourceInput{CombinedPdfSourcePdf: openapiclient.NewCombinedPdfSourcePdf("Type_example", "Id_example")}}) // CreateCombinedPdfRequest | The PDFs to merge, and the combined PDF's settings.
	wait := true // bool | Wait until the PDF has been generated before responding. Default true on sync.api hosts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.CreateCombinedPdf(context.Background()).Body(body).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateCombinedPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCombinedPdf`: CombinedPdf
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.CreateCombinedPdf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCombinedPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateCombinedPdfRequest**](CreateCombinedPdfRequest.md) | The PDFs to merge, and the combined PDF's settings. | 
 **wait** | **bool** | Wait until the PDF has been generated before responding. Default true on sync.api hosts. | 

### Return type

[**CombinedPdf**](CombinedPdf.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomFileFromUpload

> CustomFile CreateCustomFileFromUpload(ctx).Body(body).Execute()

Create a custom file from an upload



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	body := *openapiclient.NewCreateCustomFileFromUploadRequest("CacheId_example") // CreateCustomFileFromUploadRequest | The uploaded file to create the custom file from.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.CreateCustomFileFromUpload(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateCustomFileFromUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomFileFromUpload`: CustomFile
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.CreateCustomFileFromUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFileFromUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateCustomFileFromUploadRequest**](CreateCustomFileFromUploadRequest.md) | The uploaded file to create the custom file from. | 

### Return type

[**CustomFile**](CustomFile.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFolder

> Folder CreateFolder(ctx).Body(body).Execute()

Create a folder



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	body := *openapiclient.NewCreateFolderRequest("Name_example") // CreateFolderRequest | The folder to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.CreateFolder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolder`: Folder
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.CreateFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateFolderRequest**](CreateFolderRequest.md) | The folder to create. | 

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


## CreateSignatureRequestEvent

> SignatureRequestEvent CreateSignatureRequestEvent(ctx, signatureRequestId).Body(body).Execute()

Record a signature request event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	signatureRequestId := "srq_1234567890abcdef01" // string | The signature request's ID.
	body := *openapiclient.NewCreateSignatureRequestEventRequest(openapiclient.SignatureRequestEventType("send_request")) // CreateSignatureRequestEventRequest | The event to record.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.CreateSignatureRequestEvent(context.Background(), signatureRequestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateSignatureRequestEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSignatureRequestEvent`: SignatureRequestEvent
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.CreateSignatureRequestEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signatureRequestId** | **string** | The signature request's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSignatureRequestEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateSignatureRequestEventRequest**](CreateSignatureRequestEventRequest.md) | The event to record. | 

### Return type

[**SignatureRequestEvent**](SignatureRequestEvent.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSignatureRequestToken

> SignatureRequestToken CreateSignatureRequestToken(ctx, signatureRequestId).Body(body).Execute()

Create a signature request token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	signatureRequestId := "srq_1234567890abcdef01" // string | The signature request's ID.
	body := *openapiclient.NewCreateSignatureRequestTokenRequest() // CreateSignatureRequestTokenRequest | The kind of link to make. Send no body for an `api` token. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.CreateSignatureRequestToken(context.Background(), signatureRequestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateSignatureRequestToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSignatureRequestToken`: SignatureRequestToken
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.CreateSignatureRequestToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signatureRequestId** | **string** | The signature request's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSignatureRequestTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateSignatureRequestTokenRequest**](CreateSignatureRequestTokenRequest.md) | The kind of link to make. Send no body for an `api` token. | 

### Return type

[**SignatureRequestToken**](SignatureRequestToken.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubmission

> Submission CreateSubmission(ctx, templateId).Body(body).Version(version).Wait(wait).Execute()

Generate a PDF



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewCreateSubmissionRequest() // CreateSubmissionRequest | The submission's data and settings.
	version := "version_example" // string | The template version to use: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published). (optional)
	wait := true // bool | Wait until the files have been generated before responding. Default true on sync.api hosts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.CreateSubmission(context.Background(), templateId).Body(body).Version(version).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateSubmission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubmission`: Submission
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.CreateSubmission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateSubmissionRequest**](CreateSubmissionRequest.md) | The submission's data and settings. | 
 **version** | **string** | The template version to use: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published). | 
 **wait** | **bool** | Wait until the files have been generated before responding. Default true on sync.api hosts. | 

### Return type

[**Submission**](Submission.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubmissionBatch

> SubmissionBatch CreateSubmissionBatch(ctx).Body(body).Wait(wait).Execute()

Generate PDFs in a batch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	body := *openapiclient.NewCreateSubmissionBatchRequest([]openapiclient.SubmissionInput{*openapiclient.NewSubmissionInput()}) // CreateSubmissionBatchRequest | The submissions to create, and the settings they share.
	wait := true // bool | Wait until every PDF in the batch has been generated before responding. Default true on sync.api hosts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.CreateSubmissionBatch(context.Background()).Body(body).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateSubmissionBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubmissionBatch`: SubmissionBatch
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.CreateSubmissionBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubmissionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateSubmissionBatchRequest**](CreateSubmissionBatchRequest.md) | The submissions to create, and the settings they share. | 
 **wait** | **bool** | Wait until every PDF in the batch has been generated before responding. Default true on sync.api hosts. | 

### Return type

[**SubmissionBatch**](SubmissionBatch.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplate

> Template CreateTemplate(ctx).Body(body).Wait(wait).Execute()

Create a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	body := *openapiclient.NewCreateTemplateRequest(openapiclient.TemplateType("pdf")) // CreateTemplateRequest | The template to create.
	wait := true // bool | Wait until the PDF has been processed, and a `document_url` downloaded, before responding. Default true on sync.api hosts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.CreateTemplate(context.Background()).Body(body).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.CreateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateTemplateRequest**](CreateTemplateRequest.md) | The template to create. | 
 **wait** | **bool** | Wait until the PDF has been processed, and a `document_url` downloaded, before responding. Default true on sync.api hosts. | 

### Return type

[**Template**](Template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> Webhook CreateWebhook(ctx).Body(body).Execute()

Create a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	body := *openapiclient.NewCreateWebhookRequest("Url_example", []openapiclient.WebhookEventType{openapiclient.WebhookEventType("submission.processed")}) // CreateWebhookRequest | The webhook to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.CreateWebhook(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateWebhookRequest**](CreateWebhookRequest.md) | The webhook to create. | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomFile

> CustomFile DeleteCustomFile(ctx, customFileId).Execute()

Delete a custom file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	customFileId := "customFileId_example" // string | The custom file's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.DeleteCustomFile(context.Background(), customFileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.DeleteCustomFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomFile`: CustomFile
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.DeleteCustomFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFileId** | **string** | The custom file's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFile**](CustomFile.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	folderId := "folderId_example" // string | The folder's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.DeleteFolder(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.DeleteFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFolder`: Folder
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.DeleteFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The folder's ID. | 

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

> Template DeleteTemplate(ctx, templateId).Version(version).Execute()

Delete a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef02" // string | The template's ID.
	version := "version_example" // string | Delete only this published version (`1.2.0`) or the draft (`draft`, when no version has been published). Without it, the draft and every published version are deleted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.DeleteTemplate(context.Background(), templateId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.DeleteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.DeleteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | Delete only this published version (`1.2.0`) or the draft (`draft`, when no version has been published). Without it, the draft and every published version are deleted. | 

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


## DeleteWebhook

> Webhook DeleteWebhook(ctx, webhookId).Execute()

Delete a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	webhookId := "webhookId_example" // string | The webhook's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.DeleteWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | The webhook's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireCombinedPdf

> CombinedPdf ExpireCombinedPdf(ctx, combinedPdfId).Execute()

Expire a combined PDF



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	combinedPdfId := "com_1234567890abcdef05" // string | The combined PDF's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ExpireCombinedPdf(context.Background(), combinedPdfId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ExpireCombinedPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpireCombinedPdf`: CombinedPdf
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ExpireCombinedPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**combinedPdfId** | **string** | The combined PDF's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireCombinedPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CombinedPdf**](CombinedPdf.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireSubmission

> Submission ExpireSubmission(ctx, submissionId).Execute()

Expire a submission



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	submissionId := "sub_1234567890abcdef01" // string | The submission's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ExpireSubmission(context.Background(), submissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ExpireSubmission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpireSubmission`: Submission
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ExpireSubmission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionId** | **string** | The submission's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GenerateSubmissionFiles

> Submission GenerateSubmissionFiles(ctx, submissionId).Body(body).Wait(wait).Execute()

Generate more files



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	submissionId := "sub_1234567890abcdef01" // string | The submission's ID.
	body := *openapiclient.NewGenerateSubmissionFilesRequest([]openapiclient.GenerateFilesFormat{openapiclient.GenerateFilesFormat("png")}) // GenerateSubmissionFilesRequest | The formats to make.
	wait := true // bool | Wait until the files have been generated before responding. Default true on sync.api hosts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GenerateSubmissionFiles(context.Background(), submissionId).Body(body).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GenerateSubmissionFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateSubmissionFiles`: Submission
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GenerateSubmissionFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionId** | **string** | The submission's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSubmissionFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GenerateSubmissionFilesRequest**](GenerateSubmissionFilesRequest.md) | The formats to make. | 
 **wait** | **bool** | Wait until the files have been generated before responding. Default true on sync.api hosts. | 

### Return type

[**Submission**](Submission.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSubmissionPreview

> Submission GenerateSubmissionPreview(ctx, submissionId).Execute()

Generate a preview PDF



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	submissionId := "sub_1234567890abcdef03" // string | The submission's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GenerateSubmissionPreview(context.Background(), submissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GenerateSubmissionPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateSubmissionPreview`: Submission
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GenerateSubmissionPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionId** | **string** | The submission's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSubmissionPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetCombinedPdf

> CombinedPdf GetCombinedPdf(ctx, combinedPdfId).Execute()

Get a combined PDF



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	combinedPdfId := "com_1234567890abcdef05" // string | The combined PDF's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetCombinedPdf(context.Background(), combinedPdfId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetCombinedPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCombinedPdf`: CombinedPdf
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetCombinedPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**combinedPdfId** | **string** | The combined PDF's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCombinedPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CombinedPdf**](CombinedPdf.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomFile

> CustomFile GetCustomFile(ctx, customFileId).Execute()

Get a custom file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	customFileId := "customFileId_example" // string | The custom file's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetCustomFile(context.Background(), customFileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetCustomFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomFile`: CustomFile
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetCustomFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFileId** | **string** | The custom file's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFile**](CustomFile.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolder

> Folder GetFolder(ctx, folderId).Execute()

Get a folder



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	folderId := "folderId_example" // string | The folder's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetFolder(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolder`: Folder
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The folder's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderRequest struct via the builder pattern


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


## GetPresignUrl

> UploadPresign GetPresignUrl(ctx).Execute()

Get a presigned URL to upload a file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetPresignUrl(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetPresignUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPresignUrl`: UploadPresign
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetPresignUrl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPresignUrlRequest struct via the builder pattern


### Return type

[**UploadPresign**](UploadPresign.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignatureRequest

> SignatureRequest GetSignatureRequest(ctx, signatureRequestId).Execute()

Get a signature request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	signatureRequestId := "srq_1234567890abcdef01" // string | The signature request's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetSignatureRequest(context.Background(), signatureRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetSignatureRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignatureRequest`: SignatureRequest
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetSignatureRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signatureRequestId** | **string** | The signature request's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignatureRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SignatureRequest**](SignatureRequest.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubmission

> Submission GetSubmission(ctx, submissionId).Execute()

Get a submission



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	submissionId := "sub_1234567890abcdef01" // string | The submission's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetSubmission(context.Background(), submissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetSubmission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubmission`: Submission
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetSubmission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionId** | **string** | The submission's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> SubmissionBatch GetSubmissionBatch(ctx, submissionBatchId).IncludeData(includeData).Execute()

Get a batch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	submissionBatchId := "sbb_1234567890abcdef05" // string | The batch's ID.
	includeData := true // bool | Include each submission's `data` (null without it). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetSubmissionBatch(context.Background(), submissionBatchId).IncludeData(includeData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetSubmissionBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubmissionBatch`: SubmissionBatch
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetSubmissionBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionBatchId** | **string** | The batch's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubmissionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeData** | **bool** | Include each submission's `data` (null without it). | 

### Return type

[**SubmissionBatch**](SubmissionBatch.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplate

> Template GetTemplate(ctx, templateId).Version(version).Execute()

Get a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	version := "latest" // string | A published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `draft`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetTemplate(context.Background(), templateId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | A published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `draft`. | 

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


## GetTemplateDiff

> TemplateDiff GetTemplateDiff(ctx, templateId).Version(version).Execute()

Compare a version of a template with the one before it



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	version := "0.1.0" // string | A published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `draft`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetTemplateDiff(context.Background(), templateId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetTemplateDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateDiff`: TemplateDiff
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetTemplateDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | A published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `draft`. | 

### Return type

[**TemplateDiff**](TemplateDiff.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateModelCode

> TemplateModelCode GetTemplateModelCode(ctx, templateId).Lang(lang).Version(version).Execute()

Generate typed code for a template's submission data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	lang := openapiclient.ModelCodeLanguage("typescript") // ModelCodeLanguage | The language to generate the code in.
	version := "draft" // string | The template version: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published), the version a submission uses by default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetTemplateModelCode(context.Background(), templateId).Lang(lang).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetTemplateModelCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateModelCode`: TemplateModelCode
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetTemplateModelCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateModelCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lang** | [**ModelCodeLanguage**](ModelCodeLanguage.md) | The language to generate the code in. | 
 **version** | **string** | The template version: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published), the version a submission uses by default. | 

### Return type

[**TemplateModelCode**](TemplateModelCode.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatePageImages

> TemplatePageImages GetTemplatePageImages(ctx, templateId).Version(version).Execute()

Get images of a template's pages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	version := "latest" // string | A published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `draft`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetTemplatePageImages(context.Background(), templateId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetTemplatePageImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplatePageImages`: TemplatePageImages
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetTemplatePageImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatePageImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | A published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `draft`. | 

### Return type

[**TemplatePageImages**](TemplatePageImages.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateSchema

> map[string]interface{} GetTemplateSchema(ctx, templateId).Version(version).Execute()

Get the JSON Schema of a template's submission data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	version := "draft" // string | The template version: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published), the version a submission uses by default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetTemplateSchema(context.Background(), templateId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetTemplateSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetTemplateSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The template version: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published), the version a submission uses by default. | 

### Return type

**map[string]interface{}**

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> Webhook GetWebhook(ctx, webhookId).Execute()

Get a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	webhookId := "webhookId_example" // string | The webhook's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | The webhook's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCombinedPdfs

> CombinedPdfList ListCombinedPdfs(ctx).Limit(limit).Cursor(cursor).Execute()

List combined PDFs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	limit := int32(56) // int32 | How many items to return (1-100, default 50). (optional) (default to 50)
	cursor := "cursor_example" // string | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ListCombinedPdfs(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ListCombinedPdfs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCombinedPdfs`: CombinedPdfList
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ListCombinedPdfs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCombinedPdfsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many items to return (1-100, default 50). | [default to 50]
 **cursor** | **string** | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. | 

### Return type

[**CombinedPdfList**](CombinedPdfList.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomFiles

> CustomFileList ListCustomFiles(ctx).Limit(limit).Cursor(cursor).Execute()

List custom files



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	limit := int32(56) // int32 | How many items to return (1-100, default 50). (optional) (default to 50)
	cursor := "cursor_example" // string | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ListCustomFiles(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ListCustomFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomFiles`: CustomFileList
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ListCustomFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many items to return (1-100, default 50). | [default to 50]
 **cursor** | **string** | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. | 

### Return type

[**CustomFileList**](CustomFileList.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFolders

> FolderList ListFolders(ctx).ParentFolderId(parentFolderId).Limit(limit).Cursor(cursor).Execute()

List folders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	parentFolderId := "fld_1234567890abcdef02" // string | Only list the folders in this folder. Without it, the list has every folder, at any depth. (optional)
	limit := int32(2) // int32 | How many items to return (1-100, default 50). (optional) (default to 50)
	cursor := "fld_1234567890abcdef02" // string | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ListFolders(context.Background()).ParentFolderId(parentFolderId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ListFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFolders`: FolderList
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ListFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentFolderId** | **string** | Only list the folders in this folder. Without it, the list has every folder, at any depth. | 
 **limit** | **int32** | How many items to return (1-100, default 50). | [default to 50]
 **cursor** | **string** | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. | 

### Return type

[**FolderList**](FolderList.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubmissions

> SubmissionList ListSubmissions(ctx).TemplateId(templateId).Test(test).CreatedAfter(createdAfter).CreatedBefore(createdBefore).IncludeData(includeData).Limit(limit).Cursor(cursor).Execute()

List submissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | Only list this template's submissions. (optional)
	test := false // bool | Only list test submissions (true) or live submissions (false). (optional)
	createdAfter := time.Now() // time.Time | Only list submissions created after this time (RFC 3339, such as `2025-01-31T09:00:00Z`). (optional)
	createdBefore := time.Now() // time.Time | Only list submissions created before this time. (optional)
	includeData := true // bool | Include each submission's `data` (null without it). (optional)
	limit := int32(56) // int32 | How many items to return (1-100, default 50). (optional) (default to 50)
	cursor := "cursor_example" // string | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ListSubmissions(context.Background()).TemplateId(templateId).Test(test).CreatedAfter(createdAfter).CreatedBefore(createdBefore).IncludeData(includeData).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ListSubmissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubmissions`: SubmissionList
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ListSubmissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubmissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **string** | Only list this template's submissions. | 
 **test** | **bool** | Only list test submissions (true) or live submissions (false). | 
 **createdAfter** | **time.Time** | Only list submissions created after this time (RFC 3339, such as `2025-01-31T09:00:00Z`). | 
 **createdBefore** | **time.Time** | Only list submissions created before this time. | 
 **includeData** | **bool** | Include each submission's `data` (null without it). | 
 **limit** | **int32** | How many items to return (1-100, default 50). | [default to 50]
 **cursor** | **string** | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. | 

### Return type

[**SubmissionList**](SubmissionList.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplateFields

> TemplateFields ListTemplateFields(ctx, templateId).Version(version).Execute()

List a template's fields



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	version := "latest" // string | A published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `draft`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ListTemplateFields(context.Background(), templateId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ListTemplateFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTemplateFields`: TemplateFields
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ListTemplateFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplateFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | A published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `draft`. | 

### Return type

[**TemplateFields**](TemplateFields.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplateSubmissions

> SubmissionList ListTemplateSubmissions(ctx, templateId).Test(test).CreatedAfter(createdAfter).CreatedBefore(createdBefore).IncludeData(includeData).Limit(limit).Cursor(cursor).Execute()

List a template's submissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	test := true // bool | Only list test submissions (true) or live submissions (false). (optional)
	createdAfter := time.Now() // time.Time | Only list submissions created after this time (RFC 3339, such as `2025-01-31T09:00:00Z`). (optional)
	createdBefore := time.Now() // time.Time | Only list submissions created before this time. (optional)
	includeData := true // bool | Include each submission's `data` (null without it). (optional)
	limit := int32(56) // int32 | How many items to return (1-100, default 50). (optional) (default to 50)
	cursor := "cursor_example" // string | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ListTemplateSubmissions(context.Background(), templateId).Test(test).CreatedAfter(createdAfter).CreatedBefore(createdBefore).IncludeData(includeData).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ListTemplateSubmissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTemplateSubmissions`: SubmissionList
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ListTemplateSubmissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplateSubmissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **test** | **bool** | Only list test submissions (true) or live submissions (false). | 
 **createdAfter** | **time.Time** | Only list submissions created after this time (RFC 3339, such as `2025-01-31T09:00:00Z`). | 
 **createdBefore** | **time.Time** | Only list submissions created before this time. | 
 **includeData** | **bool** | Include each submission's `data` (null without it). | 
 **limit** | **int32** | How many items to return (1-100, default 50). | [default to 50]
 **cursor** | **string** | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. | 

### Return type

[**SubmissionList**](SubmissionList.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplates

> TemplateSummaryList ListTemplates(ctx).Query(query).ParentFolderId(parentFolderId).Limit(limit).Cursor(cursor).Execute()

List templates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	query := "rental" // string | Only list templates whose name contains this text (ignoring case). (optional)
	parentFolderId := "fld_1234567890abcdef01" // string | Only list the templates in this folder (not in its subfolders). (optional)
	limit := int32(1) // int32 | How many items to return (1-100, default 50). (optional) (default to 50)
	cursor := "cursor_example" // string | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ListTemplates(context.Background()).Query(query).ParentFolderId(parentFolderId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ListTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTemplates`: TemplateSummaryList
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ListTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Only list templates whose name contains this text (ignoring case). | 
 **parentFolderId** | **string** | Only list the templates in this folder (not in its subfolders). | 
 **limit** | **int32** | How many items to return (1-100, default 50). | [default to 50]
 **cursor** | **string** | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. | 

### Return type

[**TemplateSummaryList**](TemplateSummaryList.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooks

> WebhookList ListWebhooks(ctx).EventType(eventType).Mode(mode).TemplateId(templateId).FolderId(folderId).ApiTokenId(apiTokenId).Limit(limit).Cursor(cursor).Execute()

List webhooks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	eventType := openapiclient.WebhookEventType("submission.processed") // WebhookEventType | Only list webhooks that subscribe to this event. (optional)
	mode := openapiclient.WebhookMode("test") // WebhookMode | Only list webhooks set to this mode. (optional)
	templateId := "tpl_1234567890abcdef01" // string | Only list webhooks scoped to this template. (optional)
	folderId := "folderId_example" // string | Only list webhooks scoped to this folder. (optional)
	apiTokenId := "apiTokenId_example" // string | Only list webhooks scoped to this API token. (optional)
	limit := int32(56) // int32 | How many items to return (1-100, default 50). (optional) (default to 50)
	cursor := "cursor_example" // string | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ListWebhooks(context.Background()).EventType(eventType).Mode(mode).TemplateId(templateId).FolderId(folderId).ApiTokenId(apiTokenId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ListWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhooks`: WebhookList
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ListWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventType** | [**WebhookEventType**](WebhookEventType.md) | Only list webhooks that subscribe to this event. | 
 **mode** | [**WebhookMode**](WebhookMode.md) | Only list webhooks set to this mode. | 
 **templateId** | **string** | Only list webhooks scoped to this template. | 
 **folderId** | **string** | Only list webhooks scoped to this folder. | 
 **apiTokenId** | **string** | Only list webhooks scoped to this API token. | 
 **limit** | **int32** | How many items to return (1-100, default 50). | [default to 50]
 **cursor** | **string** | The previous page's `next_cursor`, to get the next page. It's the ID of that page's last item, so get the next page before you delete or change the items on this one. | 

### Return type

[**WebhookList**](WebhookList.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveFolder

> Folder MoveFolder(ctx, folderId).Body(body).Execute()

Move a folder



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	folderId := "folderId_example" // string | The folder's ID.
	body := *openapiclient.NewMoveFolderRequest("ParentFolderId_example") // MoveFolderRequest | Where to move the folder.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.MoveFolder(context.Background(), folderId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.MoveFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveFolder`: Folder
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.MoveFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The folder's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MoveFolderRequest**](MoveFolderRequest.md) | Where to move the folder. | 

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


## MoveTemplate

> Template MoveTemplate(ctx, templateId).Body(body).Execute()

Move a template to a folder



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewMoveTemplateRequest("ParentFolderId_example") // MoveTemplateRequest | Where to move the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.MoveTemplate(context.Background(), templateId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.MoveTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.MoveTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MoveTemplateRequest**](MoveTemplateRequest.md) | Where to move the template. | 

### Return type

[**Template**](Template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishTemplateVersion

> TemplateVersion PublishTemplateVersion(ctx, templateId).Body(body).Execute()

Publish a template version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewPublishTemplateVersionRequest(openapiclient.TemplateVersionType("major")) // PublishTemplateVersionRequest | The version to publish.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.PublishTemplateVersion(context.Background(), templateId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PublishTemplateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishTemplateVersion`: TemplateVersion
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PublishTemplateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishTemplateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PublishTemplateVersionRequest**](PublishTemplateVersionRequest.md) | The version to publish. | 

### Return type

[**TemplateVersion**](TemplateVersion.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTemplateFields

> TemplateFields RemoveTemplateFields(ctx, templateId).Body(body).Execute()

Remove fields from a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewRemoveTemplateFieldsRequest([]int32{int32(123)}) // RemoveTemplateFieldsRequest | The fields to remove from the template's draft.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.RemoveTemplateFields(context.Background(), templateId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.RemoveTemplateFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveTemplateFields`: TemplateFields
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.RemoveTemplateFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTemplateFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RemoveTemplateFieldsRequest**](RemoveTemplateFieldsRequest.md) | The fields to remove from the template's draft. | 

### Return type

[**TemplateFields**](TemplateFields.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameFolder

> Folder RenameFolder(ctx, folderId).Body(body).Execute()

Rename a folder



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	folderId := "folderId_example" // string | The folder's ID.
	body := *openapiclient.NewRenameFolderRequest("Name_example") // RenameFolderRequest | The folder's new name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.RenameFolder(context.Background(), folderId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.RenameFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameFolder`: Folder
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.RenameFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The folder's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RenameFolderRequest**](RenameFolderRequest.md) | The folder's new name. | 

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


## ReorderTemplateFields

> TemplateFields ReorderTemplateFields(ctx, templateId).Body(body).Execute()

Reorder a template's fields



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewReorderTemplateFieldsRequest([]int32{int32(123)}) // ReorderTemplateFieldsRequest | The new order of the fields in the template's draft.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ReorderTemplateFields(context.Background(), templateId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ReorderTemplateFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderTemplateFields`: TemplateFields
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ReorderTemplateFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReorderTemplateFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReorderTemplateFieldsRequest**](ReorderTemplateFieldsRequest.md) | The new order of the fields in the template's draft. | 

### Return type

[**TemplateFields**](TemplateFields.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreTemplateVersion

> Template RestoreTemplateVersion(ctx, templateId).Body(body).Execute()

Restore a template version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewRestoreTemplateVersionRequest("Version_example") // RestoreTemplateVersionRequest | The version to restore.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.RestoreTemplateVersion(context.Background(), templateId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.RestoreTemplateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreTemplateVersion`: Template
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.RestoreTemplateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreTemplateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RestoreTemplateVersionRequest**](RestoreTemplateVersionRequest.md) | The version to restore. | 

### Return type

[**Template**](Template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestAuthentication

> AuthenticationCheck TestAuthentication(ctx).Execute()

Test authentication



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.TestAuthentication(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.TestAuthentication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestAuthentication`: AuthenticationCheck
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.TestAuthentication`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestAuthenticationRequest struct via the builder pattern


### Return type

[**AuthenticationCheck**](AuthenticationCheck.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSignatureRequest

> SignatureRequest UpdateSignatureRequest(ctx, signatureRequestId).Body(body).Execute()

Update a signature request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	signatureRequestId := "srq_1234567890abcdef01" // string | The signature request's ID.
	body := *openapiclient.NewUpdateSignatureRequestRequest() // UpdateSignatureRequestRequest | The settings to change. Properties left out keep their values.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.UpdateSignatureRequest(context.Background(), signatureRequestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UpdateSignatureRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSignatureRequest`: SignatureRequest
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.UpdateSignatureRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signatureRequestId** | **string** | The signature request's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSignatureRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateSignatureRequestRequest**](UpdateSignatureRequestRequest.md) | The settings to change. Properties left out keep their values. | 

### Return type

[**SignatureRequest**](SignatureRequest.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplate

> Template UpdateTemplate(ctx, templateId).Body(body).Execute()

Update a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewUpdateTemplateRequest() // UpdateTemplateRequest | The settings to change. Properties left out keep their values.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.UpdateTemplate(context.Background(), templateId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UpdateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.UpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateTemplateRequest**](UpdateTemplateRequest.md) | The settings to change. Properties left out keep their values. | 

### Return type

[**Template**](Template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplateDocument

> Template UpdateTemplateDocument(ctx, templateId).Body(body).Wait(wait).Execute()

Replace a template's PDF



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewUpdateTemplateDocumentRequest() // UpdateTemplateDocumentRequest | The new PDF, as one of `cache_id`, `document_url` and `document_base64`.
	wait := true // bool | Wait until the PDF has been processed, and a `document_url` downloaded, before responding. Default true on sync.api hosts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.UpdateTemplateDocument(context.Background(), templateId).Body(body).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UpdateTemplateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTemplateDocument`: Template
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.UpdateTemplateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateTemplateDocumentRequest**](UpdateTemplateDocumentRequest.md) | The new PDF, as one of `cache_id`, `document_url` and `document_base64`. | 
 **wait** | **bool** | Wait until the PDF has been processed, and a `document_url` downloaded, before responding. Default true on sync.api hosts. | 

### Return type

[**Template**](Template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplateFields

> TemplateFields UpdateTemplateFields(ctx, templateId).Body(body).Execute()

Update a template's fields



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewUpdateTemplateFieldsRequest([]openapiclient.FieldInput{*openapiclient.NewFieldInput()}) // UpdateTemplateFieldsRequest | The fields to change in the template's draft.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.UpdateTemplateFields(context.Background(), templateId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UpdateTemplateFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTemplateFields`: TemplateFields
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.UpdateTemplateFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateTemplateFieldsRequest**](UpdateTemplateFieldsRequest.md) | The fields to change in the template's draft. | 

### Return type

[**TemplateFields**](TemplateFields.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> Webhook UpdateWebhook(ctx, webhookId).Body(body).Execute()

Update a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	webhookId := "webhookId_example" // string | The webhook's ID.
	body := *openapiclient.NewUpdateWebhookRequest() // UpdateWebhookRequest | The settings to change. Properties left out keep their values.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.UpdateWebhook(context.Background(), webhookId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | The webhook's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateWebhookRequest**](UpdateWebhookRequest.md) | The settings to change. Properties left out keep their values. | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadTemplate

> Template UploadTemplate(ctx).Document(document).Wait(wait).Name(name).Description(description).ParentFolderId(parentFolderId).AiFieldNamingLanguage(aiFieldNamingLanguage).DetectFields(detectFields).FieldSource(fieldSource).FieldNaming(fieldNaming).FieldLogic(fieldLogic).Execute()

Create a PDF template from an uploaded PDF



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	document := os.NewFile(1234, "some_file") // *os.File | The PDF.
	wait := true // bool | Wait until the PDF has been processed, and a `document_url` downloaded, before responding. Default true on sync.api hosts. (optional)
	name := "name_example" // string | The template's name. (optional)
	description := "description_example" // string | The template's description. (optional)
	parentFolderId := "parentFolderId_example" // string | The folder to put the template in. Left out, it's created at the root. (optional)
	aiFieldNamingLanguage := openapiclient.AiFieldNamingLanguage("auto") // AiFieldNamingLanguage | The language AI field setup names fields in: `auto` uses the form's language when it's written in the Latin alphabet without accents, and English otherwise; `english` always uses English; `form_language` uses the form's language and script. Default: your account's template settings. (optional)
	detectFields := true // bool | Set up the PDF's fields as DocSpring's "Automatically detect fields" upload option does: import or detect the fields, name them with AI, and set up conditional logic. `field_source`, `field_naming` and `field_logic` choose how. Default false. `document.page_processing_statuses` and `document.field_logic_status` show how it's going, and `document.state` becomes `processed` once every page's fields are set up. (optional)
	fieldSource := openapiclient.FieldSetupSource("auto") // FieldSetupSource | With `detect_fields`, where the fields come from. `auto` (the default) imports the form fields of a fillable PDF and detects the fields of a PDF without them; `pdf_detection` detects the fields from the page layout, even on a fillable PDF; `pdf_form_fields` only imports the PDF's form fields. (optional)
	fieldNaming := openapiclient.FieldSetupNaming("ai") // FieldSetupNaming | With `detect_fields`, how the fields are named. `ai` (the default) names them and sets their types and settings with AI; `preserve` keeps the names of the PDF's form fields, and AI only names the fields it detected. `preserve` needs `field_source` `auto` or `pdf_form_fields`. (optional)
	fieldLogic := openapiclient.FieldSetupLogic("ai") // FieldSetupLogic | With `detect_fields`, whether AI reads the whole form once the fields are named and sets up conditional logic and required fields (`ai`) or not (`none`). Default: your account's template settings. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.UploadTemplate(context.Background()).Document(document).Wait(wait).Name(name).Description(description).ParentFolderId(parentFolderId).AiFieldNamingLanguage(aiFieldNamingLanguage).DetectFields(detectFields).FieldSource(fieldSource).FieldNaming(fieldNaming).FieldLogic(fieldLogic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UploadTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.UploadTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **document** | ***os.File** | The PDF. | 
 **wait** | **bool** | Wait until the PDF has been processed, and a `document_url` downloaded, before responding. Default true on sync.api hosts. | 
 **name** | **string** | The template's name. | 
 **description** | **string** | The template's description. | 
 **parentFolderId** | **string** | The folder to put the template in. Left out, it's created at the root. | 
 **aiFieldNamingLanguage** | [**AiFieldNamingLanguage**](AiFieldNamingLanguage.md) | The language AI field setup names fields in: `auto` uses the form's language when it's written in the Latin alphabet without accents, and English otherwise; `english` always uses English; `form_language` uses the form's language and script. Default: your account's template settings. | 
 **detectFields** | **bool** | Set up the PDF's fields as DocSpring's "Automatically detect fields" upload option does: import or detect the fields, name them with AI, and set up conditional logic. `field_source`, `field_naming` and `field_logic` choose how. Default false. `document.page_processing_statuses` and `document.field_logic_status` show how it's going, and `document.state` becomes `processed` once every page's fields are set up. | 
 **fieldSource** | [**FieldSetupSource**](FieldSetupSource.md) | With `detect_fields`, where the fields come from. `auto` (the default) imports the form fields of a fillable PDF and detects the fields of a PDF without them; `pdf_detection` detects the fields from the page layout, even on a fillable PDF; `pdf_form_fields` only imports the PDF's form fields. | 
 **fieldNaming** | [**FieldSetupNaming**](FieldSetupNaming.md) | With `detect_fields`, how the fields are named. `ai` (the default) names them and sets their types and settings with AI; `preserve` keeps the names of the PDF's form fields, and AI only names the fields it detected. `preserve` needs `field_source` `auto` or `pdf_form_fields`. | 
 **fieldLogic** | [**FieldSetupLogic**](FieldSetupLogic.md) | With `detect_fields`, whether AI reads the whole form once the fields are named and sets up conditional logic and required fields (`ai`) or not (`none`). Default: your account's template settings. | 

### Return type

[**Template**](Template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadTemplateDocument

> Template UploadTemplateDocument(ctx, templateId).Document(document).Wait(wait).ReimportFields(reimportFields).Execute()

Replace a template's PDF with an uploaded PDF



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	document := os.NewFile(1234, "some_file") // *os.File | The new PDF.
	wait := true // bool | Wait until the PDF has been processed, and a `document_url` downloaded, before responding. Default true on sync.api hosts. (optional)
	reimportFields := true // bool | Replace all of the template's fields with the fillable fields in the new PDF, once DocSpring has processed it (`document.field_reimport_state`). As on a first upload, the default typeface, font size, bold and color are taken from the PDF's fields; the template's other settings are kept. If the PDF has no fillable fields, the template keeps its fields. A locked template can't re-import its fields. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.UploadTemplateDocument(context.Background(), templateId).Document(document).Wait(wait).ReimportFields(reimportFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UploadTemplateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadTemplateDocument`: Template
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.UploadTemplateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadTemplateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **document** | ***os.File** | The new PDF. | 
 **wait** | **bool** | Wait until the PDF has been processed, and a `document_url` downloaded, before responding. Default true on sync.api hosts. | 
 **reimportFields** | **bool** | Replace all of the template's fields with the fillable fields in the new PDF, once DocSpring has processed it (`document.field_reimport_state`). As on a first upload, the default typeface, font size, bold and color are taken from the PDF's fields; the template's other settings are kept. If the PDF has no fillable fields, the template keeps its fields. A locked template can't re-import its fields. | 

### Return type

[**Template**](Template.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSubmission

> SubmissionValidation ValidateSubmission(ctx, templateId).Body(body).Version(version).Execute()

Check a submission without creating it



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/docspring/docspring-go/v4"
)

func main() {
	templateId := "tpl_1234567890abcdef01" // string | The template's ID.
	body := *openapiclient.NewValidateSubmissionRequest() // ValidateSubmissionRequest | The submission's data and settings.
	version := "latest" // string | The template version to use: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ValidateSubmission(context.Background(), templateId).Body(body).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ValidateSubmission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateSubmission`: SubmissionValidation
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ValidateSubmission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template's ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ValidateSubmissionRequest**](ValidateSubmissionRequest.md) | The submission's data and settings. | 
 **version** | **string** | The template version to use: a published version (`1.2.0`, or `1` for the latest 1.x.x), `latest` or `draft`. Default `latest` (the draft when no version has been published). | 

### Return type

[**SubmissionValidation**](SubmissionValidation.md)

### Authorization

[api_token_basic](../README.md#api_token_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

