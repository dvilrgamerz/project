# Requires the full standalone v1.1.3 compiler source ZIP on Windows x64.
# This GitHub storefront repo does NOT contain the complete compiler source.
param([string]$Compiler = '.\installer\payload\dvl.exe')
$ErrorActionPreference = 'Stop'
$compilerPath = (Resolve-Path $Compiler).Path
$root = Split-Path -Parent (Split-Path -Parent $PSScriptRoot)
$work = Join-Path $env:TEMP ('dvilr-v113-native-' + [guid]::NewGuid().ToString('N'))
New-Item -ItemType Directory -Force $work | Out-Null
try {
    $src = Join-Path $root 'examples\v113\typed_scalar_success.dvl'
    $out = Join-Path $work 'typed_scalar_success.exe'
    & $compilerPath build $src -o $out
    if ($LASTEXITCODE -ne 0 -or !(Test-Path $out)) { throw 'Native compiler failed' }
    $output = @( & $out ) -join ','
    if ($LASTEXITCODE -ne 0 -or $output -ne '1,0,1') {throw "Incorrect Windows native output: $output; exit=$LASTEXITCODE"}
    Write-Host '[PASS] Windows native bool/i64 program output'
    foreach ($name in @('bad_bool_literal','bad_i64_literal','bad_bool_assignment','bad_bool_arithmetic')) {
        $bad = Join-Path $root ('examples\v113\' + $name + '.dvl')
        $badOut = Join-Path $work ($name + '.exe')
        & $compilerPath build $bad -o $badOut 2>$null
        if ($LASTEXITCODE -eq 0 -or (Test-Path $badOut)) {throw "Expected compiler rejection: $name"}
        Write-Host "[PASS] rejected $name"
    }
    Write-Host 'All five V1.1.3 native checks passed.'
} finally {
    Remove-Item -Recurse -Force $work -ErrorAction SilentlyContinue
}
