// Copyright (c) 2014-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
//Copyright (c) 2018 The box developers

package sebtcjson

// SessionResult models the data from the session command.
type SessionResult struct {
	SessionID uint64 `json:"sessionid"`
}

// RescannedBlock contains the hash and all discovered transactions of a single
// rescanned block.
//
// NOTE: This is a btcsuite extension ported from
// github.com/decred/dcrd/dcrjson.
type RescannedBlock struct {
	Hash         string   `json:"hash"`
	Transactions []string `json:"transactions"`
}
