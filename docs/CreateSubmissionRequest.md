# CreateSubmissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Test** | Pointer to **bool** | Test PDFs have a watermark and are free. Live PDFs use credits. Submissions made with a test API token are always test submissions, and the demo API token's are always live. | [optional] 
**Editable** | Pointer to **bool** | `true` keeps the PDF's text and check fields fillable, and `false` flattens every field. Left out, the template decides: its `editable_submissions` setting, or each field's own Editable option. A response shows the template's setting when the request left this out. Can't be `true` with `signature_requests`. | [optional] 
**Data** | Pointer to **map[string]interface{}** | The values of the template's fields, by field name. Required for PDF templates, Word templates and page designs, unless `signature_requests` collects them (PDF templates only). | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Your own data about the submission, such as an ID from your app. It isn't printed on the PDF. Two keys also change the PDF: `pdf_filename` names its file (`files.pdf.filename`, and its download), and `pdf_title` is its title in the PDF's document properties (and an HTML template's `&lt;title&gt;`, unless its HTML has one). Without `pdf_title`, the title is the template's name. In `pdf_filename`, a `.pdf` ending is removed, each run of characters other than letters, digits, `-`, `_` and `.` becomes one `_`, and the first 128 characters are kept, then `.pdf` is added. Without it, the file is named after the submission's ID (`sub_1234567890abcdef01.pdf`). | [optional] 
**Password** | Pointer to **string** | A password that the PDF needs to be opened with. It's never returned; `has_password` says if it's set. Word documents and OpenDocument files can't be password-protected, so a submission with a password can't make them. | [optional] 
**OutputFormats** | Pointer to [**[]OutputFormat**](OutputFormat.md) | What the submission makes: `pdf`, `docx` (a Word document, from a Word template), `odt` (an OpenDocument file, the format LibreOffice uses, converted from a Word template's Word document), and `png` (an image of each page of the PDF, up to 100 pages, which adds `pdf`) when you ask for it. Default `["pdf"]`, or a Word template's `output_format`. POST /submissions/{submission_id}/generate_files adds `png`, and a Word template's submission's `pdf`, `docx` or `odt`, to a processed submission. | [optional] 
**ExpiresIn** | Pointer to **int32** | Seconds after the PDF is generated until the submission expires, up to 10 years (315360000). Default the template's expiration setting. A submission that isn't processed, such as one with invalid data, is kept for at least 48 hours. The demo API token's submissions expire within an hour: up to 3600, and that's their default. | [optional] 
**FieldOverrides** | Pointer to [**map[string]FieldOverride**](FieldOverride.md) | Changes to field settings for this submission only, by field name, such as `{"name": {"required": false, "font_size": 9}}`. A setting that identifies the field (`id`, `name`, `type`), that only the template sets (such as `description`), or that doesn't apply to the field's type and display type (such as `check_character` on a field displayed as text) is a 422. With `signature_requests`, `editable: true` is a 422, and a `default` for a signature or initials field makes the submission `invalid_data`. | [optional] 
**Html** | Pointer to **string** | For an HTML template, HTML to use instead of the template's. | [optional] 
**Css** | Pointer to **string** | For an HTML template, SCSS (plain CSS works too) added after the template's styles, on the page and in its header and footer. It doesn't replace them: its rules come later, so they win over the template's rules of the same specificity. | [optional] 
**SignatureRequests** | Pointer to [**[]SignatureRequestInput**](SignatureRequestInput.md) | The people who sign the PDF, and fill in their fields first. The PDF is generated once every signature request is completed. PDF templates only. Without `role`s, the first signature request is `Signer 1`, the second `Signer 2`, and so on. To leave a signer out, give every signature request its `role` and leave that one's out. Only the signers sign: a value for a signature or initials field in `data` makes the submission `invalid_data`. | [optional] 
**SendEmail** | Pointer to **bool** | Whether DocSpring emails the signers their signing links, and sends the reminders and the completed document. Default true. Each signature request's own `send_email` overrides it. A test submission only emails your DocSpring team: every other address is a 422. `true` needs `signature_requests`. DocSpring doesn't email the signers of a form library template: leave `send_email` out, and send them their links yourself. | [optional] 
**SigningOrder** | Pointer to [**SigningOrder**](SigningOrder.md) | `parallel` (the default): everyone can sign at once. `sequential`: people sign one step at a time (`signing_step`), and DocSpring emails each step when the one before it is done. Needs `signature_requests`. | [optional] 
**EmailSubject** | Pointer to **string** | The subject of the emails that ask people to sign. You can use `{{ signer.name }}`, `{{ signer.email }}`, `{{ document.name }}`, `{{ sender.name }}` and `{{ account.name }}`. Default "&lt;sender&gt; sent you &lt;document&gt; to sign" when you leave it out (a blank subject is a 422). Links and web addresses are shown as plain text (`example[.]com`). | [optional] 
**EmailMessage** | Pointer to **string** | A message for the people who sign, shown above the signing button. It takes the same variables as `email_subject`. Links and web addresses are shown as plain text. Leave it out for none: a blank message is a 422. | [optional] 
**Cc** | Pointer to **[]string** | People who get a copy of the completed document. They don't sign. | [optional] 
**ReplyTo** | Pointer to **string** | Where replies to the emails go: a member of your DocSpring team, or an address at one of your verified email domains. Leave it out for the template's reply-to address (the account owner when it has none). | [optional] 
**NotificationEmails** | Pointer to **[]string** | Who DocSpring tells when the document is completed or declined, or an email bounces or can't be sent: members of your DocSpring team or addresses at your verified email domains. Leave it out for the template's notification addresses (the account owner when it has none); an empty list means the account owner. | [optional] 
**NotifyWhenSigned** | Pointer to **bool** | Also email `notification_emails` each time someone signs. Default false. | [optional] 
**Reminders** | Pointer to [**SignatureRequestReminders**](SignatureRequestReminders.md) | Remind people who haven't signed: every `interval_days` days, up to `max_count` times. Each reminder has a new signing link. Without it, DocSpring sends no reminders. | [optional] 

## Methods

### NewCreateSubmissionRequest

`func NewCreateSubmissionRequest() *CreateSubmissionRequest`

NewCreateSubmissionRequest instantiates a new CreateSubmissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubmissionRequestWithDefaults

`func NewCreateSubmissionRequestWithDefaults() *CreateSubmissionRequest`

NewCreateSubmissionRequestWithDefaults instantiates a new CreateSubmissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTest

`func (o *CreateSubmissionRequest) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CreateSubmissionRequest) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CreateSubmissionRequest) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CreateSubmissionRequest) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetEditable

`func (o *CreateSubmissionRequest) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *CreateSubmissionRequest) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *CreateSubmissionRequest) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *CreateSubmissionRequest) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetData

