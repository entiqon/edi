package po

import (
	"fmt"
	"strings"
)

// LineStatus represents the acknowledgment status of a purchase order line.
//
// LineStatus is canonical and intentionally independent of X12 codes (IA, IC, IR, IB).
// Conversions to/from X12 belong in transformation layers.
type LineStatus int

const (
	// LineUnknown indicates the line status is not known or not set.
	// This is the zero value.
	LineUnknown LineStatus = iota

	// LineAccepted indicates the line was fully accepted with no changes.
	LineAccepted

	// LineChanged indicates the line was accepted with changes
	// (e.g., quantity/date modifications).
	LineChanged

	// LineBackorder indicates the line was partially fulfilled and the remaining
	// quantity is backordered.
	LineBackorder

	// LineRejected indicates the line was rejected.
	LineRejected
)

// IsValid reports whether s is a recognized LineStatus value.
func (s LineStatus) IsValid() bool {
	return s >= LineUnknown && s <= LineRejected
}

// String returns the canonical string representation of the line status.
//
// Values are stable and intended for logging, debugging, and text-based encodings.
func (s LineStatus) String() string {
	switch s {
	case LineAccepted:
		return "ACCEPTED"
	case LineChanged:
		return "CHANGED"
	case LineBackorder:
		return "BACKORDER"
	case LineRejected:
		return "REJECTED"
	default:
		return "UNKNOWN"
	}
}

// MarshalText implements encoding.TextMarshaler, allowing LineStatus to be
// serialized as text in JSON, XML, and other encodings.
func (s LineStatus) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler, allowing LineStatus to be
// parsed from text in JSON, XML, and other encodings.
//
// Accepted values (case-insensitive):
//   - "ACCEPTED", "CHANGED", "BACKORDER", "REJECTED", "UNKNOWN"
func (s *LineStatus) UnmarshalText(text []byte) error {
	if s == nil {
		return fmt.Errorf("LineStatus.UnmarshalText: nil receiver")
	}
	*s = LineStatusFrom(string(text))
	return nil
}

// LineStatusFrom converts a dynamic value (string or int) into a LineStatus.
//
// Supported inputs:
//   - LineStatus (returned as-is)
//   - int (must match a valid enum value)
//   - string (case-insensitive): "ACCEPTED", "CHANGED", "BACKORDER", "REJECTED", "UNKNOWN"
//
// Any unsupported or invalid value returns LineUnknown.
func LineStatusFrom(value any) LineStatus {
	switch v := value.(type) {
	case LineStatus:
		if v.IsValid() {
			return v
		}
		return LineUnknown

	case int:
		s := LineStatus(v)
		if s.IsValid() {
			return s
		}
		return LineUnknown

	case string:
		switch strings.ToUpper(strings.TrimSpace(v)) {
		case "ACCEPTED":
			return LineAccepted
		case "CHANGED":
			return LineChanged
		case "BACKORDER":
			return LineBackorder
		case "REJECTED":
			return LineRejected
		case "UNKNOWN":
			return LineUnknown
		default:
			return LineUnknown
		}

	default:
		return LineUnknown
	}
}
