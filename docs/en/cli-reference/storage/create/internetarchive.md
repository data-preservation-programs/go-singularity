# Internet Archive

{% code fullWidth="true" %}
```
NAME:
   singularity storage create internetarchive - Internet Archive

USAGE:
   singularity storage create internetarchive [command options] <name> <path>

DESCRIPTION:
   --access_key_id
      IAS3 Access Key.
      
      Leave blank for anonymous access.
      You can find one here: https://archive.org/account/s3.php

   --secret_access_key
      IAS3 Secret Key (password).
      
      Leave blank for anonymous access.

   --endpoint
      IAS3 Endpoint.
      
      Leave blank for default value.

   --front_endpoint
      Host of InternetArchive Frontend.
      
      Leave blank for default value.

   --disable_checksum
      Don't ask the server to test against MD5 checksum calculated by rclone.
      Normally rclone will calculate the MD5 checksum of the input before
      uploading it so it can ask the server to check the object against checksum.
      This is great for data integrity checking but can cause long delays for
      large files to start uploading.

   --wait_archive
      Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
      Only enable if you need to be guaranteed to be reflected after write operations.
      0 to disable waiting. No errors to be thrown in case of timeout.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --access_key_id value      IAS3 Access Key. [$ACCESS_KEY_ID]
   --help, -h                 show help
   --secret_access_key value  IAS3 Secret Key (password). [$SECRET_ACCESS_KEY]

   Advanced

   --disable_checksum      Don't ask the server to test against MD5 checksum calculated by rclone. (default: true) [$DISABLE_CHECKSUM]
   --encoding value        The encoding for the backend. (default: "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value        IAS3 Endpoint. (default: "https://s3.us.archive.org") [$ENDPOINT]
   --front_endpoint value  Host of InternetArchive Frontend. (default: "https://archive.org") [$FRONT_ENDPOINT]
   --wait_archive value    Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish. (default: "0s") [$WAIT_ARCHIVE]

```
{% endcode %}