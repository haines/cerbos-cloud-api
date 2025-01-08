// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.3.4
// Source: cerbos/cloud/store/v1/store.proto

package storev1

import (
	hash "hash"
)

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *StringMatch) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_StringMatch_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *StringMatch_InList) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_StringMatch_InList_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *FileFilter) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_FileFilter_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ListFilesRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_ListFilesRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ListFilesResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_ListFilesResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *GetFilesRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_GetFilesRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *File) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_File_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *GetFilesResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_GetFilesResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *FileOp) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_FileOp_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ModifyFilesRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_ModifyFilesRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ModifyFilesRequest_Condition) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_ModifyFilesRequest_Condition_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ModifyFilesResponse) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_ModifyFilesResponse_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ModifyFilesResponse_Error) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_ModifyFilesResponse_Error_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ModifyFilesResponse_Failure) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_ModifyFilesResponse_Failure_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ModifyFilesResponse_Success) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_store_v1_ModifyFilesResponse_Success_hashpb_sum(m, hasher, ignore)
	}
}
