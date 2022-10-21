# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [contracts/daemon/v1/daemon.proto](#contracts_daemon_v1_daemon-proto)
    - [HeartBeatRequest](#daemon-v1-HeartBeatRequest)
    - [HeartBeatResponse](#daemon-v1-HeartBeatResponse)
    - [RegisterRequest](#daemon-v1-RegisterRequest)
    - [RegisterResponse](#daemon-v1-RegisterResponse)
  
    - [DaemonService](#daemon-v1-DaemonService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="contracts_daemon_v1_daemon-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## contracts/daemon/v1/daemon.proto



<a name="daemon-v1-HeartBeatRequest"></a>

### HeartBeatRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | The id used to identify the daemon |






<a name="daemon-v1-HeartBeatResponse"></a>

### HeartBeatResponse







<a name="daemon-v1-RegisterRequest"></a>

### RegisterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| daemon_version | [string](#string) |  | The version of the daemon |
| daemon_host | [string](#string) |  | The host (IP, hostname) to use to contact the daemon |
| daemon_port | [int32](#int32) |  | The port to use to contact the daemon |
| daemon_name | [string](#string) |  | The name of the daemon |
| owner | [string](#string) |  | the email of the owner |






<a name="daemon-v1-RegisterResponse"></a>

### RegisterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | The id used to identify the daemon |
| exp_in | [int32](#int32) |  | The duration in seconds that the daemon id is valid |





 

 

 


<a name="daemon-v1-DaemonService"></a>

### DaemonService
The daemon service definition

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterRequest](#daemon-v1-RegisterRequest) | [RegisterResponse](#daemon-v1-RegisterResponse) | Register registers a new daemon |
| HeartBeat | [HeartBeatRequest](#daemon-v1-HeartBeatRequest) | [HeartBeatResponse](#daemon-v1-HeartBeatResponse) | HeartBeat notifies the web app that a daemon is alive |

 



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

