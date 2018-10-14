# rebuilder
Periodically rebuild broker if it has gone down.  This is a temporary solution while we have not resolved the badger
crashes.  The current expectation is that brokernode code lives in /home/ubuntu/brokernode and that this repo will live in
/home/ubuntu/rebuilder, and that the status path is :3000/api/v2/status.  We should make all of this configurable if
we have to use this rebuilder for very long.

Take note that the docker containers must be built in debug mode (DEBUG=1) for this rebuilder to work.

1.  Set up
curl https://raw.githubusercontent.com/oysterprotocol/rebuilder/master/setup.sh | bash

2.  Set cron job.  This runs it every 10 minutes and writes output to rebuilder.log
crontab -e
*/10 * * * * cd /home/ubuntu/rebuilder && ./rebuilder >> /home/ubuntu/rebuilder/rebuilder.log 2>&1