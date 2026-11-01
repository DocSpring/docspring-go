# UpdateEmailActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | A name for the email action, shown in the dashboard. Null shows its recipients instead. | [optional] 
**Enabled** | Pointer to **bool** | Whether it sends emails. DocSpring turns an email action off after 3 days of failed emails, or straight away when one of the addresses in `to`, `cc` or `bcc` (not one from a field or metadata) bounces or marks an email as spam. | [optional] 
**To** | Pointer to **[]string** | Who gets each email: email addresses, and tags for an address that each submission gives, from one of the template's text or email fields (`{{ data.client_email }}`, or `{{ data["client email"] }}` for a name with spaces) or from its metadata (`{{ metadata.agent_email }}`, or `{{ metadata.agent.email }}` for a value in an object). A tag whose value is blank or isn't a valid address is skipped for that submission, and so is a field an account email action names that the template doesn't have. `{{ signer.email }}` sends each of the submission's signers their own email, to them alone, and everyone else in `to`, `cc` and `bcc` one email together, without the signers. An email needs at least one entry in `to`, `cc` or `bcc`, and at most 10 across them. | [optional] 
**Cc** | Pointer to **[]string** | Who is copied on each email, as in `to` (but not `{{ signer.email }}`). | [optional] 
**Bcc** | Pointer to **[]string** | Who is blind-copied on each email, as in `to` (but not `{{ signer.email }}`). | [optional] 
**FromName** | Pointer to **NullableString** | The sender's name. Null uses your account's name. From DocSpring's shared address, the name is followed by "via DocSpring". | [optional] 
**FromAddress** | Pointer to **NullableString** | The sender's address, at one of your verified email domains. Null sends from DocSpring's shared address. | [optional] 
**ReplyTo** | Pointer to **NullableString** | Where replies go: a team member's email address, or an address at one of your verified email domains. Null sends without a reply address. | [optional] 
**Subject** | Pointer to **NullableString** | The subject line, a Liquid template with the variables `account.name`, `template.name`, `template.id`, `submission.id`, the submission's values as `data.&lt;field name&gt;` (`data["field name"]` for a name with spaces) and its metadata as `metadata.&lt;key&gt;`, such as `Your {{ template.name }}`. Null uses the template's name followed by the file it sends, such as "Lease PDF", or by "files" when it sends several. | [optional] 
**Message** | Pointer to **NullableString** | The email's text, a Liquid template with the same variables as `subject`. `if`, `unless` and `case` tags work. Line breaks are kept. Null uses a short message that says the files are attached or ready to download. | [optional] 
**FileName** | Pointer to **NullableString** | The attached files' name, a Liquid template with the same variables as `subject`, such as `{{ data.last_name }} lease`. Each file's extension (`.pdf`, `.docx` or `.odt`) is added. Null uses the submission's `pdf_filename` metadata when it has one, and otherwise the template's name. A download link's files have the names their other downloads have: the submission's `pdf_filename`, or its ID. | [optional] 
**FileFormats** | Pointer to [**EmailActionFileFormats**](EmailActionFileFormats.md) | Which of the submission's files the email sends: `all` sends every file the submission made (its PDF, and a Word template's Word document and OpenDocument file when it made them), and `docx`, `odt` or `pdf` sends only that file. A submission that made none of them sends nothing, and its email shows as skipped. Only Word templates make Word documents and OpenDocument files, so another template's email action takes `pdf` or `all`. | [optional] 
**Delivery** | Pointer to [**EmailActionDelivery**](EmailActionDelivery.md) | `attachment` attaches the files, or sends a download link when together they're larger than `max_attachment_mb`. `link` sends a download link, which works for 7 days. Accounts with a BAA start with `link`. | [optional] 
**MaxAttachmentMb** | Pointer to **int32** | With `delivery` `attachment`, files larger than this many megabytes in all (1 MB is 1,000,000 bytes) are sent as a download link instead. Emailing a file makes it about 38% larger, and many company mail servers refuse emails over 10 MB, so the default is 7. Raise it only when every recipient's mail server accepts larger emails, such as Microsoft 365 (up to 36 MB) or Gmail (50 MB). | [optional] 
**SubmissionType** | Pointer to [**EmailActionSubmissionType**](EmailActionSubmissionType.md) | Which submissions send emails: `live`, `test` or `all`. Test submissions only email your team's members, and skip every other recipient. | [optional] 
**SendWhen** | Pointer to [**NullableEmailActionConditionsInput**](EmailActionConditionsInput.md) | Sends the email only when the submission's values meet these conditions, which work as a conditional logic rule's do. The values are the ones the PDF has: formulas worked out, and blank for a field that a rule hides. A field the template doesn't have (in an account email action's conditions) counts as blank. A submission that doesn't meet them sends nothing, and its email shows as skipped. Null sends for every submission. | [optional] 
**FailureNotificationsEnabled** | Pointer to **bool** | Whether the account's owner gets an email when the email action's emails keep failing. | [optional] 

