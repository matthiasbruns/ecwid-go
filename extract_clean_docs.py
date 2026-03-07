#!/usr/bin/env python3
"""
Extract clean markdown from Ecwid API documentation pages.
Uses html2text for better HTML to Markdown conversion.
"""

import os
import subprocess
import sys
import time
from pathlib import Path

# Base configuration
DOCS_DIR = Path("docs/ecwid")
BASE_URL = "https://docs.ecwid.com/api-reference"

# API sections to extract
sections = [
    ("", "overview"),
    ("rest-api-error-codes", "rest-api-error-codes"),
    ("rest-api/store-profile", "store-profile"),
    ("rest-api/orders", "orders"),
    ("rest-api/products", "products"),
    ("rest-api/categories", "categories"),
    ("rest-api/customers", "customers"),
    ("rest-api/discounts", "discounts"),
    ("rest-api/domains", "domains"),
    ("rest-api/dictionaries", "dictionaries"),
    ("rest-api/staff-accounts", "staff-accounts"),
    ("rest-api/application", "application"),
    ("rest-api/batch-requests", "batch-requests"),
    ("rest-api/shipping-options", "shipping-options"),
    ("rest-api/payment-options", "payment-options"),
    ("rest-api/checkout-extra-fields", "checkout-extra-fields"),
    ("rest-api/storefront-widget-details", "storefront-widget-details"),
    ("rest-api/carts", "carts"),
    ("rest-api/instant-site", "instant-site"),
    ("rest-api/reviews", "reviews"),
    ("rest-api/subscriptions", "subscriptions"),
]

def install_html2text():
    """Install html2text if not available."""
    try:
        import html2text
        return True
    except ImportError:
        print("Installing html2text...")
        subprocess.run([sys.executable, "-m", "pip", "install", "html2text"], check=True)
        import html2text  # noqa: F401
        return True

def fetch_and_convert(url, output_file):
    """Fetch URL and convert to clean markdown."""
    try:
        import html2text
        import urllib.request
        
        print(f"Fetching: {url}")
        
        # Fetch the page
        req = urllib.request.Request(url, headers={'User-Agent': 'Mozilla/5.0'})
        with urllib.request.urlopen(req, timeout=30) as response:
            html_content = response.read().decode('utf-8')
        
        # Configure html2text
        h = html2text.HTML2Text()
        h.ignore_links = False
        h.ignore_images = True
        h.ignore_emphasis = False
        h.body_width = 0  # Don't wrap lines
        h.skip_internal_links = False
        
        # Convert to markdown
        markdown = h.handle(html_content)
        
        # Write to file
        output_file.write_text(markdown, encoding='utf-8')
        print(f"  ✓ Saved to: {output_file}")
        return True
        
    except Exception as e:
        print(f"  ✗ Error: {e}")
        return False

def main():
    """Main extraction function."""
    print("Installing dependencies...")
    install_html2text()
    
    print("\nStarting Ecwid API documentation extraction...\n")
    
    # Create docs directory
    DOCS_DIR.mkdir(parents=True, exist_ok=True)
    
    success_count = 0
    fail_count = 0
    
    # Process each section
    for section_path, filename in sections:
        if section_path:
            url = f"{BASE_URL}/{section_path}"
        else:
            url = BASE_URL
        
        output_file = DOCS_DIR / f"{filename}.md"
        
        if fetch_and_convert(url, output_file):
            success_count += 1
        else:
            fail_count += 1
        
        # Be respectful to the server
        time.sleep(0.5)
    
    print("\n✓ Documentation extraction complete!")
    print(f"  Success: {success_count} files")
    print(f"  Failed: {fail_count} files")
    print(f"  Output directory: {DOCS_DIR}/")

if __name__ == "__main__":
    main()
