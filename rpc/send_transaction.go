package rpc

import (
	"context"
)

type SendTransactionResponse JsonRpcResponse[string]

type SendTransactionConfigEncoding string

const (
	SendTransactionConfigEncodingBase58 SendTransactionConfigEncoding = "base58"
	SendTransactionConfigEncodingBase64 SendTransactionConfigEncoding = "base64"
)

type SendTransactionConfig struct {
	SkipPreflight       bool                          `json:"skipPreflight,omitempty"`       // default: false
	PreflightCommitment Commitment                    `json:"preflightCommitment,omitempty"` // default: finalized
	Encoding            SendTransactionConfigEncoding `json:"encoding,omitempty"`            // default: base58
	MaxRetries          uint64                        `json:"maxRetries,omitempty"`
}

// SendTransaction submits a signed transaction to the cluster for processing
func (c *RpcClient) SendTransaction(ctx context.Context, tx string, params ...any) (JsonRpcResponse[string], error) {
	return call[JsonRpcResponse[string]](c, ctx, append([]any{"sendTransaction", tx}, params...)...)
}

// SendTransaction submits a signed transaction to the cluster for processing
func (c *RpcClient) SendTransactionWithConfig(ctx context.Context, tx string, cfg SendTransactionConfig, params ...any) (JsonRpcResponse[string], error) {
	return call[JsonRpcResponse[string]](c, ctx, append([]any{"sendTransaction", tx, cfg}, params...)...)
}

// SendTransaction submits a signed transaction to the cluster for processing
func (c *RpcClient) SendBundle(ctx context.Context, txs []string, params ...any) (JsonRpcResponse[string], error) {
	return call[JsonRpcResponse[string]](c, ctx, append([]any{"sendBundle", txs}, params...)...)
}
