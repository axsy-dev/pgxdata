// Do not edit. Automatically generated from templates.go.erb
package main

import (
	"encoding/base64"
	"fmt"
	"text/template"
)

func pkPlaceholder(n int) string {
	return fmt.Sprintf("$%d", n+1)
}

func loadTemplates() *template.Template {
	funcMap := template.FuncMap{"pkPlaceholder": pkPlaceholder}
	templates := template.New("base").Funcs(funcMap)

	var decoded []byte
	var err error

	decoded, err = base64.StdEncoding.DecodeString(`Ly8gVGhpcyBmaWxlIGlzIGF1dG9tYXRpY2FsbHkgZ2VuZXJhdGVkLgpwYWNrYWdlIHt7LlBrZ05hbWV9fQoKaW1wb3J0ICgKICAiZW5jb2RpbmcvanNvbiIKICAiZXJyb3JzIgogICJmbXQiCiAgIm5ldCIKICAic3RyY29udiIKICAidGltZSIKCiAgImdpdGh1Yi5jb20vamFja2MvcGd4IgopCgp0eXBlIFN0YXR1cyBieXRlCgpjb25zdCAoCiAgVW5kZWZpbmVkIFN0YXR1cyA9IGlvdGEKICBOdWxsCiAgUHJlc2VudAopCgpmdW5jIChzIFN0YXR1cykgU3RyaW5nKCkgc3RyaW5nIHsKICBzd2l0Y2ggcyB7CiAgY2FzZSBVbmRlZmluZWQ6CiAgICByZXR1cm4gIlVuZGVmaW5lZCIKICBjYXNlIE51bGw6CiAgICByZXR1cm4gIk51bGwiCiAgY2FzZSBQcmVzZW50OgogICAgcmV0dXJuICJQcmVzZW50IgogIH0KCiAgcmV0dXJuICJJbnZhbGlkIHN0YXR1cyIKfQoKe3tyYW5nZSAuQm94VHlwZXN9fQp0eXBlIHt7Lk5hbWV9fSBzdHJ1Y3QgewogIFZhbHVlICB7ey5WYWx1ZVR5cGV9fQogIFN0YXR1cyBTdGF0dXMKfQoKZnVuYyAoYXR0ciAqe3suTmFtZX19KSBTdHJpbmcoKSBzdHJpbmcgewogIGlmIGF0dHIuU3RhdHVzID09IFByZXNlbnQgewogICAgcmV0dXJuIGZtdC5TcHJpbnRmKCIldiIsIGF0dHIuVmFsdWUpCiAgfQogIHJldHVybiBhdHRyLlN0YXR1cy5TdHJpbmcoKQp9CgpmdW5jIChhdHRyICp7ey5OYW1lfX0pIGFkZFVwZGF0ZShjb2x1bW5OYW1lIHN0cmluZywgc2V0cyAqW11zdHJpbmcsIGFyZ3MgKnBneC5RdWVyeUFyZ3MpIHsKICBzd2l0Y2ggYXR0ci5TdGF0dXMgewogICAgY2FzZSBQcmVzZW50LCBOdWxsOgogICAgICAqc2V0cyA9IGFwcGVuZCgqc2V0cywgY29sdW1uTmFtZSsiPSIrYXJncy5BcHBlbmQoYXR0cikpCiAgfQp9CgpmdW5jIChhdHRyICp7ey5OYW1lfX0pIGFkZEluc2VydChjb2x1bW5OYW1lIHN0cmluZywgc2V0cywgdmFsdWVzICpbXXN0cmluZywgYXJncyAqcGd4LlF1ZXJ5QXJncykgewogIHN3aXRjaCBhdHRyLlN0YXR1cyB7CiAgICBjYXNlIFByZXNlbnQsIE51bGw6CiAgICAgICpzZXRzID0gYXBwZW5kKCpzZXRzLCBjb2x1bW5OYW1lKQogICAgICAqdmFsdWVzID0gYXBwZW5kKCp2YWx1ZXMsIGFyZ3MuQXBwZW5kKGF0dHIpKQogIH0KfQoKZnVuYyAoYXR0ciAqe3suTmFtZX19KSBGb3JtYXRDb2RlKCkgaW50MTYgewogIHJldHVybiB7ey5Gb3JtYXRDb2RlfX0KfQp7e2VuZH19CgpmdW5jIChhdHRyICpCb29sKSBTY2FuKHZyICpwZ3guVmFsdWVSZWFkZXIpIGVycm9yIHsKICBpZiB2ci5UeXBlKCkuRGF0YVR5cGUgIT0gcGd4LkJvb2xPaWQgewogICAgcmV0dXJuIHBneC5TZXJpYWxpemF0aW9uRXJyb3IoZm10LlNwcmludGYoIkJvb2wuU2NhbiBjYW5ub3QgZGVjb2RlIE9JRCAlZCIsIHZyLlR5cGUoKS5EYXRhVHlwZSkpCiAgfQoKICBpZiB2ci5MZW4oKSA9PSAtMSB7CiAgICBhdHRyLlZhbHVlID0gZmFsc2UKICAgIGF0dHIuU3RhdHVzID0gTnVsbAogICAgcmV0dXJuIG5pbAogIH0KCiAgZXJyIDo9IHBneC5EZWNvZGUodnIsICZhdHRyLlZhbHVlKQogIGlmIGVyciAhPSBuaWwgewogICAgcmV0dXJuIGVycgogIH0KICBhdHRyLlN0YXR1cyA9IFByZXNlbnQKCiAgcmV0dXJuIHZyLkVycigpCn0KCmZ1bmMgKGF0dHIgKkJvb2wpIEVuY29kZSh3ICpwZ3guV3JpdGVCdWYsIG9pZCBwZ3guT2lkKSBlcnJvciB7CiAgaWYgb2lkICE9IHBneC5Cb29sT2lkIHsKICAgIHJldHVybiBwZ3guU2VyaWFsaXphdGlvbkVycm9yKGZtdC5TcHJpbnRmKCJCb29sLkVuY29kZSBjYW5ub3QgZW5jb2RlIGludG8gT0lEICVkIiwgb2lkKSkKICB9CgogIHN3aXRjaCBhdHRyLlN0YXR1cyB7CiAgY2FzZSBQcmVzZW50OgogICAgcmV0dXJuIHBneC5FbmNvZGUodywgb2lkLCBhdHRyLlZhbHVlKQogIGNhc2UgTnVsbDoKICAgIHcuV3JpdGVJbnQzMigtMSkKICAgIHJldHVybiBuaWwKICBjYXNlIFVuZGVmaW5lZDoKICAgIHJldHVybiBlcnJvcnMuTmV3KCJjYW5ub3QgZW5jb2RlIHVuZGVmaW5lZCBhdHRyIikKICBkZWZhdWx0OgogICAgcGFuaWMoInVucmVhY2hhYmxlIikKICB9Cn0KCmZ1bmMgKGF0dHIgKkludDE2KSBTY2FuKHZyICpwZ3guVmFsdWVSZWFkZXIpIGVycm9yIHsKICBpZiB2ci5UeXBlKCkuRGF0YVR5cGUgIT0gcGd4LkludDJPaWQgewogICAgcmV0dXJuIHBneC5TZXJpYWxpemF0aW9uRXJyb3IoZm10LlNwcmludGYoIkludDE2LlNjYW4gY2Fubm90IGRlY29kZSBPSUQgJWQiLCB2ci5UeXBlKCkuRGF0YVR5cGUpKQogIH0KCiAgaWYgdnIuTGVuKCkgPT0gLTEgewogICAgYXR0ci5WYWx1ZSA9IDAKICAgIGF0dHIuU3RhdHVzID0gTnVsbAogICAgcmV0dXJuIG5pbAogIH0KCiAgZXJyIDo9IHBneC5EZWNvZGUodnIsICZhdHRyLlZhbHVlKQogIGlmIGVyciAhPSBuaWwgewogICAgcmV0dXJuIGVycgogIH0KICBhdHRyLlN0YXR1cyA9IFByZXNlbnQKCiAgcmV0dXJuIHZyLkVycigpCn0KCmZ1bmMgKGF0dHIgKkludDE2KSBFbmNvZGUodyAqcGd4LldyaXRlQnVmLCBvaWQgcGd4Lk9pZCkgZXJyb3IgewogIGlmIG9pZCAhPSBwZ3guSW50Mk9pZCB7CiAgICByZXR1cm4gcGd4LlNlcmlhbGl6YXRpb25FcnJvcihmbXQuU3ByaW50ZigiSW50MTYuRW5jb2RlIGNhbm5vdCBlbmNvZGUgaW50byBPSUQgJWQiLCBvaWQpKQogIH0KCiAgc3dpdGNoIGF0dHIuU3RhdHVzIHsKICBjYXNlIFByZXNlbnQ6CiAgICByZXR1cm4gcGd4LkVuY29kZSh3LCBvaWQsIGF0dHIuVmFsdWUpCiAgY2FzZSBOdWxsOgogICAgdy5Xcml0ZUludDMyKC0xKQogICAgcmV0dXJuIG5pbAogIGNhc2UgVW5kZWZpbmVkOgogICAgcmV0dXJuIGVycm9ycy5OZXcoImNhbm5vdCBlbmNvZGUgdW5kZWZpbmVkIGF0dHIiKQogIGRlZmF1bHQ6CiAgICBwYW5pYygidW5yZWFjaGFibGUiKQogIH0KfQoKZnVuYyAoYXR0ciAqSW50MzIpIFNjYW4odnIgKnBneC5WYWx1ZVJlYWRlcikgZXJyb3IgewogIGlmIHZyLlR5cGUoKS5EYXRhVHlwZSAhPSBwZ3guSW50NE9pZCB7CiAgICByZXR1cm4gcGd4LlNlcmlhbGl6YXRpb25FcnJvcihmbXQuU3ByaW50ZigiSW50MzIuU2NhbiBjYW5ub3QgZGVjb2RlIE9JRCAlZCIsIHZyLlR5cGUoKS5EYXRhVHlwZSkpCiAgfQoKICBpZiB2ci5MZW4oKSA9PSAtMSB7CiAgICBhdHRyLlZhbHVlID0gMAogICAgYXR0ci5TdGF0dXMgPSBOdWxsCiAgICByZXR1cm4gbmlsCiAgfQoKICBlcnIgOj0gcGd4LkRlY29kZSh2ciwgJmF0dHIuVmFsdWUpCiAgaWYgZXJyICE9IG5pbCB7CiAgICByZXR1cm4gZXJyCiAgfQogIGF0dHIuU3RhdHVzID0gUHJlc2VudAoKICByZXR1cm4gdnIuRXJyKCkKfQoKZnVuYyAoYXR0ciAqSW50MzIpIEVuY29kZSh3ICpwZ3guV3JpdGVCdWYsIG9pZCBwZ3guT2lkKSBlcnJvciB7CiAgaWYgb2lkICE9IHBneC5JbnQ0T2lkIHsKICAgIHJldHVybiBwZ3guU2VyaWFsaXphdGlvbkVycm9yKGZtdC5TcHJpbnRmKCJJbnQzMi5FbmNvZGUgY2Fubm90IGVuY29kZSBpbnRvIE9JRCAlZCIsIG9pZCkpCiAgfQoKICBzd2l0Y2ggYXR0ci5TdGF0dXMgewogIGNhc2UgUHJlc2VudDoKICAgIHJldHVybiBwZ3guRW5jb2RlKHcsIG9pZCwgYXR0ci5WYWx1ZSkKICBjYXNlIE51bGw6CiAgICB3LldyaXRlSW50MzIoLTEpCiAgICByZXR1cm4gbmlsCiAgY2FzZSBVbmRlZmluZWQ6CiAgICByZXR1cm4gZXJyb3JzLk5ldygiY2Fubm90IGVuY29kZSB1bmRlZmluZWQgYXR0ciIpCiAgZGVmYXVsdDoKICAgIHBhbmljKCJ1bnJlYWNoYWJsZSIpCiAgfQp9CgpmdW5jIChhdHRyICpJbnQ2NCkgU2Nhbih2ciAqcGd4LlZhbHVlUmVhZGVyKSBlcnJvciB7CiAgaWYgdnIuVHlwZSgpLkRhdGFUeXBlICE9IHBneC5JbnQ4T2lkIHsKICAgIHJldHVybiBwZ3guU2VyaWFsaXphdGlvbkVycm9yKGZtdC5TcHJpbnRmKCJJbnQ2NC5TY2FuIGNhbm5vdCBkZWNvZGUgT0lEICVkIiwgdnIuVHlwZSgpLkRhdGFUeXBlKSkKICB9CgogIGlmIHZyLkxlbigpID09IC0xIHsKICAgIGF0dHIuVmFsdWUgPSAwCiAgICBhdHRyLlN0YXR1cyA9IE51bGwKICAgIHJldHVybiBuaWwKICB9CgogIGVyciA6PSBwZ3guRGVjb2RlKHZyLCAmYXR0ci5WYWx1ZSkKICBpZiBlcnIgIT0gbmlsIHsKICAgIHJldHVybiBlcnIKICB9CiAgYXR0ci5TdGF0dXMgPSBQcmVzZW50CgogIHJldHVybiB2ci5FcnIoKQp9CgpmdW5jIChhdHRyICpJbnQ2NCkgRW5jb2RlKHcgKnBneC5Xcml0ZUJ1Ziwgb2lkIHBneC5PaWQpIGVycm9yIHsKICBpZiBvaWQgIT0gcGd4LkludDhPaWQgewogICAgcmV0dXJuIHBneC5TZXJpYWxpemF0aW9uRXJyb3IoZm10LlNwcmludGYoIkludDY0LkVuY29kZSBjYW5ub3QgZW5jb2RlIGludG8gT0lEICVkIiwgb2lkKSkKICB9CgogIHN3aXRjaCBhdHRyLlN0YXR1cyB7CiAgY2FzZSBQcmVzZW50OgogICAgcmV0dXJuIHBneC5FbmNvZGUodywgb2lkLCBhdHRyLlZhbHVlKQogIGNhc2UgTnVsbDoKICAgIHcuV3JpdGVJbnQzMigtMSkKICAgIHJldHVybiBuaWwKICBjYXNlIFVuZGVmaW5lZDoKICAgIHJldHVybiBlcnJvcnMuTmV3KCJjYW5ub3QgZW5jb2RlIHVuZGVmaW5lZCBhdHRyIikKICBkZWZhdWx0OgogICAgcGFuaWMoInVucmVhY2hhYmxlIikKICB9Cn0KCmZ1bmMgKGF0dHIgKlN0cmluZykgU2Nhbih2ciAqcGd4LlZhbHVlUmVhZGVyKSBlcnJvciB7CiAgaWYgdnIuTGVuKCkgPT0gLTEgewogICAgYXR0ci5WYWx1ZSA9ICIiCiAgICBhdHRyLlN0YXR1cyA9IE51bGwKICAgIHJldHVybiBuaWwKICB9CgogIGVyciA6PSBwZ3guRGVjb2RlKHZyLCAmYXR0ci5WYWx1ZSkKICBpZiBlcnIgIT0gbmlsIHsKICAgIHJldHVybiBlcnIKICB9CiAgYXR0ci5TdGF0dXMgPSBQcmVzZW50CgogIHJldHVybiB2ci5FcnIoKQp9CgpmdW5jIChhdHRyICpTdHJpbmcpIEVuY29kZSh3ICpwZ3guV3JpdGVCdWYsIG9pZCBwZ3guT2lkKSBlcnJvciB7CiAgc3dpdGNoIGF0dHIuU3RhdHVzIHsKICBjYXNlIFByZXNlbnQ6CiAgICByZXR1cm4gcGd4LkVuY29kZSh3LCBvaWQsIGF0dHIuVmFsdWUpCiAgY2FzZSBOdWxsOgogICAgdy5Xcml0ZUludDMyKC0xKQogICAgcmV0dXJuIG5pbAogIGNhc2UgVW5kZWZpbmVkOgogICAgcmV0dXJuIGVycm9ycy5OZXcoImNhbm5vdCBlbmNvZGUgdW5kZWZpbmVkIGF0dHIiKQogIGRlZmF1bHQ6CiAgICBwYW5pYygidW5yZWFjaGFibGUiKQogIH0KfQoKZnVuYyAoYXR0ciAqVGltZSkgU2Nhbih2ciAqcGd4LlZhbHVlUmVhZGVyKSBlcnJvciB7CiAgb2lkIDo9IHZyLlR5cGUoKS5EYXRhVHlwZQogIGlmIG9pZCAhPSBwZ3guVGltZXN0YW1wVHpPaWQgJiYgb2lkICE9IHBneC5UaW1lc3RhbXBPaWQgJiYgb2lkICE9IHBneC5EYXRlT2lkIHsKICAgIHJldHVybiBwZ3guU2VyaWFsaXphdGlvbkVycm9yKGZtdC5TcHJpbnRmKCJUaW1lLlNjYW4gY2Fubm90IGRlY29kZSBPSUQgJWQiLCB2ci5UeXBlKCkuRGF0YVR5cGUpKQogIH0KCiAgaWYgdnIuTGVuKCkgPT0gLTEgewogICAgYXR0ci5WYWx1ZSA9IHRpbWUuVGltZXt9CiAgICBhdHRyLlN0YXR1cyA9IE51bGwKICAgIHJldHVybiBuaWwKICB9CgogIGVyciA6PSBwZ3guRGVjb2RlKHZyLCAmYXR0ci5WYWx1ZSkKICBpZiBlcnIgIT0gbmlsIHsKICAgIHJldHVybiBlcnIKICB9CiAgYXR0ci5TdGF0dXMgPSBQcmVzZW50CgogIHJldHVybiB2ci5FcnIoKQp9CgpmdW5jIChhdHRyICpUaW1lKSBFbmNvZGUodyAqcGd4LldyaXRlQnVmLCBvaWQgcGd4Lk9pZCkgZXJyb3IgewogIGlmIG9pZCAhPSBwZ3guVGltZXN0YW1wVHpPaWQgJiYgb2lkICE9IHBneC5UaW1lc3RhbXBPaWQgJiYgb2lkICE9IHBneC5EYXRlT2lkIHsKICAgIHJldHVybiBwZ3guU2VyaWFsaXphdGlvbkVycm9yKGZtdC5TcHJpbnRmKCJUaW1lLkVuY29kZSBjYW5ub3QgZW5jb2RlIGludG8gT0lEICVkIiwgb2lkKSkKICB9CgogIHN3aXRjaCBhdHRyLlN0YXR1cyB7CiAgY2FzZSBQcmVzZW50OgogICAgcmV0dXJuIHBneC5FbmNvZGUodywgb2lkLCBhdHRyLlZhbHVlKQogIGNhc2UgTnVsbDoKICAgIHcuV3JpdGVJbnQzMigtMSkKICAgIHJldHVybiBuaWwKICBjYXNlIFVuZGVmaW5lZDoKICAgIHJldHVybiBlcnJvcnMuTmV3KCJjYW5ub3QgZW5jb2RlIHVuZGVmaW5lZCBhdHRyIikKICBkZWZhdWx0OgogICAgcGFuaWMoInVucmVhY2hhYmxlIikKICB9Cn0KCmZ1bmMgKGF0dHIgKklQTmV0KSBTY2FuKHZyICpwZ3guVmFsdWVSZWFkZXIpIGVycm9yIHsKICBvaWQgOj0gdnIuVHlwZSgpLkRhdGFUeXBlCiAgaWYgb2lkICE9IHBneC5JbmV0T2lkICYmIG9pZCAhPSBwZ3guQ2lkck9pZCB7CiAgICByZXR1cm4gcGd4LlNlcmlhbGl6YXRpb25FcnJvcihmbXQuU3ByaW50ZigiSVBOZXQuU2NhbiBjYW5ub3QgZGVjb2RlIE9JRCAlZCIsIHZyLlR5cGUoKS5EYXRhVHlwZSkpCiAgfQoKICBpZiB2ci5MZW4oKSA9PSAtMSB7CiAgICBhdHRyLlZhbHVlID0gbmV0LklQTmV0e30KICAgIGF0dHIuU3RhdHVzID0gTnVsbAogICAgcmV0dXJuIG5pbAogIH0KCiAgYXR0ci5TdGF0dXMgPSBQcmVzZW50CiAgZXJyIDo9IHBneC5EZWNvZGUodnIsICZhdHRyLlZhbHVlKQogIGlmIGVyciAhPSBuaWwgewogICAgcmV0dXJuIGVycgogIH0KCiAgcmV0dXJuIHZyLkVycigpCn0KCmZ1bmMgKGF0dHIgKklQTmV0KSBFbmNvZGUodyAqcGd4LldyaXRlQnVmLCBvaWQgcGd4Lk9pZCkgZXJyb3IgewogIGlmIG9pZCAhPSBwZ3guSW5ldE9pZCAmJiBvaWQgIT0gcGd4LkNpZHJPaWQgewogICAgcmV0dXJuIHBneC5TZXJpYWxpemF0aW9uRXJyb3IoZm10LlNwcmludGYoIklQTmV0LkVuY29kZSBjYW5ub3QgZW5jb2RlIGludG8gT0lEICVkIiwgb2lkKSkKICB9CgogIHN3aXRjaCBhdHRyLlN0YXR1cyB7CiAgY2FzZSBQcmVzZW50OgogICAgcmV0dXJuIHBneC5FbmNvZGUodywgb2lkLCBhdHRyLlZhbHVlKQogIGNhc2UgTnVsbDoKICAgIHcuV3JpdGVJbnQzMigtMSkKICAgIHJldHVybiBuaWwKICBjYXNlIFVuZGVmaW5lZDoKICAgIHJldHVybiBlcnJvcnMuTmV3KCJjYW5ub3QgZW5jb2RlIHVuZGVmaW5lZCBhdHRyIikKICBkZWZhdWx0OgogICAgcGFuaWMoInVucmVhY2hhYmxlIikKICB9Cn0KCnR5cGUgQnl0ZXMgc3RydWN0IHsKICBWYWx1ZSAgW11ieXRlCiAgU3RhdHVzIFN0YXR1cwp9CgpmdW5jIChhdHRyICpCeXRlcykgYWRkVXBkYXRlKGNvbHVtbk5hbWUgc3RyaW5nLCBzZXRzICpbXXN0cmluZywgYXJncyAqcGd4LlF1ZXJ5QXJncykgewogIHN3aXRjaCBhdHRyLlN0YXR1cyB7CiAgICBjYXNlIFByZXNlbnQ6CiAgICAgICpzZXRzID0gYXBwZW5kKCpzZXRzLCBjb2x1bW5OYW1lKyI9IithcmdzLkFwcGVuZChhdHRyLlZhbHVlKSkKICAgIGNhc2UgTnVsbDoKICAgICAgKnNldHMgPSBhcHBlbmQoKnNldHMsIGNvbHVtbk5hbWUrIj0iK2FyZ3MuQXBwZW5kKG5pbCkpCiAgfQp9CgpmdW5jIChhdHRyICpCeXRlcykgYWRkSW5zZXJ0KGNvbHVtbk5hbWUgc3RyaW5nLCBzZXRzLCB2YWx1ZXMgKltdc3RyaW5nLCBhcmdzICpwZ3guUXVlcnlBcmdzKSB7CiAgc3dpdGNoIGF0dHIuU3RhdHVzIHsKICAgIGNhc2UgUHJlc2VudDoKICAgICAgKnNldHMgPSBhcHBlbmQoKnNldHMsIGNvbHVtbk5hbWUpCiAgICAgICp2YWx1ZXMgPSBhcHBlbmQoKnZhbHVlcywgYXJncy5BcHBlbmQoYXR0ci5WYWx1ZSkpCiAgICBjYXNlIE51bGw6CiAgICAgICpzZXRzID0gYXBwZW5kKCpzZXRzLCBjb2x1bW5OYW1lKQogICAgICAqdmFsdWVzID0gYXBwZW5kKCp2YWx1ZXMsIGFyZ3MuQXBwZW5kKG5pbCkpCiAgfQp9CgpmdW5jIChhdHRyICpCeXRlcykgU2Nhbih2ciAqcGd4LlZhbHVlUmVhZGVyKSBlcnJvciB7CiAgaWYgdnIuTGVuKCkgPT0gLTEgewogICAgYXR0ci5WYWx1ZSA9IG5pbAogICAgYXR0ci5TdGF0dXMgPSBOdWxsCiAgICByZXR1cm4gbmlsCiAgfQoKICBhdHRyLlZhbHVlID0gdnIuUmVhZEJ5dGVzKHZyLkxlbigpKQogIGF0dHIuU3RhdHVzID0gUHJlc2VudAogIHJldHVybiB2ci5FcnIoKQp9CgoKZnVuYyAoYXR0ciBCeXRlcykgRm9ybWF0Q29kZSgpIGludDE2IHsKICByZXR1cm4gcGd4LkJpbmFyeUZvcm1hdENvZGUKfQoKZnVuYyAoYXR0ciBCeXRlcykgRW5jb2RlKHcgKnBneC5Xcml0ZUJ1Ziwgb2lkIHBneC5PaWQpIGVycm9yIHsKICBpZiBvaWQgIT0gcGd4LkJ5dGVhT2lkIHsKICAgIHJldHVybiBwZ3guU2VyaWFsaXphdGlvbkVycm9yKGZtdC5TcHJpbnRmKCJCeXRlcy5FbmNvZGUgY2Fubm90IGVuY29kZSBpbnRvIE9JRCAlZCIsIG9pZCkpCiAgfQoKICBpZiBhdHRyLlN0YXR1cyAhPSBQcmVzZW50IHsKICAgIHcuV3JpdGVJbnQzMigtMSkKICAgIHJldHVybiBuaWwKICB9CgogIHcuV3JpdGVCeXRlcyhhdHRyLlZhbHVlKQogIHJldHVybiBuaWwKfQoKCmZ1bmMgKGF0dHIgQm9vbCkgTWFyc2hhbEpTT04oKSAoW11ieXRlLCBlcnJvcikgewogIGlmIGF0dHIuU3RhdHVzICE9IFByZXNlbnQgewogICAgcmV0dXJuIFtdYnl0ZSgibnVsbCIpLCBuaWwKICB9CiAgaWYgYXR0ci5WYWx1ZSB7CiAgICByZXR1cm4gW11ieXRlKCJ0cnVlIiksIG5pbAogIH0KICByZXR1cm4gW11ieXRlKCJmYWxzZSIpLCBuaWwKfQoKZnVuYyAoYXR0ciAqQm9vbCkgVW5tYXJzaGFsSlNPTihidmFsIFtdYnl0ZSkgZXJyb3IgewogIHN2YWwgOj0gc3RyaW5nKGJ2YWwpCgogIHN3aXRjaCBzdmFsIHsKICBjYXNlICJ0cnVlIjoKICAgIGF0dHIuVmFsdWUgPSB0cnVlCiAgICBhdHRyLlN0YXR1cyA9IFByZXNlbnQKICBjYXNlICJmYWxzZSI6CiAgICBhdHRyLlZhbHVlID0gZmFsc2UKICAgIGF0dHIuU3RhdHVzID0gUHJlc2VudAogIGNhc2UgIm51bGwiOgogICAgYXR0ci5TdGF0dXMgPSBOdWxsCiAgZGVmYXVsdDoKICAgIHJldHVybiBlcnJvcnMuTmV3KCJ1bmtub3duIEJvb2wgdmFsdWUiKQogIH0KCiAgcmV0dXJuIG5pbAp9CgoKe3tyYW5nZSAuSW50Qm94VHlwZXN9fQpmdW5jIChhdHRyIHt7Lk5hbWV9fSkgTWFyc2hhbEpTT04oKSAoW11ieXRlLCBlcnJvcikgewogIGlmIGF0dHIuU3RhdHVzICE9IFByZXNlbnQgewogICAgcmV0dXJuIFtdYnl0ZSgibnVsbCIpLCBuaWwKICB9CiAgcmV0dXJuIFtdYnl0ZShzdHJjb252LkZvcm1hdEludChpbnQ2NChhdHRyLlZhbHVlKSwgMTApKSwgbmlsCn0KCmZ1bmMgKGF0dHIgKnt7Lk5hbWV9fSkgVW5tYXJzaGFsSlNPTihidmFsIFtdYnl0ZSkgZXJyb3IgewogIHN2YWwgOj0gc3RyaW5nKGJ2YWwpCgogIGlmIHN2YWwgPT0gIm51bGwiIHsKICAgIGF0dHIuU3RhdHVzID0gTnVsbAogICAgcmV0dXJuIG5pbAogIH0KCiAgbnZhbCwgZXJyIDo9IHN0cmNvbnYuUGFyc2VJbnQoc3ZhbCwgMTAsIHt7LkJpdFNpemV9fSkKICBpZiBlcnIgIT0gbmlsIHsKICAgIHJldHVybiBlcnIKICB9CgogIGF0dHIuVmFsdWUgPSBpbnR7ey5CaXRTaXplfX0obnZhbCkKICBhdHRyLlN0YXR1cyA9IFByZXNlbnQKCiAgcmV0dXJuIG5pbAp9Cnt7ZW5kfX0KCmZ1bmMgKGF0dHIgU3RyaW5nKSBNYXJzaGFsSlNPTigpIChbXWJ5dGUsIGVycm9yKSB7CiAgaWYgYXR0ci5TdGF0dXMgIT0gUHJlc2VudCB7CiAgICByZXR1cm4gW11ieXRlKCJudWxsIiksIG5pbAogIH0KCiAgcmV0dXJuIGpzb24uTWFyc2hhbChhdHRyLlZhbHVlKQp9CgpmdW5jIChhdHRyICpTdHJpbmcpIFVubWFyc2hhbEpTT04oYnZhbCBbXWJ5dGUpIGVycm9yIHsKICBzdmFsIDo9IHN0cmluZyhidmFsKQoKICBpZiBzdmFsID09ICJudWxsIiB7CiAgICBhdHRyLlN0YXR1cyA9IE51bGwKICAgIHJldHVybiBuaWwKICB9CgogIGVyciA6PSBqc29uLlVubWFyc2hhbChidmFsLCAmYXR0ci5WYWx1ZSkKICBpZiBlcnIgIT0gbmlsIHsKICAgIHJldHVybiBlcnIKICB9CgogIGF0dHIuU3RhdHVzID0gUHJlc2VudAogIHJldHVybiBuaWwKfQo=`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`attribute`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`cGFja2FnZSA9ICJ7ey5Qa2dOYW1lfX0iCgojIERhdGFiYXNlIGNvbm5lY3Rpb24gaW5mb3JtYXRpb24gY2FuIGJlIHNwZWNpZmllZCBoZXJlIG9yIGluIFBHKiBlbnZpcm9ubWVudCB2YXJpYWJsZXMKIwojIFtkYXRhYmFzZV0KIyBob3N0ID0gIjEyNy4wLjAuMSIKIyBwb3J0ID0gNTQzMgojIGRhdGFiYXNlID0gIm15YXBwX2RldmVsb3BtZW50IgojIHVzZXIgPSAibXl1c2VyIgojIHBhc3N3b3JkID0gInNlY3JldCIKCltbdGFibGVzXV0KdGFibGVfbmFtZSA9ICJjdXN0b21lciIKIyBzdHJ1Y3RfbmFtZSA9ICJDdXN0b21lciIK`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`config`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`ZnVuYyBDb3VudHt7LlN0cnVjdE5hbWV9fShkYiBRdWVyeWVyKSAoaW50NjQsIGVycm9yKSB7CiAgdmFyIG4gaW50NjQKICBzcWwgOj0gYHNlbGVjdCBjb3VudCgqKSBmcm9tICJ7ey5UYWJsZU5hbWV9fSJgCiAgZXJyIDo9IGRiLlF1ZXJ5Um93KHNxbCkuU2NhbigmbikKICByZXR1cm4gbiwgZXJyCn0K`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`count_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`Ly8gVGhpcyBmaWxlIGlzIGF1dG9tYXRpY2FsbHkgZ2VuZXJhdGVkLgpwYWNrYWdlIHt7LlBrZ05hbWV9fQoKaW1wb3J0ICgKICAiZXJyb3JzIgoKICAiZ2l0aHViLmNvbS9qYWNrYy9wZ3giCikKCmNvbnN0IFBHWERBVEFfVkVSU0lPTiA9ICJ7ey5WZXJzaW9ufX0iCgp2YXIgRXJyTm90Rm91bmQgPSBlcnJvcnMuTmV3KCJub3QgZm91bmQiKQoKdHlwZSBRdWVyeWVyIGludGVyZmFjZSB7CiAgUXVlcnkoc3FsIHN0cmluZywgYXJncyAuLi5pbnRlcmZhY2V7fSkgKCpwZ3guUm93cywgZXJyb3IpCiAgUXVlcnlSb3coc3FsIHN0cmluZywgYXJncyAuLi5pbnRlcmZhY2V7fSkgKnBneC5Sb3cKICBFeGVjKHNxbCBzdHJpbmcsIGFyZ3VtZW50cyAuLi5pbnRlcmZhY2V7fSkgKHBneC5Db21tYW5kVGFnLCBlcnJvcikKfQo=`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`db`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`ZnVuYyBEZWxldGV7ey5TdHJ1Y3ROYW1lfX0oZGIgUXVlcnllcnt7cmFuZ2UgLlByaW1hcnlLZXlDb2x1bW5zfX0sCiAge3suVmFyTmFtZX19IHt7LkdvVHlwZX19e3tlbmR9fSwKKSBlcnJvciB7CiAgYXJncyA6PSBwZ3guUXVlcnlBcmdzKG1ha2UoW11pbnRlcmZhY2V7fSwgMCwge3tsZW4gLlByaW1hcnlLZXlDb2x1bW5zfX0pKQoKICBzcWwgOj0gYGRlbGV0ZSBmcm9tICJ7ey5UYWJsZU5hbWV9fSIgd2hlcmUgYCB7eyByYW5nZSAkaSwgJGNvbHVtbiA6PSAuUHJpbWFyeUtleUNvbHVtbnN9fSArIGB7e2lmICRpfX0gYW5kIHt7ZW5kfX0ie3skY29sdW1uLkNvbHVtbk5hbWV9fSI9YCArIGFyZ3MuQXBwZW5kKHt7JGNvbHVtbi5WYXJOYW1lfX0pe3tlbmR9fQoKICBjb21tYW5kVGFnLCBlcnIgOj0gZGIuRXhlYyhzcWwsIGFyZ3MuLi4pCiAgaWYgZXJyICE9IG5pbCB7CiAgICByZXR1cm4gZXJyCiAgfQogIGlmIGNvbW1hbmRUYWcuUm93c0FmZmVjdGVkKCkgIT0gMSB7CiAgICByZXR1cm4gcGd4LkVyck5vUm93cwogIH0KICByZXR1cm4gbmlsCn0K`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`delete_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`ZnVuYyBJbnNlcnR7ey5TdHJ1Y3ROYW1lfX0oZGIgUXVlcnllciwgcm93ICp7ey5TdHJ1Y3ROYW1lfX0pIGVycm9yIHsKICBhcmdzIDo9IHBneC5RdWVyeUFyZ3MobWFrZShbXWludGVyZmFjZXt9LCAwLCB7e2xlbiAuQ29sdW1uc319KSkKCiAgdmFyIGNvbHVtbnMsIHZhbHVlcyBbXXN0cmluZwoKe3tyYW5nZSAuQ29sdW1uc319ICByb3cue3suRmllbGROYW1lfX0uYWRkSW5zZXJ0KGB7ey5Db2x1bW5OYW1lfX1gLCAmY29sdW1ucywgJnZhbHVlcywgJmFyZ3MpCnt7ZW5kfX0KCiAgc3FsIDo9IGBpbnNlcnQgaW50byAie3suVGFibGVOYW1lfX0iKGAgKyBzdHJpbmdzLkpvaW4oY29sdW1ucywgIiwgIikgKyBgKQp2YWx1ZXMoYCArIHN0cmluZ3MuSm9pbih2YWx1ZXMsICIsIikgKyBgKQpyZXR1cm5pbmcge3sgcmFuZ2UgJGksICRjb2x1bW4gOj0gLlByaW1hcnlLZXlDb2x1bW5zfX17e2lmICRpfX0sIHt7ZW5kfX0ie3skY29sdW1uLkNvbHVtbk5hbWV9fSJ7e2VuZH19CiAgYAoKICByZXR1cm4gZGIuUXVlcnlSb3coc3FsLCBhcmdzLi4uKS5TY2FuKHt7IHJhbmdlICRpLCAkY29sdW1uIDo9IC5QcmltYXJ5S2V5Q29sdW1uc319e3tpZiAkaX19LCB7e2VuZH19JnJvdy57eyRjb2x1bW4uRmllbGROYW1lfX17e2VuZH19KQp9Cg==`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`insert_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`Ly8gVGhpcyBmaWxlIGlzIGF1dG9tYXRpY2FsbHkgZ2VuZXJhdGVkLgpwYWNrYWdlIHt7LlBrZ05hbWV9fQoKaW1wb3J0ICgKICAic3RyaW5ncyIKCiAgImdpdGh1Yi5jb20vamFja2MvcGd4IgopCgp0eXBlIHt7LlN0cnVjdE5hbWV9fSBzdHJ1Y3Qgewp7e3JhbmdlIC5Db2x1bW5zfX0gIHt7LkZpZWxkTmFtZX19IHt7LkdvQm94VHlwZX19Cnt7ZW5kfX19Cgp7e3RlbXBsYXRlICJjb3VudF9mdW5jIiAufX0Ke3t0ZW1wbGF0ZSAic2VsZWN0X2FsbF9mdW5jIiAufX0Ke3t0ZW1wbGF0ZSAic2VsZWN0X2J5X3BrX2Z1bmMiIC59fQp7e3RlbXBsYXRlICJpbnNlcnRfZnVuYyIgLn19Cnt7dGVtcGxhdGUgInVwZGF0ZV9mdW5jIiAufX0Ke3t0ZW1wbGF0ZSAiZGVsZXRlX2Z1bmMiIC59fQo=`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`row`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`ZnVuYyBTZWxlY3RBbGx7ey5TdHJ1Y3ROYW1lfX0oZGIgUXVlcnllcikgKFtde3suU3RydWN0TmFtZX19LCBlcnJvcikgewogIHNxbCA6PSBgc2VsZWN0e3sgcmFuZ2UgJGksICRjb2x1bW4gOj0gLkNvbHVtbnN9fXt7aWYgJGl9fSx7e2VuZH19CiAgInt7JGNvbHVtbi5Db2x1bW5OYW1lfX0ie3tlbmR9fQpmcm9tICJ7ey5UYWJsZU5hbWV9fSJgCgogIHZhciByb3dzIFtde3suU3RydWN0TmFtZX19CgogIGRiUm93cywgZXJyIDo9IGRiLlF1ZXJ5KHNxbCkKICBpZiBlcnIgIT0gbmlsIHsKICAgIHJldHVybiBuaWwsIGVycgogIH0KCiAgZm9yIGRiUm93cy5OZXh0KCkgewogICAgdmFyIHJvdyB7ey5TdHJ1Y3ROYW1lfX0KICAgIGRiUm93cy5TY2FuKAp7e3JhbmdlIC5Db2x1bW5zfX0mcm93Lnt7LkZpZWxkTmFtZX19LAogICAge3tlbmR9fSkKICAgIHJvd3MgPSBhcHBlbmQocm93cywgcm93KQogIH0KCiAgaWYgZGJSb3dzLkVycigpICE9IG5pbCB7CiAgICByZXR1cm4gbmlsLCBkYlJvd3MuRXJyKCkKICB9CgogIHJldHVybiByb3dzLCBuaWwKfQo=`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`select_all_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`ZnVuYyBTZWxlY3R7ey5TdHJ1Y3ROYW1lfX1CeVBLKAogIGRiIFF1ZXJ5ZXJ7e3JhbmdlIC5QcmltYXJ5S2V5Q29sdW1uc319LAogIHt7LlZhck5hbWV9fSB7ey5Hb1R5cGV9fXt7ZW5kfX0sCikgKCp7ey5TdHJ1Y3ROYW1lfX0sIGVycm9yKSB7CiAgc3FsIDo9IGBzZWxlY3R7eyByYW5nZSAkaSwgJGNvbHVtbiA6PSAuQ29sdW1uc319e3tpZiAkaX19LHt7ZW5kfX0KICAie3skY29sdW1uLkNvbHVtbk5hbWV9fSJ7e2VuZH19CmZyb20gInt7LlRhYmxlTmFtZX19Igp3aGVyZSB7eyByYW5nZSAkaSwgJGNvbHVtbiA6PSAuUHJpbWFyeUtleUNvbHVtbnN9fXt7aWYgJGl9fSBhbmQge3tlbmR9fSJ7eyRjb2x1bW4uQ29sdW1uTmFtZX19Ij17e3BrUGxhY2Vob2xkZXIgJGl9fXt7ZW5kfX1gCgogIHZhciByb3cge3suU3RydWN0TmFtZX19CiAgZXJyIDo9IGRiLlF1ZXJ5Um93KHNxbCB7e3JhbmdlIC5QcmltYXJ5S2V5Q29sdW1uc319LCB7ey5WYXJOYW1lfX17e2VuZH19KS5TY2FuKAp7e3JhbmdlIC5Db2x1bW5zfX0mcm93Lnt7LkZpZWxkTmFtZX19LAogICAge3tlbmR9fSkKICBpZiBlcnIgPT0gcGd4LkVyck5vUm93cyB7CiAgICByZXR1cm4gbmlsLCBFcnJOb3RGb3VuZAogIH0gZWxzZSBpZiBlcnIgIT0gbmlsIHsKICAgIHJldHVybiBuaWwsIGVycgogIH0KCiAgcmV0dXJuICZyb3csIG5pbAp9Cg==`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`select_by_pk_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`ZnVuYyBVcGRhdGV7ey5TdHJ1Y3ROYW1lfX0oZGIgUXVlcnllcnt7cmFuZ2UgLlByaW1hcnlLZXlDb2x1bW5zfX0sCiAge3suVmFyTmFtZX19IHt7LkdvVHlwZX19e3tlbmR9fSwKICByb3cgKnt7LlN0cnVjdE5hbWV9fSwKKSBlcnJvciB7CiAgc2V0cyA6PSBtYWtlKFtdc3RyaW5nLCAwLCB7e2xlbiAuQ29sdW1uc319KQogIGFyZ3MgOj0gcGd4LlF1ZXJ5QXJncyhtYWtlKFtdaW50ZXJmYWNle30sIDAsIHt7bGVuIC5Db2x1bW5zfX0pKQoKe3tyYW5nZSAuQ29sdW1uc319ICByb3cue3suRmllbGROYW1lfX0uYWRkVXBkYXRlKGB7ey5Db2x1bW5OYW1lfX1gLCAmc2V0cywgJmFyZ3MpCnt7ZW5kfX0KCiAgaWYgbGVuKHNldHMpID09IDAgewogICAgcmV0dXJuIG5pbAogIH0KCiAgc3FsIDo9IGB1cGRhdGUgInt7LlRhYmxlTmFtZX19IiBzZXQgYCArIHN0cmluZ3MuSm9pbihzZXRzLCAiLCAiKSArIGAgd2hlcmUgYCB7eyByYW5nZSAkaSwgJGNvbHVtbiA6PSAuUHJpbWFyeUtleUNvbHVtbnN9fSArIGB7e2lmICRpfX0gYW5kIHt7ZW5kfX0ie3skY29sdW1uLkNvbHVtbk5hbWV9fSI9YCArIGFyZ3MuQXBwZW5kKHt7JGNvbHVtbi5WYXJOYW1lfX0pe3tlbmR9fQoKICBjb21tYW5kVGFnLCBlcnIgOj0gZGIuRXhlYyhzcWwsIGFyZ3MuLi4pCiAgaWYgZXJyICE9IG5pbCB7CiAgICByZXR1cm4gZXJyCiAgfQogIGlmIGNvbW1hbmRUYWcuUm93c0FmZmVjdGVkKCkgIT0gMSB7CiAgICByZXR1cm4gRXJyTm90Rm91bmQKICB9CiAgcmV0dXJuIG5pbAp9Cg==`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`update_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	return templates
}
