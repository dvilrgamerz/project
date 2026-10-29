# DVILR v0.6.6 — Windows Installer Preview

**Release status: experimental. DVILR V1.0 is not complete.** The standalone compiler is Go-bootstrapped and does not compile itself.

## What's new in the downloadable v0.6.6 package

A **clickable 64-bit Windows setup executable** was compiled in the ChatGPT build workspace and delivered as:
- `DVILR-Setup-v0.6.6.exe` (unsigned, GUI subsystem)
- `DVILR-v0.6.6-Windows-Installer-and-Full-Source.zip` (complete compiler source, installer source, tests, docs, examples and binaries)

**The binary and complete source ZIP are conversation artifacts, not uploaded to this repository.** This GitHub page records the release evidence and instructions; it is not a GitHub Release download.

SHA-256 of the installer executable:

`7af6cce477dae6e76b6d29261446e2da99c82024f15c81a6e8be7642f3dd3161`

The ZIP includes `SHA256SUMS.txt` with per-file checksums.

## One-click install (similar to Python setup)

1. Run the installer on 64-bit Windows 10/11.
2. Click Yes. It installs to `%LOCALAPPDATA%\Programs\DVILR` with **no administrator rights**.
3. It adds the installation directory to the **current user's PATH** and registers an entry under Windows Installed Apps.
4. Open a **new** terminal to pick up PATH.

```powershell
dvl version
dvl help
dvl new HelloDVILR
dvl check HelloDVILR\main.dvl
dvl run HelloDVILR\main.dvl
```

The compiler emits standalone Windows x64 PE programs. End users do not need Python, Go, LLVM or a C/C++ compiler.

To uninstall, select DVILR Programming Language (Preview) under Windows Installed Apps or run `DVILR-Setup.exe -uninstall` from its installation directory.

**Caution:** The installer is unsigned; Microsoft SmartScreen may display a warning. Verify the downloaded checksum before running.

## Local verification performed on October 29, 2026

| Test | Result |
| --- | --- |
| `go test ./...` | PASSED |
| `go test -race ./...` | PASSED |
| `go vet ./...` | PASSED |
| Cross-compile Windows 64-bit compiler and GUI setup PE | PASSED |
| Create new project, check its source, emit Windows PE | PASSED using Linux-hosted Go bootstrap |
| Export JSON lexer tokens | PASSED |
| Compile numeric WASI module and run in Node.js | PASSED (45) |
| Complete ZIP archive integrity | PASSED |
| Run installer on an actual Windows host | **NOT TESTED** |
| Execute generated Windows-native program | **NOT TESTED** |
| Run GitHub Windows end-to-end workflow | **NOT RUN** |

The ZIP contains a proposed `.github/workflows/windows-installer-e2e.yml` that can test installation, native execution and uninstall when the full compiler source is checked into a suitable repository.

## What this is not

- Not Python-compatible; no `pip`, Python import compatibility or Python-style REPL.
- Not complete memory-safe generic containers or full type system.
- Not a native graphics/game/AI GPU ecosystem.
- Not complete Wasm/WASI, native networking/concurrency or a self-hosting compiler.
- No Stage 0→1→2→3 self-hosting proof.

**Next priority:** Check the full compiler source into a canonical repository; run the Windows installer end-to-end CI before treating it as production-ready. Then develop native language/runtime and verify true self-hosting.
