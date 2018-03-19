// Copyright 2018 The box.la Authors All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sebtcjson

// Standard JSON-RPC 2.0 errors.
var (
	ErrRPCInvalidRequest = &RPCError{
		Code:    -32600,
		Message: "Invalid request",
	}
	ErrRPCMethodNotFound = &RPCError{
		Code:    -32601,
		Message: "Method not found",
	}
	ErrRPCInvalidParams = &RPCError{
		Code:    -32602,
		Message: "Invalid parameters",
	}
	ErrRPCInternal = &RPCError{
		Code:    -32603,
		Message: "Internal error",
	}
	ErrRPCParse = &RPCError{
		Code:    -32700,
		Message: "Parse error",
	}
)

// General application defined JSON errors.
const (
	ErrRPCMisc                RPCErrorCode = -1
	ErrRPCForbiddenBySafeMode RPCErrorCode = -2
	ErrRPCType                RPCErrorCode = -3
	ErrRPCInvalidAddressOrKey RPCErrorCode = -5
	ErrRPCOutOfMemory         RPCErrorCode = -7
	ErrRPCInvalidParameter    RPCErrorCode = -8
	ErrRPCDatabase            RPCErrorCode = -20
	ErrRPCDeserialization     RPCErrorCode = -22
	ErrRPCVerify              RPCErrorCode = -25
)

// Peer-to-peer client errors.
const (
	ErrRPCClientNotConnected      RPCErrorCode = -9
	ErrRPCClientInInitialDownload RPCErrorCode = -10
	ErrRPCClientNodeNotAdded      RPCErrorCode = -24
)

// Wallet JSON errors
const (
	ErrRPCWallet                    RPCErrorCode = -4
	ErrRPCWalletInsufficientFunds   RPCErrorCode = -6
	ErrRPCWalletInvalidAccountName  RPCErrorCode = -11
	ErrRPCWalletKeypoolRanOut       RPCErrorCode = -12
	ErrRPCWalletUnlockNeeded        RPCErrorCode = -13
	ErrRPCWalletPassphraseIncorrect RPCErrorCode = -14
	ErrRPCWalletWrongEncState       RPCErrorCode = -15
	ErrRPCWalletEncryptionFailed    RPCErrorCode = -16
	ErrRPCWalletAlreadyUnlocked     RPCErrorCode = -17
)

// Specific Errors related to commands.  These are the ones a user of the RPC
// server are most likely to see.  Generally, the codes should match one of the
// more general errors above.
const (
	ErrRPCBlockNotFound     RPCErrorCode = -5
	ErrRPCBlockCount        RPCErrorCode = -5
	ErrRPCBestBlockHash     RPCErrorCode = -5
	ErrRPCDifficulty        RPCErrorCode = -5
	ErrRPCOutOfRange        RPCErrorCode = -1
	ErrRPCNoTxInfo          RPCErrorCode = -5
	ErrRPCNoNewestBlockInfo RPCErrorCode = -5
	ErrRPCInvalidTxVout     RPCErrorCode = -5
	ErrRPCRawTxString       RPCErrorCode = -32602
	ErrRPCDecodeHexString   RPCErrorCode = -22
)

// Errors that are specific to btcd.
const (
	ErrRPCNoWallet      RPCErrorCode = -1
	ErrRPCUnimplemented RPCErrorCode = -1
)
