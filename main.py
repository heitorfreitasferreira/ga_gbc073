import json
import pandas as pd
import matplotlib.pyplot as plt
import sys


def plot_fitness(csv_file, json_file, output):
    # Carregar os dados do CSV
    df = pd.read_csv(csv_file)
    df = df[df["generation"] % max(1, len(df) // 100) == 0]

    # Carregar os benchmarks do arquivo JSON
    with open(json_file, "r") as f:
        benchmarks = json.load(f)

    # Obter o nome da instância a partir do nome do arquivo CSV
    instance_name = csv_file.split("/")[-1].split(".")[0]

    # Procurar a instância no arquivo JSON
    benchmark = next(
        (item for item in benchmarks if item["name"] == instance_name), None
    )

    if not benchmark:
        print(f"Instância {instance_name} não encontrada no benchmark.")
        return

    # Plotar o gráfico de fitness
    plt.figure(figsize=(10, 6))
    plt.plot(df["generation"], df["best"], label="Best", marker=",")
    plt.plot(df["generation"], df["worst"], label="Worst", marker=",")
    plt.plot(df["generation"], df["median"], label="Median", marker=",")
    plt.plot(df["generation"], df["average"], label="Average", marker=",")

    # Marcar as linhas de bounds e optimum, se existirem
    if "optimum" in benchmark and benchmark["optimum"] is not None:
        plt.axhline(y=benchmark["optimum"], color="g", linestyle="--", label="Optimum")

    if "bounds" in benchmark and benchmark["bounds"]:
        plt.axhline(
            y=benchmark["bounds"]["upper"],
            color="r",
            linestyle="--",
            label="Upper Bound",
        )
        plt.axhline(
            y=benchmark["bounds"]["lower"],
            color="b",
            linestyle="--",
            label="Lower Bound",
        )

    plt.title(f"Fitness da Instância {instance_name}")
    plt.xlabel("Generation")
    plt.ylabel("Fitness")
    plt.legend()
    plt.grid(True)

    plt.tight_layout()

    if output == "save":
        plt.savefig(f"{instance_name}.png")
    else:
        plt.show()


if __name__ == "__main__":
    if len(sys.argv) < 3 or len(sys.argv) > 4:
        print("Uso: python script.py <csv_file> <json_file> [show|save]")
    else:
        csv_file = sys.argv[1]
        json_file = sys.argv[2]
        output = sys.argv[3] if len(sys.argv) == 4 else "show"
        plot_fitness(csv_file, json_file, output)
