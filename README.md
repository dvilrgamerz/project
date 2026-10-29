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
