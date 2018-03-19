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

import (
	"encoding/json"
	"reflect"
	"testing"
)

// TestIsValidIDType ensures the IsValidIDType function behaves as expected.
func TestIsValidIDType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		id      interface{}
		isValid bool
	}{
		{"int", int(1), true},
		{"int8", int8(1), true},
		{"int16", int16(1), true},
		{"int32", int32(1), true},
		{"int64", int64(1), true},
		{"uint", uint(1), true},
		{"uint8", uint8(1), true},
		{"uint16", uint16(1), true},
		{"uint32", uint32(1), true},
		{"uint64", uint64(1), true},
		{"string", "1", true},
		{"nil", nil, true},
		{"float32", float32(1), true},
		{"float64", float64(1), true},
		{"bool", true, false},
		{"chan int", make(chan int), false},
		{"complex64", complex64(1), false},
		{"complex128", complex128(1), false},
		{"func", func() {}, false},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		if IsValidIDType(test.id) != test.isValid {
			t.Errorf("Test #%d (%s) valid mismatch - got %v, "+
				"want %v", i, test.name, !test.isValid,
				test.isValid)
			continue
		}
	}
}

// TestMarshalResponse ensures the MarshalResponse function works as expected.
func TestMarshalResponse(t *testing.T) {
	t.Parallel()

	testID := 1
	tests := []struct {
		name     string
		result   interface{}
		jsonErr  *RPCError
		expected []byte
	}{
		{
			name:     "ordinary bool result with no error",
			result:   true,
			jsonErr:  nil,
			expected: []byte(`{"result":true,"error":null,"id":1}`),
		},
		{
			name:   "result with error",
			result: nil,
			jsonErr: func() *RPCError {
				return NewRPCError(ErrRPCBlockNotFound, "123 not found")
			}(),
			expected: []byte(`{"result":null,"error":{"code":-5,"message":"123 not found"},"id":1}`),
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		_, _ = i, test
		marshalled, err := MarshalResponse(testID, test.result, test.jsonErr)
		if err != nil {
			t.Errorf("Test #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !reflect.DeepEqual(marshalled, test.expected) {
			t.Errorf("Test #%d (%s) mismatched result - got %s, "+
				"want %s", i, test.name, marshalled,
				test.expected)
		}
	}
}

// TestMiscErrors tests a few error conditions not covered elsewhere.
func TestMiscErrors(t *testing.T) {
	t.Parallel()

	// Force an error in NewRequest by giving it a parameter type that is
	// not supported.
	_, err := NewRequest(nil, "test", []interface{}{make(chan int)})
	if err == nil {
		t.Error("NewRequest: did not receive error")
		return
	}

	// Force an error in MarshalResponse by giving it an id type that is not
	// supported.
	wantErr := Error{ErrorCode: ErrInvalidType}
	_, err = MarshalResponse(make(chan int), nil, nil)
	if jerr, ok := err.(Error); !ok || jerr.ErrorCode != wantErr.ErrorCode {
		t.Errorf("MarshalResult: did not receive expected error - got "+
			"%v (%[1]T), want %v (%[2]T)", err, wantErr)
		return
	}

	// Force an error in MarshalResponse by giving it a result type that
	// can't be marshalled.
	_, err = MarshalResponse(1, make(chan int), nil)
	if _, ok := err.(*json.UnsupportedTypeError); !ok {
		wantErr := &json.UnsupportedTypeError{}
		t.Errorf("MarshalResult: did not receive expected error - got "+
			"%v (%[1]T), want %T", err, wantErr)
		return
	}
}

// TestRPCError tests the error output for the RPCError type.
func TestRPCError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   *RPCError
		want string
	}{
		{
			ErrRPCInvalidRequest,
			"-32600: Invalid request",
		},
		{
			ErrRPCMethodNotFound,
			"-32601: Method not found",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}
