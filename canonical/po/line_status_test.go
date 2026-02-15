package po_test

import (
	"testing"

	"github.com/entiqon/edi/canonical/po"
)

type lineStatusCase struct {
	status  po.LineStatus
	text    string
	isValid bool
	aliases []string // extra text inputs that should parse to `status` (or LineUnknown for invalid status)
}

var lineStatusCases = []lineStatusCase{
	{po.LineStatus(-1), "UNKNOWN", false, nil},

	{po.LineUnknown, "UNKNOWN", true, []string{
		"unknown",
		" UNKNOWN ",
		"",
		"nope",
	}},

	{po.LineAccepted, "ACCEPTED", true, []string{
		"accepted",
		" accepted ",
	}},

	{po.LineChanged, "CHANGED", true, []string{
		"changed",
		" CHANGED ",
	}},

	{po.LineBackorder, "BACKORDER", true, []string{
		"backorder",
		" BACKORDER ",
	}},

	{po.LineRejected, "REJECTED", true, []string{
		"rejected",
		" REJECTED ",
	}},

	{po.LineStatus(999), "UNKNOWN", false, nil},
}

func TestLineStatus(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		for _, tc := range lineStatusCases {
			if got := tc.status.String(); got != tc.text {
				t.Fatalf("String(%v) = %q, want %q", tc.status, got, tc.text)
			}
		}
	})

	t.Run("IsValid", func(t *testing.T) {
		for _, tc := range lineStatusCases {
			if got := tc.status.IsValid(); got != tc.isValid {
				t.Fatalf("IsValid(%v) = %v, want %v", tc.status, got, tc.isValid)
			}
		}
	})

	t.Run("MarshalText", func(t *testing.T) {
		for _, tc := range lineStatusCases {
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
		for _, tc := range lineStatusCases {
			// Expected: if the enum is invalid, parsing its text should still yield LineUnknown.
			want := tc.status
			if !tc.status.IsValid() {
				want = po.LineUnknown
			}

			// Always test canonical text + any aliases.
			inputs := append([]string{tc.text}, tc.aliases...)

			for _, in := range inputs {
				var s po.LineStatus
				if err := s.UnmarshalText([]byte(in)); err != nil {
					t.Fatalf("UnmarshalText(%q) error = %v", in, err)
				}
				if s != want {
					t.Fatalf("UnmarshalText(%q) = %v, want %v", in, s, want)
				}
			}
		}
	})

	t.Run("UnmarshalText_NilReceiver", func(t *testing.T) {
		var s *po.LineStatus
		if err := s.UnmarshalText([]byte("ACCEPTED")); err == nil {
			t.Fatalf("expected error for nil receiver, got nil")
		}
	})

	t.Run("LineStatusFrom", func(t *testing.T) {
		for _, tc := range lineStatusCases {
			want := tc.status
			if !tc.status.IsValid() {
				want = po.LineUnknown
			}

			// Base inputs derived from the same shared table.
			inputs := []any{
				tc.status,
				int(tc.status),
				tc.text,
			}

			// Also validate aliases for string parsing behavior.
			for _, a := range tc.aliases {
				inputs = append(inputs, a)
			}

			for _, in := range inputs {
				if got := po.LineStatusFrom(in); got != want {
					t.Fatalf("LineStatusFrom(%v) = %v, want %v", in, got, want)
				}
			}
		}

		// Unsupported types should always fall back to LineUnknown.
		extras := []any{true, nil, 3.14}
		for _, in := range extras {
			if got := po.LineStatusFrom(in); got != po.LineUnknown {
				t.Fatalf("LineStatusFrom(%v) = %v, want %v", in, got, po.LineUnknown)
			}
		}
	})
}
