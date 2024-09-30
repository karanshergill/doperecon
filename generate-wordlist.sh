function stripdomains() {
    local puredns_output_files=("$@")
    local output_wordlist="generated_wordlist.txt"

    for file in "${puredns_output_files[@]}"; do
        echo "Processing file: $file"
        while read -r subdomain; do
            stripped_subdomain=$(echo "$subdomain" | awk -F. '{OFS="."; NF-=2; print}')
            echo "$stripped_subdomain" >> "$output_wordlist"
        done < "$file"
    done

    sort "$output_wordlist" | uniq > temp-wordlist.txt && mv temp-wordlist.txt "$output_wordlist"
    
    echo "New wordlist generated: $output_wordlist"
}