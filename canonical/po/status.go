package po

import (
	"fmt"
	"strings"
)

// Status represents the header-level status
// of a Purchase Order Acknowledgment (EDI 855).
//
// It summarizes the overall outcome of the order after
// supplier processing.
//
// The zero value (StatusUnknown) represents an undefined
// or not-yet-determined state.
type Status int

const (
	// StatusUnknown indicates that the acknowledgment status
	// has not been determined or parsed yet.
	StatusUnknown Status = iota

	// StatusAccepted indicates all lines were fully accepted
	// without changes.
	StatusAccepted

	// StatusPartial indicates at least one line was changed,
	// backordered, or rejected while others were accepted.
	StatusPartial

	// StatusRejected indicates the entire order was rejected.
	StatusRejected
)

// IsValid reports whether the status represents
// a valid business state.
func (a Status) IsValid() bool {
	return a > StatusUnknown && a <= StatusRejected
}

// String returns the canonical string representation
// of the acknowledgment status.
//
// This method is safe for logging and serialization.
func (a Status) String() string {
	switch a {
	case StatusAccepted:
		return "ACCEPTED"
	case StatusPartial:
		return "PARTIAL"
	case StatusRejected:
		return "REJECTED"
	default:
		return "UNKNOWN"
	}
}

// MarshalText implements encoding.TextMarshaler,
// allowing the status to be serialized as text
// in JSON, XML, and other encodings.
func (a Status) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler,
// allowing the status to be parsed from textual
// JSON or XML representations.
func (a *Status) UnmarshalText(text []byte) error {
	if a == nil {
		return fmt.Errorf("Status.UnmarshalText: nil receiver")
	}
	*a = StatusFrom(string(text))
	return nil
}

// StatusFrom converts a dynamic value (string or int)
// into an Status.
//
// Supported inputs:
//   - string ("ACCEPTED", "PARTIAL", "REJECTED")
//   - int (matching enum values)
//
// Any unsupported value returns StatusUnknown.
func StatusFrom(v any) Status {
	switch x := v.(type) {
	case Status:
		if x == StatusUnknown || x.IsValid() {
			return x
		}
		return StatusUnknown

	case int:
		s := Status(x)
		if s == StatusUnknown || s.IsValid() {
			return s
		}
		return StatusUnknown

	case string:
		switch strings.ToUpper(strings.TrimSpace(x)) {
		case "ACCEPTED":
			return StatusAccepted
		case "PARTIAL":
			return StatusPartial
		case "REJECTED":
			return StatusRejected
		case "UNKNOWN", "":
			return StatusUnknown
		default:
			return StatusUnknown
		}

	default:
		return StatusUnknown
	}
}
