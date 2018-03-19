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

package serpcclient

import (
	"encoding/json"
	"errors"
	"github.com/boxproject/lib-bitcore/sebtcjson"
)

// FutureRawResult is a future promise to deliver the result of a RawRequest RPC
// invocation (or an applicable error).
type FutureRawResult chan *response

// Receive waits for the response promised by the future and returns the raw
// response, or an error if the request was unsuccessful.
func (r FutureRawResult) Receive() (json.RawMessage, error) {
	return receiveFuture(r)
}

// RawRequestAsync returns an instance of a type that can be used to get the
// result of a custom RPC request at some future time by invoking the Receive
// function on the returned instance.
//
// See RawRequest for the blocking version and more details.
func (c *Client) RawRequestAsync(method string, params []json.RawMessage) FutureRawResult {
	// Method may not be empty.
	if method == "" {
		return newFutureError(errors.New("no method"))
	}

	// Marshal parameters as "[]" instead of "null" when no parameters
	// are passed.
	if params == nil {
		params = []json.RawMessage{}
	}

	// Create a raw JSON-RPC request using the provided method and params
	// and marshal it.  This is done rather than using the sendCmd function
	// since that relies on marshalling registered btcjson commands rather
	// than custom commands.
	id := c.NextID()
	rawRequest := &sebtcjson.Request{
		Jsonrpc: "1.0",
		ID:      id,
		Method:  method,
		Params:  params,
	}
	marshalledJSON, err := json.Marshal(rawRequest)
	if err != nil {
		return newFutureError(err)
	}

	// Generate the request and send it along with a channel to respond on.
	responseChan := make(chan *response, 1)
	jReq := &jsonRequest{
		id:             id,
		method:         method,
		cmd:            nil,
		marshalledJSON: marshalledJSON,
		responseChan:   responseChan,
	}
	c.sendRequest(jReq)

	return responseChan
}

// RawRequest allows the caller to send a raw or custom request to the server.
// This method may be used to send and receive requests and responses for
// requests that are not handled by this client package, or to proxy partially
// unmarshaled requests to another JSON-RPC server if a request cannot be
// handled directly.
func (c *Client) RawRequest(method string, params []json.RawMessage) (json.RawMessage, error) {
	return c.RawRequestAsync(method, params).Receive()
}
