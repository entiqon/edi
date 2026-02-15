# Electronic Data Interchange (EDI)

[![Go Reference](https://pkg.go.dev/badge/github.com/entiqon/edi.svg)](https://pkg.go.dev/github.com/entiqon/edi)
[![Build Status](https://github.com/entiqon/edi/actions/workflows/ci.yml/badge.svg)](https://github.com/entiqon/edi/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/entiqon/edi/graph/badge.svg?token=ec14R1zZk1)](https://codecov.io/gh/entiqon/edi)
[![Go Report Card](https://goreportcard.com/badge/github.com/entiqon/edi)](https://goreportcard.com/report/github.com/entiqon/edi)
[![Latest Release](https://img.shields.io/github/v/release/entiqon/edi)](https://github.com/entiqon/edi/releases)
[![License](https://img.shields.io/github/license/entiqon/edi)](https://github.com/entiqon/edi/blob/main/LICENSE)
![Documentation](https://img.shields.io/badge/docs-coming--soon-lightgrey)

---

## Overview

A modular, ERP-independent EDI engine written in Go.

This project introduces a clean architectural separation between business logic and transport formats.

It provides:

- Canonical business document models
- X12 generation and parsing
- BOD XML integration (Infor M3 and beyond)
- Clean, strongly-typed APIs
- Extensible transaction routing

### Core Separation Principle

```
BOD XML ⇄ Canonical ⇄ X12
```

Canonical models are the stable core of the system.  
Adapters handle format transformations.

---

## ⚠️ Project Status

This project is currently in **v0.x (active development)**.

Breaking changes may occur until `v1.0.0`.

---

## Installation

Install latest version:

```bash
go get github.com/entiqon/edi
```

Install specific version:

```bash
go get github.com/entiqon/edi@v0.1.0-dev.1
```

---

## Status Legend

- ⏳ Planned
- 🚧 In Progress
- ✅ Implemented
- 🔒 Stable (post v1.0)

---

## Goals

### ERP-independent canonical models

- ⏳ 810 Invoice
- ⏳ 832 Price/Sales Catalog
- ⏳ 846 Inventory Inquiry/Advice
- ⏳ 850 Purchase Order (IN + OUT)
- 🚧 855 Purchase Order Acknowledgment (IN + OUT)
- ⏳ 856 Advance Ship Notice (IN + OUT)
- ⏳ 870 Order Status Report
- ⏳ 940 Warehouse Shipping Order
- ⏳ 943 Warehouse Stock Transfer Shipment Advice
- ⏳ 944 Warehouse Stock Transfer Receipt Advice
- ⏳ 945 Warehouse Shipping Advice

### X12 generation and parsing

- ⏳ 810
- ⏳ 832
- ⏳ 846
- ⏳ 850 (IN + OUT)
- 🚧 855 (IN + OUT)
- ⏳ 856 (IN + OUT)
- ⏳ 870
- ⏳ 940
- ⏳ 943
- ⏳ 944
- ⏳ 945

---

## Architecture

The canonical layer is format-agnostic.

Rules:

- No ERP-specific logic inside canonical.
- No X12 segment leakage into canonical.
- Adapters convert formats into canonical.
- Canonical remains stable long-term.

Adapters may evolve. Canonical must endure.

---

## Roadmap

- **v0.1.x** — BOD SalesOrder + 855
- **v0.2.x** — 850 IN/OUT canonical + X12
- **v0.3.x** — 856 ASN support
- **v1.0.0** — Stable canonical API

---

## Contributing

Contributions are welcome.

Before submitting a PR:

- Ensure tests pass
- Respect canonical isolation principles
- Avoid ERP/X12 leakage into core models
- Update documentation (README, doc.go, CHANGELOG)
- Follow semantic versioning

See full contribution guidelines in  
👉 [CONTRIBUTING.md](CONTRIBUTING.md)

---

## 📄 License

💡 Originally created by [Isidro A. Lopez G.](https://github.com/ialopezg)  
🏢 Maintained by the [Entiqon Labs Organization](https://github.com/entiqon)

[MIT](./LICENSE) — © Isidro A. López G. / Entiqon Project