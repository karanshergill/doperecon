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

while read -r domain; do
    echo "Fuzzing $domain"
    update_resolvers
    output_file="fuzz-result-$domain.txt"
    puredns bruteforce "$WORDLIST" "$domain" --resolvers "$RESOLVERS" --resolvers-trusted "$TRUSTED_RESOLVERS" --write "$output_file"
    echo "Results saved to $output_file"
done < "$DOMAIN_LIST"