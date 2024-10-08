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

func easyjson2ff71951DecodeGithubComHanchonHanchondLibTypesWeb3(in *jlexer.Lexer, out *BlockByNumberWithTransactions) {
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
			easyjson2ff71951Decode(in, &out.Result)
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
func easyjson2ff71951EncodeGithubComHanchonHanchondLibTypesWeb3(out *jwriter.Writer, in BlockByNumberWithTransactions) {
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
		easyjson2ff71951Encode(out, in.Result)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BlockByNumberWithTransactions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2ff71951EncodeGithubComHanchonHanchondLibTypesWeb3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BlockByNumberWithTransactions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2ff71951EncodeGithubComHanchonHanchondLibTypesWeb3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BlockByNumberWithTransactions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2ff71951DecodeGithubComHanchonHanchondLibTypesWeb3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BlockByNumberWithTransactions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2ff71951DecodeGithubComHanchonHanchondLibTypesWeb3(l, v)
}
func easyjson2ff71951Decode(in *jlexer.Lexer, out *struct {
	BaseFeePerGas   string `json:"baseFeePerGas"`
	Difficulty      string `json:"difficulty"`
	ExtraData       string `json:"extraData"`
	GasLimit        string `json:"gasLimit"`
	GasUsed         string `json:"gasUsed"`
	Hash            string `json:"hash"`
	LogsBloom       string `json:"logsBloom"`
	Miner           string `json:"miner"`
	MixHash         string `json:"mixHash"`
	Nonce           string `json:"nonce"`
	Number          string `json:"number"`
	ParentHash      string `json:"parentHash"`
	ReceiptsRoot    string `json:"receiptsRoot"`
	Sha3Uncles      string `json:"sha3Uncles"`
	Size            string `json:"size"`
	StateRoot       string `json:"stateRoot"`
	Timestamp       string `json:"timestamp"`
	TotalDifficulty string `json:"totalDifficulty"`
	Transactions    []struct {
		BlockHash        string `json:"blockHash"`
		BlockNumber      string `json:"blockNumber"`
		From             string `json:"from"`
		Gas              string `json:"gas"`
		GasPrice         string `json:"gasPrice"`
		Hash             string `json:"hash"`
		Input            string `json:"input"`
		Nonce            string `json:"nonce"`
		To               string `json:"to"`
		TransactionIndex string `json:"transactionIndex"`
		Value            string `json:"value"`
		Type             string `json:"type"`
		ChainID          string `json:"chainId"`
		V                string `json:"v"`
		R                string `json:"r"`
		S                string `json:"s"`
	} `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []interface{} `json:"uncles"`
}) {
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
		case "baseFeePerGas":
			out.BaseFeePerGas = string(in.String())
		case "difficulty":
			out.Difficulty = string(in.String())
		case "extraData":
			out.ExtraData = string(in.String())
		case "gasLimit":
			out.GasLimit = string(in.String())
		case "gasUsed":
			out.GasUsed = string(in.String())
		case "hash":
			out.Hash = string(in.String())
		case "logsBloom":
			out.LogsBloom = string(in.String())
		case "miner":
			out.Miner = string(in.String())
		case "mixHash":
			out.MixHash = string(in.String())
		case "nonce":
			out.Nonce = string(in.String())
		case "number":
			out.Number = string(in.String())
		case "parentHash":
			out.ParentHash = string(in.String())
		case "receiptsRoot":
			out.ReceiptsRoot = string(in.String())
		case "sha3Uncles":
			out.Sha3Uncles = string(in.String())
		case "size":
			out.Size = string(in.String())
		case "stateRoot":
			out.StateRoot = string(in.String())
		case "timestamp":
			out.Timestamp = string(in.String())
		case "totalDifficulty":
			out.TotalDifficulty = string(in.String())
		case "transactions":
			if in.IsNull() {
				in.Skip()
				out.Transactions = nil
			} else {
				in.Delim('[')
				if out.Transactions == nil {
					if !in.IsDelim(']') {
						out.Transactions = make([]struct {
							BlockHash        string `json:"blockHash"`
							BlockNumber      string `json:"blockNumber"`
							From             string `json:"from"`
							Gas              string `json:"gas"`
							GasPrice         string `json:"gasPrice"`
							Hash             string `json:"hash"`
							Input            string `json:"input"`
							Nonce            string `json:"nonce"`
							To               string `json:"to"`
							TransactionIndex string `json:"transactionIndex"`
							Value            string `json:"value"`
							Type             string `json:"type"`
							ChainID          string `json:"chainId"`
							V                string `json:"v"`
							R                string `json:"r"`
							S                string `json:"s"`
						}, 0, 0)
					} else {
						out.Transactions = []struct {
							BlockHash        string `json:"blockHash"`
							BlockNumber      string `json:"blockNumber"`
							From             string `json:"from"`
							Gas              string `json:"gas"`
							GasPrice         string `json:"gasPrice"`
							Hash             string `json:"hash"`
							Input            string `json:"input"`
							Nonce            string `json:"nonce"`
							To               string `json:"to"`
							TransactionIndex string `json:"transactionIndex"`
							Value            string `json:"value"`
							Type             string `json:"type"`
							ChainID          string `json:"chainId"`
							V                string `json:"v"`
							R                string `json:"r"`
							S                string `json:"s"`
						}{}
					}
				} else {
					out.Transactions = (out.Transactions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 struct {
						BlockHash        string `json:"blockHash"`
						BlockNumber      string `json:"blockNumber"`
						From             string `json:"from"`
						Gas              string `json:"gas"`
						GasPrice         string `json:"gasPrice"`
						Hash             string `json:"hash"`
						Input            string `json:"input"`
						Nonce            string `json:"nonce"`
						To               string `json:"to"`
						TransactionIndex string `json:"transactionIndex"`
						Value            string `json:"value"`
						Type             string `json:"type"`
						ChainID          string `json:"chainId"`
						V                string `json:"v"`
						R                string `json:"r"`
						S                string `json:"s"`
					}
					easyjson2ff71951Decode1(in, &v1)
					out.Transactions = append(out.Transactions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "transactionsRoot":
			out.TransactionsRoot = string(in.String())
		case "uncles":
			if in.IsNull() {
				in.Skip()
				out.Uncles = nil
			} else {
				in.Delim('[')
				if out.Uncles == nil {
					if !in.IsDelim(']') {
						out.Uncles = make([]interface{}, 0, 4)
					} else {
						out.Uncles = []interface{}{}
					}
				} else {
					out.Uncles = (out.Uncles)[:0]
				}
				for !in.IsDelim(']') {
					var v2 interface{}
					if m, ok := v2.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v2.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v2 = in.Interface()
					}
					out.Uncles = append(out.Uncles, v2)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson2ff71951Encode(out *jwriter.Writer, in struct {
	BaseFeePerGas   string `json:"baseFeePerGas"`
	Difficulty      string `json:"difficulty"`
	ExtraData       string `json:"extraData"`
	GasLimit        string `json:"gasLimit"`
	GasUsed         string `json:"gasUsed"`
	Hash            string `json:"hash"`
	LogsBloom       string `json:"logsBloom"`
	Miner           string `json:"miner"`
	MixHash         string `json:"mixHash"`
	Nonce           string `json:"nonce"`
	Number          string `json:"number"`
	ParentHash      string `json:"parentHash"`
	ReceiptsRoot    string `json:"receiptsRoot"`
	Sha3Uncles      string `json:"sha3Uncles"`
	Size            string `json:"size"`
	StateRoot       string `json:"stateRoot"`
	Timestamp       string `json:"timestamp"`
	TotalDifficulty string `json:"totalDifficulty"`
	Transactions    []struct {
		BlockHash        string `json:"blockHash"`
		BlockNumber      string `json:"blockNumber"`
		From             string `json:"from"`
		Gas              string `json:"gas"`
		GasPrice         string `json:"gasPrice"`
		Hash             string `json:"hash"`
		Input            string `json:"input"`
		Nonce            string `json:"nonce"`
		To               string `json:"to"`
		TransactionIndex string `json:"transactionIndex"`
		Value            string `json:"value"`
		Type             string `json:"type"`
		ChainID          string `json:"chainId"`
		V                string `json:"v"`
		R                string `json:"r"`
		S                string `json:"s"`
	} `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []interface{} `json:"uncles"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"baseFeePerGas\":"
		out.RawString(prefix[1:])
		out.String(string(in.BaseFeePerGas))
	}
	{
		const prefix string = ",\"difficulty\":"
		out.RawString(prefix)
		out.String(string(in.Difficulty))
	}
	{
		const prefix string = ",\"extraData\":"
		out.RawString(prefix)
		out.String(string(in.ExtraData))
	}
	{
		const prefix string = ",\"gasLimit\":"
		out.RawString(prefix)
		out.String(string(in.GasLimit))
	}
	{
		const prefix string = ",\"gasUsed\":"
		out.RawString(prefix)
		out.String(string(in.GasUsed))
	}
	{
		const prefix string = ",\"hash\":"
		out.RawString(prefix)
		out.String(string(in.Hash))
	}
	{
		const prefix string = ",\"logsBloom\":"
		out.RawString(prefix)
		out.String(string(in.LogsBloom))
	}
	{
		const prefix string = ",\"miner\":"
		out.RawString(prefix)
		out.String(string(in.Miner))
	}
	{
		const prefix string = ",\"mixHash\":"
		out.RawString(prefix)
		out.String(string(in.MixHash))
	}
	{
		const prefix string = ",\"nonce\":"
		out.RawString(prefix)
		out.String(string(in.Nonce))
	}
	{
		const prefix string = ",\"number\":"
		out.RawString(prefix)
		out.String(string(in.Number))
	}
	{
		const prefix string = ",\"parentHash\":"
		out.RawString(prefix)
		out.String(string(in.ParentHash))
	}
	{
		const prefix string = ",\"receiptsRoot\":"
		out.RawString(prefix)
		out.String(string(in.ReceiptsRoot))
	}
	{
		const prefix string = ",\"sha3Uncles\":"
		out.RawString(prefix)
		out.String(string(in.Sha3Uncles))
	}
	{
		const prefix string = ",\"size\":"
		out.RawString(prefix)
		out.String(string(in.Size))
	}
	{
		const prefix string = ",\"stateRoot\":"
		out.RawString(prefix)
		out.String(string(in.StateRoot))
	}
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix)
		out.String(string(in.Timestamp))
	}
	{
		const prefix string = ",\"totalDifficulty\":"
		out.RawString(prefix)
		out.String(string(in.TotalDifficulty))
	}
	{
		const prefix string = ",\"transactions\":"
		out.RawString(prefix)
		if in.Transactions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Transactions {
				if v3 > 0 {
					out.RawByte(',')
				}
				easyjson2ff71951Encode1(out, v4)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"transactionsRoot\":"
		out.RawString(prefix)
		out.String(string(in.TransactionsRoot))
	}
	{
		const prefix string = ",\"uncles\":"
		out.RawString(prefix)
		if in.Uncles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Uncles {
				if v5 > 0 {
					out.RawByte(',')
				}
				if m, ok := v6.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v6.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v6))
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson2ff71951Decode1(in *jlexer.Lexer, out *struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	Type             string `json:"type"`
	ChainID          string `json:"chainId"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}) {
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
		case "blockHash":
			out.BlockHash = string(in.String())
		case "blockNumber":
			out.BlockNumber = string(in.String())
		case "from":
			out.From = string(in.String())
		case "gas":
			out.Gas = string(in.String())
		case "gasPrice":
			out.GasPrice = string(in.String())
		case "hash":
			out.Hash = string(in.String())
		case "input":
			out.Input = string(in.String())
		case "nonce":
			out.Nonce = string(in.String())
		case "to":
			out.To = string(in.String())
		case "transactionIndex":
			out.TransactionIndex = string(in.String())
		case "value":
			out.Value = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "chainId":
			out.ChainID = string(in.String())
		case "v":
			out.V = string(in.String())
		case "r":
			out.R = string(in.String())
		case "s":
			out.S = string(in.String())
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
func easyjson2ff71951Encode1(out *jwriter.Writer, in struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	Type             string `json:"type"`
	ChainID          string `json:"chainId"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"blockHash\":"
		out.RawString(prefix[1:])
		out.String(string(in.BlockHash))
	}
	{
		const prefix string = ",\"blockNumber\":"
		out.RawString(prefix)
		out.String(string(in.BlockNumber))
	}
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix)
		out.String(string(in.From))
	}
	{
		const prefix string = ",\"gas\":"
		out.RawString(prefix)
		out.String(string(in.Gas))
	}
	{
		const prefix string = ",\"gasPrice\":"
		out.RawString(prefix)
		out.String(string(in.GasPrice))
	}
	{
		const prefix string = ",\"hash\":"
		out.RawString(prefix)
		out.String(string(in.Hash))
	}
	{
		const prefix string = ",\"input\":"
		out.RawString(prefix)
		out.String(string(in.Input))
	}
	{
		const prefix string = ",\"nonce\":"
		out.RawString(prefix)
		out.String(string(in.Nonce))
	}
	{
		const prefix string = ",\"to\":"
		out.RawString(prefix)
		out.String(string(in.To))
	}
	{
		const prefix string = ",\"transactionIndex\":"
		out.RawString(prefix)
		out.String(string(in.TransactionIndex))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.String(string(in.Value))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"chainId\":"
		out.RawString(prefix)
		out.String(string(in.ChainID))
	}
	{
		const prefix string = ",\"v\":"
		out.RawString(prefix)
		out.String(string(in.V))
	}
	{
		const prefix string = ",\"r\":"
		out.RawString(prefix)
		out.String(string(in.R))
	}
	{
		const prefix string = ",\"s\":"
		out.RawString(prefix)
		out.String(string(in.S))
	}
	out.RawByte('}')
}
func easyjson2ff71951DecodeGithubComHanchonHanchondLibTypesWeb31(in *jlexer.Lexer, out *BlockByNumber) {
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
			easyjson2ff71951Decode2(in, &out.Result)
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
func easyjson2ff71951EncodeGithubComHanchonHanchondLibTypesWeb31(out *jwriter.Writer, in BlockByNumber) {
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
		easyjson2ff71951Encode2(out, in.Result)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BlockByNumber) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2ff71951EncodeGithubComHanchonHanchondLibTypesWeb31(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BlockByNumber) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2ff71951EncodeGithubComHanchonHanchondLibTypesWeb31(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BlockByNumber) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2ff71951DecodeGithubComHanchonHanchondLibTypesWeb31(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BlockByNumber) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2ff71951DecodeGithubComHanchonHanchondLibTypesWeb31(l, v)
}
func easyjson2ff71951Decode2(in *jlexer.Lexer, out *struct {
	BaseFeePerGas    string        `json:"baseFeePerGas"`
	Difficulty       string        `json:"difficulty"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	Hash             string        `json:"hash"`
	LogsBloom        string        `json:"logsBloom"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	Nonce            string        `json:"nonce"`
	Number           string        `json:"number"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	Size             string        `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        string        `json:"timestamp"`
	TotalDifficulty  string        `json:"totalDifficulty"`
	Transactions     []string      `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []interface{} `json:"uncles"`
}) {
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
		case "baseFeePerGas":
			out.BaseFeePerGas = string(in.String())
		case "difficulty":
			out.Difficulty = string(in.String())
		case "extraData":
			out.ExtraData = string(in.String())
		case "gasLimit":
			out.GasLimit = string(in.String())
		case "gasUsed":
			out.GasUsed = string(in.String())
		case "hash":
			out.Hash = string(in.String())
		case "logsBloom":
			out.LogsBloom = string(in.String())
		case "miner":
			out.Miner = string(in.String())
		case "mixHash":
			out.MixHash = string(in.String())
		case "nonce":
			out.Nonce = string(in.String())
		case "number":
			out.Number = string(in.String())
		case "parentHash":
			out.ParentHash = string(in.String())
		case "receiptsRoot":
			out.ReceiptsRoot = string(in.String())
		case "sha3Uncles":
			out.Sha3Uncles = string(in.String())
		case "size":
			out.Size = string(in.String())
		case "stateRoot":
			out.StateRoot = string(in.String())
		case "timestamp":
			out.Timestamp = string(in.String())
		case "totalDifficulty":
			out.TotalDifficulty = string(in.String())
		case "transactions":
			if in.IsNull() {
				in.Skip()
				out.Transactions = nil
			} else {
				in.Delim('[')
				if out.Transactions == nil {
					if !in.IsDelim(']') {
						out.Transactions = make([]string, 0, 4)
					} else {
						out.Transactions = []string{}
					}
				} else {
					out.Transactions = (out.Transactions)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.Transactions = append(out.Transactions, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "transactionsRoot":
			out.TransactionsRoot = string(in.String())
		case "uncles":
			if in.IsNull() {
				in.Skip()
				out.Uncles = nil
			} else {
				in.Delim('[')
				if out.Uncles == nil {
					if !in.IsDelim(']') {
						out.Uncles = make([]interface{}, 0, 4)
					} else {
						out.Uncles = []interface{}{}
					}
				} else {
					out.Uncles = (out.Uncles)[:0]
				}
				for !in.IsDelim(']') {
					var v8 interface{}
					if m, ok := v8.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v8.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v8 = in.Interface()
					}
					out.Uncles = append(out.Uncles, v8)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson2ff71951Encode2(out *jwriter.Writer, in struct {
	BaseFeePerGas    string        `json:"baseFeePerGas"`
	Difficulty       string        `json:"difficulty"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	Hash             string        `json:"hash"`
	LogsBloom        string        `json:"logsBloom"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	Nonce            string        `json:"nonce"`
	Number           string        `json:"number"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	Size             string        `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        string        `json:"timestamp"`
	TotalDifficulty  string        `json:"totalDifficulty"`
	Transactions     []string      `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []interface{} `json:"uncles"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"baseFeePerGas\":"
		out.RawString(prefix[1:])
		out.String(string(in.BaseFeePerGas))
	}
	{
		const prefix string = ",\"difficulty\":"
		out.RawString(prefix)
		out.String(string(in.Difficulty))
	}
	{
		const prefix string = ",\"extraData\":"
		out.RawString(prefix)
		out.String(string(in.ExtraData))
	}
	{
		const prefix string = ",\"gasLimit\":"
		out.RawString(prefix)
		out.String(string(in.GasLimit))
	}
	{
		const prefix string = ",\"gasUsed\":"
		out.RawString(prefix)
		out.String(string(in.GasUsed))
	}
	{
		const prefix string = ",\"hash\":"
		out.RawString(prefix)
		out.String(string(in.Hash))
	}
	{
		const prefix string = ",\"logsBloom\":"
		out.RawString(prefix)
		out.String(string(in.LogsBloom))
	}
	{
		const prefix string = ",\"miner\":"
		out.RawString(prefix)
		out.String(string(in.Miner))
	}
	{
		const prefix string = ",\"mixHash\":"
		out.RawString(prefix)
		out.String(string(in.MixHash))
	}
	{
		const prefix string = ",\"nonce\":"
		out.RawString(prefix)
		out.String(string(in.Nonce))
	}
	{
		const prefix string = ",\"number\":"
		out.RawString(prefix)
		out.String(string(in.Number))
	}
	{
		const prefix string = ",\"parentHash\":"
		out.RawString(prefix)
		out.String(string(in.ParentHash))
	}
	{
		const prefix string = ",\"receiptsRoot\":"
		out.RawString(prefix)
		out.String(string(in.ReceiptsRoot))
	}
	{
		const prefix string = ",\"sha3Uncles\":"
		out.RawString(prefix)
		out.String(string(in.Sha3Uncles))
	}
	{
		const prefix string = ",\"size\":"
		out.RawString(prefix)
		out.String(string(in.Size))
	}
	{
		const prefix string = ",\"stateRoot\":"
		out.RawString(prefix)
		out.String(string(in.StateRoot))
	}
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix)
		out.String(string(in.Timestamp))
	}
	{
		const prefix string = ",\"totalDifficulty\":"
		out.RawString(prefix)
		out.String(string(in.TotalDifficulty))
	}
	{
		const prefix string = ",\"transactions\":"
		out.RawString(prefix)
		if in.Transactions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Transactions {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"transactionsRoot\":"
		out.RawString(prefix)
		out.String(string(in.TransactionsRoot))
	}
	{
		const prefix string = ",\"uncles\":"
		out.RawString(prefix)
		if in.Uncles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Uncles {
				if v11 > 0 {
					out.RawByte(',')
				}
				if m, ok := v12.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v12.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v12))
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
