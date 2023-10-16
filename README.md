# Caliber

- Caliber will be a linter utility which will be able to analyze Go programs and flag problematic third-party dependencies, outputting report. 
- Caliber can be run locally, or for example, within a CI/CD pipeline.

## Author

- I am Jari "Haspe" Haapasaari. I am passionate about Software Quality, and this is an attempt to create a tool which will help me and others to improve the quality of their software. 

## Version

- Every commit is tagged as a version number which is read from the git, and the version number is updated in the `Makefile` as well as the `Dockerfile`.

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

### Config

- Config can be loaded from a file, or passed in as a flag. Caliber searchs for a `.caliber.yml` file in the current working directory, and if it is not found, it will use the default config.

```yaml
exclude:
    - "github.com/example/example"
    - "github.com/example/example"
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

```bash
    # Run Linter 
    caliber -run path=<path_to_project> 
    caliber -run path=<path_to_project> load=<path_to_config> 
    caliber -run path=<path_to_project> output=<output_file_path>.json
    caliber -run path=<path_to_project> load=<path_to_config> output=<output_file_path>.json
    caliber -run path=<path_to_project> output=<output_file_path>.json load=<path_to_config>
    
    # Example Run Command
    caliber -run path=~/go/src/github.com/example/example output=~/go/src/github.com/example/example/report.json
    
    # Update Definitions
    caliber -update
    
    # Version
    caliber -version
    
    # Help
    caliber -help
```

---

## Features

- Headless: Run as a Comnmand-Line Utility, Generates Report.
- Graphical User Interface (HTMX and Go)
- Performance Metrics from Analysis
- External, mountable config.

---