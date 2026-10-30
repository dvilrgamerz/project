# Execute from full compiler's 'compiler' folder on actual Windows x64.
param([string]$Compiler = '.\installer\payload\dvl.exe')
$ErrorActionPreference = 'Stop'
$compilerPath = (Resolve-Path $Compiler).Path
$root = Split-Path -Parent $PSScriptRoot
$work = Join-Path $env:TEMP ('dvilr-v114-' + [guid]::NewGuid().ToString('N'))
New-Item -ItemType Directory -Force $work | Out-Null
try {
    $valid = Join-Path $root 'examples\v114\typed_functions.dvl'
    $binary = Join-Path $work 'typed_functions.exe'
    & $compilerPath build $valid -o $binary
    if ($LASTEXITCODE -ne 0 -or -not (Test-Path $binary)) { throw 'Compiler failed to produce PE' }
    $lines = @(& $binary | ForEach-Object { $_.Trim() } | Where-Object { $_ -ne '' })
    if ($LASTEXITCODE -ne 0 -or ($lines -join ',') -ne '1,0,42') {
        throw "Native output mismatch: $($lines -join ','); status $LASTEXITCODE"
    }
    Write-Host '[PASS] Windows native output: 1,0,42'
    foreach ($name in @('bad_bool_argument','bad_bool_return','bad_numeric_argument')) {
        $source = Join-Path $root ('examples\v114\' + $name + '.dvl')
        $output = Join-Path $work ($name + '.exe')
        $logs = & $compilerPath build $source -o $output 2>&1
        if ($LASTEXITCODE -eq 0 -or (Test-Path $output)) { throw "Expected type error: $name" }
        Write-Host "[PASS] Rejected $name"
    }
    Write-Host 'DVILR v1.1.4 Windows-native verification PASSED'
} finally {
    Remove-Item -Path $work -Recurse -Force -ErrorAction SilentlyContinue
}
