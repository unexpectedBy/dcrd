// Copyright (c) 2014-2015 The btcsuite developers
// Copyright (c) 2015 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/decred/dcrd/wire"
)

// TestFilterCLearLatest tests the MsgFilterClear API against the latest
// protocol version.
func TestFilterClearLatest(t *testing.T) {
	pver := wire.ProtocolVersion

	msg := wire.NewMsgFilterClear()

	// Ensure the command is expected value.
	wantCmd := "filterclear"
	if cmd := msg.Command(); cmd != wantCmd {
		t.Errorf("NewMsgFilterClear: wrong command - got %v want %v",
			cmd, wantCmd)
	}

	// Ensure max payload is expected value for latest protocol version.
	wantPayload := uint32(0)
	maxPayload := msg.MaxPayloadLength(pver)
	if maxPayload != wantPayload {
		t.Errorf("MaxPayloadLength: wrong max payload length for "+
			"protocol version %d - got %v, want %v", pver,
			maxPayload, wantPayload)
	}

	return
}

// TestFilterClearWire tests the MsgFilterClear wire encode and decode for
// various protocol versions.
func TestFilterClearWire(t *testing.T) {
	msgFilterClear := wire.NewMsgFilterClear()
	msgFilterClearEncoded := []byte{}

	tests := []struct {
		in   *wire.MsgFilterClear // Message to encode
		out  *wire.MsgFilterClear // Expected decoded message
		buf  []byte               // Wire encoding
		pver uint32               // Protocol version for wire encoding
	}{
		// Latest protocol version.
		{
			msgFilterClear,
			msgFilterClear,
			msgFilterClearEncoded,
			wire.ProtocolVersion,
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		// Encode the message to wire format.
		var buf bytes.Buffer
		err := test.in.BtcEncode(&buf, test.pver)
		if err != nil {
			t.Errorf("BtcEncode #%d error %v", i, err)
			continue
		}
		if !bytes.Equal(buf.Bytes(), test.buf) {
			t.Errorf("BtcEncode #%d\n got: %s want: %s", i,
				spew.Sdump(buf.Bytes()), spew.Sdump(test.buf))
			continue
		}

		// Decode the message from wire format.
		var msg wire.MsgFilterClear
		rbuf := bytes.NewReader(test.buf)
		err = msg.BtcDecode(rbuf, test.pver)
		if err != nil {
			t.Errorf("BtcDecode #%d error %v", i, err)
			continue
		}
		if !reflect.DeepEqual(&msg, test.out) {
			t.Errorf("BtcDecode #%d\n got: %s want: %s", i,
				spew.Sdump(msg), spew.Sdump(test.out))
			continue
		}
	}
}
