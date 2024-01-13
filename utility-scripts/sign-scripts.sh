#!/bin/bash

UNSIGNED_DIR="unsigned-scripts"
TEST_DIR="test-scripts"
KEY_DIR="keys"

mkdir -p "$TEST_DIR"

# Loop through all .sh scripts in the current directory
for script_file in $UNSIGNED_DIR/*.sh; do
    if [ -f "$script_file" ]; then

        # Get script base name
        script_base_name=$(basename "$script_file")
        
        # Read the content of the script excluding the first line
        script_content=$(tail -n +1 "$script_file")

        # Sign the script, convert to readable, and remove line breaks
        signature=$(openssl dgst -sha256 -sign $KEY_DIR/private-key.pem "$script_file" | base64 | tr -d '\n')
        
        # Create a new signed script with shebang, signature, and content
        echo "$signature" > $TEST_DIR/"$script_base_name"
        echo "$script_content" >> $TEST_DIR/"$script_base_name"
    fi
done

# Script 7
script_content=$(tail -n +1 "$UNSIGNED_DIR/7.sh")
signature=$(openssl dgst -sha256 -sign $KEY_DIR/private-key.pem "$UNSIGNED_DIR/7.sh" | base64 | tr -d '\n')
echo "$signature" > "$TEST_DIR/7.sh"
echo "$signature" >> "$TEST_DIR/7.sh"
echo "$script_content" >> "$TEST_DIR/7.sh"
# Script 9
script_content=$(tail -n +1 "$UNSIGNED_DIR/9.sh")
script_content="${script_content%?}X"
signature=$(openssl dgst -sha256 -sign $KEY_DIR/private-key.pem "$UNSIGNED_DIR/9.sh" | base64 | tr -d '\n')
echo "$signature" > "$TEST_DIR/9.sh"
echo "$script_content" >> "$TEST_DIR/9.sh"
# Script 10
script_content=$(tail -n +1 "$UNSIGNED_DIR/10.sh")
signature=$(openssl dgst -sha256 -sign $KEY_DIR/private-key.pem "$UNSIGNED_DIR/10.sh" | base64 | tr -d '\n')
signature="${signature%?}"
echo "$signature" > "$TEST_DIR/10.sh"
echo "$script_content" >> "$TEST_DIR/10.sh"