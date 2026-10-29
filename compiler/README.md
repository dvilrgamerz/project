# DVILR native compiler source migration

**Status: INCOMPLETE IMPORT.** This folder is not yet a standalone buildable Go module in GitHub. It contains verified v0.6.5 developer-tool source and tests only, not the complete compiler implementation.

## Present here

- [src/tooling.go](src/tooling.go): `newProject`, `parseTokenArgs`, `writeTokens` helpers.
- [src/tooling_test.go](src/tooling_test.go): Go regression tests for project creation and lexer token output. These depend on `lex`, `compile` and `makePE` supplied by the full native compiler source.

## Source of truth for the preview

The complete DVILR Standalone v0.6.5 preview has been built in the associated ChatGPT conversation and supplied as **DVILR-Standalone-v0.6.5-Developer-Tools-Source-Windows.zip**. It includes `go.mod`, `src/main.go`, `src/wasm.go`, `src/modules.go`, `src/shop.go`, `sdk/`, `examples/`, workflows and other files.

The web storefront at this repository root is a separate demonstration and not the standalone compiler.

## How to migrate into a complete buildable GitHub source tree

1. Extract the preview ZIP into a working directory.
2. Copy the complete compiler source (all files, including `go.mod`, `src/`, `sdk/`, `examples/`, `tools/`) into this `compiler/` directory in the repository. Keep `.github/workflows/` at repository root and set the workflow working directory to `compiler`.
3. Ensure none of the files in `compiler/src/` here were independently modified after the ZIP was built.
4. Run the commands from `compiler/`:

```sh
go test ./...
go test -race ./...
go vet ./...
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -o dvl-v065-preview.exe ./src
go run ./src tokens --json examples/hello.dvl
```

5. Run real compiled Windows native programs in GitHub's Windows runner. A cross-build alone does not verify generated PE execution.

## New V0.6.5 commands

```text
dvl new MyProject
dvl tokens examples/hello.dvl
dvl tokens --json examples/hello.dvl
```

Token inspection supports a future language server; it is **not** an implemented VS Code language server.

## Release truth

DVILR V0.6, V0.8, V0.9 and V1.0 are **not complete**. The compiler is written in Go, and no Stage 1→2→3 self-hosting proof or GPU execution exists.
