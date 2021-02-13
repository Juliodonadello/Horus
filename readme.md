# Facultad de Ingeniería del Ejército
## Proyecto de Promoción y Síntesis

* Integrante: CT Carlos A. MACEIRA GARCIA CONI

## Horus: Tablero de Control del JCCIC
Sistema de monitorización de KPI de los CCIC del EA.

### Requisitos
* Docker
### Instalación
```bash
git clone https://proyecto-horus-admin@bitbucket.org/proyecto-horus/horus-jccic-principal.git
docker volume create --name=grafana-data
docker-compose up --build
```