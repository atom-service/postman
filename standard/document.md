# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [standard.proto](#standard.proto)
    - [CheckVerifyCodeRequest](#standard.CheckVerifyCodeRequest)
    - [CheckVerifyCodeeResponse](#standard.CheckVerifyCodeeResponse)
    - [DestroyVerifyCodeByKeyRequest](#standard.DestroyVerifyCodeByKeyRequest)
    - [DestroyVerifyCodeByKeyResponse](#standard.DestroyVerifyCodeByKeyResponse)
    - [SendVerifyCodeByCallPhoneRequest](#standard.SendVerifyCodeByCallPhoneRequest)
    - [SendVerifyCodeByCallPhoneResponse](#standard.SendVerifyCodeByCallPhoneResponse)
    - [SendVerifyCodeByEmailRequest](#standard.SendVerifyCodeByEmailRequest)
    - [SendVerifyCodeByEmailResponse](#standard.SendVerifyCodeByEmailResponse)
    - [SendVerifyCodeBySmsRequest](#standard.SendVerifyCodeBySmsRequest)
    - [SendVerifyCodeBySmsResponse](#standard.SendVerifyCodeBySmsResponse)
    - [VerifyCode](#standard.VerifyCode)
  
    - [State](#standard.State)
  
  
    - [Sender](#standard.Sender)
  

- [Scalar Value Types](#scalar-value-types)



<a name="standard.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## standard.proto



<a name="standard.CheckVerifyCodeRequest"></a>

### CheckVerifyCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [string](#string) |  | 唯一标识符 |
| VerifyCode | [string](#string) |  | 验证码 |






<a name="standard.CheckVerifyCodeeResponse"></a>

### CheckVerifyCodeeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.DestroyVerifyCodeByKeyRequest"></a>

### DestroyVerifyCodeByKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [string](#string) |  | 唯一标识符 |






<a name="standard.DestroyVerifyCodeByKeyResponse"></a>

### DestroyVerifyCodeByKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.SendVerifyCodeByCallPhoneRequest"></a>

### SendVerifyCodeByCallPhoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Operation | [string](#string) |  | 操作 |
| PhoneNumber | [string](#string) |  | 手机号码 |
| CountryCode | [string](#string) |  | 国家代码 |
| ValidityPeriod | [int64](#int64) |  | 有效期 单位s |






<a name="standard.SendVerifyCodeByCallPhoneResponse"></a>

### SendVerifyCodeByCallPhoneResponse
发送成功失败状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Key | [string](#string) |  | 唯一标识符 |






<a name="standard.SendVerifyCodeByEmailRequest"></a>

### SendVerifyCodeByEmailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Operation | [string](#string) |  | 操作 |
| EmailAddress | [string](#string) |  | 邮件地址 |
| ValidityPeriod | [int64](#int64) |  | 有效期 单位s |






<a name="standard.SendVerifyCodeByEmailResponse"></a>

### SendVerifyCodeByEmailResponse
发送成功失败状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Key | [string](#string) |  | 唯一标识符 |






<a name="standard.SendVerifyCodeBySmsRequest"></a>

### SendVerifyCodeBySmsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Operation | [string](#string) |  | 操作 |
| PhoneNumber | [string](#string) |  | 手机号码 |
| CountryCode | [string](#string) |  | 国家代码 |
| ValidityPeriod | [int64](#int64) |  | 有效期 单位s |






<a name="standard.SendVerifyCodeBySmsResponse"></a>

### SendVerifyCodeBySmsResponse
发送成功失败状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Key | [string](#string) |  | 唯一标识符 |






<a name="standard.VerifyCode"></a>

### VerifyCode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [string](#string) |  | 数据 |
| Code | [string](#string) |  | 失效时间 |
| Operation | [string](#string) |  | 创建时间 |
| ExpireTime | [string](#string) |  | 销毁时间 |
| DeletedTime | [string](#string) |  |  |
| CreatedTime | [string](#string) |  |  |
| UpdatedTime | [string](#string) |  | 生效时间 |





 


<a name="standard.State"></a>

### State
状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | 未知 |
| SUCCESS | 1 | 成功 |
| FAILURE | 2 | 失败 |
| NOT_EXIST | 3 | 不存在 |
| UNDEFINED | 4 | 未定义 |
| PROHIBITED | 5 | 被禁止 |
| NO_SERVICE_AVAILABLE | 6 | 无可用服务 |
| DB_OPERATION_FATLURE | 7 | 数据库操作失败 |


 

 


<a name="standard.Sender"></a>

### Sender


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CheckVerifyCode | [CheckVerifyCodeRequest](#standard.CheckVerifyCodeRequest) | [CheckVerifyCodeeResponse](#standard.CheckVerifyCodeeResponse) |  |
| SendVerifyCodeBySms | [SendVerifyCodeBySmsRequest](#standard.SendVerifyCodeBySmsRequest) | [SendVerifyCodeBySmsResponse](#standard.SendVerifyCodeBySmsResponse) |  |
| DestroyVerifyCodeByKey | [DestroyVerifyCodeByKeyRequest](#standard.DestroyVerifyCodeByKeyRequest) | [DestroyVerifyCodeByKeyResponse](#standard.DestroyVerifyCodeByKeyResponse) |  |
| SendVerifyCodeByEmail | [SendVerifyCodeByEmailRequest](#standard.SendVerifyCodeByEmailRequest) | [SendVerifyCodeByEmailResponse](#standard.SendVerifyCodeByEmailResponse) |  |
| SendVerifyCodeByCallPhone | [SendVerifyCodeByCallPhoneRequest](#standard.SendVerifyCodeByCallPhoneRequest) | [SendVerifyCodeByCallPhoneResponse](#standard.SendVerifyCodeByCallPhoneResponse) |  |

 



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

