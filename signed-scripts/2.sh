X4eFo2rOzTACgOBnyQG+xR6i073AhzR/KmomujLtrJceljqtqVDoLXnDvESYn9fSJm40iQvd+qpI5x3wGNZda1e5BDYauNuh3ALpXDQIWF3CcwnUtmuHyxRAk7zDM0YDxufHZCf1CCgjknYHeRVuf1NxCDcMv+7aeXaxwWaOTkfLkorqXQ+EB2TgSPir3XULLplpVgBG9NtSQFeZw8X0G6cTMKZoJ9VK6g/et561kxhI80ARhJ5IBy0xKtUR14uLEyBlXE5y6E3+MvmkLREhfpODm1D/6/2w6lTSb95bWAaFmHCh4uuE/gEXpzGHAk1+2Pa8eQ8BVMvuv/sZvi+QqA==
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
