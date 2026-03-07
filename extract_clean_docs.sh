#!/bin/bash

# Script to extract Ecwid API documentation pages to markdown
# Uses pandoc with optimized settings for cleaner output

set -e

DOCS_DIR="docs/ecwid"
BASE_URL="https://docs.ecwid.com/api-reference"

# Create docs directory if it doesn't exist
mkdir -p "$DOCS_DIR"

echo "Starting Ecwid API documentation extraction..."

# List of main API sections with output filenames
declare -A sections=(
    [""]="overview"
    ["rest-api-error-codes"]="rest-api-error-codes"
    ["rest-api/store-profile"]="store-profile"
    ["rest-api/orders"]="orders"
    ["rest-api/products"]="products"
    ["rest-api/categories"]="categories"
    ["rest-api/customers"]="customers"
    ["rest-api/discounts"]="discounts"
    ["rest-api/domains"]="domains"
    ["rest-api/dictionaries"]="dictionaries"
    ["rest-api/staff-accounts"]="staff-accounts"
    ["rest-api/application"]="application"
    ["rest-api/batch-requests"]="batch-requests"
    ["rest-api/shipping-options"]="shipping-options"
    ["rest-api/payment-options"]="payment-options"
    ["rest-api/checkout-extra-fields"]="checkout-extra-fields"
    ["rest-api/storefront-widget-details"]="storefront-widget-details"
    ["rest-api/carts"]="carts"
    ["rest-api/instant-site"]="instant-site"
    ["rest-api/reviews"]="reviews"
    ["rest-api/subscriptions"]="subscriptions"
)

# Function to convert HTML to markdown using pandoc with clean settings
convert_page() {
    local url=$1
    local output_file=$2
    
    echo "Fetching: $url"
    
    # Fetch the page and convert to markdown with better options
    # --strip-comments: remove HTML comments
    # --wrap=none: don't wrap lines
    # -t gfm: use GitHub-flavored markdown
    curl -s "$url" | \
        pandoc \
            --from=html \
            --to=gfm \
            --wrap=none \
            --strip-comments \
            -o "$output_file"
    
    if [ -s "$output_file" ]; then
        echo "  ✓ Saved to: $output_file"
    else
        echo "  ✗ Failed (empty file)"
        rm -f "$output_file"
    fi
}

# Process each section
for section in "${!sections[@]}"; do
    filename="${sections[$section]}"
    
    if [ -z "$section" ]; then
        url="$BASE_URL"
    else
        url="$BASE_URL/$section"
    fi
    
    output="$DOCS_DIR/${filename}.md"
    convert_page "$url" "$output"
    
    # Small delay to be respectful to the server
    sleep 0.5
done

echo ""
echo "✓ Documentation extraction complete!"
echo "Files saved to: $DOCS_DIR/"
echo ""
echo "Cleaning up old files..."
ls -lh "$DOCS_DIR"
