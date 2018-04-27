// Copyright (c) 2014-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
//Copyright (c) 2018 The box developers

package sebtcjson

// VersionResult models objects included in the version response.  In the actual
// result, these objects are keyed by the program or API name.
//
// NOTE: This is a btcsuite extension ported from
// github.com/decred/dcrd/dcrjson.
type VersionResult struct {
	VersionString string `json:"versionstring"`
	Major         uint32 `json:"major"`
	Minor         uint32 `json:"minor"`
	Patch         uint32 `json:"patch"`
	Prerelease    string `json:"prerelease"`
	BuildMetadata string `json:"buildmetadata"`
}
