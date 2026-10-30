# TemplatePageProcessingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | [**TemplatePageImageStatus**](TemplatePageImageStatus.md) | The page's image, which the template editor shows. `pending` until it's rendered. | 
**Fields** | [**TemplatePageFieldsStatus**](TemplatePageFieldsStatus.md) | Setting up the page's fields: `pending` while they're imported or detected, `naming` while AI names them, then `processed` or `failed`. `disabled` when DocSpring doesn't set up the page's fields. | 
**AiNaming** | [**NullableTemplatePageAiNamingStatus**](TemplatePageAiNamingStatus.md) | `failed` or `skipped` when AI didn't name the page's fields as asked (`ai_error` says why), and the fields kept plain names. Null otherwise. | 
**AiError** | [**NullableTemplatePageAiError**](TemplatePageAiError.md) | Why AI naming failed or was skipped: `unavailable` (the AI service was unavailable: set up the fields again later), `error`, `credits` (the account had no free AI pages or credits left) or `skipped` (someone skipped AI setup in the template editor). Null when `ai_naming` is. | 
**DetectedBoxes** | **[][]float32** | While AI names the page's fields, the boxes field detection found: `[x, y, width, height]` in PDF points from the page's top-left corner. Null otherwise. | 

## Methods

### NewTemplatePageProcessingStatus

`func NewTemplatePageProcessingStatus(image TemplatePageImageStatus, fields TemplatePageFieldsStatus, aiNaming NullableTemplatePageAiNamingStatus, aiError NullableTemplatePageAiError, detectedBoxes [][]float32, ) *TemplatePageProcessingStatus`

NewTemplatePageProcessingStatus instantiates a new TemplatePageProcessingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePageProcessingStatusWithDefaults

`func NewTemplatePageProcessingStatusWithDefaults() *TemplatePageProcessingStatus`

NewTemplatePageProcessingStatusWithDefaults instantiates a new TemplatePageProcessingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *TemplatePageProcessingStatus) GetImage() TemplatePageImageStatus`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *TemplatePageProcessingStatus) GetImageOk() (*TemplatePageImageStatus, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *TemplatePageProcessingStatus) SetImage(v TemplatePageImageStatus)`

SetImage sets Image field to given value.


### GetFields

`func (o *TemplatePageProcessingStatus) GetFields() TemplatePageFieldsStatus`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TemplatePageProcessingStatus) GetFieldsOk() (*TemplatePageFieldsStatus, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TemplatePageProcessingStatus) SetFields(v TemplatePageFieldsStatus)`

SetFields sets Fields field to given value.


### GetAiNaming

`func (o *TemplatePageProcessingStatus) GetAiNaming() TemplatePageAiNamingStatus`

GetAiNaming returns the AiNaming field if non-nil, zero value otherwise.

### GetAiNamingOk

`func (o *TemplatePageProcessingStatus) GetAiNamingOk() (*TemplatePageAiNamingStatus, bool)`

GetAiNamingOk returns a tuple with the AiNaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiNaming

`func (o *TemplatePageProcessingStatus) SetAiNaming(v TemplatePageAiNamingStatus)`

SetAiNaming sets AiNaming field to given value.


### SetAiNamingNil

`func (o *TemplatePageProcessingStatus) SetAiNamingNil(b bool)`

 SetAiNamingNil sets the value for AiNaming to be an explicit nil

### UnsetAiNaming
`func (o *TemplatePageProcessingStatus) UnsetAiNaming()`

UnsetAiNaming ensures that no value is present for AiNaming, not even an explicit nil
### GetAiError

`func (o *TemplatePageProcessingStatus) GetAiError() TemplatePageAiError`

GetAiError returns the AiError field if non-nil, zero value otherwise.

### GetAiErrorOk

`func (o *TemplatePageProcessingStatus) GetAiErrorOk() (*TemplatePageAiError, bool)`

GetAiErrorOk returns a tuple with the AiError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiError

`func (o *TemplatePageProcessingStatus) SetAiError(v TemplatePageAiError)`

SetAiError sets AiError field to given value.


### SetAiErrorNil

`func (o *TemplatePageProcessingStatus) SetAiErrorNil(b bool)`

 SetAiErrorNil sets the value for AiError to be an explicit nil

### UnsetAiError
`func (o *TemplatePageProcessingStatus) UnsetAiError()`

UnsetAiError ensures that no value is present for AiError, not even an explicit nil
### GetDetectedBoxes

`func (o *TemplatePageProcessingStatus) GetDetectedBoxes() [][]float32`

GetDetectedBoxes returns the DetectedBoxes field if non-nil, zero value otherwise.

### GetDetectedBoxesOk

`func (o *TemplatePageProcessingStatus) GetDetectedBoxesOk() (*[][]float32, bool)`

GetDetectedBoxesOk returns a tuple with the DetectedBoxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedBoxes

`func (o *TemplatePageProcessingStatus) SetDetectedBoxes(v [][]float32)`

SetDetectedBoxes sets DetectedBoxes field to given value.


### SetDetectedBoxesNil

`func (o *TemplatePageProcessingStatus) SetDetectedBoxesNil(b bool)`

 SetDetectedBoxesNil sets the value for DetectedBoxes to be an explicit nil

### UnsetDetectedBoxes
`func (o *TemplatePageProcessingStatus) UnsetDetectedBoxes()`

UnsetDetectedBoxes ensures that no value is present for DetectedBoxes, not even an explicit nil

[[Back to Model list]](../README.md#methods) [[Back to API list]](../README.md#methods) [[Back to README]](../README.md)


