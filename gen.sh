#!/bin/bash
curl http://json.schemastore.org/sarif-2.1.0-rtm.4 > sarif-schema-2.1.0.json
schematyper --out-file=sarif/sarif.go --package sarif --root-type=Sarif sarif-schema-2.1.0.json
