# Instructions for Building

- Make sure the program has permission to read and write on user home directory
- Create a folder `.nepdate` on user home directory.
- `go build`
- `./nepdate [options]` for usage 


# Instructions for Installing and Using

- Download the zip file from the release section.
- Unzip
- Make sure `install.sh` has executable permissions (`chmod +x install.sh`)
- Run `install.sh`
- Put `.nepdate/bin` in your path variable: `export PATH="$HOME/.nepdate/bin:$PATH"` to `.zshrc` or `.bashrc` depending upon your shell.