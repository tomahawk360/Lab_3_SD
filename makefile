docker-broker:
        docker build . -t go-containerized-broker:latest --build-arg var=broker
        docker run -e var=broker -p 8080:8080 --name broker go-containerized-broker:latest

docker-vanguardia:
        docker build . -t go-containerized-vanguardia:latest --build-arg var=vanguardia
        docker run -e var=vanguardia -p 8020:8020 --name vanguardia go-containerized-vanguardia:latest

docker-f1:
        docker build . -t go-containerized-f1:latest --build-arg var=f1
        docker run -e var=f1 -p 50051:50051 --name f1 go-containerized-f1:latest

docker-f2:
        docker build . -t go-containerized-f2:latest --build-arg var=f2
        docker run -e var=f2 -p 50052:50052 --name f2 go-containerized-f2:latest

docker-f3:
        docker build . -t go-containerized-f3:latest --build-arg var=f3
        docker run -e var=f3 -p 50053:50053 --name f3 go-containerized-f3:latest

docker-i1:
        docker build . -t go-containerized-i1:latest --build-arg var=i1
        docker run -e var=i1 -p 8030:8030 --name i1 go-containerized-i1:latest

docker-i2:
        docker build . -t go-containerized-i2:latest --build-arg var=i2
        docker run -e var=i2 -p 8035:8035 --name i2 go-containerized-i2:latest