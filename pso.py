import json
import pandas as pd
import matplotlib.pyplot as plt
import sys
import os


def plot_stats(csv_file, json_file, output):
    # Carregar os dados do CSV
    df = pd.read_csv(csv_file)
    # Carregar os benchmarks do arquivo JSON
    with open(json_file, "r") as f:
        benchmarks = json.load(f)

    # Obter o nome da instância a partir do nome do arquivo CSV

    instance_info = csv_file.split("/")[-1].split(".")[0].split("_")
    instance_name = instance_info[0]
    instance_model = instance_info[1] if len(instance_info) > 1 else "og"

    # Procurar a instância no arquivo JSON
    benchmark = next(
        (item for item in benchmarks if item["name"] == instance_name), None
    )

    if not benchmark:
        print(f"Instância {instance_name} não encontrada no benchmark.")
        return

    # Plotar o gráfico de fitness
    plt.figure(figsize=(10, 6))
    plt.plot(df["iteration"], df["makespan"], label="Best makespan", marker=",")

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

    min_fitness = df["makespan"].min()
    max_fitness = df["makespan"].max()
    plt.figtext(0.1, 0.01, f"Min (makespan): {min_fitness}", ha="left", fontsize=10)
    plt.figtext(0.9, 0.01, f"Max (makespan): {max_fitness}", ha="right", fontsize=10)

    plt.title(f"Melhor partícula por iteração ({instance_name})")
    plt.xlabel("Iteração")
    plt.ylabel("Makespan")
    plt.legend()
    plt.grid(True)

    plt.tight_layout()

    if output == "save":
        images_path = "./images_pso"
        # Criar dir images se não existir
        if not os.path.exists(images_path):
            os.makedirs(images_path)
        plt.savefig(f"{images_path}/{instance_name}_{instance_model}.png")
    else:
        plt.show()


if __name__ == "__main__":
    if len(sys.argv) < 3 or len(sys.argv) > 4:
        print("Uso: python script.py <csv_dir> <json_file> [show|save]")
    else:
        csv_dir = sys.argv[1]
        json_file = sys.argv[2]
        output = sys.argv[3] if len(sys.argv) == 4 else "show"

        # Lista arquivos CSV no diretório e chama plot_fitness para cada um
        dir = os.listdir(csv_dir)
        for file in dir:
            if file.endswith(".csv"):
                csv_file = os.path.join(csv_dir, file)
                plot_stats(csv_file, json_file, output)
