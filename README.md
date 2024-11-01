
# Go Opportunities

Repositório criado para aplicar o conteúdo passado na live do @arthur404dev, procurando reproduzir todos os passos e evoluir a api com novas funcionalidades.

## Ferramentas utilizadas

- go
- swagger
- gin
- make

## Documentação sobre a API

#### Retorna uma oportunidade de trabalho

```http
  GET /api/v1/opening
```
| Par?metro   | Tipo       | Descri??o                           |
| :---------- | :--------- | :---------------------------------- |
| `id` | `int64` | **Obrigatório**. Informar o id da oportunidade |

#### Gravar uma oportunidade de trabalho

```http
  POST /api/v1/opening

  Body:

  {
    "role":"Engenheiro Cloud",
    "company":"RCarvalho",
    "location":"Sao Paulo",
    "remote":true,
    "link":"rcaarvalho.com.br/vagas/senior_cloud",
    "salary": 15000
}

```
#### Retorna todas as oportunidades

```http
  GET /api/v1/openings
```
#### Remove uma oportunidade

```http
  DELETE /api/v1/opening/?id={id}
```

| Par?metro   | Tipo       | Descri??o                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `int64` | **Obrigatório**. O ID para ser deletado |

#### Atualiza um cadastro

```http
  PUT /api/v1/opening

  Body:

  {
    "role":"Engenheiro Cloud",
    "company":"RCarvalho",
    "location":"Sao Paulo",
    "remote":true,
    "link":"rcaarvalho.com.br/vagas/senior_cloud",
    "salary": 25000   # Atualizamos o salário
}

```


## Referência

 - [Projeto Original](https://github.com/arthur404dev/gopportunities)
