# Mega

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add mega - Mega

USAGE:
   singularity datasource add mega [command options] <dataset_name> <source_path>

DESCRIPTION:
   --mega-debug
      Output more debug from Mega.
      
      If this flag is set (along with -vv) it will print further debugging
      information from the mega backend.

   --mega-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --mega-hard-delete
      Delete files permanently rather than putting them into the trash.
      
      Normally the mega backend will put all deletions into the trash rather
      than permanently deleting them.  If you specify this then rclone will
      permanently delete objects instead.

   --mega-pass
      Password.

   --mega-use-https
      Use HTTPS for transfers.
      
      MEGA uses plain text HTTP connections by default.
      Some ISPs throttle HTTP connections, this causes transfers to become very slow.
      Enabling this will force MEGA to use HTTPS for all transfers.
      HTTPS is normally not necesary since all data is already encrypted anyway.
      Enabling it will increase CPU usage and add network overhead.

   --mega-user
      User name.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)
   --scanning-state value   set the initial scanning state (default: ready)

   Options for mega

   --mega-debug value        Output more debug from Mega. (default: "false") [$MEGA_DEBUG]
   --mega-encoding value     The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$MEGA_ENCODING]
   --mega-hard-delete value  Delete files permanently rather than putting them into the trash. (default: "false") [$MEGA_HARD_DELETE]
   --mega-pass value         Password. [$MEGA_PASS]
   --mega-use-https value    Use HTTPS for transfers. (default: "false") [$MEGA_USE_HTTPS]
   --mega-user value         User name. [$MEGA_USER]

```
{% endcode %}
