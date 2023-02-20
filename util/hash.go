package util

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// Define salt size
const SaltSize = 16

// Generate 16 bytes randomly and securely using the
// Cryptographically secure pseudorandom number generator (CSPRNG)
// in the crypto.rand package
func GenerateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

// Combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password
// as a hex string
func HashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

// Check if two passwords match
func DoPasswordsMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var CurrPasswordHash = HashPassword(currPassword, salt)

	return hashedPassword == CurrPasswordHash
}

func HashSaltPassword (pass string) (string, error) {
	generateHashSalt, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return string(generateHashSalt), err
	}

	return string(generateHashSalt), nil
}

func CheckPasswordCredibility (notHashedPass string, HashedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(HashedPass), []byte(notHashedPass))	
} 

/*func main() {
	// First generate random 16 byte salt
	var salt = generateRandomSalt(saltSize)

	// Hash password using the salt
	var hashedPassword = hashPassword("hello", salt)

	fmt.Println("Password Hash:", hashedPassword)
	fmt.Println("Salt:", salt)

	// Check if passed password matches the original password by hashing it
	// with the original password's salt and check if the hashes match
	fmt.Println("Password Match:",
		doPasswordsMatch(hashedPassword, "hello", salt))
}*/