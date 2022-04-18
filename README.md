Engine

---

Project with docker and Amazon Elastic Kubernetes Service(EKS)



## Usage

1. Clone this repo

2. Create env.yaml file:
   
   ```yaml
   port: "8000"
   database:
     mongo_timeout: 10
     mongo_db_name: db_name
     mongo_uri:
       - mongodb://mongodb:27017
       - mongodb+srv://connection
       - mongodb://connection
   
   ```

3. Edit Mongodb uri in configs/mongodb.go with your mongo_uri env.yaml index

4. Edit Collection name in repository/repository.go with your collection

5. Save and run:
   
   ```bash
   go run main.go
   ```



## Why Docker

Container on microservices architecture becomes a crucial part nowadays so here i just want to save my time with container-based ecosystem.

### Pull image from dockerhub:

multi architecture image:

```url
https://hub.docker.com/repository/docker/fadelpm/engine-multi:release-1
```

arm-based image (tested on M1 Chip macOS Monterey):

```url
https://hub.docker.com/repository/docker/fadelpm/engine-multi:release-1
```

### Build docker image:

```bash
docker buildx create --name mybuilder --use
docker buildx use mybuilder
docker buildx build --platform linux/arm/v7,linux/arm64/v8,linux/amd64 --tag engine .
```



This images uses Mongodb atlas, persistent volume for database available asap:) ...



## Kubernetes on AWS EKS

This is the best part where i feels like i need a zero downtime server , fast rollback mechanism, and easy to update my server. AWS EKS takes you a charge to for their services see:  https://aws.amazon.com/eks/pricing/

Kubernetes need a lot of personal configurations so you need to use 'aws configure' to set your aws credentials before start. On k8s.yaml file uses Loadbalancer run port binding 8000 to 80 services and use fadelpm/engine-multi:release-1 image with 1 replica.


