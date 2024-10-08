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

func easyjson7a32972cDecodeGithubComHanchonHanchondLibTypesWeb3(in *jlexer.Lexer, out *SendRawTransactionResponse) {
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
		case "error":
			(out.Error).UnmarshalEasyJSON(in)
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
func easyjson7a32972cEncodeGithubComHanchonHanchondLibTypesWeb3(out *jwriter.Writer, in SendRawTransactionResponse) {
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
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		(in.Error).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SendRawTransactionResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7a32972cEncodeGithubComHanchonHanchondLibTypesWeb3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SendRawTransactionResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7a32972cEncodeGithubComHanchonHanchondLibTypesWeb3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SendRawTransactionResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7a32972cDecodeGithubComHanchonHanchondLibTypesWeb3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SendRawTransactionResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7a32972cDecodeGithubComHanchonHanchondLibTypesWeb3(l, v)
}
func easyjson7a32972cDecodeGithubComHanchonHanchondLibTypesWeb31(in *jlexer.Lexer, out *BroadcastError) {
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
		case "code":
			out.Code = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "data":
			out.Data = string(in.String())
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
func easyjson7a32972cEncodeGithubComHanchonHanchondLibTypesWeb31(out *jwriter.Writer, in BroadcastError) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		out.String(string(in.Data))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BroadcastError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7a32972cEncodeGithubComHanchonHanchondLibTypesWeb31(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BroadcastError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7a32972cEncodeGithubComHanchonHanchondLibTypesWeb31(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BroadcastError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7a32972cDecodeGithubComHanchonHanchondLibTypesWeb31(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BroadcastError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7a32972cDecodeGithubComHanchonHanchondLibTypesWeb31(l, v)
}
