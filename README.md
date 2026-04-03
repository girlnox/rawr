# rawr
Simple package manager and network file server, made for local file sending and receiving

# Note:
This project isn't finished. I do not expect anyone to use this and I'm really just having fun :3 
So again, this is JUST a hobby project.
I'm also still working on some big update, adding the main function of actually installing packages (wow, i know)
I will add a simple sort of scripting language that will be able to tell the rawr interpreter how to actually install the package.
Also rawr is and will always be able to just function as a simple file sharing service so if you need a really simple solution, then this might be enough!

# Installation
To install this project you could either grab a binary from the release section or
clone the repo ( `git clone https://github.com/girlnox/rawr.git` ) then make sure you have the dependencies listed in the `info/dependencies.txt` file installed and run
the export.sh script in the root of the project by simply running `bash export.sh` or `chmod +x export.sh` and then `./export.sh`

once you chose the right binary for your device's architecture from the `exports/` directory, just simply rename it to "`rawr`"
and move it to a folder in your $PATH.

(there will be an install script later on)

# Usage
`rawr receive [port]`: host an upload server for receiving packages

`rawr serve [port] [directory]`: host a download server for downloading packages 


`rawr give [file] [ip:port]`: give a package to an open upload server

`rawr get [package_path] [ip:port]`: get a package from an open download server


`rawr pack [directory]`: create a rawr package

`rawr unpack [package]`: unpack a rawr package


`rawr help/--help/(nothing)`: Display this message
