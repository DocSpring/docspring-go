# TemplateSignerRoleInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The role's ID, which never changes: fields refer to their role by it, so changing the role's key or name, or moving it in the list, keeps its fields. Leave it out for a new role. A role keeps the ID you give it, so roles copied from another template with their IDs keep the fields copied with them. | [optional] 
**Key** | Pointer to **string** | The role's name in the API and in the template's code, in snake_case: `role: tenant` on a field tag, a field's `signer_role`, a signature request's `role`, and `@@signature_requests/tenant/...` fields. Each role has its own, and `sender` is taken (it marks the fields that submission data fills in). Left out, it's made from `name`: Co-tenant 2 is `co_tenant_2`. | [optional] 
**Name** | **string** | The role's name, such as Tenant, which the template editor shows. Each role has its own, and `Signer N` can only be role N's. | 
**NameExpression** | Pointer to **NullableString** | The signer's name, for a signature request that doesn't give one: a fixed name, or text and `{{ variables }}` that name a path into the submission's data, such as `{{ tenant.name }}` or `{{ tenant.first_name }} {{ tenant.last_name }}`. No filters or tags. Each variable is one of the template's fields, unless the template allows additional properties: setting one that isn't is a `422`. A value the data doesn't have prints nothing, and a name that comes out blank is left out. Null for none. | [optional] 
**EmailExpression** | Pointer to **NullableString** | The signer's email address, for a signature request that doesn't give one: a fixed address, such as `leasing@example.com`, or one from the submission's data, written like `name_expression`: `{{ tenant.email }}`. The data has to have the value, so it's one of the template's fields (one the sender fills in) or the template allows additional properties: setting a variable that isn't a field is a `422` otherwise. Null for none. | [optional] 

## Methods

### NewTemplateSignerRoleInput

`func NewTemplateSignerRoleInput(name string, ) *TemplateSignerRoleInput`

NewTemplateSignerRoleInput instantiates a new TemplateSignerRoleInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateSignerRoleInputWithDefaults

`func NewTemplateSignerRoleInputWithDefaults() *TemplateSignerRoleInput`

NewTemplateSignerRoleInputWithDefaults instantiates a new TemplateSignerRoleInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateSignerRoleInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateSignerRoleInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateSignerRoleInput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateSignerRoleInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *TemplateSignerRoleInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TemplateSignerRoleInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TemplateSignerRoleInput) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TemplateSignerRoleInput) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *TemplateSignerRoleInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateSignerRoleInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateSignerRoleInput) SetName(v string)`

SetName sets Name field to given value.


### GetNameExpression

`func (o *TemplateSignerRoleInput) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *TemplateSignerRoleInput) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *TemplateSignerRoleInput) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.

### HasNameExpression

`func (o *TemplateSignerRoleInput) HasNameExpression() bool`

HasNameExpression returns a boolean if a field has been set.

### SetNameExpressionNil

`func (o *TemplateSignerRoleInput) SetNameExpressionNil(b bool)`

 SetNameExpressionNil sets the value for NameExpression to be an explicit nil

### UnsetNameExpression
`func (o *TemplateSignerRoleInput) UnsetNameExpression()`

UnsetNameExpression ensures that no value is present for NameExpression, not even an explicit nil
### GetEmailExpression

`func (o *TemplateSignerRoleInput) GetEmailExpression() string`

GetEmailExpression returns the EmailExpression field if non-nil, zero value otherwise.

### GetEmailExpressionOk

`func (o *TemplateSignerRoleInput) GetEmailExpressionOk() (*string, bool)`

GetEmailExpressionOk returns a tuple with the EmailExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailExpression

`func (o *TemplateSignerRoleInput) SetEmailExpression(v string)`

SetEmailExpression sets EmailExpression field to given value.

### HasEmailExpression

`func (o *TemplateSignerRoleInput) HasEmailExpression() bool`

HasEmailExpression returns a boolean if a field has been set.

### SetEmailExpressionNil

`func (o *TemplateSignerRoleInput) SetEmailExpressionNil(b bool)`

 SetEmailExpressionNil sets the value for EmailExpression to be an explicit nil

### UnsetEmailExpression
`func (o *TemplateSignerRoleInput) UnsetEmailExpression()`

UnsetEmailExpression ensures that no value is present for EmailExpression, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


