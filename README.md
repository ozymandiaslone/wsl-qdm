# wsl-qdm
### What is qdm?
- A command line tool for easily sending files in between your WSL and Windows File Systems.
- Currently, qdm is under development, and only supports a limited number of features.

# Usage
After installing qdm, you can use it to send files between your WSL and Windows File Systems.
The `send` command acts like the unix `cp` command, though only takes in the path of the source file or directory. This file/dir will be copied into the preconfigured Windows file system folder. 
By default, all files will be sent to `C:\Users\<your-windows-username>\Documents\fromWsl`
If you cannot find the fromWsl folder in your Documents folder, ensure you are not looking in your OneDrive/Documents. Search from the abs path listed above.
Example useage:
```
qdm send .

qdm send /home/example/testfile.txt

qdm send testDirectory
```

## Coming features
- Grabbing files from Windows file system to copy to WSL.
- Sending/Grabbing files while deleting the original as an option.
- Ability to change default configured Windows file system path.
- Ability to dynamically input both source and destination paths if so desired by the user.


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