`func (o *CreateSubmissionRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateSubmissionRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateSubmissionRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CreateSubmissionRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateSubmissionRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSubmissionRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSubmissionRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSubmissionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPassword

`func (o *CreateSubmissionRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateSubmissionRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateSubmissionRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateSubmissionRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOutputFormats

`func (o *CreateSubmissionRequest) GetOutputFormats() []OutputFormat`

GetOutputFormats returns the OutputFormats field if non-nil, zero value otherwise.

### GetOutputFormatsOk

`func (o *CreateSubmissionRequest) GetOutputFormatsOk() (*[]OutputFormat, bool)`

GetOutputFormatsOk returns a tuple with the OutputFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormats

`func (o *CreateSubmissionRequest) SetOutputFormats(v []OutputFormat)`

SetOutputFormats sets OutputFormats field to given value.

### HasOutputFormats

`func (o *CreateSubmissionRequest) HasOutputFormats() bool`

HasOutputFormats returns a boolean if a field has been set.

### GetExpiresIn

`func (o *CreateSubmissionRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreateSubmissionRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreateSubmissionRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreateSubmissionRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetFieldOverrides

`func (o *CreateSubmissionRequest) GetFieldOverrides() map[string]FieldOverride`

GetFieldOverrides returns the FieldOverrides field if non-nil, zero value otherwise.

### GetFieldOverridesOk

`func (o *CreateSubmissionRequest) GetFieldOverridesOk() (*map[string]FieldOverride, bool)`

GetFieldOverridesOk returns a tuple with the FieldOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldOverrides

`func (o *CreateSubmissionRequest) SetFieldOverrides(v map[string]FieldOverride)`

SetFieldOverrides sets FieldOverrides field to given value.

### HasFieldOverrides

`func (o *CreateSubmissionRequest) HasFieldOverrides() bool`

HasFieldOverrides returns a boolean if a field has been set.

### GetHtml

`func (o *CreateSubmissionRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *CreateSubmissionRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *CreateSubmissionRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *CreateSubmissionRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetCss

`func (o *CreateSubmissionRequest) GetCss() string`

GetCss returns the Css field if non-nil, zero value otherwise.

### GetCssOk

`func (o *CreateSubmissionRequest) GetCssOk() (*string, bool)`

GetCssOk returns a tuple with the Css field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCss

`func (o *CreateSubmissionRequest) SetCss(v string)`

SetCss sets Css field to given value.

### HasCss

`func (o *CreateSubmissionRequest) HasCss() bool`

HasCss returns a boolean if a field has been set.

### GetSignatureRequests

`func (o *CreateSubmissionRequest) GetSignatureRequests() []SignatureRequestInput`

GetSignatureRequests returns the SignatureRequests field if non-nil, zero value otherwise.

### GetSignatureRequestsOk

`func (o *CreateSubmissionRequest) GetSignatureRequestsOk() (*[]SignatureRequestInput, bool)`

GetSignatureRequestsOk returns a tuple with the SignatureRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureRequests

`func (o *CreateSubmissionRequest) SetSignatureRequests(v []SignatureRequestInput)`

SetSignatureRequests sets SignatureRequests field to given value.

### HasSignatureRequests

`func (o *CreateSubmissionRequest) HasSignatureRequests() bool`

HasSignatureRequests returns a boolean if a field has been set.

### GetSendEmail

`func (o *CreateSubmissionRequest) GetSendEmail() bool`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *CreateSubmissionRequest) GetSendEmailOk() (*bool, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *CreateSubmissionRequest) SetSendEmail(v bool)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *CreateSubmissionRequest) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.

