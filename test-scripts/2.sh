U973qTVC8Thr/SfiWe5jK13U2qwKx+zz7S7z6tVJe+l9OP6PZU5n7Npj/QsgpEebzqiTP+BpP7POhmr7xhZ0a+hgVJ2E2343J2uhERqjX8uMAY+IsQ31lcDHZNGgsjO1e18ILjB24iuZ8LgBETGc8w1H982CfmvVR/1jWxnU3/t+/WehDhRfIqffErncMgaIaulDPX6XQ8H9ACY5BtuyL4EBIREuaeoytLrbwM9l3ym6ARZuccFs6QO3EalhHcx5VJ/4kwu7S2HnWC1wl+k2MEBMjabKM5UAqXfJhAc6l8Y3T8BTUbBmWavqU0zp2RtjLN2Jx/7MIjISVO498FKzAw==
# Function to generate random prime numbers
generate_prime() {
  local max=$1
  shuf -i 2-$max -n 1 | factor | awk 'NF>2{print $2; exit}'
}

# Function to calculate the modular inverse
mod_inverse() {
  local a=$1 m=$2 m0=0 x0=1 x=0 q=0 temp=0
  while [ $a -gt 1 ]; do
    q=$((a / m))
    temp=$m
    m=$((a % m))
    a=$temp
    temp=$x0
    x0=$((x - q * x0))
    x=$temp
  done
  x=$((x + m0))  # Ensure x is positive
  echo $x
}

# Key generation
generate_keys() {
  local p q n phi e d
  p=$(generate_prime 1000)
  q=$(generate_prime 1000)
  n=$((p * q))
  phi=$(( (p-1)*(q-1) ))
  # e is a Fermat prime and has a low Hamming weight
  e=65537
  d=$(mod_inverse $e $phi)
  echo "Public key (e, n): ($e, $n)"
  echo "Private key (d, n): ($d, $n)"
}

# Example usage
generate_keys
