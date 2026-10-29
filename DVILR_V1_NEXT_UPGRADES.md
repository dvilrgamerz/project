# DVILR — Next upgrades to add to the V1.0 goal

**Planning document only. Nothing below is claimed implemented or tested.**
Last updated: 2026-10-29. See [DVILR V1.0 Unified Release Goal](DVILR_V1_UNIFIED_RELEASE_GOAL.md) for the language, systems, AI/game, and self-hosting release blockers.

## Goal

Deliver one **complete, independently buildable DVILR V1.0 developer toolchain**: a safe language, practical native/WebAssembly backends, useful standard libraries, networking/concurrency, CPU/GPU compute, game rendering, and a genuinely self-hosted compiler. In addition, it must be pleasant to install, edit, debug, test, package, and deploy real software.

## Priority zero: get the real compiler into version control

**Next implementation step, not a future stretch goal.**

- [ ] Create a dedicated DVILR compiler source location. Import complete source from the latest standalone preview, with appropriate rights and intact history where available. The current `project` repository is a **storefront + demo**, not the full compiler.
- [ ] Check in a reproducible Go Stage 0 bootstrap build, source and checksum; establish the canonical golden-test corpus.
- [ ] Add GitHub Actions: Windows x64 compiler + generated .exe execution; Linux x64 compiler and native programs; WebAssembly and WASI runtime tests; invalid-input/fuzz tests.
- [ ] Block merges on tests, format, static analysis, license/security scans, and exact documented performance baselines.
- [ ] Publish a truthful, machine-readable `feature-status.json` from CI instead of hand-editing "Complete".

**Exit condition:** Fresh clone + documented commands reproduce the compiler and actually execute compiled test programs on supported hosts. A cross-compile alone does not pass.

## Additional V1.0 upgrades (new mandatory developer-experience gates)

### V1-DX1 — Production CLI, builds, tests and documentation

- [ ] Stable `dvl new`, `build`, `run`, `check`, `test`, `fmt`, `doc`, `bench`, `clean` commands.
- [ ] Standard `dvl.toml` project manifest, pinned compiler/runtime version, reproducible package lockfile and clear exit codes.
- [ ] Built-in formatter, test discovery, assertions, coverage reporting, JSON diagnostics and help/man pages.
- [ ] Official project templates: CLI utility, native HTTP server, browser/WASI app, GPU tensor demo and game.
- [ ] Deterministic compile output on pinned environments and useful dependency/build-cache diagnostics.

**Acceptance:** A brand-new user can install DVILR and create/build/test/run all five templates following only the written instructions.

### V1-DX2 — Editor, debugging and diagnostics

- [ ] Language Server Protocol (LSP) for syntax diagnostics, completion, navigation, rename and symbol references.
- [ ] VS Code extension with syntax highlighting and LSP integration; versioned editor/runtime compatibility.
- [ ] Native debug info, breakpoints, step in/out, stack traces and variable inspection through an identified debugger backend.
- [ ] Clear line/column diagnostics with fix suggestions; source-map equivalent for WebAssembly where supported.

**Acceptance:** Automated editor-protocol tests and a manual debug session with screenshots/logs from supported Windows and Linux hosts.

### V1-DX3 — Packages, modules and interoperability

- [ ] Secure module/package resolver, offline installs, lockfile checks, integrity hashes and dependency-cycle diagnostics.
- [ ] Local/file/git dependencies first; public registry is an optional hosted service, not a hidden release dependency.
- [ ] Stable C ABI interoperability: `extern`, FFI types/calling convention, callback lifetime rules, explicit unsafe boundaries.
- [ ] Library versioning, semver compatibility policy, API docs and signed/checksummed distribution artifacts.

**Acceptance:** Rebuild a multi-module project offline; consume a small C library safely; produce and verify reproducible package hashes.

### V1-DX4 — Security, tooling and portability

