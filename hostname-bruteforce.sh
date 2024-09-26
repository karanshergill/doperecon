#!/bin/bash

RESOLVERS_DIR="/resolvers"
RESOLVERS="$RESOLVERS_DIR/resolvers.txt"
TRUSTED_RESOLVERS="$RESOLVERS_DIR/resolvers-trusted.txt"
WORDLIST="/lists/megadns.txt"

function updateResolvers() {
    echo "Updating resolvers from Github..."
    cd "$RESOLVERS_DIR" || exit 1
    git pull
    echo "Fresh resolvers updated."
}

updateResolvers
