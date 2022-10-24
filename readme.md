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
### Documentacion de API
Ingresando desde un navegador web a la IP donde se encuentra desplegado el sistema. Si es en su misma equipo: https://localhost/
### Para correr tests API
Si ya ha corrido el sistema debe eliminar las imagenes creadas para generar el contenedor de pruebas
```bash
docker-compose down --rmi all
```
Entonces hacer
```bash
docker-compose -f test-project.yml up
```
Y podra ver los resultados de los test en el log del contenedor api-horus. Para volver a instanciar los contentenedores de
produccion debera nuevamente utilizar el comando:
```bash
docker-compose down --rmi all
```
### Credenciales - Grafana
- URL: https://localhost:3000
- Usuario: admin
- Contraseña: admin

### Credenciales - InfluxDB
- URL: http://localhost:8086
- Usuario: jccic
- Contraseña: C4p1t4nM4c31r4