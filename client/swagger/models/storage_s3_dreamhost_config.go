// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageS3DreamhostConfig storage s3 dreamhost config
//
// swagger:model storage.S3DreamhostConfig
type StorageS3DreamhostConfig struct {

	// AWS Access Key ID.
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// Canned ACL used when creating buckets and storing or copying objects.
	ACL string `json:"acl,omitempty"`

	// Canned ACL used when creating buckets.
	// Example: private
	BucketACL string `json:"bucketAcl,omitempty"`

	// Chunk size to use for uploading.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// Cutoff for switching to multipart copy.
	CopyCutoff *string `json:"copyCutoff,omitempty"`

	// If set this will decompress gzip encoded objects.
	Decompress *bool `json:"decompress,omitempty"`

	// Don't store MD5 checksum with object metadata.
	DisableChecksum *bool `json:"disableChecksum,omitempty"`

	// Disable usage of http2 for S3 backends.
	DisableHttp2 *bool `json:"disableHttp2,omitempty"`

	// Custom endpoint for downloads.
	DownloadURL string `json:"downloadUrl,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Endpoint for S3 API.
	// Example: objects-us-east-1.dream.io
	Endpoint string `json:"endpoint,omitempty"`

	// Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	// Example: false
	EnvAuth *bool `json:"envAuth,omitempty"`

	// If true use path style access if false use virtual hosted style.
	ForcePathStyle *bool `json:"forcePathStyle,omitempty"`

	// Size of listing chunk (response list for each ListObject S3 request).
	ListChunk *int64 `json:"listChunk,omitempty"`

	// Whether to url encode listings: true/false/unset
	ListURLEncode *string `json:"listUrlEncode,omitempty"`

	// Version of ListObjects to use: 1,2 or 0 for auto.
	ListVersion int64 `json:"listVersion,omitempty"`

	// Location constraint - must be set to match the Region.
	LocationConstraint string `json:"locationConstraint,omitempty"`

	// Maximum number of parts in a multipart upload.
	MaxUploadParts *int64 `json:"maxUploadParts,omitempty"`

	// How often internal memory buffer pools will be flushed.
	MemoryPoolFlushTime *string `json:"memoryPoolFlushTime,omitempty"`

	// Whether to use mmap buffers in internal memory pool.
	MemoryPoolUseMmap *bool `json:"memoryPoolUseMmap,omitempty"`

	// Set this if the backend might gzip objects.
	MightGzip *string `json:"mightGzip,omitempty"`

	// If set, don't attempt to check the bucket exists or create it.
	NoCheckBucket *bool `json:"noCheckBucket,omitempty"`

	// If set, don't HEAD uploaded objects to check integrity.
	NoHead *bool `json:"noHead,omitempty"`

	// If set, do not do HEAD before GET when getting objects.
	NoHeadObject *bool `json:"noHeadObject,omitempty"`

	// Suppress setting and reading of system metadata
	NoSystemMetadata *bool `json:"noSystemMetadata,omitempty"`

	// Profile to use in the shared credentials file.
	Profile string `json:"profile,omitempty"`

	// Region to connect to.
	Region string `json:"region,omitempty"`

	// AWS Secret Access Key (password).
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// An AWS session token.
	SessionToken string `json:"sessionToken,omitempty"`

	// Path to the shared credentials file.
	SharedCredentialsFile string `json:"sharedCredentialsFile,omitempty"`

	// Concurrency for multipart uploads.
	UploadConcurrency *int64 `json:"uploadConcurrency,omitempty"`

	// Cutoff for switching to chunked upload.
	UploadCutoff *string `json:"uploadCutoff,omitempty"`

	// Whether to use ETag in multipart uploads for verification
	UseMultipartEtag *string `json:"useMultipartEtag,omitempty"`

	// Whether to use a presigned request or PutObject for single part uploads
	UsePresignedRequest *bool `json:"usePresignedRequest,omitempty"`

	// If true use v2 authentication.
	V2Auth *bool `json:"v2Auth,omitempty"`

	// Show file versions as they were at the specified time.
	VersionAt *string `json:"versionAt,omitempty"`

	// Include old versions in directory listings.
	Versions *bool `json:"versions,omitempty"`
}

// Validate validates this storage s3 dreamhost config
func (m *StorageS3DreamhostConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage s3 dreamhost config based on context it is used
func (m *StorageS3DreamhostConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageS3DreamhostConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageS3DreamhostConfig) UnmarshalBinary(b []byte) error {
	var res StorageS3DreamhostConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
