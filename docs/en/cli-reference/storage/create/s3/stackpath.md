# StackPath Object Storage

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 stackpath - StackPath Object Storage

USAGE:
   singularity storage create s3 stackpath [command options] <name> <path>

DESCRIPTION:
   --env_auth
      Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
      
      Only applies if access_key_id and secret_access_key is blank.

      Examples:
         | false | Enter AWS credentials in the next step.
         | true  | Get AWS credentials from the environment (env vars or IAM).

   --access_key_id
      AWS Access Key ID.
      
      Leave blank for anonymous access or runtime credentials.

   --secret_access_key
      AWS Secret Access Key (password).
      
      Leave blank for anonymous access or runtime credentials.

   --region
      Region to connect to.
      
      Leave blank if you are using an S3 clone and you don't have a region.

      Examples:
         | <unset>            | Use this if unsure.
         |                    | Will use v4 signatures and an empty region.
         | other-v2-signature | Use this only if v4 signatures don't work.
         |                    | E.g. pre Jewel/v10 CEPH.

   --endpoint
      Endpoint for StackPath Object Storage.

      Examples:
         | s3.us-east-2.stackpathstorage.com    | US East Endpoint
         | s3.us-west-1.stackpathstorage.com    | US West Endpoint
         | s3.eu-central-1.stackpathstorage.com | EU Endpoint

   --acl
      Canned ACL used when creating buckets and storing or copying objects.
      
      This ACL is used for creating objects and if bucket_acl isn't set, for creating buckets too.
      
      For more info visit https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      Note that this ACL is applied when server-side copying objects as S3
      doesn't copy the ACL from the source but rather writes a fresh one.
      
      If the acl is an empty string then no X-Amz-Acl: header is added and
      the default (private) will be used.
      

   --bucket_acl
      Canned ACL used when creating buckets.
      
      For more info visit https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      Note that this ACL is applied when only when creating buckets.  If it
      isn't set then "acl" is used instead.
      
      If the "acl" and "bucket_acl" are empty strings then no X-Amz-Acl:
      header is added and the default (private) will be used.
      

      Examples:
         | private            | Owner gets FULL_CONTROL.
         |                    | No one else has access rights (default).
         | public-read        | Owner gets FULL_CONTROL.
         |                    | The AllUsers group gets READ access.
         | public-read-write  | Owner gets FULL_CONTROL.
         |                    | The AllUsers group gets READ and WRITE access.
         |                    | Granting this on a bucket is generally not recommended.
         | authenticated-read | Owner gets FULL_CONTROL.
         |                    | The AuthenticatedUsers group gets READ access.

   --upload_cutoff
      Cutoff for switching to chunked upload.
      
      Any files larger than this will be uploaded in chunks of chunk_size.
      The minimum is 0 and the maximum is 5 GiB.

   --chunk_size
      Chunk size to use for uploading.
      
      When uploading files larger than upload_cutoff or files with unknown
      size (e.g. from "rclone rcat" or uploaded with "rclone mount" or google
      photos or google docs) they will be uploaded as multipart uploads
      using this chunk size.
      
      Note that "--s3-upload-concurrency" chunks of this size are buffered
      in memory per transfer.
      
      If you are transferring large files over high-speed links and you have
      enough memory, then increasing this will speed up the transfers.
      
      Rclone will automatically increase the chunk size when uploading a
      large file of known size to stay below the 10,000 chunks limit.
      
      Files of unknown size are uploaded with the configured
      chunk_size. Since the default chunk size is 5 MiB and there can be at
      most 10,000 chunks, this means that by default the maximum size of
      a file you can stream upload is 48 GiB.  If you wish to stream upload
      larger files then you will need to increase chunk_size.
      
      Increasing the chunk size decreases the accuracy of the progress
      statistics displayed with "-P" flag. Rclone treats chunk as sent when
      it's buffered by the AWS SDK, when in fact it may still be uploading.
      A bigger chunk size means a bigger AWS SDK buffer and progress
      reporting more deviating from the truth.
      

   --max_upload_parts
      Maximum number of parts in a multipart upload.
      
      This option defines the maximum number of multipart chunks to use
      when doing a multipart upload.
      
      This can be useful if a service does not support the AWS S3
      specification of 10,000 chunks.
      
      Rclone will automatically increase the chunk size when uploading a
      large file of a known size to stay below this number of chunks limit.
      

   --copy_cutoff
      Cutoff for switching to multipart copy.
      
      Any files larger than this that need to be server-side copied will be
      copied in chunks of this size.
      
      The minimum is 0 and the maximum is 5 GiB.

   --disable_checksum
      Don't store MD5 checksum with object metadata.
      
      Normally rclone will calculate the MD5 checksum of the input before
      uploading it so it can add it to metadata on the object. This is great
      for data integrity checking but can cause long delays for large files
      to start uploading.

   --shared_credentials_file
      Path to the shared credentials file.
      
      If env_auth = true then rclone can use a shared credentials file.
      
      If this variable is empty rclone will look for the
      "AWS_SHARED_CREDENTIALS_FILE" env variable. If the env value is empty
      it will default to the current user's home directory.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      Profile to use in the shared credentials file.
      
      If env_auth = true then rclone can use a shared credentials file. This
      variable controls which profile is used in that file.
      
      If empty it will default to the environment variable "AWS_PROFILE" or
      "default" if that environment variable is also not set.
      

   --session_token
      An AWS session token.

   --upload_concurrency
      Concurrency for multipart uploads.
      
      This is the number of chunks of the same file that are uploaded
      concurrently.
      
      If you are uploading small numbers of large files over high-speed links
      and these uploads do not fully utilize your bandwidth, then increasing
      this may help to speed up the transfers.

   --force_path_style
      If true use path style access if false use virtual hosted style.
      
      If this is true (the default) then rclone will use path style access,
      if false then rclone will use virtual path style. See [the AWS S3
      docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      for more info.
      
      Some providers (e.g. AWS, Aliyun OSS, Netease COS, or Tencent COS) require this set to
      false - rclone will do this automatically based on the provider
      setting.

   --v2_auth
      If true use v2 authentication.
      
      If this is false (the default) then rclone will use v4 authentication.
      If it is set then rclone will use v2 authentication.
      
      Use this only if v4 signatures don't work, e.g. pre Jewel/v10 CEPH.

   --list_chunk
      Size of listing chunk (response list for each ListObject S3 request).
      
      This option is also known as "MaxKeys", "max-items", or "page-size" from the AWS S3 specification.
      Most services truncate the response list to 1000 objects even if requested more than that.
      In AWS S3 this is a global maximum and cannot be changed, see [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html).
      In Ceph, this can be increased with the "rgw list buckets max chunk" option.
      

   --list_version
      Version of ListObjects to use: 1,2 or 0 for auto.
      
      When S3 originally launched it only provided the ListObjects call to
      enumerate objects in a bucket.
      
      However in May 2016 the ListObjectsV2 call was introduced. This is
      much higher performance and should be used if at all possible.
      
      If set to the default, 0, rclone will guess according to the provider
      set which list objects method to call. If it guesses wrong, then it
      may be set manually here.
      

   --list_url_encode
      Whether to url encode listings: true/false/unset
      
      Some providers support URL encoding listings and where this is
      available this is more reliable when using control characters in file
      names. If this is set to unset (the default) then rclone will choose
      according to the provider setting what to apply, but you can override
      rclone's choice here.
      

   --no_check_bucket
      If set, don't attempt to check the bucket exists or create it.
      
      This can be useful when trying to minimise the number of transactions
      rclone does if you know the bucket exists already.
      
      It can also be needed if the user you are using does not have bucket
      creation permissions. Before v1.52.0 this would have passed silently
      due to a bug.
      

   --no_head
      If set, don't HEAD uploaded objects to check integrity.
      
      This can be useful when trying to minimise the number of transactions
      rclone does.
      
      Setting it means that if rclone receives a 200 OK message after
      uploading an object with PUT then it will assume that it got uploaded
      properly.
      
      In particular it will assume:
      
      - the metadata, including modtime, storage class and content type was as uploaded
      - the size was as uploaded
      
      It reads the following items from the response for a single part PUT:
      
      - the MD5SUM
      - The uploaded date
      
      For multipart uploads these items aren't read.
      
      If an source object of unknown length is uploaded then rclone **will** do a
      HEAD request.
      
      Setting this flag increases the chance for undetected upload failures,
      in particular an incorrect size, so it isn't recommended for normal
      operation. In practice the chance of an undetected upload failure is
      very small even with this flag.
      

   --no_head_object
      If set, do not do HEAD before GET when getting objects.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --memory_pool_flush_time
      How often internal memory buffer pools will be flushed.
      
      Uploads which requires additional buffers (f.e multipart) will use memory pool for allocations.
      This option controls how often unused buffers will be removed from the pool.

   --memory_pool_use_mmap
      Whether to use mmap buffers in internal memory pool.

   --disable_http2
      Disable usage of http2 for S3 backends.
      
      There is currently an unsolved issue with the s3 (specifically minio) backend
      and HTTP/2.  HTTP/2 is enabled by default for the s3 backend but can be
      disabled here.  When the issue is solved this flag will be removed.
      
      See: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download_url
      Custom endpoint for downloads.
      This is usually set to a CloudFront CDN URL as AWS S3 offers
      cheaper egress for data downloaded through the CloudFront network.

   --use_multipart_etag
      Whether to use ETag in multipart uploads for verification
      
      This should be true, false or left unset to use the default for the provider.
      

   --use_presigned_request
      Whether to use a presigned request or PutObject for single part uploads
      
      If this is false rclone will use PutObject from the AWS SDK to upload
      an object.
      
      Versions of rclone < 1.59 use presigned requests to upload a single
      part object and setting this flag to true will re-enable that
      functionality. This shouldn't be necessary except in exceptional
      circumstances or for testing.
      

   --versions
      Include old versions in directory listings.

   --version_at
      Show file versions as they were at the specified time.
      
      The parameter should be a date, "2006-01-02", datetime "2006-01-02
      15:04:05" or a duration for that long ago, eg "100d" or "1h".
      
      Note that when using this no file write operations are permitted,
      so you can't upload files or delete them.
      
      See [the time option docs](/docs/#time-option) for valid formats.
      

   --decompress
      If set this will decompress gzip encoded objects.
      
      It is possible to upload objects to S3 with "Content-Encoding: gzip"
      set. Normally rclone will download these files as compressed objects.
      
      If this flag is set then rclone will decompress these files with
      "Content-Encoding: gzip" as they are received. This means that rclone
      can't check the size and hash but the file contents will be decompressed.
      

   --might_gzip
      Set this if the backend might gzip objects.
      
      Normally providers will not alter objects when they are downloaded. If
      an object was not uploaded with `Content-Encoding: gzip` then it won't
      be set on download.
      
      However some providers may gzip objects even if they weren't uploaded
      with `Content-Encoding: gzip` (eg Cloudflare).
      
      A symptom of this would be receiving errors like
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      If you set this flag and rclone downloads an object with
      Content-Encoding: gzip set and chunked transfer encoding, then rclone
      will decompress the object on the fly.
      
      If this is set to unset (the default) then rclone will choose
      according to the provider setting what to apply, but you can override
      rclone's choice here.
      

   --no_system_metadata
      Suppress setting and reading of system metadata


OPTIONS:
   --access_key_id value      AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                Canned ACL used when creating buckets and storing or copying objects. [$ACL]
   --endpoint value           Endpoint for StackPath Object Storage. [$ENDPOINT]
   --env_auth                 Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars). (default: false) [$ENV_AUTH]
   --help, -h                 show help
   --region value             Region to connect to. [$REGION]
   --secret_access_key value  AWS Secret Access Key (password). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket_acl value               Canned ACL used when creating buckets. [$BUCKET_ACL]
   --chunk_size value               Chunk size to use for uploading. (default: "5Mi") [$CHUNK_SIZE]
   --copy_cutoff value              Cutoff for switching to multipart copy. (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     If set this will decompress gzip encoded objects. (default: false) [$DECOMPRESS]
   --disable_checksum               Don't store MD5 checksum with object metadata. (default: false) [$DISABLE_CHECKSUM]
   --disable_http2                  Disable usage of http2 for S3 backends. (default: false) [$DISABLE_HTTP2]
   --download_url value             Custom endpoint for downloads. [$DOWNLOAD_URL]
   --encoding value                 The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force_path_style               If true use path style access if false use virtual hosted style. (default: true) [$FORCE_PATH_STYLE]
   --list_chunk value               Size of listing chunk (response list for each ListObject S3 request). (default: 1000) [$LIST_CHUNK]
   --list_url_encode value          Whether to url encode listings: true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list_version value             Version of ListObjects to use: 1,2 or 0 for auto. (default: 0) [$LIST_VERSION]
   --max_upload_parts value         Maximum number of parts in a multipart upload. (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory_pool_flush_time value   How often internal memory buffer pools will be flushed. (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory_pool_use_mmap           Whether to use mmap buffers in internal memory pool. (default: false) [$MEMORY_POOL_USE_MMAP]
   --might_gzip value               Set this if the backend might gzip objects. (default: "unset") [$MIGHT_GZIP]
   --no_check_bucket                If set, don't attempt to check the bucket exists or create it. (default: false) [$NO_CHECK_BUCKET]
   --no_head                        If set, don't HEAD uploaded objects to check integrity. (default: false) [$NO_HEAD]
   --no_head_object                 If set, do not do HEAD before GET when getting objects. (default: false) [$NO_HEAD_OBJECT]
   --no_system_metadata             Suppress setting and reading of system metadata (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  Profile to use in the shared credentials file. [$PROFILE]
   --session_token value            An AWS session token. [$SESSION_TOKEN]
   --shared_credentials_file value  Path to the shared credentials file. [$SHARED_CREDENTIALS_FILE]
   --upload_concurrency value       Concurrency for multipart uploads. (default: 4) [$UPLOAD_CONCURRENCY]
   --upload_cutoff value            Cutoff for switching to chunked upload. (default: "200Mi") [$UPLOAD_CUTOFF]
   --use_multipart_etag value       Whether to use ETag in multipart uploads for verification (default: "unset") [$USE_MULTIPART_ETAG]
   --use_presigned_request          Whether to use a presigned request or PutObject for single part uploads (default: false) [$USE_PRESIGNED_REQUEST]
   --v2_auth                        If true use v2 authentication. (default: false) [$V2_AUTH]
   --version_at value               Show file versions as they were at the specified time. (default: "off") [$VERSION_AT]
   --versions                       Include old versions in directory listings. (default: false) [$VERSIONS]

```
{% endcode %}