### GetSigningOrder

`func (o *CreateSubmissionRequest) GetSigningOrder() SigningOrder`

GetSigningOrder returns the SigningOrder field if non-nil, zero value otherwise.

### GetSigningOrderOk

`func (o *CreateSubmissionRequest) GetSigningOrderOk() (*SigningOrder, bool)`

GetSigningOrderOk returns a tuple with the SigningOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningOrder

`func (o *CreateSubmissionRequest) SetSigningOrder(v SigningOrder)`

SetSigningOrder sets SigningOrder field to given value.

### HasSigningOrder

`func (o *CreateSubmissionRequest) HasSigningOrder() bool`

HasSigningOrder returns a boolean if a field has been set.

### GetEmailSubject

`func (o *CreateSubmissionRequest) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *CreateSubmissionRequest) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *CreateSubmissionRequest) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *CreateSubmissionRequest) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetEmailMessage

`func (o *CreateSubmissionRequest) GetEmailMessage() string`

GetEmailMessage returns the EmailMessage field if non-nil, zero value otherwise.

### GetEmailMessageOk

`func (o *CreateSubmissionRequest) GetEmailMessageOk() (*string, bool)`

GetEmailMessageOk returns a tuple with the EmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMessage

`func (o *CreateSubmissionRequest) SetEmailMessage(v string)`

SetEmailMessage sets EmailMessage field to given value.

### HasEmailMessage

`func (o *CreateSubmissionRequest) HasEmailMessage() bool`

HasEmailMessage returns a boolean if a field has been set.

### GetCc

`func (o *CreateSubmissionRequest) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *CreateSubmissionRequest) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *CreateSubmissionRequest) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *CreateSubmissionRequest) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetReplyTo

`func (o *CreateSubmissionRequest) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *CreateSubmissionRequest) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *CreateSubmissionRequest) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *CreateSubmissionRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetNotificationEmails

`func (o *CreateSubmissionRequest) GetNotificationEmails() []string`

GetNotificationEmails returns the NotificationEmails field if non-nil, zero value otherwise.

### GetNotificationEmailsOk

`func (o *CreateSubmissionRequest) GetNotificationEmailsOk() (*[]string, bool)`

GetNotificationEmailsOk returns a tuple with the NotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEmails

`func (o *CreateSubmissionRequest) SetNotificationEmails(v []string)`

SetNotificationEmails sets NotificationEmails field to given value.

### HasNotificationEmails

`func (o *CreateSubmissionRequest) HasNotificationEmails() bool`

HasNotificationEmails returns a boolean if a field has been set.

### GetNotifyWhenSigned

`func (o *CreateSubmissionRequest) GetNotifyWhenSigned() bool`

GetNotifyWhenSigned returns the NotifyWhenSigned field if non-nil, zero value otherwise.

### GetNotifyWhenSignedOk

`func (o *CreateSubmissionRequest) GetNotifyWhenSignedOk() (*bool, bool)`

GetNotifyWhenSignedOk returns a tuple with the NotifyWhenSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWhenSigned

`func (o *CreateSubmissionRequest) SetNotifyWhenSigned(v bool)`

SetNotifyWhenSigned sets NotifyWhenSigned field to given value.

### HasNotifyWhenSigned

`func (o *CreateSubmissionRequest) HasNotifyWhenSigned() bool`

HasNotifyWhenSigned returns a boolean if a field has been set.

### GetReminders

`func (o *CreateSubmissionRequest) GetReminders() SignatureRequestReminders`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *CreateSubmissionRequest) GetRemindersOk() (*SignatureRequestReminders, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *CreateSubmissionRequest) SetReminders(v SignatureRequestReminders)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *CreateSubmissionRequest) HasReminders() bool`

HasReminders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


