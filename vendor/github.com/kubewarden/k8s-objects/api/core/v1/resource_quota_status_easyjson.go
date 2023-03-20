// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonBe0fda7cDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *ResourceQuotaStatus) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "hard":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Hard = make(map[string]*resource.Quantity)
				} else {
					out.Hard = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 *resource.Quantity
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(resource.Quantity)
						}
						*v1 = resource.Quantity(in.String())
					}
					(out.Hard)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "used":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Used = make(map[string]*resource.Quantity)
				} else {
					out.Used = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 *resource.Quantity
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(resource.Quantity)
						}
						*v2 = resource.Quantity(in.String())
					}
					(out.Used)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBe0fda7cEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in ResourceQuotaStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Hard) != 0 {
		const prefix string = ",\"hard\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('{')
			v3First := true
			for v3Name, v3Value := range in.Hard {
				if v3First {
					v3First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v3Name))
				out.RawByte(':')
				if v3Value == nil {
					out.RawString("null")
				} else {
					out.String(string(*v3Value))
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Used) != 0 {
		const prefix string = ",\"used\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.Used {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				if v4Value == nil {
					out.RawString("null")
				} else {
					out.String(string(*v4Value))
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResourceQuotaStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBe0fda7cEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResourceQuotaStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBe0fda7cEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResourceQuotaStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBe0fda7cDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResourceQuotaStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBe0fda7cDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
