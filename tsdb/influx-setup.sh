#!/bin/bash
influxd & disown

sleep 20
influx config
influx setup --org ccic --bucket ccic-inicial --username dashboard --password ExAmPl3PA55W0rD --token u8C0TQ4SOAQ6za2BLZO-grxyIg5NT-WEQGNE9O6pYy73rPt9t0MNx8trbAgXEdyrwXYQ9dCRs5-zr3BMeeVufw== --force
influx bucket create --name colas-espera --retention 48h
influx bucket create --name estado-servicio --retention 48h
influx bucket create --name combustible-generador --retention 24h
influx bucket create --name alimentacion-electrica --retention 1h
influx bucket create --name horus-dashboard 
influx user create --name jccic --password C4p1t4nM4c31r4
pkill -f influxd
sleep 2