# Contributing to Entiqon EDI

Welcome! 🎉

This project follows strict architectural and documentation standards.
Canonical isolation and clean layering are mandatory.

------------------------------------------------------------------------

## Testing

Minimum coverage: 80% (core packages should target 90%+).

Structure: - file → methods → cases - PascalCase naming - Use t.Run
hierarchy

Mandatory edge cases: - nil receivers - empty collections - invalid
input - malformed XML/X12

------------------------------------------------------------------------

## Architectural Rules

1.  Canonical Isolation

Canonical models: - MUST NOT import X12 - MUST NOT import BOD - MUST NOT
contain ERP-specific logic

Flow: BOD ⇄ Canonical ⇄ X12

Canonical is the stable core.

2.  No Format Leakage

-   No X12 segment names inside canonical models.
-   No ERP-specific structures inside core.

------------------------------------------------------------------------

## Documentation Requirements

Every feature must update:

-   doc.go
-   README.md
-   example_test.go
-   CHANGELOG.md

If documentation is missing, the feature is incomplete.

------------------------------------------------------------------------

## Commit Messages

We use Conventional Commits:

    feat(scope): short description

Examples:

    feat(bod/sales_order): implement ToCanonical transformation
    fix(x12/855): correct ACK segment mapping
    docs(readme): update installation instructions

------------------------------------------------------------------------

## Release Process

1.  Ensure tests pass
2.  Update CHANGELOG
3.  Tag using semver (v0.x.y)
4.  Push signed tag

Example:

    git tag -s v0.2.0 -m "Release v0.2.0"
    git push origin v0.2.0"

------------------------------------------------------------------------

## Contribution Checklist

-   [ ] Tests added
-   [ ] Edge cases covered
-   [ ] Canonical isolation respected
-   [ ] Docs updated
-   [ ] CHANGELOG updated

Keep canonical pure.
