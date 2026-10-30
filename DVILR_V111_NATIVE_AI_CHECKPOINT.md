# DVILR v1.1.1 — Native AI Score-Ranking Compiler Preview

**Status: EXPERIMENTAL COMPILER PREVIEW — NOT full DVILR v1.1.** The updated full source and installer are available as ChatGPT downloadable artifacts, not GitHub-hosted Release files. The GitHub `project` repository remains a partial compiler-source migration plus web demo.

## Actually implemented in the Go-bootstrapped native compiler

- `max_value(values)`: largest signed 64-bit integer among the current elements of an array/vector.
- `min_value(values)`: smallest signed 64-bit integer.
- `argmax(values)`: zero-based index of the **first** largest integer value (ties use earliest index).
- Runtime x64 code generation reads mutated array/vector memory and dynamic vector lengths, not a compile-time snapshot.
- Type checks reject non-array/vector arguments; empty fixed arrays are rejected at compile time. Empty vectors exit the generated program with status 8.
- Added native code generation, deterministic PE and negative-case tests.
- Updated VS Code highlighting/snippet, Windows compiler/GUI installer, and unexecuted Windows CI to check native example results.

See [the DVILR example](examples/native_ai_rank_v111.dvl) and [the compiler patch](compiler/patches/v111_native_ai_ranking.diff). The patch targets `src/main.go` of the **v1.1 Unified Goals Compiler Checkpoint** full source archive; it cannot be applied to this GitHub repository alone because the full compiler is not yet checked in.

## Example

```text
vector scores = [-10, 11, 19, 19, 7]
say argmax(scores)
say max_value(scores)
say min_value(scores)
```

Expected output on Windows:

```text
2
19
-10
```

## Local verification (October 29, 2026)

| Check | Result |
| --- | --- |
| `go test ./... -count=1` | PASS, 7 Go packages |
| `go test -race ./... -count=1` | PASS |
| `go vet ./...` | PASS |
| Compile valid/invalid ranking test source | PASS |
| Deterministic PE binary, x86-64 instruction sequence | PASS |
| Inspect generated x86-64 disassembly | PASS: signed comparison branches and indexed loads |
| Compile the new example to standalone Windows x64 PE | PASS (143,360 bytes) |
| Cross-build Windows x64 compiler and GUI installer | PASS |
| Execute generated ranking program on Windows | **NOT VERIFIED** |
| Run Windows installer or Windows CI | **NOT VERIFIED** |
| Tensor GPU, graphics/game SDK, self-hosting stages | **NOT COMPLETE** |

## Artifact SHA-256 checksums

- `DVILR-Setup-v1.1.1-AI-Preview.exe`: `c6f2313fdd8105cca0ddbea47889cd50d5adccb428aa0088fe4454bdddeb5e98`
- `DVILR-v1.1.1-Native-AI-Compiler-Source.zip`: `bb5b753e4f0aac810be2df98efaaee1bb5f0fd2523453399025eb0d63b997815`
- `DVILR-v1.1.1-VSCode.vsix`: `d250af1871bbdf53d589ec1529f4722d62bd0573c8ccf37fa336247b8817d876`

## Windows usage

Inspect and verify the **unsigned** installer locally, then run it on Windows x64 and open a new terminal:

```powershell
dvl version
dvl doctor
dvl check rank.dvl
dvl run rank.dvl
```

Save the above sample as `rank.dvl`. The standalone v1.1 compiler will *not* recognize the new ranking builtins; use v1.1.1.

## Remaining unified DVILR v1.1 blockers

- Complete type system, memory safety and generic dynamic collections.
- Networking/async APIs callable directly from general DVILR programs; broad Wasm/WASI support.
- Tensor autograd, verified GPU kernels and hardware game graphics.
- Compiler source written in DVILR and verified reproducible Stage 0→1→2→3 self-hosting.

**Verdict:** one verified native compiler-feature increment, not a completed language release.
