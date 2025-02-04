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



1. Create a New Environment
This command creates a new environment (similar to creating a new branch in Git).
Command:
        
```bash
lockr env create <environment-name>
```
. Create and Switch to a New Environment
This command creates a new environment and switches to it immediately (similar to git checkout -b).

```bash
lockr env switch -c <environment-name>
```


3. Switch to an Existing Environment
This command switches to an existing environment (similar to git checkout).

Command:
```bash
lockr env switch <environment-name>
```

4. Check the Current Environment
This command displays the currently active environment (similar to git branch showing the current branch).

```bash
lockr env current
```

6. Delete an Environment
This command deletes an environment (similar to git branch -d).

```bash
lockr env delete <environment-name>
```