# goping
Simple GO service to check if the server is reachable in case ICMP Echo Requests is disabled (standard `ping`). It reads the `/proc/uptime` file, which is available on Linux systems.

## Returns JSON response:
### OK
Status Code *200*
```json
{"status": "ok", "uptime": "45 days"}
```

### Error
Status Code *400*
```json
{"status": "failed"}
```

## Run as systemd service

1. `make build` or `make build-arm` for ARM
2. Copy `build/goping` binary to `/usr/local/bin` on server
3. Copy `systemd/goping.service` to `/etc/systemd/system` (adjust values if needed)
4. Reload systemd and enable the service
```shell
sudo systemctl daemon-reload  # Reload unit files
sudo systemctl enable goping  # Enable to start at boot
sudo systemctl start goping   # Start the service
```
5. Check systemctl status
```shell
systemctl status goping
```
