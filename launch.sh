#!/bin/bash
kubectl port-forward service/funny-jokes-frontend-service 3000:3000 & \
#kubectl port-forward service/adminer-service 8080:8080 & \
kubectl port-forward service/funnyjokes-backend-service 8000:8000