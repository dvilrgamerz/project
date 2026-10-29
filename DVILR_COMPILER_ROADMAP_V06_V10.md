# DVILR Native Compiler — V0.6 to V1.0 Implementation Tracker

## October 29, 2026 — V0.6.4 collections compiler preview

A separate downloadable ZIP (not committed into this repository) contains the updated compiler source, Windows cross-built compiler, tests, and Windows CI workflow.

- **Native backend changes:** `fill(array_or_vector, i64)` and `reverse(array_or_vector)`; only the current logical length of bounded vectors is accessed, and empty/single-element reversals are guarded.
- **Locally verified:** Go unit tests, race tests, vet, Windows x64 compiler cross-build, generated PE inspection, Node-executed numeric Wasm and WASI checks (45).
- **Not verified:** execution of generated Windows programs and GPU hardware.
- **Still incomplete:** Full V0.6 core (general memory-managed collections and complete type system); V0.8 general networking/async/WASI; V0.9 native GPU/graphics; V1.0 self-hosting Stage 1→2→3. A Go-authored native compiler does not prove self-hosting.


## October 29, 2026 — Native V0.6.4 checked vector insertion/removal

- Local source ZIP includes native compiler support for `insert(vector,index,value)` and `remove_at(vector,index)` on bounded global i64 vectors (maximum 1024 elements). Both reject invalid unsigned indices and full-capacity insertions with runtime exit code 8. Removal shifts left and clears the former final slot.
- Added Go compiler tests for valid/invalid forms, deterministic PE output, and machine-code safety checks, plus a **proposed, unexecuted** Windows CI test of expected program output and invalid-index exit codes.
- Verified locally: `go test ./...`, `go test -race ./...`, `go vet ./...`, cross-built Windows x64 compiler, generated Windows sample PE binaries; generated minimal WASI program executed in Node and returned 45.
- **Not verified:** running compiled sample programs on Windows. **Not complete:** general memory safety/type system, production networking/WASI, native GPU/hardware graphics, and any DVILR-written compiler Stage 1→2→3 self-hosting proof.
- The new compiler source and executable are packaged separately in `DVILR-Standalone-v0.6.4-Checked-Vectors-Preview.zip`, **not merged into this demo GitHub repo**.


## October 29, 2026 — Native V0.6.3 safety and collections preview

The source/binary are provided in a downloadable ZIP in the associated ChatGPT conversation; **this GitHub repository still does not contain the standalone compiler implementation.**

- Added native `resize(v,n)` for bounded i64 vectors; 0–1024 elements, zero-initialized expansion; invalid sizes exit status 8.
- Added `clear_vector(v)` (distinct from console `clear`); vectors refilled after clear are zero-initialized.
- Added signed division/modulo overflow protection for MinInt64 / -1 and division by zero; failure exit status 7.
- Added positive/negative tests and x64 opcode checks plus a proposed Windows CI runtime workflow.
- Locally verified: `go test ./...`, `go test -race ./...`, `go vet ./...`, Windows x64 compiler cross-build and PE generation; numeric WASI preview1 module ran in Node and returned 45.
- **Not verified:** running generated Windows programs under Windows or GitHub Windows CI.
- **Full V0.6, V0.8, V0.9, V1.0 remain INCOMPLETE.** Complete type/memory system, native network/async APIs, GPU/hardware graphics, and the DVILR-written compiler Stage 1→2→3 proof remain unimplemented.


## October 26, 2026 — Verified V0.6.2 local preview (source ZIP, not merged into this GitHub repo)

- The standalone Windows x64 Go-bootstrapped compiler now generates native `sum(array_or_vector)` and `dot(a,b)` loops for **i64** arrays and bounded vectors, reading values and lengths at runtime.
- Compiler diagnostics reject fixed-array dot-product length mismatches; the generated executable checks dynamic vector length mismatches and exits with status 8.
- Compiler checks also reject a subset of reserved-name and symbol collisions while retaining compatibility with existing DVILR examples.
- Added positive/negative tests for these intrinsics and namespace rules; all Go tests, race tests and vet passed, Windows x64 Go compiler cross-build succeeded, generated PE parsed correctly, x64 instruction disassembly inspected.
- Numeric WebAssembly and minimal WASI preview1 modules validated and executed under Node.js.
- **Windows execution of compiled .exe code is unverified.** No float tensor, GPU runtime, complete type checker, first-class heap-safe containers or Stage 1→2→3 compiler self-hosting proof.
- The updated compiler binary and complete source are provided as a separate V0.6.2 preview ZIP in the corresponding ChatGPT conversation; this GitHub demo repository is **not** the compiler source repository.
- Full V0.6, V0.8, V0.9 and V1.0 status remains **INCOMPLETE**.


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
