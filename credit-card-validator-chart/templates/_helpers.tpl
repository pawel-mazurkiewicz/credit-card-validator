{{- define "credit-card-validator.fullname" -}}
{{- if .Chart.Name -}}
{{ include "credit-card-validator.name" . }}-{{ .Release.Name }}
{{- else -}}
{{ .Release.Name }}
{{- end -}}
{{- end }}

{{- define "credit-card-validator.name" -}}
credit-card-validator
{{- end }}