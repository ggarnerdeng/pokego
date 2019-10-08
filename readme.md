pokmeon go demo

create database
* bash
cd db (folder)

docker build -t pokedb . 
docker run --name pokedb -d -p 5432:5432 pokedb
--

then use main.go to connect
