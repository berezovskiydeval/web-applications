VOLUME      := notes-manager-db-data
CONTAINER   := notes-manager-db
DB_PASSWORD := n0t3sMaNaG3R
HOST_PORT   := 5436

start-db:
	docker volume create $(VOLUME)
	docker run -d --name $(CONTAINER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-p $(HOST_PORT):5432 \
		-v $(VOLUME):/var/lib/postgresql/data \
		--restart unless-stopped \
		postgres

stop-db:
	docker stop $(CONTAINER)

rm-db:
	docker rm -f $(CONTAINER) || true

logs-db:
	docker logs -f $(CONTAINER)

psql:
	docker exec -it $(CONTAINER) psql -U postgres

backup:
	docker run --rm \
	  -v $(VOLUME):/data \
	  -v $$(pwd):/backup \
	  alpine tar czvf /backup/pgdata-backup.tar.gz -C /data .

restore:
	docker run --rm \
	  -v $(VOLUME):/data \
	  -v $$(pwd):/backup \
	  alpine sh -c 'cd /data && tar xzvf /backup/pgdata-backup.tar.gz --strip 1'

.PHONY: start-db stop-db rm-db logs-db psql backup restore