## Methods

### NewUpdateEmailActionRequest

`func NewUpdateEmailActionRequest() *UpdateEmailActionRequest`

NewUpdateEmailActionRequest instantiates a new UpdateEmailActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEmailActionRequestWithDefaults

`func NewUpdateEmailActionRequestWithDefaults() *UpdateEmailActionRequest`

NewUpdateEmailActionRequestWithDefaults instantiates a new UpdateEmailActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateEmailActionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEmailActionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEmailActionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateEmailActionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateEmailActionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateEmailActionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnabled

`func (o *UpdateEmailActionRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateEmailActionRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateEmailActionRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateEmailActionRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTo

`func (o *UpdateEmailActionRequest) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *UpdateEmailActionRequest) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *UpdateEmailActionRequest) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *UpdateEmailActionRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCc

`func (o *UpdateEmailActionRequest) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *UpdateEmailActionRequest) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *UpdateEmailActionRequest) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *UpdateEmailActionRequest) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *UpdateEmailActionRequest) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *UpdateEmailActionRequest) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *UpdateEmailActionRequest) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *UpdateEmailActionRequest) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetFromName

`func (o *UpdateEmailActionRequest) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *UpdateEmailActionRequest) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *UpdateEmailActionRequest) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *UpdateEmailActionRequest) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### SetFromNameNil

`func (o *UpdateEmailActionRequest) SetFromNameNil(b bool)`

 SetFromNameNil sets the value for FromName to be an explicit nil

### UnsetFromName
`func (o *UpdateEmailActionRequest) UnsetFromName()`

UnsetFromName ensures that no value is present for FromName, not even an explicit nil
### GetFromAddress

`func (o *UpdateEmailActionRequest) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *UpdateEmailActionRequest) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *UpdateEmailActionRequest) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *UpdateEmailActionRequest) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### SetFromAddressNil

`func (o *UpdateEmailActionRequest) SetFromAddressNil(b bool)`

 SetFromAddressNil sets the value for FromAddress to be an explicit nil

### UnsetFromAddress
`func (o *UpdateEmailActionRequest) UnsetFromAddress()`

UnsetFromAddress ensures that no value is present for FromAddress, not even an explicit nil
### GetReplyTo

`func (o *UpdateEmailActionRequest) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *UpdateEmailActionRequest) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *UpdateEmailActionRequest) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *UpdateEmailActionRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### SetReplyToNil

`func (o *UpdateEmailActionRequest) SetReplyToNil(b bool)`

 SetReplyToNil sets the value for ReplyTo to be an explicit nil

### UnsetReplyTo
`func (o *UpdateEmailActionRequest) UnsetReplyTo()`

UnsetReplyTo ensures that no value is present for ReplyTo, not even an explicit nil
### GetSubject

`func (o *UpdateEmailActionRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UpdateEmailActionRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UpdateEmailActionRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UpdateEmailActionRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *UpdateEmailActionRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *UpdateEmailActionRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *UpdateEmailActionRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateEmailActionRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateEmailActionRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateEmailActionRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *UpdateEmailActionRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *UpdateEmailActionRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetFileName

`func (o *UpdateEmailActionRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *UpdateEmailActionRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *UpdateEmailActionRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *UpdateEmailActionRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *UpdateEmailActionRequest) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *UpdateEmailActionRequest) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFileFormats

`func (o *UpdateEmailActionRequest) GetFileFormats() EmailActionFileFormats`

GetFileFormats returns the FileFormats field if non-nil, zero value otherwise.

### GetFileFormatsOk

