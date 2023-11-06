.SILENT:

include .env

### Hosts

main-hosts:
	echo "127.0.0.1  ${LOCAL_HOSTNAME_MAIN}"
	grep -q "127.0.0.1  ${LOCAL_HOSTNAME_MAIN}" "${HOSTS}" || echo '127.0.0.1  ${LOCAL_HOSTNAME_MAIN}' | sudo tee -a "${HOSTS}"

pma-hosts:
	echo "127.0.0.1  ${LOCAL_HOSTNAME_PMA}"
	grep -q "127.0.0.1  ${LOCAL_HOSTNAME_PMA}" "${HOSTS}" || echo '127.0.0.1  ${LOCAL_HOSTNAME_PMA}' | sudo tee -a "${HOSTS}"

traefik-hosts:
	echo "127.0.0.1  ${LOCAL_HOSTNAME_TRAEFIK}"
	grep -q "127.0.0.1  ${LOCAL_HOSTNAME_TRAEFIK}" "${HOSTS}" || echo '127.0.0.1  ${LOCAL_HOSTNAME_TRAEFIK}' | sudo tee -a "${HOSTS}"

hosts:
	make \
		main-hosts \
		pma-hosts \
		traefik-hosts

###

network:
	docker network inspect traefik_net >/dev/null 2>&1 || docker network create traefik_net

build:
	docker compose build

up: network
	docker compose up -d

restart:
	docker compose restart

down:
	docker compose down

down-v:
	- docker compose down -v --rmi local

clean-build-cache:
	- yes | docker builder prune -a

clean: clean-build-cache
	- docker compose down --rmi local

###

lde: hosts network build up
	echo "Local Docker Environment installed"

re:
	docker compose restart app-main
	docker compose restart app-bot

remain:
	docker compose restart app-main

rebot: 
	docker compose restart app-bot