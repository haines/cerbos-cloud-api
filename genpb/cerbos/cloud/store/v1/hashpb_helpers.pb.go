// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.3.5

package storev1

import (
	protowire "google.golang.org/protobuf/encoding/protowire"
	hash "hash"
)

func cerbos_cloud_store_v1_FileError_hashpb_sum(m *FileError, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.FileError.file"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetFile()))

	}
	if _, ok := ignore["cerbos.cloud.store.v1.FileError.cause"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetCause())))

	}
	if _, ok := ignore["cerbos.cloud.store.v1.FileError.details"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetDetails()))

	}
}

func cerbos_cloud_store_v1_FileFilter_hashpb_sum(m *FileFilter, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.FileFilter.path"]; !ok {
		if m.GetPath() != nil {
			cerbos_cloud_store_v1_StringMatch_hashpb_sum(m.GetPath(), hasher, ignore)
		}

	}
}

func cerbos_cloud_store_v1_FileOp_hashpb_sum(m *FileOp, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Op != nil {
		if _, ok := ignore["cerbos.cloud.store.v1.FileOp.op"]; !ok {
			switch t := m.Op.(type) {
			case *FileOp_AddOrUpdate:
				if t.AddOrUpdate != nil {
					cerbos_cloud_store_v1_File_hashpb_sum(t.AddOrUpdate, hasher, ignore)
				}

			case *FileOp_Delete:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Delete))

			}
		}
	}
}

func cerbos_cloud_store_v1_File_hashpb_sum(m *File, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.File.path"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetPath()))

	}
	if _, ok := ignore["cerbos.cloud.store.v1.File.contents"]; !ok {
		_, _ = hasher.Write(protowire.AppendBytes(nil, m.GetContents()))

	}
}

