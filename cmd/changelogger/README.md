# Change Logger

This small cmd utility that generates a complete changelog aggregating individual changelog files from a specific folder.

## Usage

```console
$ ./changelogger -h
Generates a complete changelog aggregating individual changelog files from a specific folder.
The -changelog-dir flag's value is joined with -version, resulting: ${changelog-dir}/${version}.

Usage of changelogger:
  -base-url string
    	base URL to use for links when shorthand ref is specified
  -changelog-dir string
    	path to the changelog directory (default ".changelog")
  -template string
    	template to generate the resulting changelog
  -version string
    	version for the changelog being generated. Any 'v' prefix will be stripped
```