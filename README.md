<h1 align="center">
  <br>
  <a href=""><img src="https://user-images.githubusercontent.com/31743758/132035109-cd8a145b-6e32-4d16-b9f8-f77f8de46a12.png" alt="p2prc" width="400"></a>
  <br>
</h1>

<!-- seperator -->
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Akilan1999/p2p-rendering-computation/graphs/commit-activity)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

The aim of this project is develop a WebRTC screenshare designed for streaming video games and
accepting remote inputs.<br>
There will be ansible instructions which can be executed inside into any virtual environment. This will
be a plugin which complements the project [P2PRC](https://p2prc.akilan.io)

[//]: # (# Index)

[//]: # ()
[//]: # (- [Testing with X-Plane 11]&#40;#testing-with-x-plane-11&#41;)

[//]: # (- [Laplace]&#40;#laplace&#41;)

[//]: # (- [Installation]&#40;#installation&#41;)

[//]: # (  - [Installation required to share keyboard and mouse]&#40;#installation-required-to-share-keyboard-and-mouse&#41;)

[//]: # (    - [What is x2x?]&#40;#what-is-x2x&#41;)

[//]: # (    - [What is Barrier KVM]&#40;#what-is-barrier-kvm&#41;)

[//]: # (    - [Barrier KVM build status and links to install]&#40;#barrier-kvm-build-status-and-links-to-install&#41;)

[//]: # (  - [Build from Source]&#40;#build-from-source&#41;)

[//]: # (- [Program Execution]&#40;#program-execution&#41;)

[//]: # (  - [Starting game when screen-share is triggered]&#40;#starting-game-when-screen-share-is-triggered&#41;)

[//]: # (    - [Ex: Start X-Plane 11]&#40;#ex-start-x-plane-11&#41;)

[//]: # (    - [Open config file]&#40;#open-config-file&#41;)

[//]: # (  - [Call from built script which starts the server , creates room and outputs the ID]&#40;#call-from-built-script-which-starts-the-server--creates-room-and-outputs-the-id&#41;)

[//]: # (  - [Starting the Server]&#40;#starting-server&#41;)

[//]: # (  - [Starting the Screenshare]&#40;#starting-screenshare&#41;)

[//]: # (- [Contributing]&#40;#contributing&#41;)

[//]: # (- [Discord Server Link]&#40;#or&#41;)

[//]: # ()
[//]: # (---)

## Testing with [X-Plane 11](https://www.x-plane.com/)
[![A screenshot of X-Plane 11 (flight simulator game) running smoothly via WebRTC](https://i.ytimg.com/vi/65dn7TRgzeE/hqdefault.jpg)](https://www.youtube.com/watch?v=65dn7TRgzeE "Running X-Plane 11 using WebRTC")

## Laplace
Based on the fork:
https://github.com/Akilan1999/laplace/tree/keyboard_mouse

## Installation


### Installation required to share keyboard and mouse
Currently, you can use [Barrier KVM]().<br>

[//]: # (#### What is x2x?)

[//]: # ([x2x]&#40;https://github.com/dottedmag/x2x&#41; allows the keyboard, mouse on one X display to be used to control another X)

[//]: # (display. It also shares X clipboards between the displays.)

[//]: # ()
[//]: # (Note: x2x runs on top of SSH.)

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
# Install golang 
# https://go.dev/doc/install

# Build 
go build .

# Configure (run only once)
./remotegaming -setconfig

# Migrate (run only once)
./remotegaming -Migrate

# The following are OR Based instruction.
# You could start the server in either of
# of the modes. (run every time you
# would want to start a server)

# Setting up based on the generated config 
# open config.json 
{
 "SystemUsername": "<auto-set>",
 "BarrierHostName": "<auto-set>",
 "Rooms": "<auto-set>",
 "IPAddress": "<auto-set>",
 "ScriptToExecute": "",
 "SSHPassword": "",
 "NATEscapeGameServerPort": "<auto-set>",
 "NATEscapeScreenSharePort": "<auto-set>",
 "NATEscapeBarrierPort": "",
 "BackendURL": "<auto-set> or set to a gameserver externally hosted",
 "BrowserCommand": "Requires the user to setup the command that triggers the browser from the Terminal or Command line",
 "Rate": 0,
 "ScreenName": "Entire screen (The default point) or add custom screename or tab",
 "InternalGameServerPort": "8088",
 "InternalScreenSharePort": "8888"
}

## Both Servers 
## Starts the remote gaming screenshare server and the gameserver 
./remotegaming -BothServers 
## Start only game server
./remotegaming -GameServer
## Start only screenshare
./remotegaming -tls -headless 
```

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

[//]: # (#### Open config file )

[//]: # (```bash)

[//]: # ()
[//]: # ({)

[//]: # (  "barrierhostname": "<barrier host name>",)

[//]: # (  "ipaddress": "0.0.0.0",)

[//]: # (  "rooms": "<path to room.json file>",)

[//]: # (  "scripttoexecute": "<path to script to execute &#40;In case the Xplane 11 script&#41;>",)

[//]: # (  "sshpassword": "<SSH password for x2x>",)

[//]: # (  "systemusername": "<system username>")

[//]: # (})

[//]: # (```)

[//]: # (### Call from built script which starts the server , creates room and outputs the ID)

[//]: # (```bash)

[//]: # (sh run.sh <IPV6 or Public IPV4 address of server>)

[//]: # (```)

[//]: # (Note: This script starts the server using the port 8888 by default<br>)

[//]: # (The 2 steps below are if you want to start them command by command )

[//]: # ()
[//]: # (### Starting server)

[//]: # (```bash)

[//]: # ($ ./laplace -tls -addr 0.0.0.0:8888)

[//]: # (2020/03/25 01:01:10 Listening on TLS: 0.0.0.0:8888)

[//]: # (```)

[//]: # ()
[//]: # (### Starting screenshare )

[//]: # (```bash)

[//]: # (./laplace -headless -addr <public ip address of server> )

[//]: # (```)

[//]: # ()
[//]: # (You can then open https://localhost:8888/ to view Laplace page.<br>)

[//]: # (You may need to add certificate exceptions. In Chrome, you can type `thisisunsafe`.)



## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

### or 

[![Support Server](https://discordapp.com/api/guilds/854397492795277322/widget.png?style=banner2)](https://discord.gg/b4nRGTjYqy)
