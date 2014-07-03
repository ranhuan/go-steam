// Code generated by protoc-gen-go.
// source: steammessages_cloud.steamclient.proto
// DO NOT EDIT!

package unified

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// discarding unused import google_protobuf "code.google.com/p/goprotobuf/protoc-gen-go/descriptor"
// discarding unused import steammessages_unified_base_steamclient "steammessages_unified_base.steamclient.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type CCloud_GetUploadServerInfo_Request struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_GetUploadServerInfo_Request) Reset()         { *m = CCloud_GetUploadServerInfo_Request{} }
func (m *CCloud_GetUploadServerInfo_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetUploadServerInfo_Request) ProtoMessage()    {}

func (m *CCloud_GetUploadServerInfo_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_GetUploadServerInfo_Response struct {
	ServerUrl        *string `protobuf:"bytes,1,opt,name=server_url" json:"server_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_GetUploadServerInfo_Response) Reset()         { *m = CCloud_GetUploadServerInfo_Response{} }
func (m *CCloud_GetUploadServerInfo_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetUploadServerInfo_Response) ProtoMessage()    {}

func (m *CCloud_GetUploadServerInfo_Response) GetServerUrl() string {
	if m != nil && m.ServerUrl != nil {
		return *m.ServerUrl
	}
	return ""
}

type CCloud_BeginHTTPUpload_Request struct {
	Appid                *uint32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	FileSize             *uint32  `protobuf:"varint,2,opt,name=file_size" json:"file_size,omitempty"`
	Filename             *string  `protobuf:"bytes,3,opt,name=filename" json:"filename,omitempty"`
	FileSha              *string  `protobuf:"bytes,4,opt,name=file_sha" json:"file_sha,omitempty"`
	IsPublic             *bool    `protobuf:"varint,5,opt,name=is_public" json:"is_public,omitempty"`
	PlatformsToSync      []string `protobuf:"bytes,6,rep,name=platforms_to_sync" json:"platforms_to_sync,omitempty"`
	RequestHeadersNames  []string `protobuf:"bytes,7,rep,name=request_headers_names" json:"request_headers_names,omitempty"`
	RequestHeadersValues []string `protobuf:"bytes,8,rep,name=request_headers_values" json:"request_headers_values,omitempty"`
	XXX_unrecognized     []byte   `json:"-"`
}

func (m *CCloud_BeginHTTPUpload_Request) Reset()         { *m = CCloud_BeginHTTPUpload_Request{} }
func (m *CCloud_BeginHTTPUpload_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_BeginHTTPUpload_Request) ProtoMessage()    {}

func (m *CCloud_BeginHTTPUpload_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_BeginHTTPUpload_Request) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *CCloud_BeginHTTPUpload_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CCloud_BeginHTTPUpload_Request) GetFileSha() string {
	if m != nil && m.FileSha != nil {
		return *m.FileSha
	}
	return ""
}

func (m *CCloud_BeginHTTPUpload_Request) GetIsPublic() bool {
	if m != nil && m.IsPublic != nil {
		return *m.IsPublic
	}
	return false
}

func (m *CCloud_BeginHTTPUpload_Request) GetPlatformsToSync() []string {
	if m != nil {
		return m.PlatformsToSync
	}
	return nil
}

func (m *CCloud_BeginHTTPUpload_Request) GetRequestHeadersNames() []string {
	if m != nil {
		return m.RequestHeadersNames
	}
	return nil
}

func (m *CCloud_BeginHTTPUpload_Request) GetRequestHeadersValues() []string {
	if m != nil {
		return m.RequestHeadersValues
	}
	return nil
}

type CCloud_BeginHTTPUpload_Response struct {
	Ugcid            *uint64                                        `protobuf:"fixed64,1,opt,name=ugcid" json:"ugcid,omitempty"`
	Timestamp        *uint32                                        `protobuf:"fixed32,2,opt,name=timestamp" json:"timestamp,omitempty"`
	UrlHost          *string                                        `protobuf:"bytes,3,opt,name=url_host" json:"url_host,omitempty"`
	UrlPath          *string                                        `protobuf:"bytes,4,opt,name=url_path" json:"url_path,omitempty"`
	UseHttps         *bool                                          `protobuf:"varint,5,opt,name=use_https" json:"use_https,omitempty"`
	RequestHeaders   []*CCloud_BeginHTTPUpload_Response_HTTPHeaders `protobuf:"bytes,6,rep,name=request_headers" json:"request_headers,omitempty"`
	XXX_unrecognized []byte                                         `json:"-"`
}

func (m *CCloud_BeginHTTPUpload_Response) Reset()         { *m = CCloud_BeginHTTPUpload_Response{} }
func (m *CCloud_BeginHTTPUpload_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_BeginHTTPUpload_Response) ProtoMessage()    {}

func (m *CCloud_BeginHTTPUpload_Response) GetUgcid() uint64 {
	if m != nil && m.Ugcid != nil {
		return *m.Ugcid
	}
	return 0
}

func (m *CCloud_BeginHTTPUpload_Response) GetTimestamp() uint32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *CCloud_BeginHTTPUpload_Response) GetUrlHost() string {
	if m != nil && m.UrlHost != nil {
		return *m.UrlHost
	}
	return ""
}

func (m *CCloud_BeginHTTPUpload_Response) GetUrlPath() string {
	if m != nil && m.UrlPath != nil {
		return *m.UrlPath
	}
	return ""
}

func (m *CCloud_BeginHTTPUpload_Response) GetUseHttps() bool {
	if m != nil && m.UseHttps != nil {
		return *m.UseHttps
	}
	return false
}

func (m *CCloud_BeginHTTPUpload_Response) GetRequestHeaders() []*CCloud_BeginHTTPUpload_Response_HTTPHeaders {
	if m != nil {
		return m.RequestHeaders
	}
	return nil
}

type CCloud_BeginHTTPUpload_Response_HTTPHeaders struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_BeginHTTPUpload_Response_HTTPHeaders) Reset() {
	*m = CCloud_BeginHTTPUpload_Response_HTTPHeaders{}
}
func (m *CCloud_BeginHTTPUpload_Response_HTTPHeaders) String() string {
	return proto.CompactTextString(m)
}
func (*CCloud_BeginHTTPUpload_Response_HTTPHeaders) ProtoMessage() {}

func (m *CCloud_BeginHTTPUpload_Response_HTTPHeaders) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CCloud_BeginHTTPUpload_Response_HTTPHeaders) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type CCloud_CommitHTTPUpload_Request struct {
	TransferSucceeded *bool   `protobuf:"varint,1,opt,name=transfer_succeeded" json:"transfer_succeeded,omitempty"`
	Appid             *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	FileSha           *string `protobuf:"bytes,3,opt,name=file_sha" json:"file_sha,omitempty"`
	Filename          *string `protobuf:"bytes,4,opt,name=filename" json:"filename,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *CCloud_CommitHTTPUpload_Request) Reset()         { *m = CCloud_CommitHTTPUpload_Request{} }
func (m *CCloud_CommitHTTPUpload_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_CommitHTTPUpload_Request) ProtoMessage()    {}

func (m *CCloud_CommitHTTPUpload_Request) GetTransferSucceeded() bool {
	if m != nil && m.TransferSucceeded != nil {
		return *m.TransferSucceeded
	}
	return false
}

func (m *CCloud_CommitHTTPUpload_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_CommitHTTPUpload_Request) GetFileSha() string {
	if m != nil && m.FileSha != nil {
		return *m.FileSha
	}
	return ""
}

func (m *CCloud_CommitHTTPUpload_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

type CCloud_CommitHTTPUpload_Response struct {
	FileCommitted    *bool  `protobuf:"varint,1,opt,name=file_committed" json:"file_committed,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CCloud_CommitHTTPUpload_Response) Reset()         { *m = CCloud_CommitHTTPUpload_Response{} }
func (m *CCloud_CommitHTTPUpload_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_CommitHTTPUpload_Response) ProtoMessage()    {}

func (m *CCloud_CommitHTTPUpload_Response) GetFileCommitted() bool {
	if m != nil && m.FileCommitted != nil {
		return *m.FileCommitted
	}
	return false
}

type CCloud_GetFileDetails_Request struct {
	Ugcid            *uint64 `protobuf:"varint,1,opt,name=ugcid" json:"ugcid,omitempty"`
	Appid            *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_GetFileDetails_Request) Reset()         { *m = CCloud_GetFileDetails_Request{} }
func (m *CCloud_GetFileDetails_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetFileDetails_Request) ProtoMessage()    {}

func (m *CCloud_GetFileDetails_Request) GetUgcid() uint64 {
	if m != nil && m.Ugcid != nil {
		return *m.Ugcid
	}
	return 0
}

func (m *CCloud_GetFileDetails_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_UserFile struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Ugcid            *uint64 `protobuf:"varint,2,opt,name=ugcid" json:"ugcid,omitempty"`
	Filename         *string `protobuf:"bytes,3,opt,name=filename" json:"filename,omitempty"`
	Timestamp        *uint64 `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	FileSize         *uint32 `protobuf:"varint,5,opt,name=file_size" json:"file_size,omitempty"`
	Url              *string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	SteamidCreator   *uint64 `protobuf:"fixed64,7,opt,name=steamid_creator" json:"steamid_creator,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_UserFile) Reset()         { *m = CCloud_UserFile{} }
func (m *CCloud_UserFile) String() string { return proto.CompactTextString(m) }
func (*CCloud_UserFile) ProtoMessage()    {}

func (m *CCloud_UserFile) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_UserFile) GetUgcid() uint64 {
	if m != nil && m.Ugcid != nil {
		return *m.Ugcid
	}
	return 0
}

func (m *CCloud_UserFile) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CCloud_UserFile) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *CCloud_UserFile) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *CCloud_UserFile) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *CCloud_UserFile) GetSteamidCreator() uint64 {
	if m != nil && m.SteamidCreator != nil {
		return *m.SteamidCreator
	}
	return 0
}

type CCloud_GetFileDetails_Response struct {
	Details          *CCloud_UserFile `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CCloud_GetFileDetails_Response) Reset()         { *m = CCloud_GetFileDetails_Response{} }
func (m *CCloud_GetFileDetails_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetFileDetails_Response) ProtoMessage()    {}

func (m *CCloud_GetFileDetails_Response) GetDetails() *CCloud_UserFile {
	if m != nil {
		return m.Details
	}
	return nil
}

type CCloud_EnumerateUserFiles_Request struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	ExtendedDetails  *bool   `protobuf:"varint,2,opt,name=extended_details" json:"extended_details,omitempty"`
	Count            *uint32 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	StartIndex       *uint32 `protobuf:"varint,4,opt,name=start_index" json:"start_index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_EnumerateUserFiles_Request) Reset()         { *m = CCloud_EnumerateUserFiles_Request{} }
func (m *CCloud_EnumerateUserFiles_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_EnumerateUserFiles_Request) ProtoMessage()    {}

func (m *CCloud_EnumerateUserFiles_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_EnumerateUserFiles_Request) GetExtendedDetails() bool {
	if m != nil && m.ExtendedDetails != nil {
		return *m.ExtendedDetails
	}
	return false
}

func (m *CCloud_EnumerateUserFiles_Request) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CCloud_EnumerateUserFiles_Request) GetStartIndex() uint32 {
	if m != nil && m.StartIndex != nil {
		return *m.StartIndex
	}
	return 0
}

type CCloud_EnumerateUserFiles_Response struct {
	Files            []*CCloud_UserFile `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
	TotalFiles       *uint32            `protobuf:"varint,2,opt,name=total_files" json:"total_files,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *CCloud_EnumerateUserFiles_Response) Reset()         { *m = CCloud_EnumerateUserFiles_Response{} }
func (m *CCloud_EnumerateUserFiles_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_EnumerateUserFiles_Response) ProtoMessage()    {}

func (m *CCloud_EnumerateUserFiles_Response) GetFiles() []*CCloud_UserFile {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *CCloud_EnumerateUserFiles_Response) GetTotalFiles() uint32 {
	if m != nil && m.TotalFiles != nil {
		return *m.TotalFiles
	}
	return 0
}

type CCloud_Delete_Request struct {
	Filename         *string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Appid            *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_Delete_Request) Reset()         { *m = CCloud_Delete_Request{} }
func (m *CCloud_Delete_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_Delete_Request) ProtoMessage()    {}

func (m *CCloud_Delete_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CCloud_Delete_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_Delete_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CCloud_Delete_Response) Reset()         { *m = CCloud_Delete_Response{} }
func (m *CCloud_Delete_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_Delete_Response) ProtoMessage()    {}

type CCloud_GetClientEncryptionKey_Request struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CCloud_GetClientEncryptionKey_Request) Reset()         { *m = CCloud_GetClientEncryptionKey_Request{} }
func (m *CCloud_GetClientEncryptionKey_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetClientEncryptionKey_Request) ProtoMessage()    {}

type CCloud_GetClientEncryptionKey_Response struct {
	Key              []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Crc              *int32 `protobuf:"varint,2,opt,name=crc" json:"crc,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CCloud_GetClientEncryptionKey_Response) Reset() {
	*m = CCloud_GetClientEncryptionKey_Response{}
}
func (m *CCloud_GetClientEncryptionKey_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetClientEncryptionKey_Response) ProtoMessage()    {}

func (m *CCloud_GetClientEncryptionKey_Response) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CCloud_GetClientEncryptionKey_Response) GetCrc() int32 {
	if m != nil && m.Crc != nil {
		return *m.Crc
	}
	return 0
}

type CCloud_CDNReport_Notification struct {
	Steamid          *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	Url              *string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Success          *bool   `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
	HttpStatusCode   *uint32 `protobuf:"varint,4,opt,name=http_status_code" json:"http_status_code,omitempty"`
	ExpectedBytes    *uint64 `protobuf:"varint,5,opt,name=expected_bytes" json:"expected_bytes,omitempty"`
	ReceivedBytes    *uint64 `protobuf:"varint,6,opt,name=received_bytes" json:"received_bytes,omitempty"`
	Duration         *uint32 `protobuf:"varint,7,opt,name=duration" json:"duration,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_CDNReport_Notification) Reset()         { *m = CCloud_CDNReport_Notification{} }
func (m *CCloud_CDNReport_Notification) String() string { return proto.CompactTextString(m) }
func (*CCloud_CDNReport_Notification) ProtoMessage()    {}

func (m *CCloud_CDNReport_Notification) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CCloud_CDNReport_Notification) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *CCloud_CDNReport_Notification) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *CCloud_CDNReport_Notification) GetHttpStatusCode() uint32 {
	if m != nil && m.HttpStatusCode != nil {
		return *m.HttpStatusCode
	}
	return 0
}

