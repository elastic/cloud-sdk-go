{{- /* Change definitions */}}
{{- $breaking := .ByCategory "breaking-change" }}
{{- $enhancement := .ByCategory "enhancement" }}
{{- $bug := .ByCategory "bug" }}
{{- $newAPI := .ByCategory "new-api" }}
{{- $doc := .ByCategory "doc" -}}
# Changelog

This release of the Elastic Cloud SDK Go should be used for ECE Version `{{ Env "ECE_VERSION" }}`.
{{ if $breaking }}
## Breaking changes
{{range $breaking}}
### {{ .TitleOrRef }} ([#{{.Ref}}]({{BaseURL .Ref}}))

{{ .Description }}
{{- end}}
{{- end}}

{{- if $newAPI}}
## New APIs
{{range $newAPI}}
### {{ .TitleOrRef }} ([#{{.Ref}}]({{BaseURL .Ref}}))

{{ .Description }}
{{- end}}
{{- end}}

{{- if $enhancement}}
## Enhancements
{{range $enhancement}}
### {{ .TitleOrRef }} ([#{{.Ref}}]({{BaseURL .Ref}}))

{{ .Description }}
{{- end}}
{{- end}}

{{- if $bug}}
## Bug fixes
{{range (.ByCategory "bug")}}
### {{ .TitleOrRef }} ([#{{.Ref}}]({{BaseURL .Ref}}))

{{ .Description }}
{{- end}}
{{- end}}

{{- if $doc}}
## Docs
{{range (.ByCategory "doc")}}
### {{ .TitleOrRef }} ([#{{.Ref}}]({{BaseURL .Ref}}))

{{ .Description }}
{{- end}}
{{- end}}
