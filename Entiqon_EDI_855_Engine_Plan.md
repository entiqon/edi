# Entiqon EDI Engine -- Implementation Plan (Priority: 855)

## Objective

Build a plug-and-play EDI Engine focused on:

**BOD → Canonical → JSON → EDI (855)**

The engine must be:

-   ERP-independent
-   Format-independent
-   Streaming-capable (no payload limits)
-   Vendor-neutral (no cloud lock-in)
-   Extensible for future transaction sets

------------------------------------------------------------------------

# Phase 0 -- Repository Foundation

## Goals

-   GitHub repository initialized
-   Go module initialized
-   CI pipeline (gofmt + go test)
-   Clean folder structure
-   License + README committed

## Initial Structure

    /edi
      /canonical
        /po
      /engine
      /formats
        /x12
        /bodxml
      /internal/testdata

------------------------------------------------------------------------

# Phase 1 -- Engine Core (Minimal but Extensible)

## Deliverables

### 1. Engine Document Abstraction

-   Document metadata (Type, Version, Direction, Partner)
-   Streaming-friendly (support io.Reader / io.Writer later)

### 2. Transformer Interface

    type Transformer interface {
        Name() string
        Can(meta DocMeta) bool
        Transform(ctx context.Context, in Document) (Document, error)
    }

### 3. Registry

-   Register transformers
-   Route documents dynamically

### 4. Pipeline Execution

-   Execute transformation chain
-   No hardcoded flows

------------------------------------------------------------------------

# Phase 2 -- Canonical Model (Purchase Order Acknowledgment)

## Deliverables

### canonical/po

-   PurchaseOrderAcknowledgment
-   AcknowledgmentStatus
-   LineStatus
-   Strong typing
-   JSON-friendly enum marshaling
-   Full GoDoc

## Requirements

-   No X12 references inside canonical
-   No ERP-specific fields
-   Stable public contract

------------------------------------------------------------------------

# Phase 3 -- X12 855 Generation (OUT)

## Goal

Canonical → X12 855

## Deliverables

-   Segment builder
-   ST/SE
-   BAK
-   REF (optional)
-   DTM (optional)
-   N1 loop (optional)
-   PO1 loop
-   ACK segments
-   CTT

## Requirements

-   Support large documents
-   Accept io.Writer
-   Validate with external validator

------------------------------------------------------------------------

# Phase 4 -- X12 855 Parsing (IN)

## Goal

X12 855 → Canonical

## Deliverables

-   ST detection
-   Segment parsing
-   ACK code mapping:
    -   IA → ACCEPTED
    -   IC → CHANGED
    -   IB → BACKORDER
    -   IR → REJECTED
-   Header status derivation from lines

------------------------------------------------------------------------

# Phase 5 -- BOD XML → Canonical

## Goal

Infor M3 BOD → Canonical PO Acknowledgment

## Deliverables

-   Minimal XML binding structs
-   Mapper to canonical
-   Test with real BOD samples

------------------------------------------------------------------------

# Phase 6 -- Full Pipeline

    BOD XML → Canonical → JSON → X12

## Deliverables

-   Engine pipeline configuration
-   In-memory execution
-   JSON export option

------------------------------------------------------------------------

# Phase 7 -- Transport Layer (Plug-and-Play)

## Interfaces

    type Source interface {
        Read(ctx context.Context) (io.Reader, error)
    }

    type Destination interface {
        Write(ctx context.Context, r io.Reader) error
    }

## Initial Implementations

-   Local filesystem
-   S3 (optional)
-   FTP/SFTP (future)

------------------------------------------------------------------------

# Non-Goals (For Now)

-   VAN replacement
-   Trading partner onboarding UI
-   Cloud multi-tenant SaaS
-   Network routing infrastructure

------------------------------------------------------------------------

# Architectural Principles

1.  Canonical-first
2.  No vendor lock-in
3.  Streaming support
4.  Pluggable transformers
5.  Extensible transaction registry

------------------------------------------------------------------------

# Immediate Priority

## Implement in this order:

1.  Engine skeleton
2.  Canonical PO Acknowledgment
3.  X12 855 Builder
4.  BOD → Canonical mapper

------------------------------------------------------------------------

# Long-Term Expansion

Future transaction sets:

-   810
-   832
-   846
-   850 (IN/OUT)
-   855 (IN/OUT)
-   856 (IN/OUT)
-   870
-   940
-   943
-   944
-   945

------------------------------------------------------------------------

# Vision

Entiqon EDI Engine will be:

-   Embedded
-   Self-hosted
-   High-performance
-   Canonical-driven
-   Independent from SaaS limitations
-   Foundation for future Entiqon Super Engines
