package internal

import (
	"context"
	"encoding/json"
	"runtime"
)

const (
	SDKSemverVersion = "v0.0.1"
	SDKLanguage      = "Go"
)

type Core interface {
	Init(ctx context.Context) error
	Invoke(ctx context.Context, invocation Invocation) ([]byte, error)
}

// SDKRuntimeInfo contains information required for sdk runtime
type SDKRuntimeInfo struct {
	SDKVersion         string `json:"sdkVersion"`
	SDKLanguage        string `json:"sdkLanguage"`
	SDKLanguageVersion string `json:"sdkLanguageVersion"`
	SystemOS           string `json:"os"`
	SystemOSVersion    string `json:"osVersion"`
	SystemArch         string `json:"architecture"`
}

func NewDefaultSDKRuntimeInfo() SDKRuntimeInfo {
	const defaultOSVersion = "0.0.0"
	return SDKRuntimeInfo{
		SDKVersion:         SDKSemverVersion,
		SDKLanguage:        SDKLanguage,
		SDKLanguageVersion: runtime.Version(),
		SystemOS:           runtime.GOOS,
		SystemArch:         runtime.GOARCH,
		SystemOSVersion:    defaultOSVersion,
	}
}

// Invocation holds the information required for invoking SDK functionality.
type Invocation struct {
	Method string `json:"method"`
	Args   []any  `json:"args,omitempty"`
}

// InnerSDK represents the sdk-core client on which calls will be made.
type InnerSDK struct {
	Core Core
}

func (sdk InnerSDK) Invoke(ctx context.Context, method string, args []any, returnValue any) error {
	invocationResponse, err := sdk.Core.Invoke(ctx, Invocation{
		Method: method,
		Args:   args,
	})
	if err != nil {
		return err
	}
	if err = json.Unmarshal(invocationResponse, returnValue); err != nil {
		return err
	}
	return nil
}
