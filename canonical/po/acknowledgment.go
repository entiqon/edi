// Package po contains canonical representations of
// purchase order related business documents.
//
// The types in this package are format-agnostic and
// independent of XML, JSON, X12, ERP systems, or
// transport protocols.
package po

// PurchaseOrderAcknowledgment represents a canonical
// Purchase Order Acknowledgment (EDI 855 equivalent).
//
// It is a normalized business representation and does not
// contain any X12 segment information or XML binding details.
//
// This struct is intended to be used as the internal,
// format-independent contract between:
//   - XML bindings
//   - JSON serialization
//   - X12 builders
//   - Other transformation layers
type PurchaseOrderAcknowledgment struct {
	// OrderID is the buyer's purchase order number being acknowledged.
	OrderID string

	// Status represents the overall acknowledgment status
	// derived from line statuses.
	Status Status

	// Lines contains all acknowledged purchase order lines.
	Lines []Line
}
