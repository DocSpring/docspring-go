# UpdateAccountEmailActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateIds** | Pointer to **[]string** | For an account email action: the IDs of the templates it sends for. With `folder_ids`, it sends for a template that either list includes. Both empty sends for every template. Always empty for a template's own email action. | [optional] 
**FolderIds** | Pointer to **[]string** | For an account email action: the IDs of folders whose templates it sends for, including the templates in the folders inside them. Always empty for a template's own email action. | [optional] 
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

### NewUpdateAccountEmailActionRequest

`func NewUpdateAccountEmailActionRequest() *UpdateAccountEmailActionRequest`

NewUpdateAccountEmailActionRequest instantiates a new UpdateAccountEmailActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountEmailActionRequestWithDefaults

`func NewUpdateAccountEmailActionRequestWithDefaults() *UpdateAccountEmailActionRequest`

NewUpdateAccountEmailActionRequestWithDefaults instantiates a new UpdateAccountEmailActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateIds

`func (o *UpdateAccountEmailActionRequest) GetTemplateIds() []string`

GetTemplateIds returns the TemplateIds field if non-nil, zero value otherwise.

### GetTemplateIdsOk

`func (o *UpdateAccountEmailActionRequest) GetTemplateIdsOk() (*[]string, bool)`

GetTemplateIdsOk returns a tuple with the TemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIds

`func (o *UpdateAccountEmailActionRequest) SetTemplateIds(v []string)`

SetTemplateIds sets TemplateIds field to given value.

### HasTemplateIds

`func (o *UpdateAccountEmailActionRequest) HasTemplateIds() bool`

HasTemplateIds returns a boolean if a field has been set.

### GetFolderIds

`func (o *UpdateAccountEmailActionRequest) GetFolderIds() []string`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *UpdateAccountEmailActionRequest) GetFolderIdsOk() (*[]string, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *UpdateAccountEmailActionRequest) SetFolderIds(v []string)`

SetFolderIds sets FolderIds field to given value.

### HasFolderIds

`func (o *UpdateAccountEmailActionRequest) HasFolderIds() bool`

HasFolderIds returns a boolean if a field has been set.

### GetName

`func (o *UpdateAccountEmailActionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAccountEmailActionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAccountEmailActionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAccountEmailActionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateAccountEmailActionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateAccountEmailActionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnabled

`func (o *UpdateAccountEmailActionRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateAccountEmailActionRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateAccountEmailActionRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateAccountEmailActionRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTo

`func (o *UpdateAccountEmailActionRequest) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *UpdateAccountEmailActionRequest) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *UpdateAccountEmailActionRequest) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *UpdateAccountEmailActionRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCc

`func (o *UpdateAccountEmailActionRequest) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *UpdateAccountEmailActionRequest) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *UpdateAccountEmailActionRequest) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *UpdateAccountEmailActionRequest) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *UpdateAccountEmailActionRequest) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *UpdateAccountEmailActionRequest) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *UpdateAccountEmailActionRequest) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *UpdateAccountEmailActionRequest) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetFromName

`func (o *UpdateAccountEmailActionRequest) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *UpdateAccountEmailActionRequest) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *UpdateAccountEmailActionRequest) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *UpdateAccountEmailActionRequest) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### SetFromNameNil

`func (o *UpdateAccountEmailActionRequest) SetFromNameNil(b bool)`

 SetFromNameNil sets the value for FromName to be an explicit nil

### UnsetFromName
`func (o *UpdateAccountEmailActionRequest) UnsetFromName()`

UnsetFromName ensures that no value is present for FromName, not even an explicit nil
### GetFromAddress

`func (o *UpdateAccountEmailActionRequest) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *UpdateAccountEmailActionRequest) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *UpdateAccountEmailActionRequest) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *UpdateAccountEmailActionRequest) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### SetFromAddressNil

`func (o *UpdateAccountEmailActionRequest) SetFromAddressNil(b bool)`

 SetFromAddressNil sets the value for FromAddress to be an explicit nil

### UnsetFromAddress
`func (o *UpdateAccountEmailActionRequest) UnsetFromAddress()`

UnsetFromAddress ensures that no value is present for FromAddress, not even an explicit nil
### GetReplyTo

`func (o *UpdateAccountEmailActionRequest) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *UpdateAccountEmailActionRequest) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *UpdateAccountEmailActionRequest) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *UpdateAccountEmailActionRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### SetReplyToNil

`func (o *UpdateAccountEmailActionRequest) SetReplyToNil(b bool)`

 SetReplyToNil sets the value for ReplyTo to be an explicit nil

### UnsetReplyTo
`func (o *UpdateAccountEmailActionRequest) UnsetReplyTo()`

UnsetReplyTo ensures that no value is present for ReplyTo, not even an explicit nil
### GetSubject

`func (o *UpdateAccountEmailActionRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UpdateAccountEmailActionRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UpdateAccountEmailActionRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UpdateAccountEmailActionRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *UpdateAccountEmailActionRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *UpdateAccountEmailActionRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *UpdateAccountEmailActionRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateAccountEmailActionRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateAccountEmailActionRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateAccountEmailActionRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *UpdateAccountEmailActionRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *UpdateAccountEmailActionRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetFileName

`func (o *UpdateAccountEmailActionRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *UpdateAccountEmailActionRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *UpdateAccountEmailActionRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *UpdateAccountEmailActionRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *UpdateAccountEmailActionRequest) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *UpdateAccountEmailActionRequest) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFileFormats

`func (o *UpdateAccountEmailActionRequest) GetFileFormats() EmailActionFileFormats`

GetFileFormats returns the FileFormats field if non-nil, zero value otherwise.

### GetFileFormatsOk

`func (o *UpdateAccountEmailActionRequest) GetFileFormatsOk() (*EmailActionFileFormats, bool)`

GetFileFormatsOk returns a tuple with the FileFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormats

`func (o *UpdateAccountEmailActionRequest) SetFileFormats(v EmailActionFileFormats)`

SetFileFormats sets FileFormats field to given value.

### HasFileFormats

`func (o *UpdateAccountEmailActionRequest) HasFileFormats() bool`

HasFileFormats returns a boolean if a field has been set.

### GetDelivery

`func (o *UpdateAccountEmailActionRequest) GetDelivery() EmailActionDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *UpdateAccountEmailActionRequest) GetDeliveryOk() (*EmailActionDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *UpdateAccountEmailActionRequest) SetDelivery(v EmailActionDelivery)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *UpdateAccountEmailActionRequest) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetMaxAttachmentMb

`func (o *UpdateAccountEmailActionRequest) GetMaxAttachmentMb() int32`

GetMaxAttachmentMb returns the MaxAttachmentMb field if non-nil, zero value otherwise.

### GetMaxAttachmentMbOk

`func (o *UpdateAccountEmailActionRequest) GetMaxAttachmentMbOk() (*int32, bool)`

GetMaxAttachmentMbOk returns a tuple with the MaxAttachmentMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttachmentMb

`func (o *UpdateAccountEmailActionRequest) SetMaxAttachmentMb(v int32)`

SetMaxAttachmentMb sets MaxAttachmentMb field to given value.

### HasMaxAttachmentMb

`func (o *UpdateAccountEmailActionRequest) HasMaxAttachmentMb() bool`

HasMaxAttachmentMb returns a boolean if a field has been set.

### GetSubmissionType

`func (o *UpdateAccountEmailActionRequest) GetSubmissionType() EmailActionSubmissionType`

GetSubmissionType returns the SubmissionType field if non-nil, zero value otherwise.

### GetSubmissionTypeOk

`func (o *UpdateAccountEmailActionRequest) GetSubmissionTypeOk() (*EmailActionSubmissionType, bool)`

GetSubmissionTypeOk returns a tuple with the SubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionType

`func (o *UpdateAccountEmailActionRequest) SetSubmissionType(v EmailActionSubmissionType)`

SetSubmissionType sets SubmissionType field to given value.

### HasSubmissionType

`func (o *UpdateAccountEmailActionRequest) HasSubmissionType() bool`

HasSubmissionType returns a boolean if a field has been set.

### GetSendWhen

`func (o *UpdateAccountEmailActionRequest) GetSendWhen() EmailActionConditionsInput`

GetSendWhen returns the SendWhen field if non-nil, zero value otherwise.

### GetSendWhenOk

`func (o *UpdateAccountEmailActionRequest) GetSendWhenOk() (*EmailActionConditionsInput, bool)`

GetSendWhenOk returns a tuple with the SendWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendWhen

`func (o *UpdateAccountEmailActionRequest) SetSendWhen(v EmailActionConditionsInput)`

SetSendWhen sets SendWhen field to given value.

### HasSendWhen

`func (o *UpdateAccountEmailActionRequest) HasSendWhen() bool`

HasSendWhen returns a boolean if a field has been set.

### SetSendWhenNil

`func (o *UpdateAccountEmailActionRequest) SetSendWhenNil(b bool)`

 SetSendWhenNil sets the value for SendWhen to be an explicit nil

### UnsetSendWhen
`func (o *UpdateAccountEmailActionRequest) UnsetSendWhen()`

UnsetSendWhen ensures that no value is present for SendWhen, not even an explicit nil
### GetFailureNotificationsEnabled

`func (o *UpdateAccountEmailActionRequest) GetFailureNotificationsEnabled() bool`

GetFailureNotificationsEnabled returns the FailureNotificationsEnabled field if non-nil, zero value otherwise.

### GetFailureNotificationsEnabledOk

`func (o *UpdateAccountEmailActionRequest) GetFailureNotificationsEnabledOk() (*bool, bool)`

GetFailureNotificationsEnabledOk returns a tuple with the FailureNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureNotificationsEnabled

`func (o *UpdateAccountEmailActionRequest) SetFailureNotificationsEnabled(v bool)`

SetFailureNotificationsEnabled sets FailureNotificationsEnabled field to given value.

### HasFailureNotificationsEnabled

`func (o *UpdateAccountEmailActionRequest) HasFailureNotificationsEnabled() bool`

HasFailureNotificationsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


