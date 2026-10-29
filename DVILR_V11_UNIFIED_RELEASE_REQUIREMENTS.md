# DVILR v1.1 — Unified Complete Release Requirements

**Current status: Developer Preview — NOT COMPLETE.** This document adds all formerly separate V0.6, V0.8, V0.9, and V1.0 requirements to the **final v1.1 release target**. It is not evidence that those features already work.

## A. Core language (former V0.6)
- [ ] Complete checked type system with generics, scopes, struct/enum/union layout and diagnostics.
- [ ] Real UTF-8 dynamic strings, heap-backed generic Vec/Slice/Map/Set and bounds-checked slicing.
- [ ] A documented memory-safe reference model (ownership/borrowing or GC), verified with alias, lifetime, leak and concurrency stress tests.
- [ ] Result/Option-style error handling, packages, typed I/O and portable standard library.
- [ ] Typed IR verifier, optimizations, debugging information, Windows and Linux x64 native backend.

**Existing partial subset:** bounded i64 arrays/vectors, static and limited dynamic strings, basic enums/structs/imports, guarded arithmetic and PE output.

## B. Web and systems (former V0.8)
- [ ] DVILR-callable TCP/UDP, DNS, HTTP client/server, TLS verification, streaming, cancellation and structured errors.
- [ ] DVILR-native threads/tasks, futures/channels, mutex/atomics, bounded concurrency and race-safe sharing.
- [ ] Wasm/WASI backend supporting the full implemented language core and real memory/string/call/host interfaces.
- [ ] DVILR-authored HTTP server, concurrent tasks and WASI application run in target integration tests.

**Existing partial subset:** numeric Wasm/WASI, experimental OS-specific APIs and Go SDK foundations; not a complete DVILR runtime.

## C. AI and games (former V0.9)
- [ ] DVILR tensors with dtype, shape/strides, matmul, broadcasting and autograd, tested CPU reference results.
- [ ] Working GPU backend on documented real hardware, with device detection, CPU fallback and numerical parity tests.
- [ ] Hardware graphics, 2D/3D rendering, window/input/audio, physics, resource lifecycle and playable DVILR game.
- [ ] Training + game demos built from DVILR source and reproducible hardware benchmarks.

**Existing partial subset:** Go CPU tensor/graphics scaffolding and native integer vector reductions. **GPU not implemented.**

## D. Fully self-hosting compiler (former V1.0)
- [ ] Freeze Go-built Stage 0 with source, SHA-256 and reproducible bootstrap.
- [ ] Compiler lexer/parser/semantics/IR/backend written *in DVILR*.
- [ ] Stage 0 builds Stage 1; Stage 1 independently builds Stage 2 **without Go**; Stage 2 builds Stage 3.
- [ ] Stage 2/3 deterministic equivalence, logs, binaries, hashes, repeatable fresh-host CI proof.

**Existing status:** Go-bootstrapped compiler; **no Stage 1→2→3 proof**.

## E. v1.1 developer experience and verification
- [x] Experimental Windows installer, REPL subset, formatter, doctor, project generator and VS Code syntax highlighting.
- [x] Local preview adds `dvl test [path]` with deterministic compile-and-PE checks; explicit `--run` only on Windows with five-second timeout.
- [x] Frontend rejects malformed UTF-8; valid Unicode string literals accepted.
- [ ] Full compiler source committed and built in GitHub CI (currently source is largely in separate ZIPs).
- [ ] Real Windows installation, execution, uninstall, fuzz and security tests passing in CI.
- [ ] Installer signing, package manager, LSP, debugger and proper native SDK integration.

## Acceptance gates for **v1.1 Complete**
1. Every section A–D must be proven through tests, not just an API sketch or Go-only library.
2. Publish native Windows/Linux, Wasm/WASI, CPU/GPU and graphical-game results from supported hosts.
3. Publish a reproducible Stage 1→2→3 self-hosting build.
4. Release a checked-in, reproducible source tree, documentation, installers and logs.
5. Only after *all* gates pass update status to COMPLETE. Until then use **v1.1 Developer Preview**.

## Engineering order
P0: Import full compiler source, CI and Windows native tests. P1: Typed runtime, safe containers, modules. P2: Networking/concurrency and Wasm/WASI. P3: Tensors/GPU and graphics/game systems. P4: Full compiler rewrite in DVILR and self-hosting. P5: Hardening and final release.

**Latest local checkpoint:** Go tests/race/vet pass; new `dvl test` compile-only examples pass (2/2); Windows compiler and installer cross-build pass. Windows-native runtime, GPU and self-hosting proof remain **UNVERIFIED/INCOMPLETE**.
