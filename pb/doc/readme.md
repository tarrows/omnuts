# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [minutes.proto](#minutes.proto)
    - [Group](#omnutspb.Group)
    - [Minutes](#omnutspb.Minutes)
    - [Minutes.Agendum](#omnutspb.Minutes.Agendum)
    - [Person](#omnutspb.Person)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="minutes.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## minutes.proto



<a name="omnutspb.Group"></a>

### Group



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| member | [Person](#omnutspb.Person) | repeated |  |






<a name="omnutspb.Minutes"></a>

### Minutes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| date | [google.type.Date](#google.type.Date) |  |  |
| attendees | [Person](#omnutspb.Person) | repeated |  |
| agenda | [Minutes.Agendum](#omnutspb.Minutes.Agendum) | repeated |  |
| tags | [string](#string) | repeated |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="omnutspb.Minutes.Agendum"></a>

### Minutes.Agendum



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| no | [int32](#int32) |  |  |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |






<a name="omnutspb.Person"></a>

### Person



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

