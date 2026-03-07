#!/bin/bash

# Script to convert HTML files to markdown

set -e

DOCS_DIR="docs/ecwid"

echo "Converting HTML files to markdown..."

# Convert all HTML files to markdown
for html_file in "$DOCS_DIR"/*.html; do
    if [ -f "$html_file" ]; then
        md_file="${html_file%.html}.md"
        echo "Converting: $(basename "$html_file") -> $(basename "$md_file")"
        pandoc -f html -t markdown "$html_file" -o "$md_file"
        # Remove the HTML file after successful conversion
        rm "$html_file"
    fi
done

echo ""
echo "✓ Conversion complete!"
echo "Files saved to: $DOCS_DIR/"
