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

        npm run build:dev
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

    Get-ChildItem -Path (Join-Path $PSScriptRoot 'web' 'dist')  | Copy-Item -Destination $clientDirectory -Recurse -Container -Force
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

    $command = "docker build"
    $command += " --force-rm"
    $command += " --build-arg NOW=$date"
    $command += " --build-arg REVISION=$sha1"
    $command += " --build-arg VERSION=$version"
    $command += " --file ./build/package/bff/dockerfile"

    if ($dockerTags.Length -gt 0)
    {
        Write-Output "$($dockerTags.Length)"
        foreach($tag in $dockerTags)
        {
            Write-Output "Tagging with: $($tag)/service-provisioning-ui-web:$version"
            $command += " --tag $($tag)/service-provisioning-ui-web:$version"
        }
    }
    else
    {
        $command += " --tag service-provisioning-ui-web:$version"
    }

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

    & swag init --parseInternal --output ./api --generalInfo ./internal/cmd/serve.go

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

    go build -a -installsuffix cgo -v -ldflags="-X github.com/calvinverse/service.provisioning.ui.web/internal/info.applicationName=provisioning-bff -X github.com/calvinverse/service.provisioning.ui.web/internal/info.sha1=$sha1 -X github.com/calvinverse/service.provisioning.ui.web/internal/info.buildTime=$date -X github.com/calvinverse/service.provisioning.ui.web/internal/info.version=$version" -o $outputDir/bff.exe ./cmd

    go test -cover ./... ./cmd
}

$revision = Get-Revision
Write-Output "Using revision: '$revision'"

$version = Get-Version
Write-Output "Using version: '$version'"

$date = Get-Date -UFormat '%Y-%m-%dT%T'
Write-Output "Using date: '$date'"

if ($Direct)
{
    Write-Output "Building container ..."
    New-LocalBuild -date $date -sha1 $revision -version $version
}
else
{
    Write-Output "Building locally ..."
    New-Container -date $date -sha1 $revision -version $version -dockerTags $($dockerTags.Split(',', [System.StringSplitOptions]::RemoveEmptyEntries))
}

