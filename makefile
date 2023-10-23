build:
	docker build -t lab1:latest .

docker-ONU:
ifeq ($(HOST),localhost)
	docker rm -f onu
	docker run -d -it --name onu --expose 50052 lab1:latest go run ONU/onu.go
else
	echo "Ejecutar SOLO en dist105"
endif

docker-continentes:
	docker rm -f regional
ifeq ($(HOST),localhost)
	docker run -d -it --rm --name regional -p 50052:50052 lab1:latest go run Regionales/Australia/regional.go
endif
ifeq ($(HOST),dist106)
	docker run -d -it --rm --name regional -p 50052:50052 lab1:latest go run Regionales/Asia/regional.go
endif
ifeq ($(HOST),dist107)
	docker run -d -it --rm --name regional -p 50052:50052 lab1:latest go run Regionales/Europa/regional.go
endif
ifeq ($(HOST),dist108)
	docker run -d -it --rm --name regional -p 50052:50052 lab1:latest go run Regionales/Latinoamerica/regional.go
endif

docker-OMS:
ifeq ($(HOST),dist106)
	docker rm -f oms
	docker run -d -it --name oms --expose 50053 lab1:latest go run OMS/oms.go
else
	echo "Ejecutar SOLO en dist106"
endif

docker-datanode:
	docker rm -f datanode
ifeq ($(HOST),dist107)
	docker run -d -it --name datanode --expose 50053 lab1:latest go run Data1/datanode.go
ifeq ($(HOST),dist108)
	docker run -d -it --name datanode --expose 50053 lab1:latest go run Data2/datanode.go
endif
else
	echo "Ejecutar SOLO en dist107 y dist108"
endif