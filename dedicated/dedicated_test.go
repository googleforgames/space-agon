package main

import (
	"testing"

	"github.com/googleforgames/space-agon/game/pb"
	"github.com/stretchr/testify/assert"
)

func TestIsMemoRecipient(t *testing.T) {

	tests := []struct {
		cid              int64
		memo             *pb.Memo
		name             string
		isPanic          bool
		correctRecipient bool
	}{
		{
			cid: int64(1),
			memo: &pb.Memo{
				Recipient: &pb.Memo_To{
					To: int64(1),
				},
			},
			name:             "correct Memo_to recipient",
			isPanic:          false,
			correctRecipient: true,
		},
		{
			cid: int64(1),
			memo: &pb.Memo{
				Recipient: &pb.Memo_To{
					To: int64(2),
				},
			},
			name:             "incorrect Memo_to recipient",
			isPanic:          false,
			correctRecipient: false,
		},
		{
			cid: int64(2),
			memo: &pb.Memo{
				Recipient: &pb.Memo_EveryoneBut{
					EveryoneBut: int64(1),
				},
			},
			name:             "correct Memo_EveryoneBut recipient",
			isPanic:          false,
			correctRecipient: true,
		},
		{
			cid: int64(1),
			memo: &pb.Memo{
				Recipient: &pb.Memo_EveryoneBut{
					EveryoneBut: int64(1),
				},
			},
			name:             "incorrect Memo_EveryoneBut recipient",
			isPanic:          false,
			correctRecipient: false,
		},
		{
			cid:              int64(12),
			memo:             &pb.Memo{},
			name:             "unknown memo recipient",
			isPanic:          true,
			correctRecipient: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isPanic {
				assert.Panics(t, func() {
					isMemoRecipient(test.cid, test.memo)
				})
			} else if test.correctRecipient {
				assert.True(t, isMemoRecipient(test.cid, test.memo))
			} else {
				assert.False(t, isMemoRecipient(test.cid, test.memo))
			}
		})
	}
}
