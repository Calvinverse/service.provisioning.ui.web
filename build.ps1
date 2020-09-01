[CmdletBinding()]
param(
    [switch] $Direct = $false
)

function Get-Revision
{
    [CmdletBinding()]
    param()

    return (git rev-parse HEAD).Trim()
}

function Get-Version
{
    $tempDir = Join-Path $PSScriptRoot 'temp'
    $gitVersion = Join-Path (Join-Path (Join-Path $tempDir 'GitVersion.CommandLine') 'tools') 'gitversion.exe'
    if (-not (Test-Path $gitVersion))
    {
        if (-not (Test-Path $tempDir))
        {
            New-Item -Path $tempDir -ItemType Directory | Out-Null
        }

        $nuget = Join-Path $tempDir 'nuget.exe'
        if (-not (Test-Path $nuget))
        {
            Invoke-WebRequest -Uri 'https://dist.nuget.org/win-x86-commandline/latest/nuget.exe' -OutFile $nuget
        }

        & nuget install GitVersion.CommandLine -ExcludeVersion -OutputDirectory $tempDir -NonInteractive -Source https://api.nuget.org/v3/index.json
    }

    return & $gitVersion /output json /showvariable SemVer
}

function Invoke-Npm
{
    [CmdletBinding()]
    param (
        [string] $clientDirectory
    )

    $currentDirectory = $pwd
    try
    {
        Set-Location (Join-Path $PSScriptRoot 'web')

        npm run lint
        if ($LASTEXITCODE -ne 0)
        {
            throw "NPM run lint failed with exit code: $LASTEXITCODE"
        }

        npm run build
        if ($LASTEXITCODE -ne 0)
        {
            throw "NPM run build failed with exit code: $LASTEXITCODE"
        }

        npm run test:unit
        if ($LASTEXITCODE -ne 0)
        {
            throw "NPM run test:unit failed with exit code: $LASTEXITCODE"
        }

        npm run test:e2e
        if ($LASTEXITCODE -ne 0)
        {
            throw "NPM run test:e2e failed with exit code: $LASTEXITCODE"
        }
    }
    finally
    {
        Set-Location $currentDirectory
    }

    if (-not (Test-Path $clientDirectory))
    {
        New-Item -Path $clientDirectory -ItemType Directory | Out-Null
    }

    Get-ChildItem -Path (Join-Path $PSScriptRoot 'web' 'dist')  | Copy-Item -Destination $clientDirectory -Recurse -Container
}

function New-Container
{
    [CmdletBinding()]
    param (
        [string] $sha1,
        [string] $version,
        [string] $date
    )

    Write-Output "Building Docker container with build arguments:"
    Write-Output "NOW = $date"
    Write-Output "REVISION = $sha1"
    Write-Output "VERSION = $version"
    Write-Output "Tagging with: service-provisioning:$version"

    docker build --force-rm --build-arg NOW=$date --build-arg REVISION=$sha1 --build-arg VERSION=$version --file ./build/package/server/dockerfile --tag service-provisioning:$version .
}

function New-LocalBuild
{
    [CmdletBinding()]
    param (
        [string] $sha1,
        [string] $version,
        [string] $date
    )

    $outputDir = './bin'
    $absoluteOutputDir = Join-Path $PSScriptRoot $outputDir
    if (-not (Test-Path $absoluteOutputDir))
    {
        New-Item -Path $absoluteOutputDir -ItemType Directory | Out-Null
    }

    Invoke-Npm

    Copy-Item -Path (Join-Path $PSScriptRoot "configs" "*") -Destination $absoluteOutputDir -Force
    go build -a -installsuffix cgo -v -ldflags="-X github.com/calvinverse/service.provisioning/internal/info.sha1=$sha1 -X github.com/calvinverse/service.provisioning/internal/info.buildTime=$date -X github.com/calvinverse/service.provisioning/internal/info.version=$version" -o $outputDir/server.exe ./cmd
}

$revision = Get-Revision
$version = Get-Version
$date = Get-Date -UFormat '%Y-%m-%d_%T'

if ($Direct)
{
    New-LocalBuild -date $date -sha1 $revision -version $version
}
else
{
    New-Container -date $date -sha1 $revision -version $version
}

