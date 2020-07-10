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

function New-Container
{
    [CmdletBinding()]
    param (
        [string] $sha1,
        [string] $version,
        [string] $date
    )

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

