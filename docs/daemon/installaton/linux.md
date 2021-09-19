Panopticon on Linux
===


## Installation

### Get Binary

```
sudo cp ~/Downloads/panopticon /usr/sbin/panoticon
sudo chown root:root /usr/sbin/panoticon
sudo chmod 700 /usr/sbin/panoticon
```

### Configuration

```
# your server address : grpc://... or http://...
export ENDPOINT=<your-panopticon-server>

# create configuration directories
sudo mkdir -p /etc/panopticon/daemon
sudo mkdir -p /etc/panopticon/server

# create daemon configuration
sudo echo "
period: 30s
endpoint: ${ENDPOINT}
" > /etc/panopticon/daemon/config.yaml
```

### Define Systemd Unit

```
sudo echo "[Unit]
Description=Panopticon Daemon
After=syslog.target

[Service]
ExecStart=/usr/sbin/panoticon daemon --config-path=/etc/panopticon/daemon/config.yaml
Restart=always
KillSignal=SIGQUIT

[Install]
WantedBy=multi-user.target
" > /etc/systemd/system/panopticon-daemon.service
```

### Activate and Start Unit

```
sudo systemctl enable panopticon-daemon.service
sudo systemctl start panopticon-daemon.service
```

### Check Unit Status

```
sudo systemctl is-enabled panopticon-daemon.service
```

### Check Unit Logs

```
journalctl -u panopticon-daemon.service -f
```

### Reset Error Count

resolve service `start-limit-hit`

```
sudo systemctl reset-failed panopticon-daemon.service
```
