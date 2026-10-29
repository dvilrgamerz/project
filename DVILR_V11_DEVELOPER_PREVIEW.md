# DVILR v1.1 Developer Preview — Windows Release

**Status: Built and locally tested as a developer-tools preview on October 29, 2026. NOT a complete unified DVILR V1.0 language/AI/game/self-hosting release.**

This document records a new experiment in the Windows developer experience. **The actual installer, VS Code extension, and full source ZIP are ChatGPT conversation download artifacts; they are NOT uploaded to this GitHub repository and must not be represented as GitHub Releases.**

## What is implemented in the separate full-source archive

| Feature | Evidence | Limit |
| --- | --- | --- |
| `dvl repl` | Go tests + Linux-hosted CLI smoke test | Persistent numeric `let`/`var`, checked integer arithmetic, comparisons, `say` and `:help/:vars/:reset/:quit` only; **not** a complete DVILR interpreter |
| `dvl fmt file.dvl [-w]` | Go idempotence, formatting, malformed-input, no-clobber and permissions tests | Brace indentation and lexical validation, **not** complete AST formatting |
| `dvl doctor [--json]` | Go tests + real CLI output | Read-only host OS/architecture/PATH diagnostics, does not test every runtime component |
| `dvl new ProjectName` | Go test + CLI smoke test | Creates `main.dvl`, `README.md`, `dvl.toml`; manifest is **metadata only** |
| VS Code extension | JSON/XML and VSIX archive integrity inspected | Syntax highlighting, brackets, snippets only; **no LSP/debugger** |
| Windows per-user installer | Windows x64 compiler and GUI installer cross-build; PE identified | Unsigned; **actual Windows install/execution/uninstall unverified** |

## How to use it

Run the unsigned `DVILR-Setup-v1.1-Developer-Preview.exe` only if the source is trusted. Installer targets Windows 10/11 x64, installs for the current user, registers an uninstall entry and adds `dvl` to the user PATH. Open a fresh terminal:

```powershell
dvl version
dvl doctor
dvl repl
dvl new FirstProject
dvl fmt FirstProject\main.dvl -w
dvl check FirstProject\main.dvl
dvl run FirstProject\main.dvl
```

REPL example (enter one line at a time):

```text
var score = 10
score = score * 5
say score
:vars
:quit
```

Optional editor extension: install `DVILR-v1.1-VSCode.vsix` using VS Code's **Install from VSIX** command. It is not a compiler and cannot replace `dvl.exe`.

## Artifact verification (SHA-256)

- `DVILR-Setup-v1.1-Developer-Preview.exe` — `45783ce7cf1ed9afd03597bc45c1a6b85f366fdca58c5343b297e8d0bbd120a1`
- `DVILR-v1.1-VSCode.vsix` — `ab30166b59c3d628d67c70ece75900b4a146584971f7515ed84018b4c5806884`
- `DVILR-v1.1-Developer-Preview-Full-Source-Windows.zip` — `bec103084690a909e5a2274e62f18e27d2d2363f607ddf6f33286dc0a5553fb5`

## Checks performed

1. `go test ./...` — passed all Go packages, including new REPL/parser/overflow/formatting/doctor tests.
2. `go test -race ./...` — passed all Go packages.
3. `go vet ./...` — no issues.
4. `go run ./src` smoke test for `version`, `doctor --json`, `repl`, `new`, `check`, `build`, `fmt` — passed.
5. `GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build` compiler and GUI installer — passed; both identified as Windows x64 PE.
6. Complete source ZIP and optional VS Code VSIX ZIP integrity checked — passed.
7. **NOT TESTED on real Windows:** installer run/uninstall, executable output, VS Code UI integration.
8. **NOT IMPLEMENTED / UNVERIFIED:** complete memory-safe language runtime, native GPU execution, full AI/game/Web SDK, compiler written in DVILR, Stage 1→2→3 compiler self-hosting.

## Developer source integration

The complete source is in the downloadable ZIP. The repository's existing `compiler/` folder still contains only partial source migration; full source has not been merged, so GitHub CI cannot verify this version yet.

**Next release gates:** import the whole compiler source, run the Windows installer end-to-end GitHub workflow, add REPL/editor tests on a Windows host, and continue the unified V1.0 compiler work rather than treating developer tooling as a substitute.
