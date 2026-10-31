# ValidateSubmissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Test** | Pointer to **bool** | Test PDFs have a watermark and are free. Live PDFs use credits. Submissions made with a test API token are always test submissions. | [optional] 
**Editable** | Pointer to **bool** | `true` keeps the PDF's text and check fields fillable, and `false` flattens every field. Left out, the template decides: its `editable_submissions` setting, or each field's own Editable option. A response shows the template's setting when the request left this out. Can't be `true` with `signature_requests`. | [optional] 
**Data** | Pointer to **map[string]interface{}** | The values of the template's fields, by field name. Required for PDF and Word templates, unless `signature_requests` collects them (PDF templates only). | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Your own data about the submission, such as an ID from your app. It isn't printed on the PDF. Two keys also change the PDF: `pdf_filename` names its file (`files.pdf.filename`, and its download), and `pdf_title` is its title in the PDF's document properties (and an HTML template's `&lt;title&gt;`, unless its HTML has one). Without `pdf_title`, the title is the template's name. In `pdf_filename`, a `.pdf` ending is removed, each run of characters other than letters, digits, `-`, `_` and `.` becomes one `_`, and the first 128 characters are kept, then `.pdf` is added. Without it, the file is named after the submission's ID (`sub_1234567890abcdef01.pdf`). | [optional] 
**Password** | Pointer to **string** | A password that the PDF needs to be opened with. It's never returned; `has_password` says if it's set. Word documents and OpenDocument files can't be password-protected, so a submission with a password can't make them. | [optional] 
**OutputFormats** | Pointer to [**[]OutputFormat**](OutputFormat.md) | What the submission makes: `pdf`, `docx` (a Word document, from a Word template), `odt` (an OpenDocument file, the format LibreOffice uses, converted from a Word template's Word document), and `png` (an image of each page of the PDF, up to 100 pages, which adds `pdf`) when you ask for it. Default `["pdf"]`, or a Word template's `output_format`. POST /submissions/{submission_id}/generate_files adds `png`, and a Word template's submission's `pdf`, `docx` or `odt`, to a processed submission. | [optional] 
**ExpiresIn** | Pointer to **int32** | Seconds after the PDF is generated until the submission expires, up to 10 years (315360000). Default the template's expiration setting. A submission that isn't processed, such as one with invalid data, is kept for at least 48 hours. | [optional] 
**FieldOverrides** | Pointer to [**map[string]FieldOverride**](FieldOverride.md) | Changes to field settings for this submission only, by field name, such as `{"name": {"required": false, "font_size": 9}}`. A setting that identifies the field (`id`, `name`, `type`), that only the template sets (such as `description`), or that doesn't apply to the field's type and display type (such as `check_character` on a field displayed as text) is a 422. With `signature_requests`, `editable: true` is a 422, and a `default` for a signature or initials field makes the submission `invalid_data`. | [optional] 
**Html** | Pointer to **string** | For an HTML template, HTML to use instead of the template's. | [optional] 
**Css** | Pointer to **string** | For an HTML template, SCSS (plain CSS works too) added after the template's styles, on the page and in its header and footer. It doesn't replace them: its rules come later, so they win over the template's rules of the same specificity. | [optional] 
**SignatureRequests** | Pointer to [**[]SignatureRequestInput**](SignatureRequestInput.md) | The people who sign the PDF, and fill in their fields first. The PDF is generated once every signature request is completed. PDF templates only. Without `role`s, the first signature request is `Signer 1`, the second `Signer 2`, and so on. To leave a signer out, give every signature request its `role` and leave that one's out. Only the signers sign: a value for a signature or initials field in `data` makes the submission `invalid_data`. | [optional] 

## Methods

### NewValidateSubmissionRequest

`func NewValidateSubmissionRequest() *ValidateSubmissionRequest`

NewValidateSubmissionRequest instantiates a new ValidateSubmissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateSubmissionRequestWithDefaults

`func NewValidateSubmissionRequestWithDefaults() *ValidateSubmissionRequest`

NewValidateSubmissionRequestWithDefaults instantiates a new ValidateSubmissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTest

`func (o *ValidateSubmissionRequest) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *ValidateSubmissionRequest) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *ValidateSubmissionRequest) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *ValidateSubmissionRequest) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetEditable

`func (o *ValidateSubmissionRequest) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ValidateSubmissionRequest) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ValidateSubmissionRequest) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ValidateSubmissionRequest) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetData

