# Execute on Windows x64 from a complete compiler source tree.
param([string]$Compiler = '.\installer\payload\dvl.exe')
$ErrorActionPreference = 'Stop'
$compilerPath = (Resolve-Path $Compiler).Path
$root = Split-Path -Parent $PSScriptRoot
$temp = Join-Path $env:TEMP ('dvilr-v115-' + [guid]::NewGuid().ToString('N'))
New-Item -ItemType Directory -Force $temp | Out-Null
try {
    $source = Join-Path $root 'examples\v115\vector_slice_unicode.dvl'
    $program = Join-Path $temp 'vector_slice_unicode.exe'
    & $compilerPath build $source -o $program
    if ($LASTEXITCODE -ne 0 -or !(Test-Path $program)) { throw 'Native compiler failed' }
    $output = @(& $program | ForEach-Object { $_.Trim() } | Where-Object { $_ -ne '' })
    if ($LASTEXITCODE -ne 0 -or ($output -join ',') -ne 'SLICE OK,50') { throw "Wrong native output: $($output -join ',')" }
    Write-Host '[PASS] Native program: SLICE OK,50'
    foreach ($name in @('bad_slice_negative_start', 'bad_slice_negative_count', 'bad_slice_past_end')) {
        $src = Join-Path $root ('examples\v115\' + $name + '.dvl')
        $bad = Join-Path $temp ($name + '.exe')
        & $compilerPath build $src -o $bad
        if ($LASTEXITCODE -ne 0 -or !(Test-Path $bad)) { throw "Failed to build bounds test $name" }
        & $bad | Out-Null
        if ($LASTEXITCODE -ne 8) { throw "Expected exit 8 for $name; got $LASTEXITCODE" }
        Write-Host "[PASS] Bounds exit 8: $name"
    }
    $typedSource = Join-Path $root 'examples\v115\bad_slice_bool_type.dvl'
    $typedOutput = Join-Path $temp 'bad_type.exe'
    $logs = & $compilerPath build $typedSource -o $typedOutput 2>&1
    if ($LASTEXITCODE -eq 0 -or (Test-Path $typedOutput)) { throw 'Boolean index was wrongly accepted' }
    Write-Host '[PASS] Bad boolean index rejected'
} finally {
    Remove-Item $temp -Force -Recurse -ErrorAction SilentlyContinue
}
