# EmailAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The email action's ID. | 
**TemplateId** | **NullableString** | The ID of the template it belongs to. Null for an account email action, which sends for many templates. | 
**TemplateIds** | **[]string** | For an account email action: the IDs of the templates it sends for. With `folder_ids`, it sends for a template that either list includes. Both empty sends for every template. Always empty for a template's own email action. | 
**FolderIds** | **[]string** | For an account email action: the IDs of folders whose templates it sends for, including the templates in the folders inside them. Always empty for a template's own email action. | 
**Name** | **NullableString** | A name for the email action, shown in the dashboard. Null shows its recipients instead. | 
**Enabled** | **bool** | Whether it sends emails. DocSpring turns an email action off after 3 days of failed emails, or straight away when one of the addresses in `to`, `cc` or `bcc` (not one from a field or metadata) bounces or marks an email as spam. | 
**To** | **[]string** | Who gets each email: email addresses, and tags for an address that each submission gives, from one of the template's text or email fields (`{{ data.client_email }}`, or `{{ data["client email"] }}` for a name with spaces) or from its metadata (`{{ metadata.agent_email }}`, or `{{ metadata.agent.email }}` for a value in an object). A tag whose value is blank or isn't a valid address is skipped for that submission, and so is a field an account email action names that the template doesn't have. `{{ signer.email }}` sends each of the submission's signers their own email, to them alone, and everyone else in `to`, `cc` and `bcc` one email together, without the signers. An email needs at least one entry in `to`, `cc` or `bcc`, and at most 10 across them. | 
**Cc** | **[]string** | Who is copied on each email, as in `to` (but not `{{ signer.email }}`). | 
**Bcc** | **[]string** | Who is blind-copied on each email, as in `to` (but not `{{ signer.email }}`). | 
**FromName** | **NullableString** | The sender's name. Null uses your account's name. From DocSpring's shared address, the name is followed by "via DocSpring". | 
**FromAddress** | **NullableString** | The sender's address, at one of your verified email domains. Null sends from DocSpring's shared address. | 
**ReplyTo** | **NullableString** | Where replies go: a team member's email address, or an address at one of your verified email domains. Null sends without a reply address. | 
**Subject** | **NullableString** | The subject line, a Liquid template with the variables `account.name`, `template.name`, `template.id`, `submission.id`, the submission's values as `data.&lt;field name&gt;` (`data["field name"]` for a name with spaces) and its metadata as `metadata.&lt;key&gt;`, such as `Your {{ template.name }}`. Null uses the template's name followed by the file it sends, such as "Lease PDF", or by "files" when it sends several. | 
**Message** | **NullableString** | The email's text, a Liquid template with the same variables as `subject`. `if`, `unless` and `case` tags work. Line breaks are kept. Null uses a short message that says the files are attached or ready to download. | 
**FileName** | **NullableString** | The attached files' name, a Liquid template with the same variables as `subject`, such as `{{ data.last_name }} lease`. Each file's extension (`.pdf`, `.docx` or `.odt`) is added. Null uses the submission's `pdf_filename` metadata when it has one, and otherwise the template's name. A download link's files have the names their other downloads have: the submission's `pdf_filename`, or its ID. | 
**FileFormats** | [**EmailActionFileFormats**](EmailActionFileFormats.md) | Which of the submission's files the email sends: `all` sends every file the submission made (its PDF, and a Word template's Word document and OpenDocument file when it made them), and `docx`, `odt` or `pdf` sends only that file. A submission that made none of them sends nothing, and its email shows as skipped. Only Word templates make Word documents and OpenDocument files, so another template's email action takes `pdf` or `all`. | 
**Delivery** | [**EmailActionDelivery**](EmailActionDelivery.md) | `attachment` attaches the files, or sends a download link when together they're larger than `max_attachment_mb`. `link` sends a download link, which works for 7 days. Accounts with a BAA start with `link`. | 
**MaxAttachmentMb** | **int32** | With `delivery` `attachment`, files larger than this many megabytes in all (1 MB is 1,000,000 bytes) are sent as a download link instead. Emailing a file makes it about 38% larger, and many company mail servers refuse emails over 10 MB, so the default is 7. Raise it only when every recipient's mail server accepts larger emails, such as Microsoft 365 (up to 36 MB) or Gmail (50 MB). | 
**SubmissionType** | [**EmailActionSubmissionType**](EmailActionSubmissionType.md) | Which submissions send emails: `live`, `test` or `all`. Test submissions only email your team's members, and skip every other recipient. | 
**SendWhen** | [**NullableEmailActionConditions**](EmailActionConditions.md) | Sends the email only when the submission's values meet these conditions, which work as a conditional logic rule's do. The values are the ones the PDF has: formulas worked out, and blank for a field that a rule hides. A field the template doesn't have (in an account email action's conditions) counts as blank. A submission that doesn't meet them sends nothing, and its email shows as skipped. Null sends for every submission. | 
**FailureNotificationsEnabled** | **bool** | Whether the account's owner gets an email when the email action's emails keep failing. | 
**ConsecutiveFailures** | **int32** | How many emails in a row have failed. | 
**FailingSince** | **NullableTime** | When the current run of failures started. Null when the last email was sent. | 
**LastFailureAt** | **NullableTime** | When an email last failed. | 
**LastFailureMessage** | **NullableString** | Why the last failed email failed. | 
**DisabledAt** | **NullableTime** | When the email action was turned off. Null while it is on. | 
**DisabledReason** | **NullableString** | Why DocSpring turned the email action off, such as after 3 days of failed emails or an address that bounced. | 
**Customization** | [**NullableEmailActionCustomization**](EmailActionCustomization.md) | In a template's list of email actions, its changes to an account email action, or null when it changes nothing. The email action's settings are then the ones it sends for that template. Null everywhere else. | 
**CreatedAt** | **time.Time** | When the email action was created. | 
**UpdatedAt** | **time.Time** | When the email action was last changed. | 

