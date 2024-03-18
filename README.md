# VK Internship 2024

### Setup

Before running the service, you need to copy the .env file template into project root:
```
cp config/template.env .env
```
After that, configure all required variables in .env file by replacing <> placeholders with desired values.

### Starting and stopping the service

For starting the servise, use:
```
sudo scripts/start.sh
```
For shutting down the service gracefully, use:
```
scripts/stop.sh
```