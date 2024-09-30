#!/bin/bash

function update_resolvers() {
    RESOLVERS="resolvers.txt"
    TRUSTED_RESOLVERS="trusted_resolvers.txt"

    echo "Generating resolvers..."
    publicresolvers --resolvers > "$RESOLVERS"
    publicresolvers --trusted > "$TRUSTED_RESOLVERS"
    
    if [[ -s "$RESOLVERS" && -s "$TRUSTED_RESOLVERS" ]]; then
        echo "Resolvers generated successfully."
    else
        echo "Error: Failed to generate resolvers!"
        exit 1
    fi
}
