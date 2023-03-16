# Monitor de serviços HTTP

Olá! 🧑‍💻

Este repositório contém um monitor de serviços HTTP feito em GO, com ele você consegue fazer o monitoramento rápido de uma lista de aplicações HTTP passadas através de um arquivo .csv. O Script devolve os servidores que estão UP e cria uma nova lista informando os servidores que estão down.

## Como usar

Para isso basta ir até a pasta referente ao seu sistema operacional dentro da pasta build e rodar o executável, você deve passar dois parâmetros ao executar o Script, o primeiro é a lista de servidores que será monitorada e o segundo é o nome da lista de servidores que será criada para armazenar os servidores que não estiverem no ar ou retornarem algum código diferente de 200. 

Exemplo: `./httpmon server-list.csv serversdown.csv`

Lembre-se de passar o caminho exato do arquivo.csv.

A lista para servidores Down não precisa necessariamente já estar criada, pois caso ela não exista o próprio Script irá criar.

## Tech

![go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)