## Methods

### NewEmailAction

`func NewEmailAction(id string, templateId NullableString, templateIds []string, folderIds []string, name NullableString, enabled bool, to []string, cc []string, bcc []string, fromName NullableString, fromAddress NullableString, replyTo NullableString, subject NullableString, message NullableString, fileName NullableString, fileFormats EmailActionFileFormats, delivery EmailActionDelivery, maxAttachmentMb int32, submissionType EmailActionSubmissionType, sendWhen NullableEmailActionConditions, failureNotificationsEnabled bool, consecutiveFailures int32, failingSince NullableTime, lastFailureAt NullableTime, lastFailureMessage NullableString, disabledAt NullableTime, disabledReason NullableString, customization NullableEmailActionCustomization, createdAt time.Time, updatedAt time.Time, ) *EmailAction`

NewEmailAction instantiates a new EmailAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailActionWithDefaults

`func NewEmailActionWithDefaults() *EmailAction`

NewEmailActionWithDefaults instantiates a new EmailAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailAction) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *EmailAction) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *EmailAction) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *EmailAction) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### SetTemplateIdNil

`func (o *EmailAction) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *EmailAction) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetTemplateIds

`func (o *EmailAction) GetTemplateIds() []string`

GetTemplateIds returns the TemplateIds field if non-nil, zero value otherwise.

### GetTemplateIdsOk

`func (o *EmailAction) GetTemplateIdsOk() (*[]string, bool)`

GetTemplateIdsOk returns a tuple with the TemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIds

`func (o *EmailAction) SetTemplateIds(v []string)`

SetTemplateIds sets TemplateIds field to given value.


### GetFolderIds

`func (o *EmailAction) GetFolderIds() []string`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *EmailAction) GetFolderIdsOk() (*[]string, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *EmailAction) SetFolderIds(v []string)`

SetFolderIds sets FolderIds field to given value.


### GetName

`func (o *EmailAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailAction) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *EmailAction) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EmailAction) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnabled

`func (o *EmailAction) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EmailAction) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EmailAction) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTo

`func (o *EmailAction) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailAction) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailAction) SetTo(v []string)`

SetTo sets To field to given value.


### GetCc

`func (o *EmailAction) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *EmailAction) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *EmailAction) SetCc(v []string)`

SetCc sets Cc field to given value.


### GetBcc

`func (o *EmailAction) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *EmailAction) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *EmailAction) SetBcc(v []string)`

SetBcc sets Bcc field to given value.


### GetFromName

`func (o *EmailAction) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *EmailAction) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *EmailAction) SetFromName(v string)`

SetFromName sets FromName field to given value.


### SetFromNameNil

`func (o *EmailAction) SetFromNameNil(b bool)`

 SetFromNameNil sets the value for FromName to be an explicit nil

### UnsetFromName
`func (o *EmailAction) UnsetFromName()`

UnsetFromName ensures that no value is present for FromName, not even an explicit nil
### GetFromAddress

`func (o *EmailAction) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *EmailAction) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *EmailAction) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### SetFromAddressNil

`func (o *EmailAction) SetFromAddressNil(b bool)`

 SetFromAddressNil sets the value for FromAddress to be an explicit nil

### UnsetFromAddress
`func (o *EmailAction) UnsetFromAddress()`

UnsetFromAddress ensures that no value is present for FromAddress, not even an explicit nil
### GetReplyTo

`func (o *EmailAction) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *EmailAction) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *EmailAction) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.


