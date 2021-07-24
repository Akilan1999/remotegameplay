export REMOTEGAMEPLAY=$PWD

go build .

./laplace -setconfig
./laplace -tls -addr 0.0.0.0:8888 &
./laplace -headless -addr 0.0.0.0:8888 &

sleep 2

./laplace -headless -roomInfo





