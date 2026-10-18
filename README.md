# DocSpring Go SDK

The official Go client for the [DocSpring API](https://docspring.com/docs/api/). DocSpring fills PDF
templates with your data, generates PDFs from HTML templates, combines PDFs and collects e-signatures.

[Documentation](https://docspring.com/docs/) · [API reference](https://docspring.com/docs/api/) ·
[pkg.go.dev](https://pkg.go.dev/github.com/docspring/docspring-go/v3) · [Changelog](https://github.com/DocSpring/docspring-go/blob/main/CHANGELOG.md)

Version 3.3.0. Requires Go 1.23 or later.

## Install and configure

Install the package:

```bash
go get github.com/docspring/docspring-go/v3
```

Configure authentication via environment variables:

```bash
export DOCSPRING_TOKEN_ID="API_TOKEN_ID"
export DOCSPRING_TOKEN_SECRET="API_TOKEN_SECRET"
```

You can also configure the region or host. The host takes priority over the region:

```bash
export DOCSPRING_REGION="US" # US or EU (defaults to US)
export DOCSPRING_HOST="docspring.yourcompany.com" # For on-premise deployments
```

Import and initialize the client:

```go
import (
  client "github.com/docspring/docspring-go/v3"
)

apiClient := client.NewClient("", "")
```

Or provide credentials directly:

```go
apiClient := client.NewClient("API_TOKEN_ID", "API_TOKEN_SECRET")
```

## Generate a PDF

Fill a template with data (use your template's ID in place of `tpl_1234567890abcdef01`):

```go
templateID := "tpl_1234567890abcdef01"
submission := client.CreatePdfSubmissionData{Data: map[string]interface{}{"first_name": "John", "last_name": "Smith", "phone_number": "+11234567890"}, Metadata: map[string]interface{}{"foo": 123, "bar": "baz"}}
request := apiClient.GeneratePdf(context.Background(), templateID).Submission(submission)
result, _, err := request.Execute()
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(result)
}
```

By default the call waits for the PDF, for up to 95 seconds, and the submission in the response has its
`download_url`. Set `wait` to false to get a pending submission straight away.
[Sync vs. async processing](https://docspring.com/docs/api-guide/synchronous-requests/) covers the time limit, status
checks and webhooks.

Optional fields are pointers. Set them with the helpers `client.PtrString`, `client.PtrInt32`, `client.PtrBool` and so on, and read them with the model's getters (`GetDownloadUrl()`).

## Methods

[docs/ClientAPI.md](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md) has each method's parameters and return type, and
[docs/](https://github.com/DocSpring/docspring-go/tree/main/docs) has the models. The [API reference](https://docspring.com/docs/api/) has
Go code for every method.

| Method | Description |
| ------ | ----------- |
| [`AddFieldsToTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#addfieldstotemplate) | Add new fields to a Template |
| [`BatchGeneratePdfs`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#batchgeneratepdfs) | Generate multiple PDFs |
| [`CombinePdfs`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#combinepdfs) | Merge submission PDFs, template PDFs, or custom files |
| [`CopyTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#copytemplate) | Copy a template |
| [`CreateCustomFileFromUpload`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createcustomfilefromupload) | Create a new custom file from a cached S3 upload |
| [`CreateDataRequestEvent`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createdatarequestevent) | Create a new event for emailing a signee a request for signature |
| [`CreateDataRequestToken`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createdatarequesttoken) | Create a new data request token for form authentication |
| [`CreateFolder`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createfolder) | Create a folder |
| [`CreateHtmlTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createhtmltemplate) | Create a new HTML template |
| [`CreatePdfTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createpdftemplate) | Create a new PDF template with a form POST file upload |
| [`CreatePdfTemplateFromUpload`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createpdftemplatefromupload) | Create a new PDF template from a cached S3 file upload |
| [`CreateWebhook`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createwebhook) | Create a webhook |
| [`DeleteFolder`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#deletefolder) | Delete a folder |
| [`DeleteTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#deletetemplate) | Delete a template |
| [`DeleteWebhook`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#deletewebhook) | Delete a webhook |
| [`ExpireCombinedSubmission`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#expirecombinedsubmission) | Expire a combined submission |
| [`ExpireSubmission`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#expiresubmission) | Expire a PDF submission |
| [`GeneratePdf`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#generatepdf) | Generate a PDF |
| [`GeneratePreview`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#generatepreview) | Generate a preview PDF for partially completed data requests |
| [`GetCombinedSubmission`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getcombinedsubmission) | Check the status of a combined submission (merged PDFs) |
| [`GetDataRequest`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getdatarequest) | Look up a submission data request |
| [`GetFullTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getfulltemplate) | Fetch the full attributes for a PDF template |
| [`GetPresignUrl`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getpresignurl) | Get a presigned S3 URL for direct file upload |
| [`GetSubmission`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getsubmission) | Check the status of a PDF |
| [`GetSubmissionBatch`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getsubmissionbatch) | Check the status of a submission batch job |
| [`GetTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#gettemplate) | Check the status of an uploaded template |
| [`GetTemplateModelCode`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#gettemplatemodelcode) | Generate typed model code for a template |
| [`GetTemplateSchema`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#gettemplateschema) | Fetch the JSON schema for a template |
| [`GetWebhook`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getwebhook) | Get a webhook |
| [`ListCombinedSubmissions`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listcombinedsubmissions) | Get a list of all combined submissions |
| [`ListFolders`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listfolders) | Get a list of all folders |
| [`ListSubmissions`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listsubmissions) | List all submissions |
| [`ListTemplates`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listtemplates) | Get a list of all templates |
| [`ListTemplateSubmissions`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listtemplatesubmissions) | List all submissions for a given template |
| [`ListWebhooks`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listwebhooks) | Get a list of all webhooks |
| [`MoveFolderToFolder`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#movefoldertofolder) | Move a folder |
| [`MoveTemplateToFolder`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#movetemplatetofolder) | Move Template to folder |
| [`PublishTemplateVersion`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#publishtemplateversion) | Publish a template version |
| [`RenameFolder`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#renamefolder) | Rename a folder |
| [`RestoreTemplateVersion`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#restoretemplateversion) | Restore a template version |
| [`TestAuthentication`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#testauthentication) | Test authentication |
| [`UpdateDataRequest`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#updatedatarequest) | Update a submission data request |
| [`UpdateTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#updatetemplate) | Update a Template |
| [`UpdateTemplateDocument`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#updatetemplatedocument) | Update a template's document with a form POST file upload |
| [`UpdateTemplateDocumentFromUpload`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#updatetemplatedocumentfromupload) | Update a template's document with a cached S3 file upload |
| [`UpdateWebhook`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#updatewebhook) | Update a webhook |

## Errors

`Execute()` returns an `error`. For an error response it's a `*client.GenericOpenAPIError`: `Body()` is the raw body and `Model()` the decoded error, and the `*http.Response` that `Execute()` also returns has the status code.

## API changes

DocSpring adds fields and enum values to API v1 without releasing a new API version. Changes that would break
existing code, such as removing or renaming a field, go in a new API version
([API changes](https://docspring.com/docs/api-guide/openapi-schema/#api-changes)). So this SDK keeps working with
responses that are newer than it is:

- The model keeps a response field added to the API after this release in its `AdditionalProperties` map.
- An enum value added to the API after this release comes back as a string, so handle values you don't recognize.

Upgrade the SDK to read new fields and enum values. The [changelog](https://github.com/DocSpring/docspring-go/blob/main/CHANGELOG.md) lists
what each release adds.

## About this SDK

This SDK is generated from DocSpring's [OpenAPI spec](https://docspring.com/api-docs/v1/openapi.json) with
[DocSpring's fork of OpenAPI Generator](https://github.com/DocSpring/openapi-generator/tree/docspring/all_patches) 7.16.0.
The fork changes the Go generator:

- Models don't reject a response field added to the API after this release. Stock OpenAPI Generator decodes them with `DisallowUnknownFields`.

On top of the generated code, we add:

- A client constructor that reads your credentials, region and host from the environment.

Because the code is generated, we can't merge pull requests that change it. Please
[open an issue](https://github.com/DocSpring/docspring-go/issues) or email [support@docspring.com](mailto:support@docspring.com) instead, and
we'll fix it in the generator. [AGENTS.md](https://github.com/DocSpring/docspring-go/blob/main/AGENTS.md) is a short guide for AI coding agents.

## License

MIT. See [LICENSE](https://github.com/DocSpring/docspring-go/blob/main/LICENSE).
