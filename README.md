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
