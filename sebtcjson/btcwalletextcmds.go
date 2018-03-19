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

// NOTE: This file is intended to house the RPC commands that are supported by
// a wallet server with btcwallet extensions.

package sebtcjson

// CreateNewAccountCmd defines the createnewaccount JSON-RPC command.
type CreateNewAccountCmd struct {
	Account string
}

// NewCreateNewAccountCmd returns a new instance which can be used to issue a
// createnewaccount JSON-RPC command.
func NewCreateNewAccountCmd(account string) *CreateNewAccountCmd {
	return &CreateNewAccountCmd{
		Account: account,
	}
}

// DumpWalletCmd defines the dumpwallet JSON-RPC command.
type DumpWalletCmd struct {
	Filename string
}

// NewDumpWalletCmd returns a new instance which can be used to issue a
// dumpwallet JSON-RPC command.
func NewDumpWalletCmd(filename string) *DumpWalletCmd {
	return &DumpWalletCmd{
		Filename: filename,
	}
}

// ImportAddressCmd defines the importaddress JSON-RPC command.
type ImportAddressCmd struct {
	Address string
	Lable string
	Rescan  *bool `jsonrpcdefault:"true"`
}

// NewImportAddressCmd returns a new instance which can be used to issue an
// importaddress JSON-RPC command.
func NewImportAddressCmd(address,lable string, rescan *bool) *ImportAddressCmd {
	return &ImportAddressCmd{
		Address: address,
		Lable:lable,
		Rescan:  rescan,
	}
}

// ImportPubKeyCmd defines the importpubkey JSON-RPC command.
type ImportPubKeyCmd struct {
	PubKey string
	Rescan *bool `jsonrpcdefault:"true"`
}

// NewImportPubKeyCmd returns a new instance which can be used to issue an
// importpubkey JSON-RPC command.
func NewImportPubKeyCmd(pubKey string, rescan *bool) *ImportPubKeyCmd {
	return &ImportPubKeyCmd{
		PubKey: pubKey,
		Rescan: rescan,
	}
}

// ImportWalletCmd defines the importwallet JSON-RPC command.
type ImportWalletCmd struct {
	Filename string
}

// NewImportWalletCmd returns a new instance which can be used to issue a
// importwallet JSON-RPC command.
func NewImportWalletCmd(filename string) *ImportWalletCmd {
	return &ImportWalletCmd{
		Filename: filename,
	}
}

// RenameAccountCmd defines the renameaccount JSON-RPC command.
type RenameAccountCmd struct {
	OldAccount string
	NewAccount string
}

// NewRenameAccountCmd returns a new instance which can be used to issue a
// renameaccount JSON-RPC command.
func NewRenameAccountCmd(oldAccount, newAccount string) *RenameAccountCmd {
	return &RenameAccountCmd{
		OldAccount: oldAccount,
		NewAccount: newAccount,
	}
}

func init() {
	// The commands in this file are only usable with a wallet server.
	flags := UFWalletOnly

	MustRegisterCmd("createnewaccount", (*CreateNewAccountCmd)(nil), flags)
	MustRegisterCmd("dumpwallet", (*DumpWalletCmd)(nil), flags)
	MustRegisterCmd("importaddress", (*ImportAddressCmd)(nil), flags)
	MustRegisterCmd("importpubkey", (*ImportPubKeyCmd)(nil), flags)
	MustRegisterCmd("importwallet", (*ImportWalletCmd)(nil), flags)
	MustRegisterCmd("renameaccount", (*RenameAccountCmd)(nil), flags)
}
