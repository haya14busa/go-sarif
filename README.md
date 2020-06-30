# Go structs for the SARIF 2.1.0 object model

## Description
This module contains Go structs for the object model defined by the [Static Analysis Results Interchange Format (SARIF) Version 2.1.0](https://docs.oasis-open.org/sarif/sarif/v2.1.0/cs01) file format, an OASIS Committee [Specification](https://www.oasis-open.org/news/announcements/static-analysis-results-interchange-format-sarif-v2-1-0-from-the-sarif-tc-is-an-a).

To learn more about SARIF and find resources for working with it, you can visit the [SARIF Home Page](http://sarifweb.azurewebsites.net/).

## Generation

The Go structs in this module were generated from the [SARIF JSON schema](https://docs.oasis-open.org/sarif/sarif/v2.1.0/cs01/schemas/sarif-schema-2.1.0.json) by the [quicktype](https://github.com/quicktype/quicktype) code generator.

```shell
$ npm install
$ npm run build
```