func (m *CCloud_CDNReport_Notification) GetExpectedBytes() uint64 {
	if m != nil && m.ExpectedBytes != nil {
		return *m.ExpectedBytes
	}
	return 0
}

func (m *CCloud_CDNReport_Notification) GetReceivedBytes() uint64 {
	if m != nil && m.ReceivedBytes != nil {
		return *m.ReceivedBytes
	}
	return 0
}

func (m *CCloud_CDNReport_Notification) GetDuration() uint32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

type CCloud_ExternalStorageTransferReport_Notification struct {
	Host             *string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Path             *string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	IsUpload         *bool   `protobuf:"varint,3,opt,name=is_upload" json:"is_upload,omitempty"`
	Success          *bool   `protobuf:"varint,4,opt,name=success" json:"success,omitempty"`
	HttpStatusCode   *uint32 `protobuf:"varint,5,opt,name=http_status_code" json:"http_status_code,omitempty"`
	BytesExpected    *uint64 `protobuf:"varint,6,opt,name=bytes_expected" json:"bytes_expected,omitempty"`
	BytesActual      *uint64 `protobuf:"varint,7,opt,name=bytes_actual" json:"bytes_actual,omitempty"`
	DurationMs       *uint32 `protobuf:"varint,8,opt,name=duration_ms" json:"duration_ms,omitempty"`
	Cellid           *uint32 `protobuf:"varint,9,opt,name=cellid" json:"cellid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_ExternalStorageTransferReport_Notification) Reset() {
	*m = CCloud_ExternalStorageTransferReport_Notification{}
}
func (m *CCloud_ExternalStorageTransferReport_Notification) String() string {
	return proto.CompactTextString(m)
}
func (*CCloud_ExternalStorageTransferReport_Notification) ProtoMessage() {}

func (m *CCloud_ExternalStorageTransferReport_Notification) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *CCloud_ExternalStorageTransferReport_Notification) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *CCloud_ExternalStorageTransferReport_Notification) GetIsUpload() bool {
	if m != nil && m.IsUpload != nil {
		return *m.IsUpload
	}
	return false
}

func (m *CCloud_ExternalStorageTransferReport_Notification) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *CCloud_ExternalStorageTransferReport_Notification) GetHttpStatusCode() uint32 {
	if m != nil && m.HttpStatusCode != nil {
		return *m.HttpStatusCode
	}
	return 0
}

func (m *CCloud_ExternalStorageTransferReport_Notification) GetBytesExpected() uint64 {
	if m != nil && m.BytesExpected != nil {
		return *m.BytesExpected
	}
	return 0
}

func (m *CCloud_ExternalStorageTransferReport_Notification) GetBytesActual() uint64 {
	if m != nil && m.BytesActual != nil {
		return *m.BytesActual
	}
	return 0
}

func (m *CCloud_ExternalStorageTransferReport_Notification) GetDurationMs() uint32 {
	if m != nil && m.DurationMs != nil {
		return *m.DurationMs
	}
	return 0
}

func (m *CCloud_ExternalStorageTransferReport_Notification) GetCellid() uint32 {
	if m != nil && m.Cellid != nil {
		return *m.Cellid
	}
	return 0
}

type CCloud_ClientBeginFileUpload_Request struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	FileSize         *uint32 `protobuf:"varint,2,opt,name=file_size" json:"file_size,omitempty"`
	RawFileSize      *uint32 `protobuf:"varint,3,opt,name=raw_file_size" json:"raw_file_size,omitempty"`
	FileSha          []byte  `protobuf:"bytes,4,opt,name=file_sha" json:"file_sha,omitempty"`
	TimeStamp        *uint64 `protobuf:"varint,5,opt,name=time_stamp" json:"time_stamp,omitempty"`
	Filename         *string `protobuf:"bytes,6,opt,name=filename" json:"filename,omitempty"`
	PlatformsToSync  *uint32 `protobuf:"varint,7,opt,name=platforms_to_sync,def=4294967295" json:"platforms_to_sync,omitempty"`
	CellId           *uint32 `protobuf:"varint,9,opt,name=cell_id" json:"cell_id,omitempty"`
	CanEncrypt       *bool   `protobuf:"varint,10,opt,name=can_encrypt" json:"can_encrypt,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_ClientBeginFileUpload_Request) Reset()         { *m = CCloud_ClientBeginFileUpload_Request{} }
func (m *CCloud_ClientBeginFileUpload_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_ClientBeginFileUpload_Request) ProtoMessage()    {}

const Default_CCloud_ClientBeginFileUpload_Request_PlatformsToSync uint32 = 4294967295

func (m *CCloud_ClientBeginFileUpload_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_ClientBeginFileUpload_Request) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *CCloud_ClientBeginFileUpload_Request) GetRawFileSize() uint32 {
	if m != nil && m.RawFileSize != nil {
		return *m.RawFileSize
	}
	return 0
}

