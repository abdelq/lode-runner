#!/bin/sh

echo '{"event": "kill", "data": "'"$1"'"}' | nc -w0 localhost 1337
