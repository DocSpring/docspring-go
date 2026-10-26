# DocSpring Go SDK

The official Go client for the [DocSpring API](https://docspring.com/docs/api/). DocSpring fills PDF
templates with your data, generates PDFs from HTML templates, combines PDFs and collects e-signatures.

[Documentation](https://docspring.com/docs/) · [API reference](https://docspring.com/docs/api/) ·
[pkg.go.dev](https://pkg.go.dev/github.com/docspring/docspring-go/v4) · [Changelog](https://github.com/DocSpring/docspring-go/blob/main/CHANGELOG.md)

The code in this repository is generated from DocSpring's OpenAPI spec ([About this SDK](#about-this-sdk)). Issues and
pull requests are welcome.

Version 4.1.0. Requires Go 1.23 or later.

## Install and configure

Install the package:

```bash
go get github.com/docspring/docspring-go/v4
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
  client "github.com/docspring/docspring-go/v4"
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
body := client.CreateSubmissionRequest{
	Data:     map[string]interface{}{"first_name": "John", "last_name": "Smith", "phone_number": "+11234567890"},
	Metadata: map[string]interface{}{"user_id": 42},
	Test:     client.PtrBool(true),
}
request := apiClient.CreateSubmission(context.Background(), templateID).Body(body)
result, _, err := request.Execute()
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(result)
}
```

By default the call waits for the PDF, for up to 95 seconds, and the submission in the response has a link to
download it (`files.pdf.download_url`). Set `wait` to false to get a pending submission straight away.
[Sync vs. async processing](https://docspring.com/docs/api-guide/synchronous-requests/) covers the time limit, status
checks and webhooks.

Optional fields are pointers. Set them with the helpers `client.PtrString`, `client.PtrInt32`, `client.PtrBool` and so on, and read them with the model's getters (`GetDownloadUrl()`).

## Methods

[docs/ClientAPI.md](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md) has each method's parameters and return type, and
[docs/](https://github.com/DocSpring/docspring-go/tree/main/docs) has the models. The [API reference](https://docspring.com/docs/api/) has
Go code for every method.

| Method | Description |
| ------ | ----------- |
| [`AddTemplateFields`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#addtemplatefields) | Add fields to a template |
| [`CopyTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#copytemplate) | Copy a template |
| [`CreateCombinedPdf`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createcombinedpdf) | Combine PDFs |
| [`CreateCustomFileFromUpload`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createcustomfilefromupload) | Create a custom file from an upload |
| [`CreateFolder`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createfolder) | Create a folder |
| [`CreateSignatureRequestEvent`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createsignaturerequestevent) | Record a signature request event |
| [`CreateSignatureRequestToken`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createsignaturerequesttoken) | Create a signature request token |
| [`CreateSubmission`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createsubmission) | Generate a PDF |
| [`CreateSubmissionBatch`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createsubmissionbatch) | Generate PDFs in a batch |
| [`CreateTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createtemplate) | Create a template |
| [`CreateWebhook`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#createwebhook) | Create a webhook |
| [`DeleteCustomFile`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#deletecustomfile) | Delete a custom file |
| [`DeleteFolder`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#deletefolder) | Delete a folder |
| [`DeleteTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#deletetemplate) | Delete a template |
| [`DeleteWebhook`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#deletewebhook) | Delete a webhook |
| [`ExpireCombinedPdf`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#expirecombinedpdf) | Expire a combined PDF |
| [`ExpireSubmission`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#expiresubmission) | Expire a submission |
| [`GenerateSubmissionFiles`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#generatesubmissionfiles) | Generate more files |
| [`GenerateSubmissionPreview`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#generatesubmissionpreview) | Generate a preview PDF |
| [`GetCombinedPdf`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getcombinedpdf) | Get a combined PDF |
| [`GetCustomFile`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getcustomfile) | Get a custom file |
| [`GetFolder`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getfolder) | Get a folder |
| [`GetPresignUrl`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getpresignurl) | Get a presigned URL to upload a file |
| [`GetSignatureRequest`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getsignaturerequest) | Get a signature request |
| [`GetSubmission`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getsubmission) | Get a submission |
| [`GetSubmissionBatch`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getsubmissionbatch) | Get a batch |
| [`GetTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#gettemplate) | Get a template |
| [`GetTemplateDiff`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#gettemplatediff) | Compare a version of a template with the one before it |
| [`GetTemplateModelCode`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#gettemplatemodelcode) | Generate typed code for a template's submission data |
| [`GetTemplatePageImages`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#gettemplatepageimages) | Get images of a template's pages |
| [`GetTemplateSchema`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#gettemplateschema) | Get the JSON Schema of a template's submission data |
| [`GetWebhook`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#getwebhook) | Get a webhook |
| [`ListCombinedPdfs`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listcombinedpdfs) | List combined PDFs |
| [`ListCustomFiles`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listcustomfiles) | List custom files |
| [`ListFolders`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listfolders) | List folders |
| [`ListSubmissions`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listsubmissions) | List submissions |
| [`ListTemplateFields`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listtemplatefields) | List a template's fields |
| [`ListTemplates`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listtemplates) | List templates |
| [`ListTemplateSubmissions`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listtemplatesubmissions) | List a template's submissions |
| [`ListWebhooks`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#listwebhooks) | List webhooks |
| [`MoveFolder`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#movefolder) | Move a folder |
| [`MoveTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#movetemplate) | Move a template to a folder |
| [`PublishTemplateVersion`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#publishtemplateversion) | Publish a template version |
| [`RemoveTemplateFields`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#removetemplatefields) | Remove fields from a template |
| [`RenameFolder`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#renamefolder) | Rename a folder |
| [`ReorderTemplateFields`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#reordertemplatefields) | Reorder a template's fields |
| [`RestoreTemplateVersion`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#restoretemplateversion) | Restore a template version |
| [`TestAuthentication`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#testauthentication) | Test authentication |
| [`UpdateSignatureRequest`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#updatesignaturerequest) | Update a signature request |
| [`UpdateTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#updatetemplate) | Update a template |
| [`UpdateTemplateDocument`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#updatetemplatedocument) | Replace a template's PDF |
| [`UpdateTemplateFields`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#updatetemplatefields) | Update a template's fields |
| [`UpdateWebhook`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#updatewebhook) | Update a webhook |
| [`UploadTemplate`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#uploadtemplate) | Create a PDF template from an uploaded PDF |
| [`UploadTemplateDocument`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#uploadtemplatedocument) | Replace a template's PDF with an uploaded PDF |
| [`ValidateSubmission`](https://github.com/DocSpring/docspring-go/blob/main/docs/ClientAPI.md#validatesubmission) | Check a submission without creating it |

## Errors

`Execute()` returns an `error`. For an error response it's a `*client.GenericOpenAPIError`: `Body()` is the raw body and `Model()` the decoded error, and the `*http.Response` that `Execute()` also returns has the status code.

The client checks each request against the API's request schemas before sending it. For a request the API would reject, `Execute()` returns a `*client.ValidationError` without sending it: its `Details` have a `Code`, a `Message` and a `Pointer` into the body (or a `Parameter`). Set `SkipRequestValidation` on the `Configuration` to turn the check off.
The details are the ones the API's 422 response would have, so you get them without a round trip.

## API changes

DocSpring adds fields and enum values to API v2 without releasing a new API version. Changes that would break
existing code, such as removing or renaming a field, go in a new API version
([API changes](https://docspring.com/docs/api-guide/openapi-schema/#api-changes)). So this SDK keeps working with
responses that are newer than it is:

- The model keeps a response field added to the API after this release in its `AdditionalProperties` map. The objects defined inside a model, such as a template's `document`, leave it out.
- Enum values in responses are strings, so a value added to the API after this release comes back as it is: handle values you don't recognize.

Upgrade the SDK to read new fields and enum values. The [changelog](https://github.com/DocSpring/docspring-go/blob/main/CHANGELOG.md) lists
what each release adds.

## About this SDK

This SDK is generated from DocSpring's [OpenAPI spec](https://docspring.com/api-docs/v2/openapi.json) with
[DocSpring's fork of OpenAPI Generator](https://github.com/DocSpring/openapi-generator/tree/docspring/all_patches) 7.26.0.
The fork changes the Go generator:

- Models don't reject a response field added to the API after this release. Stock OpenAPI Generator decodes them with `DisallowUnknownFields`.
- An enum value added to the API after this release decodes as the API sent it, instead of failing the whole response. `IsValid()` says whether it's one this release knows.

On top of the generated code, we add:

- A client constructor that reads your credentials, region and host from the environment.

Each release regenerates every file, so a change made only here would be lost in the next one. Issues and pull
requests are welcome all the same: [open an issue](https://github.com/DocSpring/docspring-go/issues) or a pull request, or email
[support@docspring.com](mailto:support@docspring.com), and we'll make the fix in the generator so that every release
keeps it. [AGENTS.md](https://github.com/DocSpring/docspring-go/blob/main/AGENTS.md) is a guide for AI coding agents that integrate this SDK.

## License

MIT. See [LICENSE](https://github.com/DocSpring/docspring-go/blob/main/LICENSE).
