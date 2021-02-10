### Requisitos
* Docker
### Instalación
```bash
docker volume create --name=grafana-data
docker volume create --name=influxdb-data
docker-compose up --build
```