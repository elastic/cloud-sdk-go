{{- /* Change definitions */}}
{{- $breaking := .ByCategory "breaking-change" }}
{{- $enhancement := .ByCategory "enhancement" }}
{{- $bug := .ByCategory "bug" }}
{{- $newAPI := .ByCategory "new-api" }}
{{- $doc := .ByCategory "doc" -}}
# Changelog

{{ if $breaking }}
BREAKING:
{{range $breaking}}
* {{ .TitleOrRef }} ([#{{.Ref}}]({{BaseURL .Ref}}))
{{- end}}
{{- end}}

{{- if $newAPI}}
NEW API:
{{range $newAPI}}
* {{ .TitleOrRef }} ([#{{.Ref}}]({{BaseURL .Ref}}))
{{- end}}
{{- end}}

{{- if $enhancement}}
ENHANCEMENT:
{{range $enhancement}}
* {{ .TitleOrRef }} ([#{{.Ref}}]({{BaseURL .Ref}}))
{{- end}}
{{- end}}

{{- if $bug}}
BUG:
{{range (.ByCategory "bug")}}
* {{ .TitleOrRef }} ([#{{.Ref}}]({{BaseURL .Ref}}))
{{- end}}
{{- end}}

{{- if $doc}}
DOCS:
{{range (.ByCategory "doc")}}
* {{ .TitleOrRef }} ([#{{.Ref}}]({{BaseURL .Ref}}))
{{- end}}
{{- end}}
