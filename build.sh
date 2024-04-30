#!/bin/bash
eval $(minikube docker-env) 
docker build -t dbms_task-backend . && docker run --name dbms_task-backend -p 8000:8000 -d dbms_task-backend
docker build -t dbms_task-frontend ./frontend && docker run --name dbms_task-frontend -p 3000:3000 -d dbms_task-frontend
