#!/bin/bash

RESOLVERS_DIR="/resolvers"
RESOLVERS="$RESOLVERS_DIR/resolvers.txt"
TRUSTED_RESOLVERS="$RESOLVERS_DIR/resolvers-trusted.txt"
WORDLIST="/lists/megadns.txt"

function update_resolvers() {
    echo "Updating resolvers from Github..."
    cd "$RESOLVERS_DIR" || exit 1
    git pull
    echo "Fresh resolvers updated."
}

function usage() {
    echo "Usage: $0 -dl|--domain-list <domain_list_file>"
    exit 1
}

while [[ "$#" -gt 0 ]]; do
    case $1 in
        -dl|--domain-list)
        DOMAIN_LIST="$2"
        shift
        ;;
    *)
        echo "Unknown parameter $1"
        usage
        ;;
    esac
    shift
done

if [[ -z "$DOMAIN_LIST" ]]; then
  echo "Error: Domain list file not provided."
  usage
fi

if [[ ! -f "$DOMAIN_LIST" ]]; then
  echo "Error: Domain list file '$DOMAIN_LIST' not found."
  exit 1
fi

while read -r DOMAIN; do
    echo "Fuzzing $DOMAIN"
    update_resolvers
    PUREDNS_OUTPUT_FILE="fuzz-result-$DOMAIN.txt"
    puredns bruteforce "$WORDLIST" "$DOMAIN" --resolvers "$RESOLVERS" --resolvers-trusted "$TRUSTED_RESOLVERS" --write "$PUREDNS_OUTPUT_FILE"
    echo "Results saved to $PUREDNS_OUTPUT_FILE"
done < "$DOMAIN_LIST"