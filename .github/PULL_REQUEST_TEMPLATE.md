# 📝 Contributor Quick Reference (Entiqon EDI)

## ✅ Before Commit

### Tests

- Follow `file → methods → cases` structure.
- Use **PascalCase** naming.
- Always use `t.Run` hierarchy.
- Cover:
  - nil receivers
  - empty collections
  - invalid input
  - malformed XML/X12
  - unknown transaction types

Target coverage:
- Minimum: **80%**
- Core canonical packages: **90%+ recommended**

---

### Documentation (Required)

Every feature must update:

- [ ] `doc.go`
- [ ] `README.md`
- [ ] `example_test.go`
- [ ] `CHANGELOG.md`

If documentation is missing, the feature is incomplete.

---

### Architectural Rules (Critical)

Before committing, verify:

- [ ] No X12 segment leakage into canonical
- [ ] No ERP/BOD types inside canonical models
- [ ] Adapters handle transformation logic
- [ ] Canonical remains format-agnostic

Core principle:

```
BOD ⇄ Canonical ⇄ X12
```

Canonical is the stable core.

---

## 💬 Commit Rules

We use **Conventional Commits**:

- `feat(scope): ...` → new feature
- `fix(scope): ...` → bug fix
- `docs(scope): ...` → documentation
- `refactor(scope): ...` → no behavior change
- `test(scope): ...` → test improvements

### Detailed Commit (Preferred)

```text
feat(bod/sales_order): implement ToCanonical transformation

- Added SalesOrder BOD parsing
- Mapped to canonical model
- Edge cases covered (nil, invalid XML)
- Updated README, doc.go, example_test, CHANGELOG

🚧 855 groundwork in progress.
```

### Squash Commit

Keep it concise:

```text
feat(bod/sales_order): implement ToCanonical transformation
```

---

## 🚀 Release Flow (v0.x)

1. Ensure tests pass.
2. Verify canonical isolation.
3. Update docs + examples.
4. Update CHANGELOG.
5. Tag using semver:

```bash
git tag -s v0.x.y -m "Release v0.x.y"
git push origin v0.x.y
```

⚠️ Breaking changes are allowed before `v1.0.0`, but must be documented.

---

## ✨ Style Guidelines

- Use clear status markers:
  - ⏳ Planned
  - 🚧 In Progress
  - ✅ Implemented
- Keep commits professional and consistent.
- Avoid unnecessary verbosity.
- Canonical purity over convenience.

Optional personality is welcome — but architecture comes first.

---

## 🔍 Notes for Reviewers

Please describe:

- Architectural impact
- Canonical changes
- Backward compatibility considerations
- Future follow-up tasks