# Changelog

## Breaking changes

Removes the previously deprecated `create-es-cluster` command.

## Notable changes

### New Commands

- `add-blueprinter-blessing` (**PUT** /platform/infrastructure/blueprinter/roles/{blueprinter_role_id}/blessings/{runner_id})
- `cancel-apm-pending-plan` (**DELETE** /clusters/apm/{cluster_id}/plan/pending)
- `cancel-deployment-resource-pending-plan` (**DELETE** /deployments/{deployment_id}/{resource_kind}/{ref_id}/plan/pending)

### Bug fixes

Fixed a bug where the injected Asset Template loader wasn't being used, resulting on the inability to extend the formatter with further 3rd party templates provided by consuming clients.

### Docs

Added a new "Workflow" section which covers all the contributing workflow guidelines for contributors, particularly focusing on commit messages.
Additionally PR and Issue templates have been added and a few re-wording changes have been made.