### SetReplyToNil

`func (o *EmailAction) SetReplyToNil(b bool)`

 SetReplyToNil sets the value for ReplyTo to be an explicit nil

### UnsetReplyTo
`func (o *EmailAction) UnsetReplyTo()`

UnsetReplyTo ensures that no value is present for ReplyTo, not even an explicit nil
### GetSubject

`func (o *EmailAction) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailAction) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailAction) SetSubject(v string)`

SetSubject sets Subject field to given value.


### SetSubjectNil

`func (o *EmailAction) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *EmailAction) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *EmailAction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EmailAction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EmailAction) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *EmailAction) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *EmailAction) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetFileName

`func (o *EmailAction) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *EmailAction) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *EmailAction) SetFileName(v string)`

SetFileName sets FileName field to given value.


### SetFileNameNil

`func (o *EmailAction) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *EmailAction) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFileFormats

`func (o *EmailAction) GetFileFormats() EmailActionFileFormats`

GetFileFormats returns the FileFormats field if non-nil, zero value otherwise.

### GetFileFormatsOk

`func (o *EmailAction) GetFileFormatsOk() (*EmailActionFileFormats, bool)`

GetFileFormatsOk returns a tuple with the FileFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormats

`func (o *EmailAction) SetFileFormats(v EmailActionFileFormats)`

SetFileFormats sets FileFormats field to given value.


### GetDelivery

`func (o *EmailAction) GetDelivery() EmailActionDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *EmailAction) GetDeliveryOk() (*EmailActionDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *EmailAction) SetDelivery(v EmailActionDelivery)`

SetDelivery sets Delivery field to given value.


### GetMaxAttachmentMb

`func (o *EmailAction) GetMaxAttachmentMb() int32`

GetMaxAttachmentMb returns the MaxAttachmentMb field if non-nil, zero value otherwise.

### GetMaxAttachmentMbOk

`func (o *EmailAction) GetMaxAttachmentMbOk() (*int32, bool)`

GetMaxAttachmentMbOk returns a tuple with the MaxAttachmentMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttachmentMb

`func (o *EmailAction) SetMaxAttachmentMb(v int32)`

SetMaxAttachmentMb sets MaxAttachmentMb field to given value.


### GetSubmissionType

`func (o *EmailAction) GetSubmissionType() EmailActionSubmissionType`

GetSubmissionType returns the SubmissionType field if non-nil, zero value otherwise.

### GetSubmissionTypeOk

`func (o *EmailAction) GetSubmissionTypeOk() (*EmailActionSubmissionType, bool)`

GetSubmissionTypeOk returns a tuple with the SubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionType

`func (o *EmailAction) SetSubmissionType(v EmailActionSubmissionType)`

SetSubmissionType sets SubmissionType field to given value.


### GetSendWhen

`func (o *EmailAction) GetSendWhen() EmailActionConditions`

GetSendWhen returns the SendWhen field if non-nil, zero value otherwise.

### GetSendWhenOk

`func (o *EmailAction) GetSendWhenOk() (*EmailActionConditions, bool)`

GetSendWhenOk returns a tuple with the SendWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendWhen

`func (o *EmailAction) SetSendWhen(v EmailActionConditions)`

SetSendWhen sets SendWhen field to given value.


### SetSendWhenNil

`func (o *EmailAction) SetSendWhenNil(b bool)`

 SetSendWhenNil sets the value for SendWhen to be an explicit nil

### UnsetSendWhen
`func (o *EmailAction) UnsetSendWhen()`

UnsetSendWhen ensures that no value is present for SendWhen, not even an explicit nil
### GetFailureNotificationsEnabled

`func (o *EmailAction) GetFailureNotificationsEnabled() bool`

GetFailureNotificationsEnabled returns the FailureNotificationsEnabled field if non-nil, zero value otherwise.

### GetFailureNotificationsEnabledOk

`func (o *EmailAction) GetFailureNotificationsEnabledOk() (*bool, bool)`

GetFailureNotificationsEnabledOk returns a tuple with the FailureNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureNotificationsEnabled

`func (o *EmailAction) SetFailureNotificationsEnabled(v bool)`

SetFailureNotificationsEnabled sets FailureNotificationsEnabled field to given value.


### GetConsecutiveFailures

`func (o *EmailAction) GetConsecutiveFailures() int32`

GetConsecutiveFailures returns the ConsecutiveFailures field if non-nil, zero value otherwise.

### GetConsecutiveFailuresOk

`func (o *EmailAction) GetConsecutiveFailuresOk() (*int32, bool)`

GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveFailures

`func (o *EmailAction) SetConsecutiveFailures(v int32)`

SetConsecutiveFailures sets ConsecutiveFailures field to given value.


