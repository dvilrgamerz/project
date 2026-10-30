# DVILR v1.1.5 — Unicode lexer and native checked vector slices

**Status: Developer Preview; full P0–P6 goals remain incomplete.**

## Implemented

- Lexing uses Unicode code points for token columns and accepts UTF-8 identifier characters, with Unicode strings preserved.
- `slice_vector(destination, source, start, count)` in generated Windows x64 code validates an integer start/count and uses unsigned bounds checks to reject negatives and end-overruns before memory access (native exit status 8).
- Destination owns a copy of the selected i64 elements. Old discarded destination elements are zeroed. In-place slicing (`source == destination`) is supported.
- Native vectors still hold at most 1024 i64 elements in statically allocated backing memory. No heap, generics, GC or ownership system is claimed.
- Added Go lexer/codegen acceptance/rejection tests; Windows PowerShell executable verification script and Windows CI workflow.
- VS Code syntax highlighting and code snippets recognize the new intrinsic.

## Locally verified on Linux build host

- `go test ./... -count=1` — passed, 8 packages.
- `go test -race ./... -count=1` — passed, 8 packages.
- `go vet ./...` — passed.
- Native compile-to-PE of `examples/v115/vector_slice_unicode.dvl` — passed, 152,064-byte Windows x64 PE.
- Checked emitted x64 unsigned branch and indexed memory instructions via `objdump` — passed static code inspection, NOT runtime execution.
- Compile-time rejection of a boolean start index — passed.
- Cross-built Windows amd64 compiler and GUI installer; source archive integrity tested.

## What is NOT verified

- Running a generated `.exe` on a real Windows machine or GitHub Windows runner.
- Windows GUI installer install/uninstall operation. The installer is unsigned.
- Full generics/type checker, dynamic safe heap, native networking/Wasm completion, GPU/game SDK, and self-hosting.

## Verify with a Windows machine

From the full source root:

```powershell
.\tools\verify-v114-windows.ps1 -Compiler .\installer\payload\dvl.exe
.\tools\verify-v115-windows.ps1 -Compiler .\installer\payload\dvl.exe
```

The v1.1.5 script requires output `SLICE OK,50`, checks three out-of-range programs exit with code 8 and rejects a boolean index at compile time.

The GitHub Actions workflow is in `.github/workflows/v115-native-windows.yml` in this source package, but the user's current GitHub `project` repository contains only a partial compiler import. A successful GitHub run requires committing the entire buildable tree first.

Source-archive SHA-256: `50eab8c5e04114f10255cfee7b10cf07f76f67af00335ce570d97a84b0d719f2`.
Unsigned installer SHA-256: `0281442aa9159850fb8e2f814aff013de02b1c7340cce79f379b5faaa8b95317`.

These artifacts are attached to the ChatGPT conversation, **not** hosted on GitHub. The complete V1.1.5 compiler source has not been imported into this repository.
