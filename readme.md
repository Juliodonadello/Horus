# Facultad de Ingeniería del Ejército
## Proyecto de Promoción y Síntesis

* Integrante: CT Carlos A. MACEIRA GARCIA CONI

## Horus: Tablero de Control del JCCIC
Sistema de monitorización de indicadores clave de performance de los Centros de Comunicaciones e Informáticad de Campaña del Ejército Argentino.

### Arquitectura
![Arquitectura Horus](./img/Sistema%20de%20Gestión%20Informático%20para%20las%20facilidades%20del%20CCIC%20-%20Arquitectura%20HORUS.jpeg)

### Requisitos
* Ubuntu 20.04
* Docker
* Usuario sudoer

### Tecnologías
* Docker
* Go
* PostgreSQL
* InfluxDB
* Grafana

### Instalación
```bash
git clone https://proyecto-horus-admin@bitbucket.org/proyecto-horus/horus-jccic-principal.git
cd horus-jccic-principal
docker volume create --name=grafana-data
docker network create --subnet 10.20.0.0/24 horus-ccic
docker-compose up --build
```
### Para dar TLS al dashboard
```bash
docker volume create --name=grafana-data
docker network create --subnet 10.20.0.0/24 horus-ccic
docker-compose up --build
```