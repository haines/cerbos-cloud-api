// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.1.0

package bootstrapv1

import (
	v1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/bundle/v1"
	protowire "google.golang.org/protobuf/encoding/protowire"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	hash "hash"
)

func cerbos_cloud_bootstrap_v1_PDPConfig_Meta_hashpb_sum(m *PDPConfig_Meta, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.bootstrap.v1.PDPConfig.Meta.created_at"]; !ok {
		if m.CreatedAt != nil {
			google_protobuf_Timestamp_hashpb_sum(m.CreatedAt, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.cloud.bootstrap.v1.PDPConfig.Meta.commit_hash"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.CommitHash))

	}
}

func cerbos_cloud_bootstrap_v1_PDPConfig_hashpb_sum(m *PDPConfig, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.bootstrap.v1.PDPConfig.meta"]; !ok {
		if m.Meta != nil {
			cerbos_cloud_bootstrap_v1_PDPConfig_Meta_hashpb_sum(m.Meta, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.cloud.bootstrap.v1.PDPConfig.bundle_info"]; !ok {
		if m.BundleInfo != nil {
			cerbos_cloud_bundle_v1_BundleInfo_hashpb_sum(m.BundleInfo, hasher, ignore)
		}

	}
}

func cerbos_cloud_bundle_v1_BundleInfo_Segment_hashpb_sum(m *v1.BundleInfo_Segment, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.bundle.v1.BundleInfo.Segment.segment_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.SegmentId)))

	}
	if _, ok := ignore["cerbos.cloud.bundle.v1.BundleInfo.Segment.checksum"]; !ok {
		_, _ = hasher.Write(protowire.AppendBytes(nil, m.Checksum))

	}
	if _, ok := ignore["cerbos.cloud.bundle.v1.BundleInfo.Segment.download_urls"]; !ok {
		if len(m.DownloadUrls) > 0 {
			for _, v := range m.DownloadUrls {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_cloud_bundle_v1_BundleInfo_hashpb_sum(m *v1.BundleInfo, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.cloud.bundle.v1.BundleInfo.label"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Label))

	}
	if _, ok := ignore["cerbos.cloud.bundle.v1.BundleInfo.bundle_hash"]; !ok {
		_, _ = hasher.Write(protowire.AppendBytes(nil, m.BundleHash))

	}
	if _, ok := ignore["cerbos.cloud.bundle.v1.BundleInfo.segments"]; !ok {
		if len(m.Segments) > 0 {
			for _, v := range m.Segments {
				if v != nil {
					cerbos_cloud_bundle_v1_BundleInfo_Segment_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Timestamp_hashpb_sum(m *timestamppb.Timestamp, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Timestamp.seconds"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Seconds)))

	}
	if _, ok := ignore["google.protobuf.Timestamp.nanos"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Nanos)))

	}
}
