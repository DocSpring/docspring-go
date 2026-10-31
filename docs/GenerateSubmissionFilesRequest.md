# GenerateSubmissionFilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formats** | [**[]GenerateFilesFormat**](GenerateFilesFormat.md) | The formats to add to the submission's files: `pdf`, the PDF of a Word template's submission that didn't make one (converted from its Word document); `docx`, the Word document of a Word template's submission that didn't make one (made again from the submission's data and the template version it used); `odt`, the OpenDocument file of a Word template's submission that didn't make one (converted from its Word document); and `png`, an image of each page of the PDF (up to 100 pages), which makes the PDF first if the submission doesn't have one. A submission without a Word document makes it again for `pdf` and `odt`, without keeping it. A format the submission already has (or is making) is left as it is. | 

## Methods

### NewGenerateSubmissionFilesRequest

`func NewGenerateSubmissionFilesRequest(formats []GenerateFilesFormat, ) *GenerateSubmissionFilesRequest`

NewGenerateSubmissionFilesRequest instantiates a new GenerateSubmissionFilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateSubmissionFilesRequestWithDefaults

`func NewGenerateSubmissionFilesRequestWithDefaults() *GenerateSubmissionFilesRequest`

NewGenerateSubmissionFilesRequestWithDefaults instantiates a new GenerateSubmissionFilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormats

`func (o *GenerateSubmissionFilesRequest) GetFormats() []GenerateFilesFormat`

GetFormats returns the Formats field if non-nil, zero value otherwise.

### GetFormatsOk

`func (o *GenerateSubmissionFilesRequest) GetFormatsOk() (*[]GenerateFilesFormat, bool)`

GetFormatsOk returns a tuple with the Formats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormats

`func (o *GenerateSubmissionFilesRequest) SetFormats(v []GenerateFilesFormat)`

SetFormats sets Formats field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


