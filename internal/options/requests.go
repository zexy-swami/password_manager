package options

// request for option switcher
var optionRequest = `What do you want to do?
    e - encrypt new data from file
    d - decrypt data from file

> `

// requests for encryption option
var (
	encFilePathRequest      = "Enter path to file with data to encrypt: "
	encStoragePathRequest   = "Enter storage path to keep encrypted data: "
	firstPassphraseRequest  = "Enter passphrase for encryption process: "
	secondPassphraseRequest = "Re-enter passphrase for encryption process: "
)

// requests for decryption option
var (
	decFilePathRequest   = "Enter path to file with data to decrypt: "
	decPassphraseRequest = "Enter passphrase for decryption process: "
)
