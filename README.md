<h1 align="center">
  <br>
  <a href=""><img src="https://user-images.githubusercontent.com/31743758/132035109-cd8a145b-6e32-4d16-b9f8-f77f8de46a12.png" alt="p2prc" width="400"></a>
  <br>
</h1>

<!-- seperator -->
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Akilan1999/p2p-rendering-computation/graphs/commit-activity)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

The aim of this project is develop a WebRTC screenshare designed for streaming video games and
accepting remote inputs.
There will be ansible instructions which can be executed inside into any virtual environment. This will
be a plugin which complements the project [P2PRC](https://p2prc.akilan.io)

## Testing with Xplane 11
[![IMAGE ALT TEXT](https://i.ytimg.com/vi/65dn7TRgzeE/hqdefault.jpg)](https://www.youtube.com/watch?v=65dn7TRgzeE "Running Xplane 11 using WebRTC")

## Laplace
Based on the fork:
https://github.com/Akilan1999/laplace/tree/keyboard_mouse

## Installation


### Installation required to share keyboard and mouse
We need to ensure that the client has SSH client installed.
We use the popular open repository known as [x2x](https://github.com/dottedmag/x2x). 

Note: x2x runs on top of SSH. 

#### What is x2x?
x2x allows the keyboard, mouse on one X display to be used to control another X
display. It also shares X clipboards between the displays.


### Build from source

```bash
# Assumes to be running on ubuntu 20.04
# Ports required to be allocated internally:
# - 8888 (laplace server)

# Updating and installing go compiler
apt update
apt install -y golang
apt install -y jq
## Installing git
apt install -y git
# Installing OpenSSH server 
apt install -y openssh-server
## Installing x2x
apt install -y x2x
## Installing chromium
wget https://github.com/RobRich999/Chromium_Clang/releases/download/v94.0.4585.0-r904940-linux64-deb-avx/chromium-browser-unstable_94.0.4585.0-1_amd64.deb
apt install -y ./chromium-browser-unstable_94.0.4585.0-1_amd64.deb
#
## clone remotegameplay distribution
git clone https://github.com/Akilan1999/remotegameplay
## enter cloned directory
cd remotegameplay
# sets REMOTEGAMEPLAY path
export REMOTEGAMING=$PWD

# Build laplace binary file
go build .

# Set configuration laplace file
./laplace -setconfig

# Open config.json file and add SSH password for x2x
```

### Program Execution

Executing this project basically serves an HTTP server that will host the frontend and the WebSocket implementation.
Note that you sometimes need to run HTTPs in order for browser to connect to websocket.

```bash
$ ./laplace --help
  -BinaryToExecute string
    	Providing path (i.e Absolute path) of binary to execute
  -addr string
    	Listen address (default "0.0.0.0:443")
  -certFile string
    	TLS cert file (default "files/server.crt")
  -headless
    	Creating screenshare using headless mode
  -keyFile string
    	TLS key file (default "files/server.key")
  -killChromium
    	Kills all chromuim
  -killServer
    	Kills the laplace
  -roomInfo
    	Getting room id of headless server
  -setconfig
    	Generates a config file
  -tls
    	Use TLS
```

By default, you can run the executable without any argument to listen to TLS port 443.
A self-signed certificate files are provided to ease up development. If you want to run 
with barrier KVM. Run as non-root. 

### Starting game when screen-share is triggered
This requires creating a bash script to trigger when the screenshare begins.
#### Ex: Start Xplane 11
Let's call this script xplane11.sh
```bash
# Navigating to the directory where XPlane11 is present 
cd /path/.local/share/Steam/steamapps/common/X-Plane\ 11/

# Execute Xplane 11 binary 
./X-Plane-x86_64
```
#### Open config file 
```bash
{
  "barrierhostname": "<barrier host name>",
  "ipaddress": "0.0.0.0",
  "rooms": "<path to room.json file>",
  "scripttoexecute": "<path to script to execute (In case the Xplane 11 script)>",
  "systemusername": "<system username>"
}
```
### Call from built script which starts the server , creates room and outputs the ID
```bash
sh run.sh <IPV6 or Public IPV4 address of server>
```
Note: This script starts the server using the port 8888 by default
The 2 steps below are if you want to start them command by command 

### Starting server
```bash
$ ./laplace -tls -addr 0.0.0.0:8888
2020/03/25 01:01:10 Listening on TLS: 0.0.0.0:8888
```

### Starting screenshare 
```bash
./laplace -headless -addr <public ip address of server> 
```

You can then open https://localhost:8888/ to view Laplace page.
You may need to add certificate exceptions. In Chrome, you can type `thisisunsafe`.



## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

### or 

[![Support Server](https://discordapp.com/api/guilds/854397492795277322/widget.png?style=banner2)](https://discord.gg/b4nRGTjYqy)


