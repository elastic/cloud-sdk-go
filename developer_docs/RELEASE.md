# Releasing a new version

This guide aims to provide guidance on how to release new versions of `cloud-sdk-go` as well as updating all the necessary parts to make it successful.

- [Releasing a new version](#releasing-a-new-version)
  - [Prerequisites](#prerequisites)
    - [Make sure `VERSION`, `ECE_VERSION` and `ECE_BRANCH` are updated](#make-sure-version-ece_version-and-ece_branch-are-updated)
    - [Update the `apidocs.json` and `apidocs-user.json` files](#update-the-apidocsjson-and-apidocs-userjson-files)
    - [Generating a changelog for the new version](#generating-a-changelog-for-the-new-version)
  - [Executing the release](#executing-the-release)

## Prerequisites

Releasing a new version implies that there have been changes in the source code which are meant to be released for wider consumption. Before releasing a new version there's some prerequisites that have to be checked.

### Make sure `VERSION`, `ECE_VERSION` and `ECE_BRANCH` are updated

The important part is to make sure that `ECE_BRANCH` matches the milestone that the ECE release tracks. The variable is defined in the `Makefile`. Once updated, you can move on to the next step.

**Since the `VERSION` and `ECE_VERSION`  are now updated via github actions, just double check that these are updated, and if not, manually do so**.

Since the source has changed, we need to update the current committed version to a higher version so that the release is published.

The version is currently defined in the [Makefile](./Makefile) as an exported environment variable called `VERSION` in the [SEMVER](https://semver.org) format: `MAJOR.MINOR.PATCH`

```Makefile
SHELL := /bin/bash
export VERSION ?= v1.0.0
```

Say we want to perform a minor version release (i.e. no breaking changes and only new features and bug fixes are being included); in which case we'll update the _MINOR_ part of the version, this can be done with the `make minor` target, but it should have been updated automatically via GitHub actions.

```Makefile
SHELL := /bin/bash
export VERSION ?= v1.1.0
```

If a patch version needs to be released, the release will be done from the minor branch. For example, if we want to release `v1.5.1`, we will check out the `1.5` branch and perform any changes in that branch. The VERSION variable in the Makefile should already be up to date, but in case it's not, it can be bumped with the `make patch` target.

### Update the `apidocs.json` and `apidocs-user.json` files

A new release will be required every time a new ECE minor or major version is released, the ECE major tracks an `ms-<id>` branch, this identifier needs to be updated in the Makefile `ECE_BRANCH` variable, then calling `make update-swagger` will update the two OpenAPI spec files.

By default, it will try and clone the cloud repository into `/tmp/cloud`, but you might choose to override that location via the `CLOUD_SOURCE_REPO` Makefile variable: `make CLOUD_SOURCE_REPO=/a/path/to/cloud update-swagger` or alternatively you can copy your **clean** local copy of the cloud repo to `/tmp/cloud`: `cp -R /a/path/to/cloud /tmp/cloud` since cloning a brand new copy of the cloud repo can take a long time.

After `make update-swagger` has been run, `make swagger` needs to be run to re-generate the API and model Go code which will generate the following files:

* `pkg/client`: All the generated API code.
* `pkg/models`: All the generated API models.
* `api/version/<ECE_VERSION>.md`: A Markdown document containing all the API endpoints for the version.

After the changes have been made, open a pull request with all the changes targeting the relevant branch (minor versions will target `master`, while patch versions will target the relevant branch). For example, if `1.6.0` needs to be released, then the pull request will be opened against master, but if `1.5.1` needs to be released, the pull request will target the `1.5` branch.

### Generating a changelog for the new version

The changelog should be automatically generated on each push to `master` or the relevant branch, if no changelog file is available in `notes/<VERSION>.md`, the target will fail, this means that no changelog entries have been created and some need to be generated. See previous versions under `.changelog/<VERSION>` for some examples on changelog entries, the folder name should match the version. To read more information on how to generate a changelog [see the changelogger README](../cmd/changelogger/README.md).

## Executing the release

After all the prerequisites have been ticked off, the only thing remaining is to run `make tag`. The target which will attempt to release a new version.