`func (o *UpdateEmailActionRequest) GetFileFormatsOk() (*EmailActionFileFormats, bool)`

GetFileFormatsOk returns a tuple with the FileFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormats

`func (o *UpdateEmailActionRequest) SetFileFormats(v EmailActionFileFormats)`

SetFileFormats sets FileFormats field to given value.

### HasFileFormats

`func (o *UpdateEmailActionRequest) HasFileFormats() bool`

HasFileFormats returns a boolean if a field has been set.

### GetDelivery

`func (o *UpdateEmailActionRequest) GetDelivery() EmailActionDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *UpdateEmailActionRequest) GetDeliveryOk() (*EmailActionDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *UpdateEmailActionRequest) SetDelivery(v EmailActionDelivery)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *UpdateEmailActionRequest) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetMaxAttachmentMb

`func (o *UpdateEmailActionRequest) GetMaxAttachmentMb() int32`

GetMaxAttachmentMb returns the MaxAttachmentMb field if non-nil, zero value otherwise.

### GetMaxAttachmentMbOk

`func (o *UpdateEmailActionRequest) GetMaxAttachmentMbOk() (*int32, bool)`

GetMaxAttachmentMbOk returns a tuple with the MaxAttachmentMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttachmentMb

`func (o *UpdateEmailActionRequest) SetMaxAttachmentMb(v int32)`

SetMaxAttachmentMb sets MaxAttachmentMb field to given value.

### HasMaxAttachmentMb

`func (o *UpdateEmailActionRequest) HasMaxAttachmentMb() bool`

HasMaxAttachmentMb returns a boolean if a field has been set.

### GetSubmissionType

`func (o *UpdateEmailActionRequest) GetSubmissionType() EmailActionSubmissionType`

GetSubmissionType returns the SubmissionType field if non-nil, zero value otherwise.

### GetSubmissionTypeOk

`func (o *UpdateEmailActionRequest) GetSubmissionTypeOk() (*EmailActionSubmissionType, bool)`

GetSubmissionTypeOk returns a tuple with the SubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionType

`func (o *UpdateEmailActionRequest) SetSubmissionType(v EmailActionSubmissionType)`

SetSubmissionType sets SubmissionType field to given value.

### HasSubmissionType

`func (o *UpdateEmailActionRequest) HasSubmissionType() bool`

HasSubmissionType returns a boolean if a field has been set.

### GetSendWhen

`func (o *UpdateEmailActionRequest) GetSendWhen() EmailActionConditionsInput`

GetSendWhen returns the SendWhen field if non-nil, zero value otherwise.

### GetSendWhenOk

`func (o *UpdateEmailActionRequest) GetSendWhenOk() (*EmailActionConditionsInput, bool)`

GetSendWhenOk returns a tuple with the SendWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendWhen

`func (o *UpdateEmailActionRequest) SetSendWhen(v EmailActionConditionsInput)`

SetSendWhen sets SendWhen field to given value.

### HasSendWhen

`func (o *UpdateEmailActionRequest) HasSendWhen() bool`

HasSendWhen returns a boolean if a field has been set.

### SetSendWhenNil

`func (o *UpdateEmailActionRequest) SetSendWhenNil(b bool)`

 SetSendWhenNil sets the value for SendWhen to be an explicit nil

### UnsetSendWhen
`func (o *UpdateEmailActionRequest) UnsetSendWhen()`

UnsetSendWhen ensures that no value is present for SendWhen, not even an explicit nil
### GetFailureNotificationsEnabled

`func (o *UpdateEmailActionRequest) GetFailureNotificationsEnabled() bool`

GetFailureNotificationsEnabled returns the FailureNotificationsEnabled field if non-nil, zero value otherwise.

### GetFailureNotificationsEnabledOk

`func (o *UpdateEmailActionRequest) GetFailureNotificationsEnabledOk() (*bool, bool)`

GetFailureNotificationsEnabledOk returns a tuple with the FailureNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureNotificationsEnabled

`func (o *UpdateEmailActionRequest) SetFailureNotificationsEnabled(v bool)`

SetFailureNotificationsEnabled sets FailureNotificationsEnabled field to given value.

### HasFailureNotificationsEnabled

`func (o *UpdateEmailActionRequest) HasFailureNotificationsEnabled() bool`

HasFailureNotificationsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


