#!/bin/bash

HOST="192.168.0.205" #ip addr of storage server

ansible-playbook -i "$HOST," getresults.yaml