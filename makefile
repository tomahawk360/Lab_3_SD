docker-broker
        docker build . -t go-containerized-broker:latest --build-arg var=broker
        docker run -e var=broker -p 8050:8050 --name broker go-containerized-broker:latest

docker-vanguardia
        docker build . -t go-containerized-vanguardia:latest --build-arg var=vanguardia
        docker run -e var=vanguardia -p 8050:8050 --name broker go-containerized-broker:vanguardia

docker-f1
        docker build . -t go-containerized-f1:latest --build-arg var=f1
        docker run -e var=f1 -p 8050:8050 --name f1 go-containerized-broker:f1

docker-f2
        docker build . -t go-containerized-f2:latest --build-arg var=f2
        docker run -e var=f2 -p 8050:8050 --name f2 go-containerized-broker:f2

docker-f3
        docker build . -t go-containerized-f3:latest --build-arg var=f3
        docker run -e var=f3 -p 8050:8050 --name f3 go-containerized-broker:f3

docker-i1
        docker build . -t go-containerized-i1:latest --build-arg var=i1
        docker run -e var=i1 -p 8050:8050 --name i1 go-containerized-broker:i1

docker-i2
        docker build . -t go-containerized-i2:latest --build-arg var=i2
        docker run -e var=i2 -p 8050:8050 --name i2 go-containerized-broker:i2