# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [contracts/photo/v1/photo.proto](#contracts_photo_v1_photo-proto)
    - [ContentByHashRequest](#photo-v1-ContentByHashRequest)
    - [ExistsByHashRequest](#photo-v1-ExistsByHashRequest)
    - [ExistsByHashResponse](#photo-v1-ExistsByHashResponse)
    - [GetByHashRequest](#photo-v1-GetByHashRequest)
    - [GetByHashResponse](#photo-v1-GetByHashResponse)
    - [GetPhotosRequest](#photo-v1-GetPhotosRequest)
    - [GetPhotosResponse](#photo-v1-GetPhotosResponse)
    - [Photo](#photo-v1-Photo)
    - [PhotoServiceContentByHashResponse](#photo-v1-PhotoServiceContentByHashResponse)
    - [PhotoServiceThumbnailByHashResponse](#photo-v1-PhotoServiceThumbnailByHashResponse)
    - [ThumbnailByHashRequest](#photo-v1-ThumbnailByHashRequest)
  
    - [PhotoService](#photo-v1-PhotoService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="contracts_photo_v1_photo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## contracts/photo/v1/photo.proto
Photo related messages and service


<a name="photo-v1-ContentByHashRequest"></a>

### ContentByHashRequest
ContentByHashRequest represent a search query with a photo hash


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  | a photo hash |






<a name="photo-v1-ExistsByHashRequest"></a>

### ExistsByHashRequest
ExistsByHashRequest represent an existence search query with a photo hash


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  | a photo hash |






<a name="photo-v1-ExistsByHashResponse"></a>

### ExistsByHashResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exists | [bool](#bool) |  |  |






<a name="photo-v1-GetByHashRequest"></a>

### GetByHashRequest
GetByHashRequest represent a search query with a photo hash


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
GetPhotosRequest represent a search query with pagination options to
indicate which results to include in the response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [uint32](#uint32) |  | the number of element to retrieve |
| offset | [uint32](#uint32) |  | the position of the first element to retrieve |






<a name="photo-v1-GetPhotosResponse"></a>

### GetPhotosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| photos | [Photo](#photo-v1-Photo) | repeated | the photos |
| total | [uint32](#uint32) |  | the total number of photos |






<a name="photo-v1-Photo"></a>

### Photo
Photo contains basic info of a photo and its metadata

Main


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  | the photo hash |
| path | [string](#string) |  | the photo path on the daemon fs |
| date_time | [string](#string) |  | the date |
| iso | [uint32](#uint32) |  | the ISO of the image |
| exposure_time | [string](#string) |  | the exposure as a fraction |
| x_dimension | [uint32](#uint32) |  | the x dimension of the image |
| y_dimension | [uint32](#uint32) |  | the y dimension of the image |
| model | [string](#string) |  | the model of the camera that took the image |
| f_number | [string](#string) |  | the f Number |






<a name="photo-v1-PhotoServiceContentByHashResponse"></a>

### PhotoServiceContentByHashResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  |  |
| content_type | [string](#string) |  |  |






<a name="photo-v1-PhotoServiceThumbnailByHashResponse"></a>

### PhotoServiceThumbnailByHashResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  |  |
| content_type | [string](#string) |  |  |






<a name="photo-v1-ThumbnailByHashRequest"></a>

### ThumbnailByHashRequest
ThumbnailByHashRequest represent a search query with a photo hash, a width or an height.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  | a photo hash |
| width | [uint32](#uint32) |  | the requested width for the thumbnail |
| height | [uint32](#uint32) |  | the requested height for the thumbnail |





 

 

 


<a name="photo-v1-PhotoService"></a>

### PhotoService
The photo service definition

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPhotos | [GetPhotosRequest](#photo-v1-GetPhotosRequest) | [GetPhotosResponse](#photo-v1-GetPhotosResponse) | GetPhotos returns all the photos |
| GetByHash | [GetByHashRequest](#photo-v1-GetByHashRequest) | [GetByHashResponse](#photo-v1-GetByHashResponse) | GetByHash returns a photo by its hash |
| ExistsByHash | [ExistsByHashRequest](#photo-v1-ExistsByHashRequest) | [ExistsByHashResponse](#photo-v1-ExistsByHashResponse) | ExistsByHash returns true if a photo with the given hash exists |
| ContentByHash | [ContentByHashRequest](#photo-v1-ContentByHashRequest) | [PhotoServiceContentByHashResponse](#photo-v1-PhotoServiceContentByHashResponse) stream | ContentByHash returns the photo content by its hash |
| ThumbnailByHash | [ThumbnailByHashRequest](#photo-v1-ThumbnailByHashRequest) | [PhotoServiceThumbnailByHashResponse](#photo-v1-PhotoServiceThumbnailByHashResponse) stream | ThumbnailByHash returns the photo thumbnail by its hash A width or a height can be provided, if so the thumbnail will have one of the property. If neither the width nor the height is specified, a width of 200 will be used. If a width and a height is specified, only the width will be used. |

 



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

