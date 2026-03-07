#!/bin/bash

# Script to extract Ecwid API documentation pages to markdown
# Uses pandoc to convert HTML to markdown

set -euo pipefail

DOCS_DIR="docs/ecwid"
BASE_URL="https://docs.ecwid.com/api-reference"

# Create docs directory if it doesn't exist
mkdir -p "$DOCS_DIR"

echo "Starting Ecwid API documentation extraction..."

# List of main API sections based on navigation menu
declare -a sections=(
    ""                                  # REST API overview
    "rest-api-error-codes"
    "rest-api/store-profile"
    "rest-api/orders"
    "rest-api/products"
    "rest-api/categories"
    "rest-api/customers"
    "rest-api/discounts"
    "rest-api/domains"
    "rest-api/dictionaries"
    "rest-api/staff-accounts"
    "rest-api/application"
    "rest-api/batch-requests"
    "rest-api/shipping-options"
    "rest-api/payment-options"
    "rest-api/checkout-extra-fields"
    "rest-api/storefront-widget-details"
    "rest-api/carts"
    "rest-api/instant-site"
    "rest-api/reviews"
    "rest-api/subscriptions"
)

# Function to convert HTML to markdown using pandoc
convert_page() {
    local url=$1
    local output_file=$2
    
    echo "Fetching: $url"
    
    # Fetch the page and convert to markdown
    if command -v pandoc &> /dev/null; then
        curl -fsS "$url" | pandoc -f html -t markdown -o "$output_file"
        if [ -s "$output_file" ]; then
            echo "  ✓ Saved to: $output_file"
        else
            echo "  ✗ Failed (empty file)"
            rm -f "$output_file"
        fi
    else
        echo "  ✗ pandoc not found, downloading HTML only"
        curl -fsS "$url" -o "${output_file%.md}.html"
    fi
}

# Process each section
for section in "${sections[@]}"; do
    if [ -z "$section" ]; then
        # Main overview page
        url="$BASE_URL"
        output="$DOCS_DIR/overview.md"
    else
        url="$BASE_URL/$section"
        # Convert URL path to filename
        filename=$(echo "$section" | sed 's/rest-api\///g' | sed 's/\//-/g')
        output="$DOCS_DIR/${filename}.md"
    fi
    
    convert_page "$url" "$output"
    
    # Small delay to be respectful to the server
    sleep 0.5
done

echo ""
echo "✓ Documentation extraction complete!"
echo "Files saved to: $DOCS_DIR/"
