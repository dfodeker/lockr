# Lockr
Lockr is a secure and lightweight CLI tool for managing sensitive data, such as secrets, API keys, and credentials. Inspired by tools like Git, Lockr provides a simple and modular way to store and manage secrets in a hidden, secure directory (.lockr-vault). It ensures that sensitive data is kept out of version control by automatically updating your .gitignore file.
## Installation
Prerequisites
> Go (1.20 or higher)

Build from Source
Clone the repository:
```bash
git clone https://github.com/dfodeker/lockr.git
```
Navigate to the project directory:

```bash
cd lockr
```
Build the project:
    
```bash
go build -o lockr
```
Move the binary to your PATH:

```bash
sudo mv lockr /usr/local/bin/
```

## Usage
Initialize a New Lockr Vault
To start using Lockr, initialize a new vault in your project:
    
```bash
lockr init
```
This creates a hidden .lockr-vault directory and updates your .gitignore file to exclude it from version control.

