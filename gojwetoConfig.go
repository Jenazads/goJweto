package gojweto

import (
  "os";
  "log";
  "crypto/rsa";
)

// CertToken is an Struct to encapsulate username and expires as parameter
type CertToken struct {
  Username string
  Expires int64
}

type CredentialsAuth struct {
  Token   string  `json:"Token"`
  Logged  bool    `json:"Logged"`
}

type CredentialsNoAuth struct {
  Logged  bool    `json:"Logged"`
}

const (
  privKey    = "/tls-ssl/jwtkeys/rsakey.pem"
  pubKey     = "/tls-ssl/jwtkeys/rsakey.pem.pub"
)

var (
  verifyKey   *rsa.PublicKey
  signingKey  *rsa.PrivateKey
  secretKey   = "Jnzads"
  headerKey   = "Jnzads-JWT"
  pwd, _      = os.Getwd()
)

func fatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

// Header authorization
func SetHeaderKey(name string){
  headerKey = name
}

func GetHeaderKey()(string){
  return headerKey
}

// method HS256
func SetSecretKey(name string){
  secretKey = name
}

func GetSecretKey()(string){
  return secretKey
}

func GetSecretByte()([]byte){
  return []byte(secretKey)
}
