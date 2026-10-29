# DVILR V1.0 — Unified Release Goal

**Status: IN DEVELOPMENT — NOT RELEASED.**
**Baseline reviewed:** DVILR Standalone v0.6.4 Preview (Go-bootstrapped Windows x64 compiler). Existing repository contains a storefront/Wasm demo and a progress tracker, **not the current standalone compiler source**.

## Mission

Ship **one unified DVILR V1.0** which includes all previously proposed V0.6, V0.8, V0.9 and self-hosting requirements. These are **acceptance criteria for a future release, not finished features**. No renaming previews to V1.0 to make progress appear complete.

## Required capabilities — all are V1.0 blockers

### L — Complete language and safe runtime (former V0.6)

- [ ] Complete typed AST/HIR/MIR; name resolution, generic functions/types, deterministic diagnostics and error recovery.
- [ ] UTF-8 dynamic strings with lengths, indexing/slicing rules, encoding validation and text APIs.
- [ ] Dynamic `Vec<T>`, slices, maps/sets, structs, enums and tagged unions with tested layout/ABI and generics.
- [ ] Memory-safe allocation model: ownership/borrowing **or** tracing GC with safe references; lifetimes or reachability, bounds, use-after-free, leak and aliasing tests.
- [ ] `Result<T,E>` / `Option<T>`, `?` propagation or explicit equivalent; deterministic cleanup and error codes.
- [ ] Modules/packages/imports, dependency resolution, separate compilation and incremental rebuilds.
- [ ] Standard library for math, Unicode text, files, buffered streams, time, data collections and test assertions.
- [ ] Optimized compilation: typed IR verifier, constant folding, dead-code elimination, benchmark regression checks and debug symbols.
- [ ] Windows and Linux x64 native release acceptance; stable ABI and minimum-version compatibility specification.

**Acceptance:** Positive + negative compiler test suites, mutation/property/fuzz tests, ASAN-equivalent/managed-runtime memory checks, deterministic diagnostics and compatibility tests must pass on supported native targets. Current bounded i64 vectors are not a substitute for general-purpose `Vec<T>`.

### W — Networking, concurrency and WebAssembly/WASI (former V0.8)

- [ ] DVILR-callable HTTP/1.1 client and server; robust URL parsing, TLS certificate verification, configurable timeouts, streaming, size limits and cancellation.
- [ ] TCP/UDP sockets, DNS, IPv4/IPv6 and portable error results; loopback/integration tests without outside network assumptions.
- [ ] Native task/thread primitives, futures, channels, mutex/atomics, structured concurrency and cancellation; race/deadlock stress tests.
- [ ] Defined `Send`/`Sync` or equivalent guarantees preventing unsafe task sharing.
- [ ] Wasm codegen for the supported language core, including strings, calls, memory, imports/exports, traps and control flow (not numeric subset only).
- [ ] WASI preview target explicitly defined, documented host permissions and conformance tests using a supported runtime.
- [ ] Native networking and concurrency API exercised by compiled **DVILR**, not only Go SDK examples.

**Acceptance:** Build a DVILR HTTP client and server; run concurrent request load and cancellation tests; validate and run real Wasm+WASI modules against conformance fixtures.

### A — AI and games with actual execution (former V0.9)

- [ ] Native DVILR-callable tensor library: dtypes, shape/broadcasting, views/strides, matmul/reductions, memory safety, CPU reference results.
- [ ] Reverse-mode autograd and training example, seeded numerical-gradient validation and model save/load.
- [ ] At least one GPU backend (e.g. Vulkan compute or CUDA) with device detection, safe fallback, supported hardware documentation, CPU/GPU parity and actual GPU integration tests.
- [ ] Window and graphics abstraction (e.g. SDL/OpenGL/Vulkan binding), keyboard/mouse/gamepad input, 2D rendering, frame pacing and resource lifecycle.
- [ ] 3D camera/mesh/rendering basics, physics/collision foundations, scene/entity model and sound; simple playable DVILR-authored demo.
- [ ] Representative performance benchmarks with published hardware and reproducible numbers; no unsupported speed claims.