func cerbos_cloud_store_v1_GetFilesRequest_hashpb_sum(m *GetFilesRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.GetFilesRequest.store_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetStoreId()))

	}
	if _, ok := ignore["cerbos.cloud.store.v1.GetFilesRequest.files"]; !ok {
		if len(m.Files) > 0 {
			for _, v := range m.Files {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_cloud_store_v1_GetFilesResponse_hashpb_sum(m *GetFilesResponse, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.GetFilesResponse.store_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetStoreVersion())))

	}
	if _, ok := ignore["cerbos.cloud.store.v1.GetFilesResponse.files"]; !ok {
		if len(m.Files) > 0 {
			for _, v := range m.Files {
				if v != nil {
					cerbos_cloud_store_v1_File_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_cloud_store_v1_ListFilesRequest_hashpb_sum(m *ListFilesRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.ListFilesRequest.store_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetStoreId()))

	}
	if _, ok := ignore["cerbos.cloud.store.v1.ListFilesRequest.filter"]; !ok {
		if m.GetFilter() != nil {
			cerbos_cloud_store_v1_FileFilter_hashpb_sum(m.GetFilter(), hasher, ignore)
		}

	}
}

func cerbos_cloud_store_v1_ListFilesResponse_hashpb_sum(m *ListFilesResponse, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.ListFilesResponse.store_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetStoreVersion())))

	}
	if _, ok := ignore["cerbos.cloud.store.v1.ListFilesResponse.files"]; !ok {
		if len(m.Files) > 0 {
			for _, v := range m.Files {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_cloud_store_v1_ModifyFilesRequest_Condition_hashpb_sum(m *ModifyFilesRequest_Condition, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.ModifyFilesRequest.Condition.store_version_must_equal"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetStoreVersionMustEqual())))

	}
}

func cerbos_cloud_store_v1_ModifyFilesRequest_hashpb_sum(m *ModifyFilesRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.ModifyFilesRequest.store_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetStoreId()))

	}
	if _, ok := ignore["cerbos.cloud.store.v1.ModifyFilesRequest.condition"]; !ok {
		if m.GetCondition() != nil {
			cerbos_cloud_store_v1_ModifyFilesRequest_Condition_hashpb_sum(m.GetCondition(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.cloud.store.v1.ModifyFilesRequest.operations"]; !ok {
		if len(m.Operations) > 0 {
			for _, v := range m.Operations {
				if v != nil {
					cerbos_cloud_store_v1_FileOp_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_cloud_store_v1_ModifyFilesResponse_Failure_hashpb_sum(m *ModifyFilesResponse_Failure, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.ModifyFilesResponse.Failure.errors"]; !ok {
		if len(m.Errors) > 0 {
			for _, v := range m.Errors {
				if v != nil {
					cerbos_cloud_store_v1_FileError_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_cloud_store_v1_ModifyFilesResponse_Success_hashpb_sum(m *ModifyFilesResponse_Success, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.ModifyFilesResponse.Success.new_store_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetNewStoreVersion())))

	}
}

func cerbos_cloud_store_v1_ModifyFilesResponse_hashpb_sum(m *ModifyFilesResponse, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Result != nil {
		if _, ok := ignore["cerbos.cloud.store.v1.ModifyFilesResponse.result"]; !ok {
			switch t := m.Result.(type) {
			case *ModifyFilesResponse_Success_:
				if t.Success != nil {
					cerbos_cloud_store_v1_ModifyFilesResponse_Success_hashpb_sum(t.Success, hasher, ignore)
				}

			case *ModifyFilesResponse_Failure_:
				if t.Failure != nil {
					cerbos_cloud_store_v1_ModifyFilesResponse_Failure_hashpb_sum(t.Failure, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_cloud_store_v1_ReplaceFilesRequest_Condition_hashpb_sum(m *ReplaceFilesRequest_Condition, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.ReplaceFilesRequest.Condition.store_version_must_equal"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetStoreVersionMustEqual())))

	}
}

func cerbos_cloud_store_v1_ReplaceFilesRequest_hashpb_sum(m *ReplaceFilesRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.ReplaceFilesRequest.store_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetStoreId()))

	}
	if _, ok := ignore["cerbos.cloud.store.v1.ReplaceFilesRequest.condition"]; !ok {
		if m.GetCondition() != nil {
			cerbos_cloud_store_v1_ReplaceFilesRequest_Condition_hashpb_sum(m.GetCondition(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.cloud.store.v1.ReplaceFilesRequest.zipped_contents"]; !ok {
		_, _ = hasher.Write(protowire.AppendBytes(nil, m.GetZippedContents()))

	}
}

func cerbos_cloud_store_v1_ReplaceFilesResponse_Failure_hashpb_sum(m *ReplaceFilesResponse_Failure, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.ReplaceFilesResponse.Failure.errors"]; !ok {
		if len(m.Errors) > 0 {
			for _, v := range m.Errors {
				if v != nil {
					cerbos_cloud_store_v1_FileError_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_cloud_store_v1_ReplaceFilesResponse_Success_hashpb_sum(m *ReplaceFilesResponse_Success, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.ReplaceFilesResponse.Success.new_store_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetNewStoreVersion())))

	}
}

func cerbos_cloud_store_v1_ReplaceFilesResponse_hashpb_sum(m *ReplaceFilesResponse, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Result != nil {
		if _, ok := ignore["cerbos.cloud.store.v1.ReplaceFilesResponse.result"]; !ok {
			switch t := m.Result.(type) {
			case *ReplaceFilesResponse_Success_:
				if t.Success != nil {
					cerbos_cloud_store_v1_ReplaceFilesResponse_Success_hashpb_sum(t.Success, hasher, ignore)
				}

			case *ReplaceFilesResponse_Failure_:
				if t.Failure != nil {
					cerbos_cloud_store_v1_ReplaceFilesResponse_Failure_hashpb_sum(t.Failure, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_cloud_store_v1_StringMatch_InList_hashpb_sum(m *StringMatch_InList, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.store.v1.StringMatch.InList.values"]; !ok {
		if len(m.Values) > 0 {
			for _, v := range m.Values {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_cloud_store_v1_StringMatch_hashpb_sum(m *StringMatch, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Match != nil {
		if _, ok := ignore["cerbos.cloud.store.v1.StringMatch.match"]; !ok {
			switch t := m.Match.(type) {
			case *StringMatch_Equals:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Equals))

			case *StringMatch_Like:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Like))

			case *StringMatch_In:
				if t.In != nil {
					cerbos_cloud_store_v1_StringMatch_InList_hashpb_sum(t.In, hasher, ignore)
				}

			}
		}
	}
}
