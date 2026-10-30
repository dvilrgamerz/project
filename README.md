# DVILR Market — experimental DVILR shopping website

This is a **demo storefront**, designed to be deployed as a static website on Vercel. Product declarations originate from DVILR-style source in `examples/shopping_site.dvl`; the current web compiler emits HTML, CSS and JavaScript for ordinary browsers.

**Not an actual store:** products are fictional, cart state stays in the browser, and checkout is a demo. Do not submit payment information.

## Run
Open `index.html` in a browser. No Python, C++, Node, or package installation is needed to view the static site.

## Deploy
Import this repository into Vercel. Use Framework: Other; Root Directory: project root; Build Command: none; Output Directory: project root. Vercel deploys `index.html`.

## Compiler status
The DVILR V0.5 production compiler is still authored in Go; compiler self-hosting (stage 1→2→3) is **not achieved**. The graphical game/AI SDKs and general Wasm SDK are **not finished**. This repository currently contains the storefront, DVILR catalog sample, and demo Wasm artifacts, not the full Go compiler implementation.

## DVILR V1.0 unified release goal

The unified goal for the standalone DVILR compiler and native SDKs is documented in [DVILR V1.0 Unified Release Goal](DVILR_V1_UNIFIED_RELEASE_GOAL.md). This is a **development specification**, not a completed compiler release. The standalone compiler source is currently distributed separately from this storefront demo repository.

## DVILR V1.0 next upgrades

Additional V1.0 acceptance requirements for the developer toolchain (CLI, package manager, LSP/debugger, safety, distribution and proof-driven showcase apps) are in [DVILR V1.0 Next Upgrades](DVILR_V1_NEXT_UPGRADES.md). These are **planned**, not implemented.

## DVILR Standalone v0.6.5 developer-tooling preview

The experimental compiler gained `dvl new ProjectName` and `dvl tokens --json file.dvl`, with Go unit/race tests and a Windows x64 cross-build. The new [compiler source migration folder](compiler/README.md) holds the tool implementation and its tests. **The full compiler source is still distributed in the V0.6.5 ZIP from the ChatGPT conversation, and is not yet checked into this GitHub repository.** This is not self-hosted and is not a DVILR V1.0 release.

## Windows one-click installer (v0.6.6 preview)

See [DVILR Windows Installer v0.6.6](DVILR_WINDOWS_INSTALLER_V066.md) for setup instructions, SHA-256 verification and test status. The unsigned Windows setup executable and full compiler source ZIP were supplied as separate ChatGPT download artifacts and **are not currently hosted in this GitHub repository**. The compiler is still Go-bootstrapped and V1.0 is unfinished.

## DVILR v1.1 Developer Preview (October 29, 2026)

See [DVILR v1.1 Developer Preview release notes](DVILR_V11_DEVELOPER_PREVIEW.md) for the REPL, formatter, doctor, starter manifest and VS Code highlighting package, verification checks and artifact SHA-256 values. This is an **experimental Windows installer build**, not a completed V1.0 core/AI/GPU/self-hosting release. The complete compiler source and executable are provided as ChatGPT conversation artifacts, not GitHub-hosted downloads.

## DVILR v1.1 unified release requirements

All previously proposed V0.6 core language, V0.8 web/systems, V0.9 AI/GPU/game and V1.0 self-hosting features are now included in the **final v1.1 release goal**. See [DVILR v1.1 Unified Release Requirements](DVILR_V11_UNIFIED_RELEASE_REQUIREMENTS.md). They are **not all implemented**. A new local preview includes `dvl test [path]`, optional Windows `--run` and UTF-8 input validation. The complete updated source is available in a separate downloadable ZIP, not fully committed into this repository.

The latest [v1.1 implementation checkpoint](DVILR_V11_IMPLEMENTATION_CHECKPOINT.md) records the tested `dvl test` command, UTF-8 validation, Windows cross-build, checksums, and unfinished release gates.

## DVILR v1.1.1 native AI score ranking (preview)

New native `max_value`, `min_value`, and `argmax` over signed i64 arrays/vectors are available in the separate v1.1.1 compiler source ZIP. See [implementation checkpoint](DVILR_V111_NATIVE_AI_CHECKPOINT.md), [compiler patch](compiler/patches/v111_native_ai_ranking.diff), and [source example](examples/native_ai_rank_v111.dvl). Go tests, native PE generation and x64 disassembly checks passed; **actual Windows runtime execution and full V1.1 completion are not verified**. The full compiler remains in the downloadable ZIP, not merged into this repository.

## DVILR v1.1.2 — Native memory and type-safety preview

A separate full-source ZIP and unsigned Windows installer now include `checked_add`, `checked_sub`, `checked_mul` (native overflow exit 10), `copy_vector` (bounded deep copy with static vector-type checks), and zeroization for `pop`/`clear_vector`. Read the [v1.1.2 implementation checkpoint](DVILR_V112_TYPE_MEMORY_CHECKPOINT.md), [compiler source patch](compiler/patches/v112_native_safety.patch), and [DVILR sample](examples/v112_memory_safe_copy.dvl). **Go tests and Windows PE cross-build passed; actual Windows execution and full language/self-hosting are NOT yet verified.** The full updated compiler source remains in the downloadable ZIP, not the GitHub demo repository.