func (m *CCloud_ClientBeginFileUpload_Request) GetFileSha() []byte {
	if m != nil {
		return m.FileSha
	}
	return nil
}

func (m *CCloud_ClientBeginFileUpload_Request) GetTimeStamp() uint64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *CCloud_ClientBeginFileUpload_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CCloud_ClientBeginFileUpload_Request) GetPlatformsToSync() uint32 {
	if m != nil && m.PlatformsToSync != nil {
		return *m.PlatformsToSync
	}
	return Default_CCloud_ClientBeginFileUpload_Request_PlatformsToSync
}

func (m *CCloud_ClientBeginFileUpload_Request) GetCellId() uint32 {
	if m != nil && m.CellId != nil {
		return *m.CellId
	}
	return 0
}

func (m *CCloud_ClientBeginFileUpload_Request) GetCanEncrypt() bool {
	if m != nil && m.CanEncrypt != nil {
		return *m.CanEncrypt
	}
	return false
}

type ClientCloudFileUploadBlockDetails struct {
	UrlHost          *string                                          `protobuf:"bytes,1,opt,name=url_host" json:"url_host,omitempty"`
	UrlPath          *string                                          `protobuf:"bytes,2,opt,name=url_path" json:"url_path,omitempty"`
	UseHttps         *bool                                            `protobuf:"varint,3,opt,name=use_https" json:"use_https,omitempty"`
	HttpMethod       *int32                                           `protobuf:"varint,4,opt,name=http_method" json:"http_method,omitempty"`
	RequestHeaders   []*ClientCloudFileUploadBlockDetails_HTTPHeaders `protobuf:"bytes,5,rep,name=request_headers" json:"request_headers,omitempty"`
	BlockOffset      *uint64                                          `protobuf:"varint,6,opt,name=block_offset" json:"block_offset,omitempty"`
	BlockLength      *uint32                                          `protobuf:"varint,7,opt,name=block_length" json:"block_length,omitempty"`
	ExplicitBodyData []byte                                           `protobuf:"bytes,8,opt,name=explicit_body_data" json:"explicit_body_data,omitempty"`
	MayParallelize   *bool                                            `protobuf:"varint,9,opt,name=may_parallelize" json:"may_parallelize,omitempty"`
	XXX_unrecognized []byte                                           `json:"-"`
}

