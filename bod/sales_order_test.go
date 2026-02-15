package bod_test

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/entiqon/edi/bod"
	"github.com/entiqon/edi/canonical/po"
)

func Test_ToCanonicalAck_FromRealXML(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(filename), "..")
	path := filepath.Join(root, "data", "0011528843.xml")

	// Load XML file
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read XML file: %v", err)
	}

	var doc bod.SalesOrder
	if err := xml.Unmarshal(data, &doc); err != nil {
		t.Fatalf("failed to unmarshal XML: %v", err)
	}

	ack := doc.ToCanonicalAck()

	// --- Header Assertions ---
	if ack.OrderID != "0011528843" {
		t.Fatalf("unexpected OrderID: %s", ack.OrderID)
	}

	// --- Lines Assertions ---
	if len(ack.Lines) == 0 {
		t.Fatal("expected lines but got none")
	}

	for _, line := range ack.Lines {
		if !line.Status.IsAccepted() {
			t.Fatalf("expected Approved line status, got %v", line.Status)
		}
	}

	if ack.Status != po.StatusAccepted {
		t.Fatalf("expected overall status Approved, got %v", ack.Status)
	}
}
