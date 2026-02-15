# Stability Policy

## Canonical Core

The canonical layer is the long‑term stable core of this project.

Rules:

-   No ERP‑specific logic inside canonical.
-   No X12 segment leakage into canonical.
-   Canonical types must remain format‑agnostic.

## Adapter Layer

Adapters (BOD, X12, JSON, APIs):

-   May evolve independently.
-   May change when external standards change.
-   Must not impact canonical stability.

## Breaking Changes

Breaking changes are:

-   Allowed before v1.0.0.
-   Restricted after v1.0.0.

All breaking changes must: - Be documented in CHANGELOG.md. - Include
migration notes when applicable.

Stability is intentional, not accidental.
