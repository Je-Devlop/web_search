start_deps:
	docker compose up

stop_deps:
	docker compose down

logstash_log:
	docker compose log logstash

elastic_log:
	docker compose log elasticsearch

POHNY: start_deps stop_deps logstash_log