package commit

import (
	"testing"
	"time"

	"github.com/containerd/nerdctl/v2/pkg/api/types"
)

func TestCommitLeaseDuration(t *testing.T) {
	tests := []struct {
		name  string
		value time.Duration
		want  time.Duration
	}{
		{name: "default", want: types.DefaultCommitLeaseDuration},
		{name: "zero", value: 0, want: types.DefaultCommitLeaseDuration},
		{name: "custom", value: 12 * time.Hour, want: 12 * time.Hour},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := commitLeaseDuration(&Opts{
				DevboxOptions: types.DevboxOptions{CommitLeaseDuration: tt.value},
			})
			if got != tt.want {
				t.Fatalf("commitLeaseDuration() = %s, want %s", got, tt.want)
			}
		})
	}
}
