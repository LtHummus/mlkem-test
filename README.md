# ML-KEM Test

I had never had the opportunity to use any Post Quantum Computing algorithms, so I figured this would be a fun way to play around and get a feel for ML-KEM.

This is a simple CLI app that can generate an ML-KEM 768 keypair, as well as encrypt and decrypt some files. The encryption generates a shared secret and encapsulates it with the encapsulation key, uses that secret to derive an key for ChaCha20-Poly1305 and encrypts a file in 64 kb chunks. The encrypted file is written to disk with a header that contains the encrypted secret and the salt used for the key derivation.

The decryption side reverses everything. We read the encrypted secret and use the deencapsuation key to get back the shared secret, uses that (+ the salt) to derive the same ChaCha20-Poly1305 key and then decrypts everything.

## WARNING WARNING WARNING

THIS IS A TOY EXAMPLE. IT HAS NOT BEEN CHECKED BY ANYONE OR ANYTHING. DO NOT USE IT FOR CRITICAL DATA. DO NOT USE IT FOR IMPORTANT DATA. DO NOT USE IT FOR ANYTHING YOU CARE ABOUT AT ALL. THIS IS ME LEARNING AND PLAYING. THERE IS NO LICENSE. DO NOT USE THIS.

### Notes for future me

- We use KHDF for key derivation (w/ SHA256)
- Nonces contain chunk numbers (safe since keys are generated per file and will not be reused)
- AAD contains chunk number with a sentinel byte indicating if it is the last chunk or not (to detect truncated input)

#### Header format

- `MLKEMTEST` magic number
- 1-byte Version (currently 1)
- 8-byte Encapsulated key length
- 8-byte salt length
- Encapsulated key
- Salt
- Encrypted data

(yes i know 8 bytes is excessive, but sometimes nothing succeeds like excess...and the parser will reject values greater than 16kB anyway)
