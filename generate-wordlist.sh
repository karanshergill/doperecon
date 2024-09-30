#!/bin/bash

# remove the top-level and the second-level domain from the puredns output files
function strip_eTLD() {
    local puredns_output_files=("$@")
    local output_wordlist="generated_wordlist.txt"

    for file in "${puredns_output_files[@]}"; do
        echo "Processing file: $file"
        eTLD=$(cat "$file" | domainer)
        grep -v "^$eTLD$" "$file" | sed "s/.$eTLD//" >> "$output_wordlist"
    done

    sort "$output_wordlist" | uniq > temp-wordlist.txt && mv temp-wordlist.txt "$output_wordlist"

    echo "New wordlist generated: $output_wordlist"
}