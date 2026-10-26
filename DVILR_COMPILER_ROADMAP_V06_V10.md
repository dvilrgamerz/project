# DVILR Native Compiler — V0.6 to V1.0 Implementation Tracker

## October 26, 2026 — Local native compiler progress (source package separate from this repo)

**A downloadable V0.6+ foundations preview has been created and tested locally. Its compiler source/binary are packaged separately; this GitHub repository still contains the storefront/WebAssembly demo, not the updated native compiler.**

- V0.6 foundations: basic global `struct` / `record` with i64 fields, integer `enum`, relative `import "name.dvl"` with cycle/path checks, `assert`, plus earlier checked global i64 arrays, `len`, and static text.
- V0.8 foundations: Windows-only `http_get("https://...", "local.file")` via URLMON, as a download-to-file *prototype*; Wasm numeric subset now has booleans, `if/else` and `while`. A generated Wasm loop was actually run and returned 55.
- V0.9 foundations: CPU-only native integer `sum(array)` and `dot(a,b)`. No floating tensor engine, GPU, autograd or full graphics engine.
- V1.0: introduced a fail-closed Stage 0→1→2→3 check, but the full DVILR-written compiler and self-hosting proof remain **missing**.

Verification: 121 passing Go tests, `go vet` passed, Windows x64 cross-build passed, and generated Wasm executed in Node. **Windows native execution has NOT been verified.** Major version completion statuses remain `false`.



**This is a roadmap, not a claim of complete implementation.** The `project` repository contains a storefront/WebAssembly demo and does not currently carry the complete native compiler sources.

## V0.6 — Core language

| Capability | Status | Next engineering gate |
| --- | --- | --- |
| Strings + arrays | PARTIAL LOCAL PREVIEW | Native V0.6 preview adds fixed-size global i64 arrays (max 1,024), checked index reads/writes, len(array), and named static printable text; dynamic strings/vectors still TODO |
| Structs + enums | TODO | Parser, type checker, layout, ABI, negative tests |
| File I/O | PARTIAL V0.5 | Existing byte load/save; typed safe streaming/error handling TODO |
| Modules + imports | TODO | Resolution, dependency graph, separate compilation |
| Error handling | TODO | Result/Option, diagnostics, cleanup behavior |
| Memory safety | TODO | Ownership/lifetime or memory-managed model; fuzz/bounds tests |
| Standard library | TODO | Math/text/collection/file/time APIs |
| Optimization | TODO | Typed IR, constant folding, DCE, benchmark gates |

## V0.7 — Compiler hardening

Typed AST/IR, diagnostic recovery, stable calling convention, native debug support and reproducible CI builds.

## V0.8 — Web and systems (TODO)

HTTP clients/servers, TLS through reviewed platform libraries, networking, channels/tasks/cancellation, race/stress testing, expanded Wasm/WASI interoperability. The existing numeric-subset Wasm backend does **not** satisfy this milestone.

## V0.9 — AI and game foundations (TODO)

CPU reference tensors, dtype/shape/stride validation, GPU backend abstraction with parity tests, graphics/window/input API, real-time math and timing. Existing demos do **not** constitute a production SDK.

## V1.0 — True self-hosting (TODO)

Freeze Go-built Stage 0; implement lexer/parser/typechecker/IR/backend/PE writer in DVILR. Stage 0 builds Stage 1; Stage 1 builds Stage 2 (without Go); Stage 2 builds Stage 3; verify behavior and reproducible hashes in CI. **No Stage 1→2→3 proof exists yet.**

## V0.6 standalone preview progress (not merged into this repo)

A local Windows x64 compiler preview has been built and tested on Linux (Go test suite and native PE generation). New language subset: `array a=[1,2,3]`, `a[i]`, `a[i]=n`, `len(a)`, `text s="hello"`, `say s`. The executable is a Windows PE, but executing it under Windows has not been tested in the current environment.

**Integration requirement:** import the complete updated compiler source into a dedicated DVILR language repository, add Windows CI and runtime tests, then tackle the remaining items incrementally. Do not mark V0.6/V0.8/V0.9/V1.0 complete until their exit criteria are met.
