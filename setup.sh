## Install golang
apt install golang-go -y

## Set up go envs
mkdir ~/.go
echo "GOPATH=$HOME/.go" >> ~/.bashrc
echo "export GOPATH" >> ~/.bashrc
echo "PATH=\$PATH:\$GOPATH/bin # Add GOPATH/bin to PATH for scripting" >> ~/.bashrc
source ~/.bashrc

## Clone repo
cd /home/ubuntu
git clone https://github.com/oysterprotocol/rebuilder.git

## Build
cd rebuilder
go build