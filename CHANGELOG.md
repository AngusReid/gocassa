# Change Log
All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## v1.0.0 - 2015-11-13

### Added
 - Allow creating tables with compound keys
 - Added the `MultimapMultiKeyTable` recipe which allows for CRUD operations on rows filtered by equality of multiple fields.

### Changed
 - Improved how gocassa handles encoding+decoding, no longer uses the `encoding/json` package and now supports embedded types and type aliases.

### Fixed
 - Mock tables are now safe for concurrent use
 - `uint` types are now supported, when generating tables the cassandra `varint` type is used.
 - Fixed gocassa using `json` tags when decoding results into structs
