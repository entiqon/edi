package po_test

import (
	"testing"

	"github.com/entiqon/edi/canonical/po"
)

type statusCase struct {
	status  po.Status
	text    string
	isValid bool
	aliases []string
}

var statusCases = []statusCase{
	{po.Status(-1), "UNKNOWN", false, nil},

	{po.StatusUnknown, "UNKNOWN", false, []string{
		"unknown",
		" UNKNOWN ",
		"",
		"nope",
	}},

	{po.StatusAccepted, "ACCEPTED", true, []string{
		"accepted",
		" accepted ",
	}},

	{po.StatusPartial, "PARTIAL", true, []string{
		"partial",
		" PARTIAL ",
	}},

	{po.StatusRejected, "REJECTED", true, []string{
		"rejected",
		" REJECTED ",
	}},

	{po.Status(999), "UNKNOWN", false, nil},
}

func TestStatus(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		for _, tc := range statusCases {
			if got := tc.status.String(); got != tc.text {
				t.Fatalf("String(%v) = %q, want %q", tc.status, got, tc.text)
			}
		}
	})

	t.Run("IsValid", func(t *testing.T) {
		for _, tc := range statusCases {
			if got := tc.status.IsValid(); got != tc.isValid {
				t.Fatalf("IsValid(%v) = %v, want %v", tc.status, got, tc.isValid)
			}
		}
	})

	t.Run("MarshalText", func(t *testing.T) {
		for _, tc := range statusCases {
			b, err := tc.status.MarshalText()
			if err != nil {
				t.Fatalf("MarshalText(%v) error = %v", tc.status, err)
			}
			if got := string(b); got != tc.text {
				t.Fatalf("MarshalText(%v) = %q, want %q", tc.status, got, tc.text)
			}
		}
	})

	t.Run("UnmarshalText", func(t *testing.T) {
		t.Run("Default", func(t *testing.T) {
			for _, tc := range statusCases {
				want := tc.status
				if !tc.status.IsValid() {
					want = po.StatusUnknown
				}

				inputs := append([]string{tc.text}, tc.aliases...)
				for _, in := range inputs {
					var s po.Status
					if err := s.UnmarshalText([]byte(in)); err != nil {
						t.Fatalf("UnmarshalText(%q) error = %v", in, err)
					}
					if s != want {
						t.Fatalf("UnmarshalText(%q) = %v, want %v", in, s, want)
					}
				}
			}
		})

		t.Run("NilReceiver", func(t *testing.T) {
			var s *po.Status
			if err := s.UnmarshalText([]byte("ACCEPTED")); err == nil {
				t.Fatalf("expected error for nil receiver, got nil")
			}
		})
	})

	t.Run("StatusFrom", func(t *testing.T) {
		for _, tc := range statusCases {
			want := tc.status
			if !tc.status.IsValid() {
				want = po.StatusUnknown
			}

			inputs := []any{
				tc.status,
				int(tc.status),
				tc.text,
			}
			for _, a := range tc.aliases {
				inputs = append(inputs, a)
			}

			for _, in := range inputs {
				if got := po.StatusFrom(in); got != want {
					t.Fatalf("StatusFrom(%v) = %v, want %v", in, got, want)
				}
			}
		}

		// Unsupported types should always fall back to StatusUnknown.
		extras := []any{true, nil, 3.14}
		for _, in := range extras {
			if got := po.StatusFrom(in); got != po.StatusUnknown {
				t.Fatalf("StatusFrom(%v) = %v, want %v", in, got, po.StatusUnknown)
			}
		}
	})
}
