package options

// request for option switcher
var optionRequest = `What do you want to do?
    e - encryption new data from file
    d - decryption data from file

> `

// requests for encryption option
var (
	encFilePathRequest      = "Enter path to file with data to encryption: "
	encStoragePathRequest   = "Enter storage path to keep encrypted data: "
	firstPassphraseRequest  = "Enter passphrase for encryption process: "
	secondPassphraseRequest = "Re-enter passphrase for encryption process: "
)

// requests for decryption option
var (
	decFilePathRequest   = "Enter path to file with data to decryption: "
	decPassphraseRequest = "Enter passphrase for decryption process: "
)
