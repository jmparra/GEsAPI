package main

import (
  "fmt"
  //"os"
  // use for hash from scrypt
  "crypto/cipher"
  "golang.org/x/crypto/blowfish"
  //"encoding/base64"
  // get config file
  //"github.com/BurntSushi/toml"
)

type Config struct {
  Auth AuthConfig
}

type AuthConfig struct {
  Salt string `toml:"salt"`
}

func checkSize(pt []byte) []byte {
  modules := len(pt) % blowfish.BlockSize
  if modules != 0 {
    padlen := blowfish.BlockSize - modules
    for i := 0; i < padlen; i++ {
      pt = append(pt, 0)
    }
  }
  return pt
}

func encrypt(ppt, key []byte) []byte {
  ecipher, err := blowfish.NewCipher(key)
  if err != nil {
    panic(err)
  }
  ciphertext := make([]byte, blowfish.BlockSize + len(ppt))
  eiv := ciphertext[:blowfish.BlockSize]
  ecbc := cipher.NewCBCEncrypter(ecipher, eiv)
  ecbc.CryptBlocks(ciphertext[blowfish.BlockSize:], ppt)
  return ciphertext
}

func main() {
  fmt.Print("Passphrase: ")
  var pass string
  fmt.Scanln(&pass)
  fmt.Println(pass)
}
