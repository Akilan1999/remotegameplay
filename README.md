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

# Index

- [Testing with X-Plane 11](#testing-with-x-plane-11)
- [Laplace](#laplace)
- [Installation](#installation)
  - [Installation required to share keyboard and mouse](#installation-required-to-share-keyboard-and-mouse)
    - [What is x2x?](#what-is-x2x)
    - [What is Barrier KVM](#what-is-barrier-kvm)
    - [Barrier KVM build status and links to install](#barrier-kvm-build-status-and-links-to-install)
  - [Build from Source](#build-from-source)
- [Program Execution](#program-execution)
  - [Starting game when screen-share is triggered](#starting-game-when-screen-share-is-triggered)
    - [Ex: Start X-Plane 11](#ex-start-x-plane-11)
    - [Open config file](#open-config-file)
  - [Call from built script which starts the server , creates room and outputs the ID](#call-from-built-script-which-starts-the-server--creates-room-and-outputs-the-id)
  - [Starting the Server](#starting-server)
  - [Starting the Screenshare](#starting-screenshare)
- [Contributing](#contributing)
- [Discord Server Link](#or)

---

## Testing with [X-Plane 11](https://www.x-plane.com/)
[![A screenshot of X-Plane 11 (flight simulator game) running smoothly via WebRTC](https://i.ytimg.com/vi/65dn7TRgzeE/hqdefault.jpg)](https://www.youtube.com/watch?v=65dn7TRgzeE "Running X-Plane 11 using WebRTC")

## Laplace
Based on the fork:
https://github.com/Akilan1999/laplace/tree/keyboard_mouse

## Installation


### Installation required to share keyboard and mouse
Currently, you can either use [x2x](https://github.com/dottedmag/x2x) or [Barrier KVM]()
We need to ensure that the client has SSH client installed or Barrierc.

#### What is x2x?
[x2x](https://github.com/dottedmag/x2x) allows the keyboard, mouse on one X display to be used to control another X
display. It also shares X clipboards between the displays.

Note: x2x runs on top of SSH.

#### What is Barrier KVM?

Barrier is software that mimics the functionality of a KVM switch, which historically would allow you to use a single
keyboard and mouse to control multiple computers by physically turning a dial on the box to switch the machine you're
controlling at any given moment. Barrier does this in software, allowing you to tell it which machine to control by
moving your mouse to the edge of the screen, or by using a keypress to switch focus to a different system.

#### Barrier KVM build status and links to install
|Platform       |Build Status|
|            --:|:--         |
|Linux          |[![Build Status](https://dev.azure.com/debauchee/Barrier/_apis/build/status/debauchee.barrier?branchName=master&jobName=Linux%20Build)](https://dev.azure.com/debauchee/Barrier/_build/latest?definitionId=1&branchName=master)|
|Mac            |[![Build Status](https://dev.azure.com/debauchee/Barrier/_apis/build/status/debauchee.barrier?branchName=master&jobName=Mac%20Build)](https://dev.azure.com/debauchee/Barrier/_build/latest?definitionId=1&branchName=master)|
|Windows Debug  |[![Build Status](https://dev.azure.com/debauchee/Barrier/_apis/build/status/debauchee.barrier?branchName=master&jobName=Windows%20Build&configuration=Windows%20Build%20Debug)](https://dev.azure.com/debauchee/Barrier/_build/latest?definitionId=1&branchName=master)|
|Windows Release|[![Build Status](https://dev.azure.com/debauchee/Barrier/_apis/build/status/debauchee.barrier?branchName=master&jobName=Windows%20Build&configuration=Windows%20Build%20Release%20with%20Release%20Installer)](https://dev.azure.com/debauchee/Barrier/_build/latest?definitionId=1&branchName=master)|
|Snap           |[![Snap Status](https://build.snapcraft.io/badge/debauchee/barrier.svg)](https://build.snapcraft.io/user/debauchee/barrier)|


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
## Installing barrier
apt install -y barrier
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

## Program Execution

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
#### Ex: Start X-Plane 11
Let's call this script `xplane11.sh`
```bash
# Navigating to the directory where X-Plane11 is present 
cd /path/.local/share/Steam/steamapps/common/X-Plane\ 11/

# Execute X-Plane 11 binary 
./X-Plane-x86_64
```
#### Open config file 
```bash

{
  "barrierhostname": "<barrier host name>",
  "ipaddress": "0.0.0.0",
  "rooms": "<path to room.json file>",
  "scripttoexecute": "<path to script to execute (In case the Xplane 11 script)>",
  "sshpassword": "<SSH password for x2x>",
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