**Acceptance:** DVILR-written CPU and GPU tensor demo verified on supported GPU hardware, and DVILR-written graphical game launched and tested on supported OS. Go-only SDKs do not count.

### S — True self-hosting (V1.0)

- [ ] Freeze documented Stage 0 seed (Go implementation), source archive, compiler binary, SHA-256 and bootstrapping instructions.
- [ ] Write the complete lexer, parser, type checker, IR lowering, optimizer, code emitter, linker/object writer and diagnostics in `.dvl`; no Go compiler logic in the finished compiler.
- [ ] Stage 0 compiles **DVILR source compiler** into a runnable Stage 1.
- [ ] Stage 1 independently compiles the same compiler source into Stage 2 **without Go**.
- [ ] Stage 2 independently compiles compiler source into Stage 3 **without Go**.
- [ ] Stage 2 / Stage 3 produce identical outputs on a canonical test corpus (and deterministic byte-identical compiler artifacts where feasible and specified).
- [ ] Hermetic CI repeats full bootstrap on supported hosts with pinned dependencies and publishes logs/hashes; fresh-install instructions pass.

**Acceptance:** Artifacts, commands, logs and hashes for Stage 0→1→2→3 exist and are independently reproducible. The current Go-authored compiler does NOT pass this gate.

## Execution order / dependency path

1. **M0 — Import actual compiler:** Commit the v0.6.4 compiler's full Go source, tests, examples and CI to a **dedicated DVILR compiler repository** (or explicitly structured `compiler/` folder); run baseline native Windows tests. Preserve web demo separately. Do not rely on a roadmap-only GitHub repo.
2. **M1 — Language frontend & types:** AST/HIR/MIR, generics, modules, diagnostics. Lock syntax conformance suite.
3. **M2 — Safe core runtime:** Strings, Vec/Map, memory model, errors, file I/O and standard library. Add memory-stress/fuzz corpus.
4. **M3 — Portable native backend:** Windows+Linux target compatibility, precise semantics, ABI, optimizations and debugger support.
5. **M4 — Web/systems:** HTTP/TLS, sockets, task runtime and async safety. Build compiled DVILR client/server examples.
6. **M5 — WebAssembly:** Expanded Wasm/WASI features and conformance tests, including cross-target behavior.
7. **M6 — AI/game runtime:** CPU tensor reference, GPU backend proven on hardware, renderer/input/audio and playable demo.
8. **M7 — Compiler in DVILR:** Lexer→parser→typechecker→IR→backend port, then Stage 1/2/3 bootstrap and independently reproduced proof.
9. **M8 — Release candidate:** CI matrix, security audit, compatibility tests, documentation, package/install/uninstall, signed/checksummed distribution.

The milestones may run in parallel where independent, but **the V1.0 label is blocked by every acceptance gate**. If some gates are removed, call it a smaller release and publish scope changes transparently.

## Evidence policy

Use `planned`, `implemented unverified`, `tested on target`, `complete` statuses only when they reflect observed evidence. A cross-compile is not proof of Windows runtime behavior. Compiling a minimal WASI module is not full WASI support. Running Go libraries is not DVILR library integration. No self-hosting claim before independently reproduced Stage 1/2/3.

## Current baseline snapshot (from v0.6.4 STATUS.json)

- Compiler: Go bootstrap, Windows x64 output; basic bounded arrays/vectors, native sum/dot/fill/reverse, guarded arithmetic and minimal structs/enums/imports.
- Verified in supplied preview: Go tests/race/vet, Windows x64 compiler cross-build and PE generation, numeric Wasm/WASI subset in Node.
- Not verified: generated Windows executable execution on a Windows host; actual GPU execution; self-hosting.
- **Release verdict: V1.0 NOT READY.**

## Immediate next step

Import and build the actual native compiler source in GitHub CI so future milestones can be tracked against executable compiler artifacts, not only against a storefront repository. Require a successful Windows native-runtime job before changing release readiness.
