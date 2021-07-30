# Note: This script is only be executed inside the docker container
# Assumes to be running on ubuntu 20.04
# Ports required to be allocated internally:
# - 8888 (laplace server)
# - 24800 (barrier server)

# Updating and installing go compiler
apt update
apt install -y golang

# Installing git
apt install -y git

# Installing barrier
apt install -y barrier

# Installing chromium
wget https://github.com/RobRich999/Chromium_Clang/releases/download/v94.0.4585.0-r904940-linux64-deb-avx/chromium-browser-unstable_94.0.4585.0-1_amd64.deb
apt install -y ./v94.0.4585.0-r904940-linux64-deb-avx/chromium-browser-unstable_94.0.4585.0-1_amd64.deb

# clone remotegameplay distribution
git clone https://github.com/Akilan1999/remotegameplay
# enter cloned directory
cd remotegameplay

export REMOTEGAMEPLAY=$PWD

# Build laplace binary file
go build .

./laplace -setconfig
./laplace -tls -addr 0.0.0.0:8888 &
./laplace -headless -addr `$1` &

sleep 2

./laplace -headless -roomInfo >> /tmp/output.txt

# Installation game script here and start remote gameplay





