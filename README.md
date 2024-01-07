#  Go encryptor

Golang data encryptor  package  by Saparov

#  Installing
Install from repository
<pre>
<code>go get github.com/Kesha005/go_encryptor</code>
</pre>

# Where to use

Encryption and decryption are used to secure data in auth,message systems and in cybersecurity

#  Usage

Firstly we need config our .env file and add there "SECRET_KEY" and "IV_16_KEY"

<h4>.ENV file</h4>
<pre>
    <code>
    SECRET_KEY="thismustbe16or24digitkey."
    IV_16_KEY="thisis16digitkey"
    </code>
</pre>

<h4>Example:</h4>

<p>Import package</p>

<pre>
    <code>import (
	"github.com/Kesha005/go_encryptor"
	"fmt"
	"github.com/joho/godotenv"

)</code>
</pre>

<p>Encryption and decryption</p>
<pre>
    <code>
    StringToEncrypt := "Encrypting this string"
	godotenv.Load(".env")
	fmt.Println(StringToEncrypt)
	encText, err := go_encryptor.Encrypt(StringToEncrypt)
	if err != nil {
		fmt.Println("error encrypting your classified text: ", err)
	}
	fmt.Println(encText)
	// To decrypt the original StringToEncrypt
	decText, err := go_encryptor.Decrypt(encText)
	if err != nil {
		fmt.Println("error decrypting your encrypted text: ", err)
	}
	fmt.Println(decText)
    </code>
</pre>

