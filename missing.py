#!/bin/python
"""
missing.py tells if the permission service implements all required handlers and
if not, which are missing.

To work, it requires the backend and the permission-service to be running at
localhost at there default ports.

No output means, there is nothing missing.

Requires at least python 3.9.
"""

import sys
import yaml
import requests

BACKEND_URL = "http://localhost:9002/health"
PERMISSION_URL = "http://localhost:9005/internal/permission/health"
MODELS_URL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/openslides4-dev/docs/models.yml"


def actions() -> set[str]:
    """
    actions returns all actions from the backend.
    """
    return set(requests.get(BACKEND_URL).json()["healthinfo"]["actions"].keys())


def collections() -> set[str]:
    """
    collections returns the names of all collections.
    """
    data = yaml.load(requests.get(MODELS_URL).text, Loader=yaml.CLoader)
    return set(data.keys())


def implemented() -> tuple[set[str], set[str]]:
    """
    implemented returns the collection- and action-routes the permission-service
    has implemented.
    """

    routes = requests.get(PERMISSION_URL).json()["healthinfo"]["routes"]
    return set(routes["collections"]), set(routes["actions"])


if __name__ == "__main__":
    printed = False

    try:
        implemented_collections, impelemented_actions = implemented()
    except requests.exceptions.ConnectionError:
        print("Can not connect to the permission service. Run:\n\n\tgo build ./cmd/permission && ./permission\n\n")
        sys.exit(2)

    try:
        actions = actions()
    except requests.exceptions.ConnectionError:
        print("Can not connect to the backend. Go to the backend repo and Run:\n\n\tmake run-prod\n\n")
        sys.exit(2)

    collections = collections()

    missing_actions = actions - impelemented_actions
    if missing_actions:
        printed = True
        print("Missing Actions:")
        for mw in sorted(missing_actions):
            print(f"* {mw}")

    unknown_actions = impelemented_actions - actions
    if unknown_actions:
        if printed:
            print()
        printed = True
        print("Unknown Actions:")
        for mw in sorted(unknown_actions):
            print(f"* {mw}")

    missing_collections = collections - implemented_collections
    if missing_collections:
        if printed:
            print()
        printed = True

        print("Missing Collections:")
        for mr in sorted(missing_collections):
            print(f"* {mr}")

    unknown_collections = implemented_collections - collections
    if unknown_collections:
        if printed:
            print()

        print("Unknown Collections:")
        for mr in sorted(unknown_collections):
            print(f"* {mr}")

    if missing_actions or missing_collections:
        sys.exit(1)