### GetFailingSince

`func (o *EmailAction) GetFailingSince() time.Time`

GetFailingSince returns the FailingSince field if non-nil, zero value otherwise.

### GetFailingSinceOk

`func (o *EmailAction) GetFailingSinceOk() (*time.Time, bool)`

GetFailingSinceOk returns a tuple with the FailingSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingSince

`func (o *EmailAction) SetFailingSince(v time.Time)`

SetFailingSince sets FailingSince field to given value.


### SetFailingSinceNil

`func (o *EmailAction) SetFailingSinceNil(b bool)`

 SetFailingSinceNil sets the value for FailingSince to be an explicit nil

### UnsetFailingSince
`func (o *EmailAction) UnsetFailingSince()`

UnsetFailingSince ensures that no value is present for FailingSince, not even an explicit nil
### GetLastFailureAt

`func (o *EmailAction) GetLastFailureAt() time.Time`

GetLastFailureAt returns the LastFailureAt field if non-nil, zero value otherwise.

### GetLastFailureAtOk

`func (o *EmailAction) GetLastFailureAtOk() (*time.Time, bool)`

GetLastFailureAtOk returns a tuple with the LastFailureAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureAt

`func (o *EmailAction) SetLastFailureAt(v time.Time)`

SetLastFailureAt sets LastFailureAt field to given value.


### SetLastFailureAtNil

`func (o *EmailAction) SetLastFailureAtNil(b bool)`

 SetLastFailureAtNil sets the value for LastFailureAt to be an explicit nil

### UnsetLastFailureAt
`func (o *EmailAction) UnsetLastFailureAt()`

UnsetLastFailureAt ensures that no value is present for LastFailureAt, not even an explicit nil
### GetLastFailureMessage

`func (o *EmailAction) GetLastFailureMessage() string`

GetLastFailureMessage returns the LastFailureMessage field if non-nil, zero value otherwise.

### GetLastFailureMessageOk

`func (o *EmailAction) GetLastFailureMessageOk() (*string, bool)`

GetLastFailureMessageOk returns a tuple with the LastFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureMessage

`func (o *EmailAction) SetLastFailureMessage(v string)`

SetLastFailureMessage sets LastFailureMessage field to given value.


### SetLastFailureMessageNil

`func (o *EmailAction) SetLastFailureMessageNil(b bool)`

 SetLastFailureMessageNil sets the value for LastFailureMessage to be an explicit nil

### UnsetLastFailureMessage
`func (o *EmailAction) UnsetLastFailureMessage()`

UnsetLastFailureMessage ensures that no value is present for LastFailureMessage, not even an explicit nil
### GetDisabledAt

`func (o *EmailAction) GetDisabledAt() time.Time`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *EmailAction) GetDisabledAtOk() (*time.Time, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *EmailAction) SetDisabledAt(v time.Time)`

SetDisabledAt sets DisabledAt field to given value.


### SetDisabledAtNil

`func (o *EmailAction) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *EmailAction) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetDisabledReason

`func (o *EmailAction) GetDisabledReason() string`

GetDisabledReason returns the DisabledReason field if non-nil, zero value otherwise.

### GetDisabledReasonOk

`func (o *EmailAction) GetDisabledReasonOk() (*string, bool)`

GetDisabledReasonOk returns a tuple with the DisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReason

`func (o *EmailAction) SetDisabledReason(v string)`

SetDisabledReason sets DisabledReason field to given value.


### SetDisabledReasonNil

`func (o *EmailAction) SetDisabledReasonNil(b bool)`

 SetDisabledReasonNil sets the value for DisabledReason to be an explicit nil

### UnsetDisabledReason
`func (o *EmailAction) UnsetDisabledReason()`

UnsetDisabledReason ensures that no value is present for DisabledReason, not even an explicit nil
### GetCustomization

`func (o *EmailAction) GetCustomization() EmailActionCustomization`

GetCustomization returns the Customization field if non-nil, zero value otherwise.

### GetCustomizationOk

`func (o *EmailAction) GetCustomizationOk() (*EmailActionCustomization, bool)`

GetCustomizationOk returns a tuple with the Customization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomization

`func (o *EmailAction) SetCustomization(v EmailActionCustomization)`

SetCustomization sets Customization field to given value.


### SetCustomizationNil

`func (o *EmailAction) SetCustomizationNil(b bool)`

 SetCustomizationNil sets the value for Customization to be an explicit nil

### UnsetCustomization
`func (o *EmailAction) UnsetCustomization()`

UnsetCustomization ensures that no value is present for Customization, not even an explicit nil
### GetCreatedAt

`func (o *EmailAction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailAction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailAction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EmailAction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EmailAction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EmailAction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


