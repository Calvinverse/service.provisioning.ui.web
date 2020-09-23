[CmdletBinding()]
param(
    [string] $dockerTags = '',
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

        <#
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
        #>
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
        [string] $date,
        [string[]] $dockerTags = @()
    )

    Write-Output "Building Docker container with build arguments:"
    Write-Output "NOW = $date"
    Write-Output "REVISION = $sha1"
    Write-Output "VERSION = $version"
    Write-Output "Tagging with: service-provisioning:$version"

    $command = "docker build"
    $command += " --force-rm"
    $command += " --build-arg NOW=$date"
    $command += " --build-arg REVISION=$sha1"
    $command += " --build-arg VERSION=$version"
    $command += " --file ./build/package/server/dockerfile"

    if ($dockerTags.Length -gt 0)
    {
        foreach($tag in $dockerTags)
        {
            $command += " --tag $($tag)/service-provisioning:$version"
        }
    }

    $command += " --tag service-provisioning:$version"
    $command += " ."

    Write-Output "Invoking: $command"
    Invoke-Expression -Command $command
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
    $absoluteOutputDir = [System.IO.Path]::GetFullPath($(Join-Path $PSScriptRoot $outputDir))
    if (-not (Test-Path $absoluteOutputDir))
    {
        New-Item -Path $absoluteOutputDir -ItemType Directory | Out-Null
    }

    $clientDirectory = Join-Path $absoluteOutputDir 'client'
    Invoke-Npm -clientDirectory $clientDirectory

    Copy-Item -Path (Join-Path $PSScriptRoot "configs" "*") -Destination $absoluteOutputDir -Force

    & swag init --parseInternal --generatedTime --output ./api --generalInfo ./internal/cmd/server.go

    $docDirectory = Join-Path $absoluteOutputDir 'api'
    if (-not (Test-Path $docDirectory))
    {
        New-Item -Path $docDirectory -ItemType Directory | Out-Null
    }

    Copy-Item -Path (Join-Path $PSScriptRoot 'api' '*') -Destination $docDirectory -Force

    $configPath = Join-Path $absoluteOutputDir 'config.yaml'
    Add-Content -Path $configPath -Value 'doc:'
    Add-Content -Path $configPath -Value "  path: $docDirectory"

    Add-Content -Path $configPath -Value 'service:'
    Add-Content -Path $configPath -Value '  port: 8080'

    Add-Content -Path $configPath -Value 'ui:'
    Add-Content -Path $configPath -Value "  path: $clientDirectory"

    go build -a -installsuffix cgo -v -ldflags="-X github.com/calvinverse/service.provisioning/internal/info.sha1=$sha1 -X github.com/calvinverse/service.provisioning/internal/info.buildTime=$date -X github.com/calvinverse/service.provisioning/internal/info.version=$version" -o $outputDir/server.exe ./cmd

    go test -cover ./... ./cmd
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
    New-Container -date $date -sha1 $revision -version $version -dockerTags $($dockerTags.Split(','))
}

