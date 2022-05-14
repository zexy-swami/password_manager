# Password Manager
Software implementation of password manager in [go](https://go.dev) language.<br>

## Base information
Password manager allows users to keep encrypted information in the file system.<br>
Encryption and decryption processes realized with using [AES](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard) and 
[Galois Counter Mode](https://en.wikipedia.org/wiki/Galois/Counter_Mode).<br>
AES key is presented as a result of hashing user's passphrase with [SHA256](https://en.wikipedia.org/wiki/SHA-2).

## Input data format
Password manager takes data from file.<br>
Data should contain:
```
* name of file for saving encrypted data;
* login (id, nickname, email and etc);
* password;
```
Examples:<br>
<p align="center">
  <img width="505" height="64" src="https://github.com/zexy-swami/password_manager/blob/main/images/input_file_format.png">
</p>

## Encryption examples
### Linux (Ubuntu 20.04)
<p align="center">
  <img width="739" height="663" src="https://github.com/zexy-swami/password_manager/blob/main/images/encryption_linux.png">
</p>

### Windows
<p align="center">
  <img width="546" height="632" src="https://github.com/zexy-swami/password_manager/blob/main/images/encryption_windows.png">
</p>

## Decryption examples
### Linux (Ubuntu 20.04)
<p align="center">
  <img width="731" height="596" src="https://github.com/zexy-swami/password_manager/blob/main/images/decryption_linux.png">
</p>

### Windows
<p align="center">
  <img width="510" height="600" src="https://github.com/zexy-swami/password_manager/blob/main/images/decryption_windows.png">
</p>