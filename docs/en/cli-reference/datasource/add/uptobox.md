# Uptobox

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add uptobox - Uptobox

USAGE:
   singularity datasource add uptobox [command options] <dataset_name> <source_path>

DESCRIPTION:
   --uptobox-access-token
      Your access token.
      
      Get it from https://uptobox.com/my_account.

   --uptobox-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for uptobox

   --uptobox-access-token value  Your access token. [$UPTOBOX_ACCESS_TOKEN]
   --uptobox-encoding value      The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot") [$UPTOBOX_ENCODING]

```
{% endcode %}