func (m *ClientCloudFileUploadBlockDetails) Reset()         { *m = ClientCloudFileUploadBlockDetails{} }
func (m *ClientCloudFileUploadBlockDetails) String() string { return proto.CompactTextString(m) }
func (*ClientCloudFileUploadBlockDetails) ProtoMessage()    {}

func (m *ClientCloudFileUploadBlockDetails) GetUrlHost() string {
	if m != nil && m.UrlHost != nil {
		return *m.UrlHost
	}
	return ""
}

func (m *ClientCloudFileUploadBlockDetails) GetUrlPath() string {
	if m != nil && m.UrlPath != nil {
		return *m.UrlPath
	}
	return ""
}

func (m *ClientCloudFileUploadBlockDetails) GetUseHttps() bool {
	if m != nil && m.UseHttps != nil {
		return *m.UseHttps
	}
	return false
}

func (m *ClientCloudFileUploadBlockDetails) GetHttpMethod() int32 {
	if m != nil && m.HttpMethod != nil {
		return *m.HttpMethod
	}
	return 0
}

func (m *ClientCloudFileUploadBlockDetails) GetRequestHeaders() []*ClientCloudFileUploadBlockDetails_HTTPHeaders {
	if m != nil {
		return m.RequestHeaders
	}
	return nil
}

