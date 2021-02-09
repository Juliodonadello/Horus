#!/bin/bash
influxd & disown
sleep 5
influx setup --org ccic --bucket mensajes --username dashboard --password ExAmPl3PA55W0rD --token u8C0TQ4SOAQ6za2BLZO-grxyIg5NT-WEQGNE9O6pYy73rPt9t0MNx8trbAgXEdyrwXYQ9dCRs5-zr3BMeeVufw== --force
influx bucket create --name cola-cmd --retention 48h
influx bucket create --name mensajes-ccic --retention 48h
influx bucket create --name gpos-rtef --retention 48h
influx bucket create --name estado-servicio --retention 48h
influx bucket create --name combustible-generador --retention 48h
influx user create --name jccic --password ExAmPl3PA55W0rD
pkill -f influxd
sleep 2