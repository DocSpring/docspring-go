# SubmissionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Test** | Pointer to **bool** | Test PDFs have a watermark and are free. Live PDFs use credits. Submissions made with a test API token are always test submissions. | [optional] 
**Editable** | Pointer to **bool** | `true` keeps the PDF's text and check fields fillable, and `false` flattens every field. Left out, the template decides: its `editable_submissions` setting, or each field's own Editable option. A response shows the template's setting when the request left this out. Can't be `true` with `signature_requests`. | [optional] 
**Data** | Pointer to **map[string]interface{}** | The values of the template's fields, by field name. Required for a PDF template. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Your own data about the submission, such as an ID from your app. It isn't printed on the PDF. Two keys also change the PDF: `pdf_filename` names its file (`files.pdf.filename`, and its download), and `pdf_title` is its title in the PDF's document properties (and an HTML template's `&lt;title&gt;`, unless its HTML has one). Without `pdf_title`, the title is the template's name. In `pdf_filename`, a `.pdf` ending is removed, each run of characters other than letters, digits, `-`, `_` and `.` becomes one `_`, and the first 128 characters are kept, then `.pdf` is added. Without it, the file is named after the submission's ID (`sub_1234567890abcdef01.pdf`). | [optional] 
**Password** | Pointer to **string** | A password that the PDF needs to be opened with. It's never returned; `has_password` says if it's set. | [optional] 
**OutputFormats** | Pointer to [**[]OutputFormat**](OutputFormat.md) | What the submission makes: `pdf`, and `png` (an image of each page of the PDF, up to 100 pages) when you ask for it. Default `["pdf"]`. POST /submissions/{submission_id}/generate_files adds `png` to a processed submission. | [optional] 
**ExpiresIn** | Pointer to **int32** | Seconds after the PDF is generated until the submission expires, up to 10 years (315360000). Default the template's expiration setting. A submission that isn't processed, such as one with invalid data, is kept for at least 48 hours. | [optional] 
**FieldOverrides** | Pointer to [**map[string]FieldOverride**](FieldOverride.md) | Changes to field settings for this submission only, by field name, such as `{"name": {"required": false, "font_size": 9}}`. A setting that identifies the field (`id`, `name`, `type`), that only the template sets (such as `description`), or that doesn't apply to the field's type and display type (such as `check_character` on a field displayed as text) is a 422. With `signature_requests`, `editable: true` is a 422, and a `default` for a signature or initials field makes the submission `invalid_data`. | [optional] 
**Html** | Pointer to **string** | For an HTML template, HTML to use instead of the template's. | [optional] 
**Css** | Pointer to **string** | For an HTML template, SCSS (plain CSS works too) added after the template's styles, on the page and in its header and footer. It doesn't replace them: its rules come later, so they win over the template's rules of the same specificity. | [optional] 
**TemplateId** | Pointer to **string** | The submission's template. Default the batch's `template_id`. | [optional] 
**TemplateVersion** | Pointer to **string** | The version of the submission's template. Default the batch's `template_version`. | [optional] 

## Methods

### NewSubmissionInput

`func NewSubmissionInput() *SubmissionInput`

NewSubmissionInput instantiates a new SubmissionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionInputWithDefaults

`func NewSubmissionInputWithDefaults() *SubmissionInput`

NewSubmissionInputWithDefaults instantiates a new SubmissionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTest

`func (o *SubmissionInput) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *SubmissionInput) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *SubmissionInput) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *SubmissionInput) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetEditable

`func (o *SubmissionInput) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *SubmissionInput) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *SubmissionInput) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *SubmissionInput) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetData

`func (o *SubmissionInput) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubmissionInput) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubmissionInput) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SubmissionInput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMetadata

`func (o *SubmissionInput) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubmissionInput) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubmissionInput) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubmissionInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPassword

`func (o *SubmissionInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SubmissionInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SubmissionInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SubmissionInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOutputFormats

`func (o *SubmissionInput) GetOutputFormats() []OutputFormat`

GetOutputFormats returns the OutputFormats field if non-nil, zero value otherwise.

### GetOutputFormatsOk

`func (o *SubmissionInput) GetOutputFormatsOk() (*[]OutputFormat, bool)`

GetOutputFormatsOk returns a tuple with the OutputFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormats

`func (o *SubmissionInput) SetOutputFormats(v []OutputFormat)`

SetOutputFormats sets OutputFormats field to given value.

### HasOutputFormats

`func (o *SubmissionInput) HasOutputFormats() bool`

HasOutputFormats returns a boolean if a field has been set.

### GetExpiresIn

`func (o *SubmissionInput) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *SubmissionInput) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *SubmissionInput) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *SubmissionInput) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetFieldOverrides

`func (o *SubmissionInput) GetFieldOverrides() map[string]FieldOverride`

GetFieldOverrides returns the FieldOverrides field if non-nil, zero value otherwise.

### GetFieldOverridesOk

`func (o *SubmissionInput) GetFieldOverridesOk() (*map[string]FieldOverride, bool)`

GetFieldOverridesOk returns a tuple with the FieldOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldOverrides

`func (o *SubmissionInput) SetFieldOverrides(v map[string]FieldOverride)`

SetFieldOverrides sets FieldOverrides field to given value.

### HasFieldOverrides

`func (o *SubmissionInput) HasFieldOverrides() bool`

HasFieldOverrides returns a boolean if a field has been set.

### GetHtml

`func (o *SubmissionInput) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *SubmissionInput) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *SubmissionInput) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *SubmissionInput) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetCss

`func (o *SubmissionInput) GetCss() string`

GetCss returns the Css field if non-nil, zero value otherwise.

### GetCssOk

`func (o *SubmissionInput) GetCssOk() (*string, bool)`

GetCssOk returns a tuple with the Css field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCss

`func (o *SubmissionInput) SetCss(v string)`

SetCss sets Css field to given value.

### HasCss

`func (o *SubmissionInput) HasCss() bool`

HasCss returns a boolean if a field has been set.

### GetTemplateId

`func (o *SubmissionInput) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SubmissionInput) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SubmissionInput) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *SubmissionInput) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateVersion

`func (o *SubmissionInput) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *SubmissionInput) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *SubmissionInput) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *SubmissionInput) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


