Listar Conectores
curl -X GET http://10.10.10.1:8083/connectors/ | jq
Filtrar Conectores
curl -X GET http://10.10.10.1:8083/connectors/ | grep -i
Status do conector
curl -X GET http://10.10.10.1:8083/connectors/connector-name/status | jq
Exibir config do conector
curl -X GET http://localhost:8083/connectors/conn-in-configuracao-grupo/config | jq
Deletar Conector
curl -X DELETE http://10.10.10.1:8083/connectors/conn-in-excecao-ajuste | jq
Pausar Conector
curl -X PUT http://10.10.10.1:8083/connectors/<NOME>/pause | jq
Implantar Conector:
curl -s -H "Content-Type: application/json" -X POST -d @<JSON>>.json  http://10.10.10.1:8083/connectors/ | jq .