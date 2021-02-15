#!/bin/bash

contract_name=$1
contract_sanitized=$(echo "$contract_name" | rev | cut -d / -f1 | rev |  cut -d . -f1)
contract_lowercase=$(echo "$contract_sanitized" | awk '{print tolower($0)}')
curdir=$(pwd)

abigen --sol="$contract_name" --pkg="$contract_lowercase" --out="$curdir"/abi/golang/"$contract_lowercase".go
