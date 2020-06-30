#!/bin/bash
curl https://docs.oasis-open.org/sarif/sarif/v2.1.0/cs01/schemas/sarif-schema-2.1.0.json > sarif-schema-2.1.0.json
schematyper --out-file=sarif/sarif.go --package sarif --root-type=Sarif sarif-schema-2.1.0.json
