# wsl-qdm

# Installation instructions
## Ubuntu - WSL (only for WSL/WSL2)
- 1 `git clone https://github.com/ozymandiaslone/wsl-qdm`
- 2 `cd wsl-qdm`
- 3 `sudo cp ./build/qdm /usr/local/bin/qdm`
It should now be installed on your system. If it does not work, ensure the `qdm` command is in your PATH.
- 1 `sudo chmod +x /usr/local/bin/qdm`
- 2 `vi ~/.bashrc`
Add the following line to the end of the file:
```
export PATH="/usr/local/bin:$PATH"
```
Back in the terminal, run:
- 3 `source ~/.bashrc`

