#!/bin/bash

contract_name=$1
contract_sanitized=$(echo "$contract_name" | rev | cut -d / -f1 | rev |  cut -d . -f1)
curdir=$(pwd)

abigen --sol="$contract_name" --pkg="$contract_sanitized" --out="$curdir"/abi/golang/"$contract_sanitized".go
