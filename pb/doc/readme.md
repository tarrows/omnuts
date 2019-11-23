# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [backlog.proto](#backlog.proto)
    - [Backlog](#omnuts.api.v1.Backlog)
    - [ListBacklogsRequest](#omnuts.api.v1.ListBacklogsRequest)
    - [ListBacklogsResponse](#omnuts.api.v1.ListBacklogsResponse)
  
    - [Backlog.BacklogType](#omnuts.api.v1.Backlog.BacklogType)
    - [Backlog.Status](#omnuts.api.v1.Backlog.Status)
  
  
    - [BacklogService](#omnuts.api.v1.BacklogService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="backlog.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## backlog.proto



<a name="omnuts.api.v1.Backlog"></a>

### Backlog



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| title | [string](#string) |  |  |
| type | [Backlog.BacklogType](#omnuts.api.v1.Backlog.BacklogType) |  |  |
| status | [Backlog.Status](#omnuts.api.v1.Backlog.Status) |  |  |
| vote | [int32](#int32) |  |  |
| created_on | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_on | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="omnuts.api.v1.ListBacklogsRequest"></a>

### ListBacklogsRequest







<a name="omnuts.api.v1.ListBacklogsResponse"></a>

### ListBacklogsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| backlogs | [Backlog](#omnuts.api.v1.Backlog) | repeated |  |





 


<a name="omnuts.api.v1.Backlog.BacklogType"></a>

### Backlog.BacklogType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PRODUCT | 0 |  |
| SPRINT | 1 |  |



<a name="omnuts.api.v1.Backlog.Status"></a>

### Backlog.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| NOT_STARTED | 0 |  |
| IN_PROGRESS | 1 |  |
| COMPLETED | 2 |  |


 

 


<a name="omnuts.api.v1.BacklogService"></a>

### BacklogService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListBacklogsRequest](#omnuts.api.v1.ListBacklogsRequest) | [ListBacklogsResponse](#omnuts.api.v1.ListBacklogsResponse) |  |

 



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

