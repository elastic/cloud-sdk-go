# Change Logger

This small cmd utility that generates a complete changelog aggregating individual changelog files from a specific folder.

## Usage

```console
$ ./changelogger -h
Generates a complete changelog aggregating individual changelog files from a specific folder.
The -changelog-dir flag's value is joined with -version, resulting: ${changelog-dir}/${version}.

Usage of changelogger:
  -base-url string
    	base URL to use for each of the changes
  -changelog-dir string
    	path to the changelog directory (default ".changelog")
  -template string
    	template to generate the resulting changelog
  -version string
    	version for the changelog being generated. Any 'v' prefix will be stripped
```

### Changelog entries

Each of the changelog entries can be `JSON` or `YAML` formatted.

#### Reference

```yaml
category: <type of change, could be anything but generally you want to specify the type of change>
title: <short description of the change, it can be used for short / compacted changelogs or as the title of more complex ones>
ref: <optional identifier for the change if unset, the name of the file will be used. The full URL will be creating by interpolating the base-url flag value with this value>
description: <long description that can be used for long or descriptive changelogs combined with the title field>
```

#### Examples

##### Enhancement

```yaml
category: enhancement
title: Improve `allocatorapi.Vacate` API calls
description: |
    Optimizes the API calls that the Vacate function performs in order to reduce the amount of work
    the API has to perform to calculate a cluster's move.
```

##### Bug

```yaml
category: bug
title: Fix ignored "order" field in `DeploymentTemplate` structures
description: |
    Updates the field type of the `DeploymentTemplate` structures to `*int32` so the `0` value is not lost when the JSON is decoded.
```
