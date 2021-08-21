# Note: This script is only be executed inside the docker container
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
export REMOTEGAMEPLAY=$PWD

# Build laplace binary file
go build .

# Set configuration laplace file
./laplace -setconfig