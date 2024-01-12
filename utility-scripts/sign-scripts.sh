#!/bin/bash

UNSIGNED_DIR="unsigned-scripts"
SIGNED_DIR="signed-scripts"
KEY_DIR="keys"

# Loop through all .sh scripts in the current directory
for script_file in $UNSIGNED_DIR/*.sh; do
    if [ -f "$script_file" ]; then

        # Get script base name
        script_base_name=$(basename "$script_file")
        
        # Read the content of the script excluding the shebang line
        script_content=$(tail -n +1 "$script_file")

        # Sign the script, convert to readable, and remove line breaks
        signature=$(openssl dgst -sha256 -sign $KEY_DIR/private-key.pem "$script_file" | base64 | tr -d '\n')
        
        # Create a new signed script with shebang, signature, and content
        echo "$signature" > $SIGNED_DIR/"$script_base_name"
        echo "$script_content" >> $SIGNED_DIR/"$script_base_name"
    fi
done
