# DVILR v1.1 unified release — locally verified implementation checkpoint (October 29, 2026)

**Release level: Developer Preview. This does not complete V0.6/V0.8/V0.9/V1.0 or the final unified v1.1 specification.**

## Local implementation in the separate downloadable source ZIP

- New `dvl test [path]` command enumerates a bounded, sorted collection of up to 128 `.dvl` files, compiles each, validates Windows PE generation, and reports per-file pass/fail and totals.
- `dvl test [path] --run` explicitly requests native execution on Windows amd64, creates per-test temporary binaries and terminates tests after 5 seconds. It declines to run on Linux and non-Windows systems.
- Compiler lexical input now rejects invalid UTF-8, with valid Unicode still supported in string literals.
- New valid program fixtures, invalid program regression tests and compiler CLI tests; updated experimental Windows installer includes the code.
- `V11_UNIFIED_RELEASE_REQUIREMENTS.md` contains the full release requirements from former V0.6, V0.8, V0.9 and V1.0 milestones.

## Verification

| Check | Result |
| --- | --- |
| `go test ./...` | PASS |
| `go test -race ./...` | PASS |
| `go vet ./...` | PASS |
| `go run ./src test tests/dvl` | PASS (2/2 compile-and-PE checks) |
| Cross-compile standalone Windows amd64 compiler | PASS |
| Cross-compile Windows amd64 graphical installer | PASS |
| Source ZIP integrity | PASS (114 entries) |
| Run installer on real Windows | **NOT TESTED** |
| Run `dvl test --run` on real Windows | **NOT TESTED** |
| Native DVILR GPU workloads | **NOT IMPLEMENTED** |
| Compiler Stage 1→2→3 self-hosting | **NOT IMPLEMENTED** |

## Artifact names and SHA-256 (separate ChatGPT downloads, not uploaded to GitHub)

- `DVILR-v1.1-Unified-Goals-Compiler-Checkpoint.zip` — `a482bb80883f6d1657b2d0c5323e37cab04eff1f7c4fcd7a22a5309073947c3d`
- `DVILR-Setup-v1.1-Unified-Developer-Preview.exe` — `3f3eca17164c939d9e5b05f99979ef86309bb621e8101ba69a79695571c235fa`

**Source distribution note:** Updated compiler code, installer and tests are inside the ZIP. This GitHub storefront/demo repository still does not contain all Go compiler source, and the new test command source has not been merged here. The source ZIP is the runnable/buildable snapshot; do not use the GitHub docs alone to attempt a compiler build.

## Immediate next gates

1. Import full Go compiler source into a dedicated source tree/branch on GitHub and run Windows native end-to-end CI.
2. Finish the compiler's type system, memory safety, generics and dynamic collections.
3. Implement DVILR-level networking/async and complete Wasm/WASI.
4. Integrate CPU/GPU tensors and verified native game/graphics code.
5. Rewrite the compiler in DVILR and independently verify Stage 1→2→3.
