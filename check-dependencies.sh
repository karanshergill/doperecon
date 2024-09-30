#!/bin/bash

# This script checks if the required tools are installed and available in the system's PATH.
# If any of these tools are missing, it will display an error and prompt the user to install them.

function check_command() {
    local cmd=$1
    local name=$2
    if ! command -v "$cmd" &> /dev/null; then
        echo "Error: $name is not installed or not available in the system's PATH."
        echo "Please install $name and ensure it is accessible."
        exit 1
    fi
}

# Check for required tools
check_command puredns "puredns"
check_command publicresolvers "publicresolvers"
check_command domainer "domainer"
check_command alterx "AlterX"