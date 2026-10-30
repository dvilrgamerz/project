# DVILR v1.1.3 — Typed Scalars and Reference Memory Checkpoint

**Current release: Developer Preview — NOT a finished v1.1 language, runtime, or self-hosting compiler.** Date: 2026-10-30.

The updated compiler source, Windows compiler executable, GUI installer and Windows verification scripts are packaged in a separate downloadable ZIP in ChatGPT: `DVILR-v1.1.3-Type-Memory-Compiler-Full-Source.zip`.

**This GitHub repository remains the website/demo and partial compiler migration, not the full compiler source.** The full archive is not published as a GitHub Release. Do not claim this repository builds v1.1.3 from a fresh clone.

## Newly implemented and locally verified

### Native Windows compiler (Go-bootstrap)

- Explicit scalar annotations for `let name: bool = expression` and `let name: i64 = expression`, also `var`.
- Type checking of annotated declarations and assignments: logical operators require booleans; arithmetic requires integers; comparisons produce boolean results.
- Scalar type metadata for locals is restored correctly after lexical block shadowing.
- The underlying native representation of boolean values remains one i64 slot: `true=1`, `false=0`.
- Old unannotated code remains backward compatible. This is **not** whole-program static type checking: function signature checking, generics, full AST/IR, heap references and all unannotated contexts still need work.

### Go-hosted memory reference design

- `sdk/memory.Arena[T]` uses thread-safe generational handles, rejects expired/cross-arena handles and zeroes released slots.
- `sdk/memory.Vector[T]` supports bounded, checked resize/mutate/remove, zeroes discarded elements, snapshots and clones. Pointer-bearing generic values are **shallow copies**, not deep-owned data.
- `sdk/memory.ByteArena` deep-copies bytes on allocation, reads and writes with per-object bounds.
- `sdk/memory.Text` validates UTF-8, checks byte capacity, supports rune-indexed slice/replace; runes are not grapheme clusters.
- All memory code is Go reference code: **not** a runtime called by compiled DVILR executables.

## Test evidence

| Check | Result |
| --- | --- |
| `go test ./... -count=1` (8 packages) | PASSED |
| `go test -race ./... -count=1` (8 packages) | PASSED |
| `go vet ./...` | PASSED |
| 10,000-operation randomized vector reference test | PASSED |
| Invalid/stale handles, UTF-8/bounds and concurrent access tests | PASSED |
| Compile + PE generation for typed scalar example | PASSED |
| Static rejection of four invalid typed programs | PASSED |
| Windows x64 compiler cross-build | PASSED |
| Windows x64 GUI installer cross-build | PASSED |
| Source ZIP archive integrity (147 members) | PASSED |
| **Execution on a Windows host** | **NOT RUN/NOT VERIFIED** |
| Windows CI workflow with complete compiler source | **NOT RUN** |

### Downloads and checksums

- `DVILR-v1.1.3-Type-Memory-Compiler-Full-Source.zip` SHA-256: `bdf1da71e1488a2e2cacc6d09d7a2c5144ae0b1641f5d9d778ddc435c589d608`
- `DVILR-Setup-v1.1.3-Type-Memory-Preview.exe` SHA-256: `05b64ad74a48fcb8fa6d253f8a21de356860f30fb8f124d65de1fe60ae544ff6`

The installer is **unsigned** and has not yet run on actual Windows. Verify the checksum and review the source before running.

## Native Windows verification script

The full-source ZIP contains `tools/verify-v113-windows.ps1` and an unrun `.github/workflows/v113-native-windows.yml`. After extracting on Windows, from the ZIP root:

```powershell
.\installer\payload\dvl.exe version
powershell -NoProfile -ExecutionPolicy Bypass -File .\tools\verify-v113-windows.ps1 -Compiler .\installer\payload\dvl.exe
```

The script builds and **runs** a generated Windows PE and verifies output `1,0,1`; it checks that four ill-typed programs fail to compile. Do not call Windows execution verified until that script successfully runs on Windows and its output is retained.

## Prioritized upgrades after v1.1.3

1. **P0, Windows verification + repository migration:** check in the full compiler source, run Windows native test workflow and installer install/uninstall tests.
2. **P1, complete type system:** typed AST/IR, typed function parameters/returns, generics, struct layouts, strict checking of *all* source expressions, deterministic diagnostics and conformance tests.
3. **P2, language-level memory safety:** choose ownership/borrowing or tracing GC, integrate safe native heap references and true resizable generic containers; test use-after-free, double-free, aliasing and concurrency.
4. **P3, standard library/UX:** checked file I/O, packages/lockfiles, formatter + LSP/editor support.
5. **P4, systems:** DVILR-callable networking, concurrency and a complete Wasm/WASI core.
6. **P5, AI and graphics:** CPU tensor correctness, verified GPU kernels, graphics and a playable DVILR game.
7. **P6, compiler self-hosting:** complete compiler authored in DVILR and independently reproducible Stage 0 → Stage 1 → Stage 2 → Stage 3 verification.

**Release verdict:** full type/memory model = INCOMPLETE; native GPU = NOT IMPLEMENTED; self-hosting = NOT ACHIEVED.
