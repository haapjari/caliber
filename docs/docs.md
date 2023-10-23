# Notes 

- `.git` folder contains history and structure of the repository. When we are dealing with offline analysis, this is the tool that we have.
    - Folder contains commit history, all the commits made in the repository, including commit messages, authors, dates.
        - Commit message formatting and conventions.
        - Frequency of commits.
        - Time since last commit.
    - Branch Information
    - Tag Information
    - Configuration Data
    - Staging area
    - Objects
    - Refs
        - Last commit of each branch.
    - Commit Graph
    - Blame Data
        - Areas of code, that haven't been touched for a long time.
- Current state of directory can be analyzed.
    - File Structure and Organization
        - Directory structure and depth.
        - Presence or absence of certain essential files (README.md, LICENSE, .gitignore)
    - File Metadata
        - File size, which can help identify, for example large files.
        - File permissions.
        - Last modified data.
    - Code Analysis:
        - Lines of code, in each file.
        - Comment-to-code ratio.
        - Complexity metrics, such as cyclomatic complexity.
        - Unused variables, methods, or imports.
        - Potential security vulnerabilities, such as hardcoded secrets or insecure functions.
    - Analyzed dependency definitions.
    - Configuration files.
    - Documentation.
    - Test Files.
        - Presence of test files and their distribution relative to source files.
        - Code coverage metrics, if applicable.
        - Test-to-code ratio.
    - File Types and Binary Data.

## Activity Plugin Algorithm

- TBD

    // Get Path
    // Analyze Path
    // How do you measure activity from .git repository (?)                               ---

    // TODO: This could be configurable (?)

    // Measurements:

    // Low 
   
    // Calculate total amount of commits.
    // Check the date of "first" commit.
    // Calculate amount of days, since the first commits.
     
    // First Iteration, completely without definitions. 
    // If there has not been commits, during the last "half".
    // = Low
    // If there has been commits on last month.
    // If there has not been commits during last "half".
    
    // Second Iteration, with definitions. 
    // What data can be pulled from definitions.



    // There are no recents commits. (Within 3 months.)
    // There are no commits happened last half quarter.

    // TODO: GitHub Stuff

    // There recent commits within 3 months.
    // Calculate total amount of commits.
    // Threshold of commits at least, some of the commits happened during last quarter of 
    // Threshold of commits at least 

---
