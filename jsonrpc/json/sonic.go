//go:build sonic && avx && (linux || windows || darwin) && amd64

package json

import (
	jsonStd "encoding/json"
	"github.com/bytedance/sonic"
)

var (
	json          = sonic.ConfigFastest
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
)

type RawMessage = jsonStd.RawMessage
