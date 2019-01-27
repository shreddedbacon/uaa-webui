# uaa-webui
Webui for uaa


# Docker
## Docker Compose
```
docker-compose up -d
## OR to rebuild
docker-compose up -d --build
```

## Build
```
docker build -t shreddedbacon/uaa-webui .
```
## Run
```
docker run -p 8443:8443 \
  -e CLIENT_ID=uaa-admin \
  -e CLIENT_SECRET=uaa-adminsecret \
  -e UI_URL=https://localhost:8443 \
  -e UAA_SERVER=https://<ip>:<port> \
  shreddedbacon/uaa-webui
```
E.g
```
docker run -p 8443:8443 \
  -e CLIENT_ID=uaa-admin \
  -e CLIENT_SECRET=uaa-adminsecret \
  -e UI_URL=https://localhost:8444 \
  -e UAA_SERVER=https://192.168.50.6:8443 \
  shreddedbacon/uaa-webui
```
