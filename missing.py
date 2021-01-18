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
    implemented returns the read and write routes the permission-service has
    implemented.
    """

    routes = requests.get(PERMISSION_URL).json()["healthinfo"]["routes"]
    return set(routes["read"]), set(routes["write"])


if __name__ == "__main__":
    printed = False

    try:
        implemented_read, impelemented_write = implemented()
    except requests.exceptions.ConnectionError:
        print("Can not connect to the permission service. Run:\n\n\tgo build ./cmd/permission && ./permission\n\n")
        sys.exit(2)

    try:
        actions = actions()
    except requests.exceptions.ConnectionError:
        print("Can not connect to the backend. Go to the backend repo and Run:\n\n\tmake run-prod\n\n")
        sys.exit(2)

    collections = collections()

    missing_write = actions - impelemented_write
    if missing_write:
        printed = True
        print("Missing write:")
        for mw in sorted(missing_write):
            print(f"* {mw}")

    unknown_write = impelemented_write - actions
    if unknown_write:
        if printed:
            print()
        printed = True
        print("Unknown actions:")
        for mw in sorted(unknown_write):
            print(f"* {mw}")

    missing_read = collections - implemented_read
    if missing_read:
        if printed:
            print()
        printed = True

        print("Missing read:")
        for mr in sorted(missing_read):
            print(f"* {mr}")

    unknown_read = implemented_read - collections
    if unknown_read:
        if printed:
            print()

        print("Unknown Collections:")
        for mr in sorted(unknown_read):
            print(f"* {mr}")

    if missing_write or missing_read:
        sys.exit(1)
