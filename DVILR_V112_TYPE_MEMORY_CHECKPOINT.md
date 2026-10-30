# DVILR v1.1.2 — Native Type and Memory Safety Checkpoint

**Status: Developer Preview; full DVILR language and memory system are NOT complete.** October 30, 2026.

This checkpoint answers the requested sequence: **test generated Windows executables → improve type/memory safety → defer GPU/self-hosting**. A Windows runtime is **not available** in the build environment, and the full compiler source still lives in the downloadable release ZIP. Actual Windows execution remains unverified; no runtime pass is claimed.

## Source changes in the separately packaged Go-bootstrapped compiler

- `checked_add(a,b)`, `checked_sub(a,b)`, `checked_mul(a,b)`: signed i64 arithmetic with emitted native x86-64 overflow-flag checks. Program exits with status **10** if signed arithmetic overflows; ordinary operators keep their previous semantics.
- `copy_vector(destination, source)`: compile-time requirement that both arguments are declared bounded i64 vectors (capacity 1,024); emits a deep-copy of active elements. No alias created. Destination length is updated and any discarded old destination elements are zeroed.
- `clear_vector(v)` now overwrites active elements with zero before setting length to zero.
- `pop(v)` now zeros the removed element to prevent stale backing-storage contents.
- Added Go compiler regressions for successful/invalid types, arithmetic machine instructions, deterministic PE generation, bounds behavior and zeroization.
- Added Windows PowerShell output/exit-code test script and a proposed GitHub Actions `windows-2022` native-execution job inside the source ZIP.

Source patch relative to the **v1.1.1 full compiler archive**: [v112_native_safety.patch](compiler/patches/v112_native_safety.patch). Sample: [v112_memory_safe_copy.dvl](examples/v112_memory_safe_copy.dvl).

**Important:** This GitHub repository is still the storefront/demo plus partial compiler migration. The patch **cannot compile on its own** here; apply it to the full v1.1.1 source, or use the v1.1.2 source ZIP downloaded in ChatGPT.

## Local verification

| Verification | Result |
| --- | --- |
| `go test ./... -count=1` (7 Go packages) | PASS |
| `go test -race ./... -count=1` | PASS |
| `go vet ./...` | PASS |
| Go Windows x64 compiler cross-build | PASS |
| Go Windows x64 GUI installer cross-build | PASS |
| Native source example compiled into a 160,768-byte Windows x64 PE | PASS |
| `llvm-objdump`: indexed copies, zeroing, signed arithmetic `jno` guard | PASS (static inspection only) |
| Source ZIP integrity test (129 members) | PASS |
| Windows PE actually **executed on Windows** | NOT VERIFIED |
| Installer actually installed/uninstalled on Windows | NOT VERIFIED |
| Windows GitHub Actions job | WRITTEN, NOT RUN |

## Windows executable test procedure

After extracting the **full source ZIP** on a Windows 10/11 x64 machine:

```powershell
.\installer\payload\dvl.exe version
.\installer\payload\dvl.exe build .\examples\v112\memory_safe_copy.dvl -o memory-safe-copy.exe
.\memory-safe-copy.exe
powershell -NoProfile -ExecutionPolicy Bypass -File .\tools\verify-v112-windows.ps1 -Compiler .\installer\payload\dvl.exe
```

Expected output for the successful program:

```text
COPY AND ARITHMETIC OK
4
12
```

The verification script additionally tests **three intentional arithmetic overflows**, each requiring native exit code **10**, and checks compile-time rejection of a vector/array type mismatch. Its packaged Windows CI workflow can perform the same test after the entire compiler tree is committed to a GitHub repository.

## Artifact hashes (downloads delivered in ChatGPT, not GitHub Releases)

- `DVILR-v1.1.2-Type-Memory-Safety-Full-Source-Windows.zip` SHA-256: `55746ab96b4dc4c48fa79a26789644a00f5df09b9ff9ecf63b7f05997e69690b`
- `DVILR-Setup-v1.1.2-Memory-Safety-Preview.exe` SHA-256: `7f0743ddeb9eab19f434ef93aa170e2b6d991c635161cf276ede254f3a9403ab`

The Windows installer is unsigned. Verify file hashes and review its source before running.

## Remaining critical engineering blockers

1. Complete semantic/type system: typed AST, generics, type inference/contracts, borrow/reachability rules.
2. General-purpose safe dynamic heap collections, allocation/reference discipline and memory-safety validation under stress/fuzzing.
3. Windows runtime tests must actually pass. Linux-target native backend and compatibility tests remain missing.
4. Native DVILR networking/concurrency and broader Wasm/WASI.
5. GPU execution and game SDK, then verified Stage 0→1→2→3 DVILR self-hosting.

**No GPU or self-hosting code was added in this checkpoint.** Those milestones should wait until compiler/runtime safety is established.
