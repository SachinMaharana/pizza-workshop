# Pizza Workshop v0.1

WIP

# Steps

```
docker pull sachinnicky/pizza-server:v0.2
```

```
docker run -p 5000:5000 -d sachinnicky/pizza-server:v0.2
```

# Examples

In Postman or other client.

localhost:5000/buy_pizza?pizzaType=veg&contact=sachin.nicky@gmail.com&name=sachin&contactType=email

localhost:5000/get_status?name=sachin

# Deployment

In Kubernetes.

1. clone it in cluster
2. kubectl apply -f manifests/web-lb.yaml
3. kubectl apply -f manifests/web-deploy.yaml

get lb point: k get svc | awk 'NR == 3 {print \$4}'
//unstable (need to debug)

http://35.184.34.164/buy_pizza?pizzaType=veg&contact=sachin.nicky@gmail.com&name=sachin&contactType=email

http://35.184.34.164/get_status?name=sachin
