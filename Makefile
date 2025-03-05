# FASE 1
# Local enviroment with docker compose
fiap-run:
	docker-compose up -d

fiap-stop:
	docker-compose down

fiap-logs:
	docker-compose logs -f

run-kitchen-api:
	go run cmd/kitchencontrol/main.go

run-payment-api:
	go run cmd/fakepaymentservice/main.go	

# Local Dev support
test-coverage:
	go test -coverprofile=coverage.out ./...
coverage: test-coverage
	go tool cover -func=coverage.out
coverage-html: test-coverage
	go tool cover -html=coverage.out
test:
	go test -v -cover ./...

docs:
	#rm -rf docs
	swag init -g ./cmd/kitchencontrol/main.go -o ./docs

# cria os mocks para testes
mocks:
	mockgen -source=internal/domain/repository/product_repository.go -destination=internal/domain/repository/mocks/mock_product_repository.go -package=mocksrepository
	mockgen -source=internal/domain/repository/customer_repository.go -destination=internal/domain/repository/mocks/mock_customer_repository.go -package=mocksrepository
	mockgen -source=internal/domain/repository/order_repository.go -destination=internal/domain/repository/mocks/mock_order_repository.go -package=mocksrepository
	mockgen -source=internal/domain/repository/checkout_repository.go -destination=internal/domain/repository/mocks/mock_checkout_repository.go -package=mocksrepository
	mockgen -source=internal/domain/repository/kitchen_repository.go -destination=internal/domain/repository/mocks/mock_kitchen_repository.go -package=mocksrepository
	mockgen -source=internal/domain/repository/transaction_manager.go -destination=internal/domain/repository/mocks/mock_transaction_manager.go -package=mocksrepository

# FASE 2
# Local enviroment with kind (kubernetes)

# CONFIGURAÇÃO do K8S e seus recursos
# cria as imagens que usaremos no kind
create-image:
	docker build -t caiojorge/fiap-rocks .

# faz o login no docker hub para subir as imagens
login:
	docker login

# sobe as imagens no docker hub
push-image:
	docker push caiojorge/fiap-rocks

# Usei o kind para testar o projeto no k8s, sendo assim, é necessário criar o cluster
setup-cluster:
	kind delete cluster
	kind create cluster --config=k8s/cluster-config.yml	
	
# cria o configmap e sobre o arquivo de scripts de inicialização do banco
setup-configmap:
	kubectl create configmap db-init-scripts --from-file=./db-init
	kubectl get configmaps

# cria os recursos no k8s
setup-k8s:
	kubectl apply -f k8s/mysql-secret.yaml
	kubectl apply -f k8s/mysql-pvc.yaml
	kubectl apply -f k8s/mysql-deployment.yaml
	kubectl apply -f k8s/mysql-service.yaml
	kubectl apply -f k8s/adminer-deployment.yaml
	kubectl apply -f k8s/adminer-service.yaml
	kubectl apply -f k8s/server-deployment.yaml
	kubectl apply -f k8s/server-service.yaml
	kubectl apply -f k8s/server-hpa.yaml

	kubectl get pods
	kubectl get svc

# faz o setup completo
setup-all: setup-cluster setup-configmap setup-k8s
	@echo "Setup concluído!"

# acesso aos logs dos pods no k8s
log-k8s:
	kubectl logs -f $(shell kubectl get pods -l app=fiap-rocks-server -o jsonpath='{.items[0].metadata.name}')

# verifica os pods e serviços no k8s
get-k8s:
	kubectl get pods
	kubectl get svc

# deleta todos os recursos no k8s
shutdown:
	kind delete cluster


# Teste do endpoint de pagamento
# Variáveis de ambiente
COLLECTOR_ID := collector123
POS_ID := pos456
URL := http://localhost:30090/instore/orders/qr/seller/collectors/$(COLLECTOR_ID)/pos/$(POS_ID)/qrs

send-payment:
	curl -X POST "$(URL)" \
	-H "Content-Type: application/json" \
	-H "Authorization: $(TOKEN)" \
	-d '{"external_reference": "721ece5d-62c3-49a2-bc12-919a2486cefd", "title": "Compra de Produtos", "description": "Pagamento na loja física", "notification_url": "http://localhost:30080/kitchencontrol/api/v1/checkouts/confirmation/payment", "total_amount": 200.50, "items": [{"sku_number": "sku-001", "category": "eletronico", "title": "Fone de Ouvido", "description": "Fone de ouvido com cancelamento de ruído", "unit_price": 100.25, "quantity": 2, "unit_measure": "unidade", "total_amount": 200.50}], "sponsor": {"id": 1}, "cash_out": {"enabled": true, "amount": 50.00, "receiver": "receiver-001"}}'

converage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -func=coverage.out | grep -E ".*[0-9]+\.[0-9]+%" | awk '{sum+=$3; count++} END {if (count > 0) print "Cobertura Média:", sum/count"%"; else print "Nenhuma cobertura"}'
