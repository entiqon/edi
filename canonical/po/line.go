package po

// AcknowledgmentLine represents the acknowledgment
// details of a single purchase order line.
type AcknowledgmentLine struct {
	// LineID is the buyer's line number.
	LineID string

	// ItemID is the product identifier.
	ItemID string

	// OrderedQty is the original ordered quantity.
	OrderedQty float64

	// AcceptedQty is the quantity confirmed by the supplier.
	AcceptedQty float64

	// BackorderQty is the quantity placed on backorder.
	BackorderQty float64

	// RejectedQty is the quantity rejected.
	RejectedQty float64

	// Status represents the acknowledgment status of this line.
	Status LineStatus
}
