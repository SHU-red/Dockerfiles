# Docker
## DockerHub
https://hub.docker.com/u/shur3d
## Disclaimer
- GitHub Repository for managing all self-created Docker-Containers
- Each Dockerfile with additional artifacts in separate folders
- Additional artifacts could be bash-scripts for Contaier-Build or Executables written & built in other languages such as GO
- Build scripts for Dockerfile and additional artifacts
## Build
1. Make sure `buildx` is working on your machine
2. Get repository
Inside the Container-folder which is to be built:
3. Build eventually necessary artifacts if you e.g. change something in the go-scripts
4. Change credentials in `build_docker.sh` to your values
5. Start terminal
6. execute `$ bash build_docker.sh`
Build should be ran and container be pushed to your DockerHub account
