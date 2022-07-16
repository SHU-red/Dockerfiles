# :snowflake: hddfancontrol

## DockerHub

- https://hub.docker.com/r/shur3d/hddfancontrol

## Build

See general description at https://github.com/SHU-red/Docker

## Disclaimer

- Using [hddfancontrol](https://github.com/desbma/hddfancontrol)
- Dockerfile and Scripts based on [fightforlife/docker_hddfancontrol](https://github.com/fightforlife/docker_hddfancontrol)
- Controlling fans, based on HDD and CPU temperatures
- Usage for homeserver

## ToDo / Ideas

[ ] -

## Docker-Compose
Example how i use it to control fan speed based on CPU and HDD temperature
```
version: "3"
services:
  hddfancontrol:
    image: shur3d/hddfancontrol:latest
    container_name: hddfancontrol
    restart: unless-stopped
    volumes:
      - /lib/modules:/lib/modules:ro
    privileged: true
    cap_add:
      - SYS_MODULE
    environment:
      - DEVICES=/dev/sda1 /dev/sdb1 /dev/sdc1 /dev/sdc2
      - PWM_DEVICES=/sys/class/hwmon/hwmon2/pwm1 /sys/class/hwmon/hwmon2/pwm2
      - PWM_START=70 80
      - PWM_STOP=20 30
      - MIN_TEMP=25
      - MAX_TEMP=35
      - MIN_FAN=20
      - INTERVALL=10
      - SPINDOWN_TIME=900
      - LOG_PATH=/var/log/hddfancontrol.log
      - CPU_TEMP_RANGE=35 60
      - CPU_TEMP_PROBE=/sys/devices/platform/coretemp.0/hwmon/hwmon1/temp1_input
```

## Installation / Configuration
1. I had to find the paths to my devices (fans, HDDs, CPU)
2. Use these paths in the docker-compose as above
3. As my homeserver boots i execute the following command to load sensor kernel modules for my `ASROCK Mainboard`
```
$ modprobe coretemp; modprobe nct6775;
```