- [ ] Document DVILR's ownership/GC model and exact memory/thread safety invariants; implement its checks in compiler + runtime.
- [ ] Safe-by-default stdlib networking: HTTPS verification, timeout/cancellation, structured errors, resource caps and fuzz tests.
- [ ] Windows and Linux x64 release installers, portable archives and fresh-machine test matrix. macOS and arm64 only when tested.
- [ ] Compiler crash/invalid-source fuzzing, dependency/license review, reproducibility checks, SBOM and coordinated disclosure policy.

**Acceptance:** Automated sanitizer-equivalent tests, reproducible installers, protected CI release workflow and published test results.

### V1-DX5 — Real integrated showcase apps

- [ ] **DVILR HTTP server**: compiled .dvl code responds to JSON requests under concurrent load.
- [ ] **DVILR AI model**: train a small MLP on CPU and demonstrate GPU-vs-CPU numerical parity where supported GPU hardware exists.
- [ ] **DVILR native 2D game**: input, rendering, audio, physics/collision, saves, frame-time statistics and reliable shutdown.
- [ ] **DVILR browser/WASI app**: same language library compiles and executes under documented host/WASI permissions.
- [ ] **DVILR builds DVILR**: independently reproduced Stage 0→1→2→3 compiler builds with hashes, test parity and logs.

**Acceptance:** Each showcase is source-controlled, built from scratch on supported targets, and recorded in CI or a documented test report. Go-only prototypes and numerical-only Wasm demos do not qualify.

## Staged delivery (do not mark the overall release complete early)

| Workstream | Priority | Dependency | Release gate |
| --- | --- | --- | --- |
| Import compiler source + CI | P0 | Latest verified compiler ZIP | Windows native runtime tests pass |
| Typed IR, safe heap, dynamic generics | P0 | Compiler repository | Positive/negative/memory tests pass |
| CLI, formatter, project layout | P0 | Reproducible compiler builds | New-project template compiles |
| Modules, packages, standard library | P1 | Safe type/runtime model | Offline reproducible build |
| HTTP/TLS, sockets, threads, channels | P1 | Runtime safety + stdlib | DVILR server stress test |
| Complete Wasm/WASI target | P1 | Typed IR + stdlib boundary | Real compiled modules execute |
| LSP + debugger + VS Code | P1 | Compiler diagnostics and debug info | Editor/debug integration tests |
| CPU tensors + autograd | P1 | Generic collections + math runtime | Gradient parity test |
| GPU compute + graphics/audio/game SDK | P2 | FFI/device/runtime foundation | Actual hardware demos pass |
| DVILR compiler self-hosting | P0 release blocker | Complete language + codegen | Stage 1→2→3 reproducibly pass |
| Release hardening and distribution | P0 release blocker | All above | Install/test/security + docs |

P0/P1/P2 are **execution order**, not permission to drop any V1.0 requirement. To release earlier, explicitly reduce scope and call it a preview or a lower version.

## Possible upgrades *after* V1.0 (not added as blockers)

- **V1.1:** Fast REPL, hot reload and interactive notebook kernel.
- **V1.2:** Official package registry, docs portal, security advisories and dependency auditing.
- **V1.3:** Additional validated targets: macOS, arm64, mobile and extra graphics/GPU backends.
- **V2.0:** Distributed runtimes, higher-performance JIT/AOT hybrid optimization, richer autograd/ML compiler and large-scale game tooling.

These are ideas, not promised release dates.

## Test-first rules

1. Before implementing a feature, add a spec: syntax, types, semantics, unsupported behavior and error codes.
2. Add passing examples **and failing examples** including malformed input, edge values, empty containers and concurrency hazards.
3. Verify generated *native* binaries on each named OS; verify Wasm modules in a real engine; verify GPU kernels on real supported hardware.
4. Track a test artifact/commit/hash for each completed item. `Done` means verified, not "code exists".
5. Do not claim "self-hosting" until Stage 0→1→2→3 actually passes independently with published build outputs.

## Immediate recommended implementation sprint

**Sprint A — Bootstrap a real compiler repository**: commit the V0.6.4+ source, add repeatable Windows/native build-and-run CI, write diagnostic expectations and freeze the Stage 0 seed. Then implement the typed IR and safe generic collection model, which unblock almost everything else.
