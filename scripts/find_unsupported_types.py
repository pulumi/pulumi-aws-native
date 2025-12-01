#!/usr/bin/env python3
"""
List CloudFormation resources we have schema files for but have not yet marked
as supported in provider/cmd/pulumi-gen-aws-native/supported-types.txt.
"""

from __future__ import annotations

import json
import sys
from pathlib import Path


def main() -> int:
    repo_root = Path(__file__).resolve().parents[1]
    supported_path = (
        repo_root / "provider" / "cmd" / "pulumi-gen-aws-native" / "supported-types.txt"
    )
    schema_dir = repo_root / "aws-cloudformation-schema"

    try:
        supported_types = {
            line.strip()
            for line in supported_path.read_text().splitlines()
            if line.strip()
            and not line.strip().startswith("#")
            and line.strip().startswith("AWS::")
        }
    except FileNotFoundError:
        print(f"Unable to locate supported types file at {supported_path}", file=sys.stderr)
        return 1

    schema_type_names: set[str] = set()
    unsupported_types: set[str] = set()
    missing_type_name_files: list[Path] = []

    if not schema_dir.is_dir():
        print(f"Schema directory {schema_dir} does not exist", file=sys.stderr)
        return 1

    for schema_file in sorted(schema_dir.glob("*.json")):
        try:
            data = json.loads(schema_file.read_text())
        except json.JSONDecodeError as err:
            print(f"Failed to parse {schema_file}: {err}", file=sys.stderr)
            return 1

        type_name = data.get("typeName")
        if not type_name:
            missing_type_name_files.append(schema_file)
            continue

        schema_type_names.add(type_name)

        if type_name not in supported_types:
            unsupported_types.add(type_name)

    if missing_type_name_files:
        print("Warning: the following schema files do not define typeName:", file=sys.stderr)
        for path in missing_type_name_files:
            print(f"  - {path.relative_to(repo_root)}", file=sys.stderr)

    print(
        "Summary: supported types: {supported} | CFN resources: {cfn} | unsupported: {unsupported}".format(
            supported=len(supported_types),
            cfn=len(schema_type_names),
            unsupported=len(unsupported_types),
        )
    )

    if unsupported_types:
        print(f"Unsupported schema types ({len(unsupported_types)}):")
        for type_name in sorted(unsupported_types):
            print(type_name)
    else:
        print("All schema types are represented in supported-types.txt.")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
