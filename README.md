# rebuilder
Periodically rebuild broker if it has gone down.  This is a temporary solution while we have not resolved the badger
crashes.  The current expectation is that brokernode code lives in /home/ubuntu/brokernode and that this repo will live in
/home/ubuntu/rebuilder, and that the status path is :3000/api/v2/status.  We should make all of this configurable if
we have to use this rebuilder for very long.

Take note that the docker containers must be built in debug mode (DEBUG=1) for this rebuilder to work.

1.  Install golang
apt install golang-go -y

2.  Set up go envs
mkdir ~/.go
echo "GOPATH=$HOME/.go" >> ~/.bashrc
echo "export GOPATH" >> ~/.bashrc
echo "PATH=\$PATH:\$GOPATH/bin # Add GOPATH/bin to PATH for scripting" >> ~/.bashrc
source ~/.bashrc

3.  Clone repo
cd /home/ubuntu
git clone https://github.com/oysterprotocol/rebuilder.git

4.  Build
cd rebuilder
go build

5.  Set up cron job.  This runs it every 10 minutes and writes output to rebuilder.log
crontab -e
*/1 * * * * cd /home/ubuntu/rebuilder && ./rebuilder >> /home/ubuntu/rebuilder/rebuilder.log 2>&1