## Goals

- ERP-independent canonical models
  - [ ] 810 Invoice
  - [ ] 832 Price/Sales Catalog
  - [ ] 846 Inventory Inquiry/Advice
  - [ ] 850 Purchase Order (IN + OUT)
  - [ ] 855 Purchase Order Acknowledgment (IN + OUT)
  - [ ] 856 Advance Ship Notice (IN + OUT)
  - [ ] 870 Order Status Report
  - [ ] 940 Warehouse Shipping Order
  - [ ] 943 Warehouse Stock Transfer Shipment Advice
  - [ ] 944 Warehouse Stock Transfer Receipt Advice
  - [ ] 945 Warehouse Shipping Advice

- X12 generation and parsing
  - [ ] 810
  - [ ] 832
  - [ ] 846
  - [ ] 850 (IN + OUT)
  - [ ] 855 (IN + OUT)
  - [ ] 856 (IN + OUT)
  - [ ] 870
  - [ ] 940
  - [ ] 943
  - [ ] 944
  - [ ] 945

- BOD XML integration
  - [ ] Infor M3 BOD → Canonical → X12
  - [ ] Canonical → (optional) BOD XML

- Clean and typed APIs
  - [ ] Stable canonical types (no X12/BOD leakage)
  - [ ] JSON-friendly enums (string + int tolerant decoding)
  - [ ] Transaction registry (detect ST*xxx and route)