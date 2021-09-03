# Remote game play
The aim of this project is develop a WebRTC screenshare designed for streaming video games and
accepting remote inputs.
There will be ansible instructions which can be executed inside into any virtual environment. This will
a plugin which complements the project [P2PRC](https://p2prc.akilan.io)

## Laplace
Based on the fork:
https://github.com/Akilan1999/laplace/tree/keyboard_mouse

## Installation


### Installation required to share keyboard and mouse
To do this we ensure that the client either has has a IPV6 
address or a public IPV4 address. 
We use the popular open repository known as [Barrier KVM](https://github.com/debauchee/barrier). 

#### What is Barrier kvm?

Barrier is software that mimics the functionality of a KVM switch, which historically would allow you to use a single keyboard and mouse to control multiple computers by physically turning a dial on the box to switch the machine you're controlling at any given moment. Barrier does this in software, allowing you to tell it which machine to control by moving your mouse to the edge of the screen, or by using a keypress to switch focus to a different system.

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
# - 24800 (barrier server)

# Updating and installing go compiler
apt update
apt install -y golang
apt install -y jq
## Installing git
apt install -y git
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


