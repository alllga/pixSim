# PixSim

- Projeto feito em conjunto com o [Imersão FullCycle][linkF3]

- O projeto deverá:

  - Simular o funcionamento do PIX
  - Realizar simulações de transferencias entre chaves
  - Transação não pode ser perdida mesmo que o codePix esteja desligado
  - Um banco pode mandar uma transação para outro banco contanto que tenha uma chave válida
  - Uma transação não pode ser perdida mesmo que o banco esteja desligado

- A mesma aplicação ira simular vários bancos, somente com mudança entre cores, nome e código.

- Ferramentas:
  - Nest.js no backend
  - Next.js no frontend
  - Docker
  - Kubernetes
  - Golang
  - gRPC
  - Apache Kafka

## Rodando o Projeto Localmente

- Para que você consiga rodar o projeto na sua máquina, é necessário:
  - 1.  rodar um `docker compose up` na pasta apacheKafka e pixSim -
  - 2.  É necessário ter `127.0.0.1 host.docker.internal` como uma opção de comunicação para que os containers possam se comunicar entre si. No linux isso fica em /etc/hosts e para Windows o caminho é C:\Windows\system32\drivers\etc\hosts
  - 3. Para iniciar o servidor do pixSim, após o container estar ligado, `docker exec -it pixsim_app bash` e então `go run cmd/codepix/main.go`

[linkF3]: https://github.com/codeedu/imersao-fullstack-fullcycle
