package bod

import (
	"encoding/xml"

	"github.com/entiqon/edi/canonical/po"
)

type SalesOrder struct {
	XMLName          xml.Name         `xml:"M3EDISalesOrder"`
	SalesOrderHeader SalesOrderHeader `xml:"SalesOrderHeader"`
	SalesOrderLine   []SalesOrderLine `xml:"SalesOrderLine"`
}

type SalesOrderHeader struct {
	DisplayID string `xml:"DisplayID"`
	Status    struct {
		Code string `xml:"Code"`
	} `xml:"Status"`
}

type SalesOrderLine struct {
	LineNumber string `xml:"LineNumber"`
	Status     struct {
		Code string `xml:"Code"`
	} `xml:"Status"`

	Item struct {
		ItemID struct {
			ID string `xml:"ID"`
		} `xml:"ItemID"`
	} `xml:"Item"`

	Quantity float64 `xml:"Quantity"`
}

func (s *SalesOrder) ToCanonicalAck() po.PurchaseOrderAcknowledgment {
	ack := po.PurchaseOrderAcknowledgment{
		OrderID: s.SalesOrderHeader.DisplayID,
	}

	for _, line := range s.SalesOrderLine {

		ordered := line.Quantity

		cLine := po.Line{
			LineID:     line.LineNumber,
			ItemID:     line.Item.ItemID.ID,
			OrderedQty: ordered,
			Status:     po.LineStatusFrom(line.Status.Code),
		}

		ack.Lines = append(ack.Lines, cLine)
	}

	ack.Status = po.StatusFrom(s.SalesOrderHeader.Status.Code)
	return ack
}
