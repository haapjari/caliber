# Caliber

- Caliber will be a linter utility which will be able to analyze Go programs and flag problematic third-party dependencies, outputting report. 
- Caliber can be run locally, or for example, within a CI/CD pipeline.

## Author

- I am Jari "Haspe" Haapasaari. I am passionate about Software Quality, and this is an attempt to create a tool which will help me and others to improve the quality of their software. 

## Version

- Every commit is tagged as a version number which is read from the git, and the version number is updated in the `Makefile` as well as the `Dockerfile`.
- If you want to clone the project, neat info can be printed with `make info` command. (Requires make)

### Example Output

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

## Supported Commands

### Run

Run, with default config and output to stdout.

```bash
    caliber run path=<path_to_project>
```

### Version 

Version of Caliber.

```bash
    caliber version
```

### Help

Help for Caliber.

```bash
    caliber help
```

---

### MVP Version (v1.0.0)

- run with activity plugin (no configs, output to report & stdout)
- version
- help 
- run with 

---

### Roadmap

- run command with custom config
    - Config can be loaded from a file, or passed in as a flag. Caliber searchs for a `.caliber.yml` file in the current working directory, and if it is not found, it will use the default config.

```yaml
exclude:
    - "github.com/example/example"
    - "github.com/example/example"
```

- support plugin
- popularity plugin
- maintainability plugin
- vulnerability plugin
- local definitions database support

#### TODO 

- run command
- activity plugin
    - online and offline mode.
- output to json file

---

