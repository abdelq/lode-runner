#!/usr/bin/env python

from socket import socket, AF_INET, SOCK_STREAM
from ai import name, room

if __name__ == '__main__':
    s = socket(AF_INET, SOCK_STREAM)
    s.connect(('127.0.0.1', 443))
    print("Connected to {}:{}".format(s.getpeername()[0], s.getpeername()[1]))
