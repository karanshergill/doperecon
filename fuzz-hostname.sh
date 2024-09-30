#!/bin/bash

source ./check-dependencies.sh
source ./update-resolvers.sh
source ./generate-wordlist.sh

WORDLIST_DIR="/home/superuser/wordlists"
RESOLVERS="resolvers.txt"
TRUSTED_RESOLVERS="trusted_resolvers.txt"

function usage() {
    echo "Usage: $0 -dl <domain-list> [-wl <wordlist>]"
    exit 1
}

while [[ "$#" -gt 0 ]]; do
    case $1 in
        -dl|--domain-list)
            DOMAINLIST="$2"
            shift
        ;;
        -wl|--wordlist)
            WORDLIST_NAME="$2"
            WORDLIST="$WORDLIST_DIR/$WORDLIST_NAME"
            shift
        ;;
        *)
            echo "Unknown parameter $1"
            usage
        ;;
    esac
    shift
done

if [[ -z "$DOMAINLIST" ]]; then
    echo "Error: Domain list file not provided."
    usage
fi

if [[ ! -f "$DOMAINLIST" ]]; then
    echo "Error: Domain list file '$DOMAINLIST' not found."
    exit 1
fi

if [[ -z "$WORDLIST" ]]; then
    WORDLIST="$WORDLIST_DIR/megadns.txt"
fi

if [[ ! -f "$WORDLIST" ]]; then
    echo "Error: Wordlist file '$WORDLIST' not found."
    exit 1
fi

trap 'rm -f "$RESOLVERS" "$TRUSTED_RESOLVERS" "$TEMP_OUTPUT_FILE"' EXIT

IFS=$'\n'
while read -r DOMAIN; do
    DOMAIN=$(echo "$DOMAIN" | sed 's/^[ \t]*//;s/[ \t]*$//;s/\r//g')
    echo "Fuzzing $DOMAIN"

    update_resolvers

    TEMP_OUTPUT_FILE="temp-fuzz-result-${DOMAIN}.txt"
    PUREDNS_OUTPUT_FILE="fuzz-result-${DOMAIN}.txt"
    
    puredns bruteforce "$WORDLIST" "$DOMAIN" --resolvers "$RESOLVERS" --resolvers-trusted "$TRUSTED_RESOLVERS" --write "$TEMP_OUTPUT_FILE" --quiet
    
    sort "$TEMP_OUTPUT_FILE" | uniq > "$PUREDNS_OUTPUT_FILE"

    echo "Results saved to $PUREDNS_OUTPUT_FILE"
    PUREDNS_OUTPUT_FILES+=("$PUREDNS_OUTPUT_FILE")
done < "$DOMAINLIST"

echo "Processing completed. Generating report..."
for FILE in "${PUREDNS_OUTPUT_FILES[@]}"; do
    SUBDOMAIN_COUNT=$(wc -l < "$FILE")
    echo "File: $FILE, Results: $SUBDOMAIN_COUNT"
done

echo "Generating wordlist..."
strip_eTLD "${PUREDNS_OUTPUT_FILES[@]}"
echo "Fuzzlist updated"
echo "Processing completed."