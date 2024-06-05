#!/bin/sh

if command -v air > /dev/null; then
    air;
else
    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice;
    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then
        go install github.com/cosmtrek/air@latest;
        air;
    else
        echo "You chose not to install air. Exiting...";
        exit 1;
    fi;
fi
