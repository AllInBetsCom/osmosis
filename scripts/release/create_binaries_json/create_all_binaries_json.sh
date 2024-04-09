#!/bin/bash

tags=(
    "v3.1.0" 
    "v4.2.0" 
    "v6.4.0" 
    "v8.0.0" 
    "v10.1.1" 
    "v11.0.1" 
    "v12.3.0" 
    "v13.1.2" 
    "v14.0.1" 
    "v15.2.0" 
    "v16.1.1"
    "v17.0.0"
    "v18.0.0"
    "v19.2.0"
    "v20.5.0"
    "v21.1.4"
    "v22.0.5"
    "v23.0.0"
    "v24.0.0"
)

echo "## Upgrade binaries"

for tag in ${tags[@]}; do
    echo
    echo "### ${tag}"
    echo
    echo '```json'
    python create_binaries_json.py --tag $tag
    echo '```'
done
