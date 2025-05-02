# How to test
## 1.Docker Build
```
docker build -t cloudrun .
```
## 2.Docker Run & Load env
```
docker run --env-file .env -p 8080:8080 cloudrun
```
## 3.cd on wsl
```
cd /mnt/c/Users/sannz/fleet/ict15-tools/CloudRun/test/
```
