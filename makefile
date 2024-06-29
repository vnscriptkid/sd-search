up:
	docker compose up -d

down:
	docker compose down --remove-orphans --volumes

psql:
	docker compose exec postgres psql -U postgres -d debezium

logs_kafka:
	docker compose logs -f kafka

logs_zoo:
	docker compose logs -f zookeeper

logs_connect:
	docker compose logs -f connect

consume:
	docker compose exec kafka /kafka/bin/kafka-console-consumer.sh \
    --bootstrap-server kafka:9092 \
    --from-beginning \
    --property print.key=true \
    --topic cdc_.public.products