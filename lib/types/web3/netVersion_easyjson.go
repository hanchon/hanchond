// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package web3

import (
	json "encoding/json"
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

func easyjsonB4c28bdfDecodeGithubComHanchonHanchondLibTypesWeb3(in *jlexer.Lexer, out *NetVersionResponse) {
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
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		case "id":
			out.ID = int(in.Int())
		case "result":
			out.Result = string(in.String())
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
func easyjsonB4c28bdfEncodeGithubComHanchonHanchondLibTypesWeb3(out *jwriter.Writer, in NetVersionResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"jsonrpc\":"
		out.RawString(prefix[1:])
		out.String(string(in.Jsonrpc))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"result\":"
		out.RawString(prefix)
		out.String(string(in.Result))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NetVersionResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB4c28bdfEncodeGithubComHanchonHanchondLibTypesWeb3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NetVersionResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB4c28bdfEncodeGithubComHanchonHanchondLibTypesWeb3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NetVersionResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB4c28bdfDecodeGithubComHanchonHanchondLibTypesWeb3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NetVersionResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB4c28bdfDecodeGithubComHanchonHanchondLibTypesWeb3(l, v)
}
