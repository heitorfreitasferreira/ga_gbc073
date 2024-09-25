# ga_gbc073

Repositório para implementação de um algoritmo genético para encontrar soluções para o problema de escalonamento Job Shop.

## Compilação e execução

Tenha Go 1.22.5 ou superior instalado e execute o seguinte comando no diretório do projeto:

```bash
go build -o ./ga.out
```

Então basta executar o binário gerado (nomeado como ga.out no comando anterior):

```bash
./ga.out runExp

```

O argumento runExp faz com que o algoritmo seja executado com benchmarks pré definidos e salve os resultados na pasta ./benchmark/stats.

## Visualização dos resultados

Na raíz do projeto há um arquivo main.py que pode ser executado para ler os arquivos de csv gerados em ./benchmark/stats e gerar gráficos com os dados dos fitness da população ao longo das gerações. As imagens geradas são salvas em ./benchmark/images.

Para executar o script é necessário passar o diretório dos dados csv:

```bash
python main.py ./benchmark/stats ./benchmark/instances.json save

```
