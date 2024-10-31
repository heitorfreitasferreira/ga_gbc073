import json

import os
import pandas as pd
import matplotlib.pyplot as plt
import numpy as np


def plot_stats(stats_path):
    modification_name = stats_path.split("/")[-1]

    # Read all csv files inside stats_path
    items = os.listdir(stats_path)
    # Remove directories
    items = [item for item in items if os.path.isfile(os.path.join(stats_path, item))]
    if len(items) % 2 != 0:
        items = items[:-1]
    # Sort files by name
    items = sorted(items)

    for ga_file, pso_file in zip(items[::2], items[1::2]):
        ga_csv = os.path.join(stats_path, ga_file)
        pso_csv = os.path.join(stats_path, pso_file)

        instance_name = pso_file.split("_")[0]
        # Create {stats_path}/images folder
        img_out_path = f"./images_h/{modification_name}"
        if not os.path.exists(img_out_path):
            os.makedirs(img_out_path)
        img_out_path = f"{img_out_path}/{instance_name}.png"

        instances_json_path = "./benchmark/instances.json"

        with open(instances_json_path, "r") as f:
            benchmarks = json.load(f)

        benchmark = next(
            (item for item in benchmarks if item["name"] == instance_name), None
        )
        if not benchmark:
            print(f"Instância {instance_name} não encontrada no benchmark.")
            os._exit(1)

        optimun = benchmark.get("optimum")
        bounds = benchmark.get("bounds")
        upper_bound = bounds.get("upper") if bounds else None
        lower_bound = bounds.get("lower") if bounds else None

        pso_df = pd.read_csv(pso_csv)
        ga_df = pd.read_csv(ga_csv)

        y_column = "makespan"
        y = np.concatenate((pso_df[y_column], ga_df[y_column]))

        vert_line = len(pso_df)

        plt.title(f"{instance_name} - {modification_name}")
        plt.plot(y, label=y_column)
        plt.axvline(x=vert_line, color="r", linestyle="--", label="Troca de algoritmo")
        min_fitness = y.min()
        max_fitness = y.max()
        plt.figtext(0.1, 0.01, f"Min (best): {min_fitness}", ha="left", fontsize=10)
        plt.figtext(0.9, 0.01, f"Max (worst): {max_fitness}", ha="right", fontsize=10)
        if optimun:
            plt.axhline(y=optimun, color="g", linestyle="--", label="Ótimo global")
        if upper_bound:
            plt.axhline(
                y=upper_bound, color="y", linestyle="--", label="Limite superior"
            )
        if lower_bound:
            plt.axhline(
                y=lower_bound, color="y", linestyle="--", label="Limite inferior"
            )

        plt.legend()
        plt.savefig(img_out_path)
        plt.close()


def plot_all_stats():
    stats_folder = "./benchmark/stats"

    # List folder in stats folder
    for item in os.listdir(stats_folder):
        item_path = os.path.join(stats_folder, item)
        if os.path.isdir(item_path):
            plot_stats(item_path)


if __name__ == "__main__":
    plot_all_stats()