func (m *ClientCloudFileUploadBlockDetails) GetBlockOffset() uint64 {
	if m != nil && m.BlockOffset != nil {
		return *m.BlockOffset
	}
	return 0
}

func (m *ClientCloudFileUploadBlockDetails) GetBlockLength() uint32 {
	if m != nil && m.BlockLength != nil {
		return *m.BlockLength
	}
	return 0
}

func (m *ClientCloudFileUploadBlockDetails) GetExplicitBodyData() []byte {
	if m != nil {
		return m.ExplicitBodyData
	}
	return nil
}

func (m *ClientCloudFileUploadBlockDetails) GetMayParallelize() bool {
	if m != nil && m.MayParallelize != nil {
		return *m.MayParallelize
	}
	return false
}

type ClientCloudFileUploadBlockDetails_HTTPHeaders struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClientCloudFileUploadBlockDetails_HTTPHeaders) Reset() {
	*m = ClientCloudFileUploadBlockDetails_HTTPHeaders{}
}
func (m *ClientCloudFileUploadBlockDetails_HTTPHeaders) String() string {
	return proto.CompactTextString(m)
}
func (*ClientCloudFileUploadBlockDetails_HTTPHeaders) ProtoMessage() {}

func (m *ClientCloudFileUploadBlockDetails_HTTPHeaders) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ClientCloudFileUploadBlockDetails_HTTPHeaders) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type CCloud_ClientBeginFileUpload_Response struct {
	EncryptFile      *bool                                `protobuf:"varint,1,opt,name=encrypt_file" json:"encrypt_file,omitempty"`
	BlockRequests    []*ClientCloudFileUploadBlockDetails `protobuf:"bytes,2,rep,name=block_requests" json:"block_requests,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *CCloud_ClientBeginFileUpload_Response) Reset()         { *m = CCloud_ClientBeginFileUpload_Response{} }
func (m *CCloud_ClientBeginFileUpload_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_ClientBeginFileUpload_Response) ProtoMessage()    {}

func (m *CCloud_ClientBeginFileUpload_Response) GetEncryptFile() bool {
	if m != nil && m.EncryptFile != nil {
		return *m.EncryptFile
	}
	return false
}

func (m *CCloud_ClientBeginFileUpload_Response) GetBlockRequests() []*ClientCloudFileUploadBlockDetails {
	if m != nil {
		return m.BlockRequests
	}
	return nil
}

type CCloud_ClientCommitFileUpload_Request struct {
	TransferSucceeded *bool   `protobuf:"varint,1,opt,name=transfer_succeeded" json:"transfer_succeeded,omitempty"`
	Appid             *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	FileSha           []byte  `protobuf:"bytes,3,opt,name=file_sha" json:"file_sha,omitempty"`
	Filename          *string `protobuf:"bytes,4,opt,name=filename" json:"filename,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *CCloud_ClientCommitFileUpload_Request) Reset()         { *m = CCloud_ClientCommitFileUpload_Request{} }
func (m *CCloud_ClientCommitFileUpload_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_ClientCommitFileUpload_Request) ProtoMessage()    {}

func (m *CCloud_ClientCommitFileUpload_Request) GetTransferSucceeded() bool {
	if m != nil && m.TransferSucceeded != nil {
		return *m.TransferSucceeded
	}
	return false
}

func (m *CCloud_ClientCommitFileUpload_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_ClientCommitFileUpload_Request) GetFileSha() []byte {
	if m != nil {
		return m.FileSha
	}
	return nil
}

func (m *CCloud_ClientCommitFileUpload_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

type CCloud_ClientCommitFileUpload_Response struct {
	FileCommitted    *bool  `protobuf:"varint,1,opt,name=file_committed" json:"file_committed,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CCloud_ClientCommitFileUpload_Response) Reset() {
	*m = CCloud_ClientCommitFileUpload_Response{}
}
func (m *CCloud_ClientCommitFileUpload_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_ClientCommitFileUpload_Response) ProtoMessage()    {}

func (m *CCloud_ClientCommitFileUpload_Response) GetFileCommitted() bool {
	if m != nil && m.FileCommitted != nil {
		return *m.FileCommitted
	}
	return false
}

type CCloud_ClientFileDownload_Request struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Filename         *string `protobuf:"bytes,2,opt,name=filename" json:"filename,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_ClientFileDownload_Request) Reset()         { *m = CCloud_ClientFileDownload_Request{} }
func (m *CCloud_ClientFileDownload_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_ClientFileDownload_Request) ProtoMessage()    {}

func (m *CCloud_ClientFileDownload_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_ClientFileDownload_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

type CCloud_ClientFileDownload_Response struct {
	Appid            *uint32                                           `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	FileSize         *uint32                                           `protobuf:"varint,2,opt,name=file_size" json:"file_size,omitempty"`
	RawFileSize      *uint32                                           `protobuf:"varint,3,opt,name=raw_file_size" json:"raw_file_size,omitempty"`
	ShaFile          []byte                                            `protobuf:"bytes,4,opt,name=sha_file" json:"sha_file,omitempty"`
	TimeStamp        *uint64                                           `protobuf:"varint,5,opt,name=time_stamp" json:"time_stamp,omitempty"`
	IsExplicitDelete *bool                                             `protobuf:"varint,6,opt,name=is_explicit_delete" json:"is_explicit_delete,omitempty"`
	UrlHost          *string                                           `protobuf:"bytes,7,opt,name=url_host" json:"url_host,omitempty"`
	UrlPath          *string                                           `protobuf:"bytes,8,opt,name=url_path" json:"url_path,omitempty"`
	UseHttps         *bool                                             `protobuf:"varint,9,opt,name=use_https" json:"use_https,omitempty"`
	RequestHeaders   []*CCloud_ClientFileDownload_Response_HTTPHeaders `protobuf:"bytes,10,rep,name=request_headers" json:"request_headers,omitempty"`
	Encrypted        *bool                                             `protobuf:"varint,11,opt,name=encrypted" json:"encrypted,omitempty"`
	XXX_unrecognized []byte                                            `json:"-"`
}

func (m *CCloud_ClientFileDownload_Response) Reset()         { *m = CCloud_ClientFileDownload_Response{} }
func (m *CCloud_ClientFileDownload_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_ClientFileDownload_Response) ProtoMessage()    {}

func (m *CCloud_ClientFileDownload_Response) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_ClientFileDownload_Response) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *CCloud_ClientFileDownload_Response) GetRawFileSize() uint32 {
	if m != nil && m.RawFileSize != nil {
		return *m.RawFileSize
	}
	return 0
}

func (m *CCloud_ClientFileDownload_Response) GetShaFile() []byte {
	if m != nil {
		return m.ShaFile
	}
	return nil
}

func (m *CCloud_ClientFileDownload_Response) GetTimeStamp() uint64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *CCloud_ClientFileDownload_Response) GetIsExplicitDelete() bool {
	if m != nil && m.IsExplicitDelete != nil {
		return *m.IsExplicitDelete
	}
	return false
}

func (m *CCloud_ClientFileDownload_Response) GetUrlHost() string {
	if m != nil && m.UrlHost != nil {
		return *m.UrlHost
	}
	return ""
}

func (m *CCloud_ClientFileDownload_Response) GetUrlPath() string {
	if m != nil && m.UrlPath != nil {
		return *m.UrlPath
	}
	return ""
}

func (m *CCloud_ClientFileDownload_Response) GetUseHttps() bool {
	if m != nil && m.UseHttps != nil {
		return *m.UseHttps
	}
	return false
}

func (m *CCloud_ClientFileDownload_Response) GetRequestHeaders() []*CCloud_ClientFileDownload_Response_HTTPHeaders {
	if m != nil {
		return m.RequestHeaders
	}
	return nil
}

func (m *CCloud_ClientFileDownload_Response) GetEncrypted() bool {
	if m != nil && m.Encrypted != nil {
		return *m.Encrypted
	}
	return false
}

type CCloud_ClientFileDownload_Response_HTTPHeaders struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_ClientFileDownload_Response_HTTPHeaders) Reset() {
	*m = CCloud_ClientFileDownload_Response_HTTPHeaders{}
}
func (m *CCloud_ClientFileDownload_Response_HTTPHeaders) String() string {
	return proto.CompactTextString(m)
}
func (*CCloud_ClientFileDownload_Response_HTTPHeaders) ProtoMessage() {}

func (m *CCloud_ClientFileDownload_Response_HTTPHeaders) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CCloud_ClientFileDownload_Response_HTTPHeaders) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type CCloud_ClientDeleteFile_Request struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Filename         *string `protobuf:"bytes,2,opt,name=filename" json:"filename,omitempty"`
	IsExplicitDelete *bool   `protobuf:"varint,3,opt,name=is_explicit_delete" json:"is_explicit_delete,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_ClientDeleteFile_Request) Reset()         { *m = CCloud_ClientDeleteFile_Request{} }
func (m *CCloud_ClientDeleteFile_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_ClientDeleteFile_Request) ProtoMessage()    {}

func (m *CCloud_ClientDeleteFile_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_ClientDeleteFile_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CCloud_ClientDeleteFile_Request) GetIsExplicitDelete() bool {
	if m != nil && m.IsExplicitDelete != nil {
		return *m.IsExplicitDelete
	}
	return false
}

type CCloud_ClientDeleteFile_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CCloud_ClientDeleteFile_Response) Reset()         { *m = CCloud_ClientDeleteFile_Response{} }
func (m *CCloud_ClientDeleteFile_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_ClientDeleteFile_Response) ProtoMessage()    {}

func init() {
}
