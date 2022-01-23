#!/bin/sh
set -e

usage()
{
    echo "usage: start.sh [OPTION]"
    echo "-ut, --unit-test Run unit test"
    echo "-s, --start Start the application through docker"
    echo "-h, --help Get help for the command"
}


# Run local app for developing
start() {
  docker-compose -f docker-compose.yaml up -d
}

unittest() {
  go test ./...
}

case $1 in
    -ut | --unit-test )
      unittest ;;
    -s | --start )
      start ;;
    -h | --help )
      usage ;;
    * )
      usage
exit 1
esac
shift