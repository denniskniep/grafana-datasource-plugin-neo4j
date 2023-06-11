# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] - YYYY-MM-DD

## [1.2.0] - 2023-06-11
### Changed
 - Updated Neo4J Driver to [neo4j-go-driver v5.9.0](https://github.com/neo4j/neo4j-go-driver/releases/tag/v5.9.0)


### Fixed 
 - #14: Error for null-value attribute

## [1.1.0] - 2022-02-28
### Changed
- Use official name Neo4j instead of Neo4J
- Use neo4j logo with blue background to support both dark and light theme. Logo was barely visible with the light theme.
- Catch errors with connection details and prevent errors with internal network informations
- Caching Neo4j driver in the backend

## [1.1.0-beta]
### Added
- Changed to Backend Plugin
- Queries can now be executed in variables (Issue #3)
- neo4j password is stored and encrypted in secureJsonData

## [1.0.1]
### Changed
- Change Organisation name

## [1.0.0]

Initial release.
