# Como Rodar
- Criar o arquivo `repos.txt` nessa pasta, contendo em cada linha o nome e a URL do repositório da aluna no formato `Rafaela Kreusch | https://github.com/rafack/presenca-digacademy-rafa`
- Criar um token do GitHub
- Executar o script de sua preferência:

```shell
GITHUB_TOKEN=ghp_xxxSeuTokenAqui AULA=1 go run presenca.go
```

```shell
GITHUB_TOKEN=seu_token AULA=1 node presenca.js
```

Exemplo de output:
```
Fulaninha Silva         0
Rafaela Kreusch         1
```

> 💡 Indico colocar as alunas no arquivo repos.txt na mesma ordem em que elas aparecem na chamada (alfabética)

- Copiar o output (remover o nome e pegar só o número - 0 ou 1 -, deixei inicialmente para facilitar a conferência) e colar na chamada
