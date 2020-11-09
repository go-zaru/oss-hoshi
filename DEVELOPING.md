
Principles

Trunk Base Development | HEAD

- Feature Flags 
    | default false / per env 
    | Any branching strategies

- "Commit to Main" development 
    | hooks to: locally build, lint, test

- Release from Trunk 
    | on Commit run CI/CD
    | on build fail development stops


Design Decisions

- CI/CD
    | for Builds, Bazel on Docker
    | for CI, Github Actions
    | for Sharing, Github Packages
