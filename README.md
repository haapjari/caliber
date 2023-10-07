# Caliber

- Caliber will be a linter utility which will be able to analyze Go programs and flag problematic third-party dependencies, outputting a `json` report. 
- Caliber can be run locally, or for example, within a CI/CD pipeline.

## Example Report

```json
{
    "project": "project",
    "report": [
        {
            "dependency": "url",
            "details": {
                "name": "example",
                "version": "0.0.1",
                "lastUpdated": "2021-01-01",
                "supportLevel": "high",
                "activity": "low",
                "popularity": "high",
                "license": "MIT",
                "vulnerabilityTrend": "medium",
                "transitiveDependencies": "high"
            }
        },
        {
            "dependency": "url",
            "details": {
                "name": "example",
                "version": "0.0.1",
                "lastUpdated": "2021-01-01",
                "supportLevel": "high",
                "activity": "low",
                "popularity": "high",
                "license": "MIT",
                "vulnerabilityTrend": "medium",
                "transitiveDependencies": "high"
            }
        }
    ]
}
```

---

## How-To: Use

- TBD

### Required Configuraton Variables

```
VERSION=
```

---

## How-To: Contribute

- TBD

---

## Architecture

- TBD

---

## Notes

- Command-Line Interface
- Exclude Standard Library
- Support Level: What level of Support is available for this library?
- Activity: Is the library actively developed?
    - Commit Insights
    - Contributor Count (Recentness)
    - Releases
- Popularity: How many people use this library?
    - Stargazers, if this is a GitHub hosted repository.
- License: What kind of licensing does this library have?
- Vulnerability: Has there been a trend of vulnerabilities with this library?
- Maintainability: Does the library have a lot of transitive dependencies?
- Activity
    - Commit Insights 
    - Contributor Count
- Maturity
    - Creation Date
    - Contributor Count
- How-To: Generate Reports (?)
- Local DB: 
    - (Redis)
- Centralized Database
    - Update Definitions Functionality

---

## Commands

- Generate Config File: `caliber init`
- Pull Definitions: `caliber --pull-local` or `caliber --pull-central`
- Push Definitions: `caliber --push-local` or `caliber --push-central`
- Run Linter: `caliber --output=<output_file_path.json> --path=<path_to_project> run` 
- Help: `caliber --help` or `caliber -h`
- Version: `caliber --version` or `caliber -v`

---

## Features

- Headless: Run as a Comnmand-Line Utility, Generates Report.
- Graphical User Interface (HTMX and Go)
- Performance Metrics from Analysis
- External, mountable config.

---