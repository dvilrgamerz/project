# DVILR v1.1.4 Engineering Status — P0 to P6

**Status: Developer Preview, not a complete v1.1 language release.** Checked October 30, 2026.

## This checkpoint: actual native compiler change

The separate v1.1.4 full-source archive includes an updated Go-bootstrapped standalone compiler that now supports user-defined `fn` signatures containing explicit `bool` and `i64` parameters and `bool` returns, with static checks at function call sites and return expressions. Forward references are supported. Compatibility with historical numeric return values (boolean-as-0/1 for `->i64`) is retained, which is **not** strict whole-language typing.

**Locally executed and passed:** `go test ./... -count=1`, `go test -race ./... -count=1`, `go vet ./...`, compiler acceptance/rejection fixtures, Windows x64 cross-build of compiler and installer, generated PE inspection and ZIP integrity validation. **Real Windows PE execution, installation/uninstallation and Windows GitHub Actions CI remain NOT VERIFIED.**

Separate local artifacts provided in this chat:
- `DVILR-v1.1.4-Typed-Functions-Compiler-Source-Windows.zip` — source, test suite, Windows runner workflow, compiler payload.
- `DVILR-Setup-v1.1.4-Typed-Functions-Preview.exe` — unsigned Windows GUI installer.

SHA-256 (source ZIP): `a0f5e0c46fdb1790c7fb2680f614a6dafc1222f18be389df49bc4b93a14cfce8`
SHA-256 (installer): `5cea4883dac6a335629a51b15ac3a82695390109370e50d295f04254d047def2`

## All P0–P6 priorities tracked as actionable issues

| Priority | Scope | GitHub | Current status |
|---|---|---|---|
| P0 | Import complete compiler source + run native Windows CI and installer | [#1](../../issues/1) | **Blocked:** GitHub source tree is incomplete |
| P1 | Complete type system, typed structs and generics | [#2](../../issues/2) | **Partial**; bool functions added in v1.1.4 |
| P2 | Native safe memory, heap and resizable generic collections | [#3](../../issues/3) | **Go-hosted reference only** |
| P3 | Standard library, package manager, LSP and debugger | [#4](../../issues/4) | **Partial** |
| P4 | DVILR HTTP/network/concurrency and expanded Wasm/WASI | [#5](../../issues/5) | **Partial** |
| P5 | GPU tensors, audio/graphics and playable game SDK | [#6](../../issues/6) | **Not complete** |
| P6 | DVILR-authored compiler; reproducible self-hosting stages | [#7](../../issues/7) | **Not implemented** |

## Windows verification already committed

- [Native example](compiler/examples/v114/typed_functions.dvl) expects three output lines `1`, `0`, `42`.
- [Windows verification script](compiler/tools/verify-v114-windows.ps1) compiles and **executes** that Windows PE, checks output and rejects three deliberately ill-typed programs.
- [Manual Windows CI workflow](.github/workflows/dvilr-native-windows-v114.yml) runs Go tests, builds the compiler, executes the PE, and silently installs/uninstalls the GUI setup.

### Critical missing prerequisite

The `project` repository contains only a **partial** compiler import under `compiler/`, not the source files from the complete v1.1.4 ZIP. The workflow deliberately refuses to run without `compiler/go.mod` and `compiler/src/main.go`. **No successful GitHub Windows execution run exists yet.** Do not trigger the workflow expecting success until the full source is copied under `compiler/`.

After full source import, on GitHub open **Actions → DVILR Native Windows v1.1.4 → Run workflow** and inspect the actual logs, native output and installer test. A compiled Windows executable running under Windows is necessary before P0 can be marked done.

## Release integrity policy

Do not interpret a roadmap, GitHub issue, Go SDK, generated PE file, test script, or Linux cross-build as a completed feature. Track each priority as COMPLETE only after its actual runtime acceptance test passes. P6 specifically requires a demonstrable Go-free Stage 1→2→3 sequence.