`func (o *ValidateSubmissionRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ValidateSubmissionRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ValidateSubmissionRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ValidateSubmissionRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMetadata

`func (o *ValidateSubmissionRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ValidateSubmissionRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ValidateSubmissionRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ValidateSubmissionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPassword

`func (o *ValidateSubmissionRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ValidateSubmissionRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ValidateSubmissionRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ValidateSubmissionRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOutputFormats

`func (o *ValidateSubmissionRequest) GetOutputFormats() []OutputFormat`

GetOutputFormats returns the OutputFormats field if non-nil, zero value otherwise.

### GetOutputFormatsOk

`func (o *ValidateSubmissionRequest) GetOutputFormatsOk() (*[]OutputFormat, bool)`

GetOutputFormatsOk returns a tuple with the OutputFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormats

`func (o *ValidateSubmissionRequest) SetOutputFormats(v []OutputFormat)`

SetOutputFormats sets OutputFormats field to given value.

### HasOutputFormats

`func (o *ValidateSubmissionRequest) HasOutputFormats() bool`

HasOutputFormats returns a boolean if a field has been set.

### GetExpiresIn

`func (o *ValidateSubmissionRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ValidateSubmissionRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ValidateSubmissionRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *ValidateSubmissionRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetFieldOverrides

`func (o *ValidateSubmissionRequest) GetFieldOverrides() map[string]FieldOverride`

GetFieldOverrides returns the FieldOverrides field if non-nil, zero value otherwise.

### GetFieldOverridesOk

`func (o *ValidateSubmissionRequest) GetFieldOverridesOk() (*map[string]FieldOverride, bool)`

GetFieldOverridesOk returns a tuple with the FieldOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldOverrides

`func (o *ValidateSubmissionRequest) SetFieldOverrides(v map[string]FieldOverride)`

SetFieldOverrides sets FieldOverrides field to given value.

### HasFieldOverrides

`func (o *ValidateSubmissionRequest) HasFieldOverrides() bool`

HasFieldOverrides returns a boolean if a field has been set.

### GetHtml

`func (o *ValidateSubmissionRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *ValidateSubmissionRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *ValidateSubmissionRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *ValidateSubmissionRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetCss

`func (o *ValidateSubmissionRequest) GetCss() string`

GetCss returns the Css field if non-nil, zero value otherwise.

### GetCssOk

`func (o *ValidateSubmissionRequest) GetCssOk() (*string, bool)`

GetCssOk returns a tuple with the Css field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCss

`func (o *ValidateSubmissionRequest) SetCss(v string)`

SetCss sets Css field to given value.

### HasCss

`func (o *ValidateSubmissionRequest) HasCss() bool`

HasCss returns a boolean if a field has been set.

### GetSignatureRequests

`func (o *ValidateSubmissionRequest) GetSignatureRequests() []SignatureRequestInput`

GetSignatureRequests returns the SignatureRequests field if non-nil, zero value otherwise.

### GetSignatureRequestsOk

`func (o *ValidateSubmissionRequest) GetSignatureRequestsOk() (*[]SignatureRequestInput, bool)`

GetSignatureRequestsOk returns a tuple with the SignatureRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureRequests

`func (o *ValidateSubmissionRequest) SetSignatureRequests(v []SignatureRequestInput)`

SetSignatureRequests sets SignatureRequests field to given value.

### HasSignatureRequests

`func (o *ValidateSubmissionRequest) HasSignatureRequests() bool`

HasSignatureRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


