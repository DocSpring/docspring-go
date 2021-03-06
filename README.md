# Go API client for docspring

DocSpring is a service that helps you fill out and sign PDF templates.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.1
- Build package: com.docspring.codegen.DocSpringGoClientCodegen

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./docspring"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.docspring.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PDFApi* | [**AddFieldsToTemplate**](docs/PDFApi.md#addfieldstotemplate) | **Put** /templates/{template_id}/add_fields | Add new fields to a Template
*PDFApi* | [**BatchGeneratePdfV1**](docs/PDFApi.md#batchgeneratepdfv1) | **Post** /templates/{template_id}/submissions/batch | Generates multiple PDFs
*PDFApi* | [**BatchGeneratePdfs**](docs/PDFApi.md#batchgeneratepdfs) | **Post** /submissions/batches | Generates multiple PDFs
*PDFApi* | [**CombinePdfs**](docs/PDFApi.md#combinepdfs) | **Post** /combined_submissions?v&#x3D;2 | Merge submission PDFs, template PDFs, or custom files
*PDFApi* | [**CombineSubmissions**](docs/PDFApi.md#combinesubmissions) | **Post** /combined_submissions | Merge generated PDFs together
*PDFApi* | [**CopyTemplate**](docs/PDFApi.md#copytemplate) | **Post** /templates/{template_id}/copy | Copy a Template
*PDFApi* | [**CreateCustomFileFromUpload**](docs/PDFApi.md#createcustomfilefromupload) | **Post** /custom_files | Create a new custom file from a cached presign upload
*PDFApi* | [**CreateDataRequestToken**](docs/PDFApi.md#createdatarequesttoken) | **Post** /data_requests/{data_request_id}/tokens | Creates a new data request token for form authentication
*PDFApi* | [**CreateFolder**](docs/PDFApi.md#createfolder) | **Post** /folders/ | Create a folder
*PDFApi* | [**CreateHTMLTemplate**](docs/PDFApi.md#createhtmltemplate) | **Post** /templates?desc&#x3D;html | Create a new HTML template
*PDFApi* | [**CreatePDFTemplate**](docs/PDFApi.md#createpdftemplate) | **Post** /templates | Create a new PDF template with a form POST file upload
*PDFApi* | [**CreatePDFTemplateFromUpload**](docs/PDFApi.md#createpdftemplatefromupload) | **Post** /templates?desc&#x3D;cached_upload | Create a new PDF template from a cached presign upload
*PDFApi* | [**DeleteFolder**](docs/PDFApi.md#deletefolder) | **Delete** /folders/{folder_id} | Delete a folder
*PDFApi* | [**ExpireCombinedSubmission**](docs/PDFApi.md#expirecombinedsubmission) | **Delete** /combined_submissions/{combined_submission_id} | Expire a combined submission
*PDFApi* | [**ExpireSubmission**](docs/PDFApi.md#expiresubmission) | **Delete** /submissions/{submission_id} | Expire a PDF submission
*PDFApi* | [**GeneratePDF**](docs/PDFApi.md#generatepdf) | **Post** /templates/{template_id}/submissions | Generates a new PDF
*PDFApi* | [**GetCombinedSubmission**](docs/PDFApi.md#getcombinedsubmission) | **Get** /combined_submissions/{combined_submission_id} | Check the status of a combined submission (merged PDFs)
*PDFApi* | [**GetDataRequest**](docs/PDFApi.md#getdatarequest) | **Get** /data_requests/{data_request_id} | Look up a submission data request
*PDFApi* | [**GetFullTemplate**](docs/PDFApi.md#getfulltemplate) | **Get** /templates/{template_id}?full&#x3D;true | Fetch the full template attributes
*PDFApi* | [**GetPresignUrl**](docs/PDFApi.md#getpresignurl) | **Get** /uploads/presign | Get a presigned URL so that you can upload a file to our AWS S3 bucket
*PDFApi* | [**GetSubmission**](docs/PDFApi.md#getsubmission) | **Get** /submissions/{submission_id} | Check the status of a PDF
*PDFApi* | [**GetSubmissionBatch**](docs/PDFApi.md#getsubmissionbatch) | **Get** /submissions/batches/{submission_batch_id} | Check the status of a submission batch job
*PDFApi* | [**GetTemplate**](docs/PDFApi.md#gettemplate) | **Get** /templates/{template_id} | Check the status of an uploaded template
*PDFApi* | [**GetTemplateSchema**](docs/PDFApi.md#gettemplateschema) | **Get** /templates/{template_id}/schema | Fetch the JSON schema for a template
*PDFApi* | [**ListFolders**](docs/PDFApi.md#listfolders) | **Get** /folders/ | Get a list of all folders
*PDFApi* | [**ListSubmissions**](docs/PDFApi.md#listsubmissions) | **Get** /submissions | List all submissions
*PDFApi* | [**ListSubmissions_0**](docs/PDFApi.md#listsubmissions_0) | **Get** /templates/{template_id}/submissions | List all submissions for a given template
*PDFApi* | [**ListTemplates**](docs/PDFApi.md#listtemplates) | **Get** /templates | Get a list of all templates
*PDFApi* | [**MoveFolderToFolder**](docs/PDFApi.md#movefoldertofolder) | **Post** /folders/{folder_id}/move | Move a folder
*PDFApi* | [**MoveTemplateToFolder**](docs/PDFApi.md#movetemplatetofolder) | **Post** /templates/{template_id}/move | Move Template to folder
*PDFApi* | [**RenameFolder**](docs/PDFApi.md#renamefolder) | **Post** /folders/{folder_id}/rename | Rename a folder
*PDFApi* | [**TestAuthentication**](docs/PDFApi.md#testauthentication) | **Get** /authentication | Test Authentication
*PDFApi* | [**UpdateDataRequest**](docs/PDFApi.md#updatedatarequest) | **Put** /data_requests/{data_request_id} | Update a submission data request
*PDFApi* | [**UpdateTemplate**](docs/PDFApi.md#updatetemplate) | **Put** /templates/{template_id} | Update a Template


## Documentation For Models

 - [AddFieldsData](docs/AddFieldsData.md)
 - [AddFieldsTemplateResponse](docs/AddFieldsTemplateResponse.md)
 - [AuthenticationError](docs/AuthenticationError.md)
 - [AuthenticationSuccessResponse](docs/AuthenticationSuccessResponse.md)
 - [CombinePdfsData](docs/CombinePdfsData.md)
 - [CombinedSubmission](docs/CombinedSubmission.md)
 - [CombinedSubmissionAction](docs/CombinedSubmissionAction.md)
 - [CombinedSubmissionData](docs/CombinedSubmissionData.md)
 - [CopyTemplateData](docs/CopyTemplateData.md)
 - [CreateCombinedSubmissionResponse](docs/CreateCombinedSubmissionResponse.md)
 - [CreateCustomFileData](docs/CreateCustomFileData.md)
 - [CreateCustomFileResponse](docs/CreateCustomFileResponse.md)
 - [CreateFolderData](docs/CreateFolderData.md)
 - [CreateHtmlTemplateData](docs/CreateHtmlTemplateData.md)
 - [CreateSubmissionBatchResponse](docs/CreateSubmissionBatchResponse.md)
 - [CreateSubmissionBatchSubmissionsResponse](docs/CreateSubmissionBatchSubmissionsResponse.md)
 - [CreateSubmissionDataRequestData](docs/CreateSubmissionDataRequestData.md)
 - [CreateSubmissionDataRequestTokenResponse](docs/CreateSubmissionDataRequestTokenResponse.md)
 - [CreateSubmissionDataRequestTokenResponseToken](docs/CreateSubmissionDataRequestTokenResponseToken.md)
 - [CreateSubmissionResponse](docs/CreateSubmissionResponse.md)
 - [CreateTemplateFromUploadData](docs/CreateTemplateFromUploadData.md)
 - [CustomFile](docs/CustomFile.md)
 - [Folder](docs/Folder.md)
 - [FoldersFolder](docs/FoldersFolder.md)
 - [HtmlTemplateData](docs/HtmlTemplateData.md)
 - [InvalidRequest](docs/InvalidRequest.md)
 - [ListSubmissionsResponse](docs/ListSubmissionsResponse.md)
 - [ModelError](docs/ModelError.md)
 - [MoveFolderData](docs/MoveFolderData.md)
 - [MoveTemplateData](docs/MoveTemplateData.md)
 - [PendingTemplate](docs/PendingTemplate.md)
 - [RenameFolderData](docs/RenameFolderData.md)
 - [Submission](docs/Submission.md)
 - [SubmissionAction](docs/SubmissionAction.md)
 - [SubmissionBatch](docs/SubmissionBatch.md)
 - [SubmissionBatchData](docs/SubmissionBatchData.md)
 - [SubmissionData](docs/SubmissionData.md)
 - [SubmissionDataBatchRequest](docs/SubmissionDataBatchRequest.md)
 - [SubmissionDataRequest](docs/SubmissionDataRequest.md)
 - [Template](docs/Template.md)
 - [Template1](docs/Template1.md)
 - [Template1Defaults](docs/Template1Defaults.md)
 - [TemplateData](docs/TemplateData.md)
 - [TemplatestemplateIdaddFieldsFields](docs/TemplatestemplateIdaddFieldsFields.md)
 - [UpdateDataRequestResponse](docs/UpdateDataRequestResponse.md)
 - [UpdateSubmissionDataRequestData](docs/UpdateSubmissionDataRequestData.md)
 - [UpdateTemplateData](docs/UpdateTemplateData.md)
 - [UpdateTemplateResponse](docs/UpdateTemplateResponse.md)
 - [UploadTemplateData](docs/UploadTemplateData.md)
 - [UploadTemplateDataDocument](docs/UploadTemplateDataDocument.md)
 - [UploadTemplateDataDocumentMetadata](docs/UploadTemplateDataDocumentMetadata.md)


## Documentation For Authorization

## api_token_basic
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author



