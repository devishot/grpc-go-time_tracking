package api

import (
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestTimeRecordMarshaling(t *testing.T) {
	r := &TimeRecord{}
	r.Amount = 30 // 30 minutes = 0.5 hour
	r.Description = "Read laws"

	out, err := proto.Marshal(r)
	if err != nil {
		t.Fatalf("Failed to encode time record: %s", err)
		return
	}

	r2 := &TimeRecord{}
	err2 := proto.Unmarshal(out, r2)
	if err2 != nil {
		t.Fatalf("Failed to decode time record: %s", err2)
	}

	t.Logf("decoded data: %s", r2.String())
}
