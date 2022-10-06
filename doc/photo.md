# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [contracts/photo/v1/photo.proto](#contracts_photo_v1_photo-proto)
    - [ContentByHashRequest](#photo-v1-ContentByHashRequest)
    - [GetByHashRequest](#photo-v1-GetByHashRequest)
    - [GetByHashResponse](#photo-v1-GetByHashResponse)
    - [GetPhotosRequest](#photo-v1-GetPhotosRequest)
    - [GetPhotosResponse](#photo-v1-GetPhotosResponse)
    - [Photo](#photo-v1-Photo)
    - [PhotoServiceContentByHashResponse](#photo-v1-PhotoServiceContentByHashResponse)
  
    - [PhotoService](#photo-v1-PhotoService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="contracts_photo_v1_photo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## contracts/photo/v1/photo.proto



<a name="photo-v1-ContentByHashRequest"></a>

### ContentByHashRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  | a photo hash |






<a name="photo-v1-GetByHashRequest"></a>

### GetByHashRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  | a photo hash |






<a name="photo-v1-GetByHashResponse"></a>

### GetByHashResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| photo | [Photo](#photo-v1-Photo) |  |  |






<a name="photo-v1-GetPhotosRequest"></a>

### GetPhotosRequest







<a name="photo-v1-GetPhotosResponse"></a>

### GetPhotosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| photos | [Photo](#photo-v1-Photo) | repeated | the photos |






<a name="photo-v1-Photo"></a>

### Photo
Main


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  | the photo hash |
| path | [string](#string) |  | the photo path on the daemon fs |
| date_time | [string](#string) |  | the date |
| iso | [int32](#int32) |  | the ISO of the image |
| exposure_time | [string](#string) |  | the exposure as a fraction |
| x_dimension | [int32](#int32) |  | the x dimension of the image |
| y_dimension | [int32](#int32) |  | the y dimension of the image |
| model | [string](#string) |  | the model of the camera that took the image |
| f_number | [string](#string) |  | the f Number |






<a name="photo-v1-PhotoServiceContentByHashResponse"></a>

### PhotoServiceContentByHashResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  |  |





 

 

 


<a name="photo-v1-PhotoService"></a>

### PhotoService
The photo service definition

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPhotos | [GetPhotosRequest](#photo-v1-GetPhotosRequest) | [GetPhotosResponse](#photo-v1-GetPhotosResponse) | gets all photos |
| GetByHash | [GetByHashRequest](#photo-v1-GetByHashRequest) | [GetByHashResponse](#photo-v1-GetByHashResponse) | gets a photo by its hash |
| ContentByHash | [ContentByHashRequest](#photo-v1-ContentByHashRequest) | [PhotoServiceContentByHashResponse](#photo-v1-PhotoServiceContentByHashResponse) stream | gets the photo content by its hash |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

