#!/bin/bash

# Script to convert HTML files to markdown

set -euo pipefail

DOCS_DIR="docs/ecwid"

if ! command -v pandoc &> /dev/null; then
    echo "Error: pandoc is not installed. Please install it first."
    exit 1
fi

echo "Converting HTML files to markdown..."

# Convert all HTML files to markdown
for html_file in "$DOCS_DIR"/*.html; do
    if [ -f "$html_file" ]; then
        md_file="${html_file%.html}.md"
        echo "Converting: $(basename "$html_file") -> $(basename "$md_file")"
        pandoc -f html -t markdown "$html_file" -o "$md_file"
        if [ -s "$md_file" ]; then
            rm "$html_file"
        else
            echo "Conversion failed for $(basename "$html_file"): empty output"
            rm -f "$md_file"
        fi
    fi
done

echo ""
echo "✓ Conversion complete!"
echo "Files saved to: $DOCS_DIR/"
