docker container stop tradingwebserv
docker container rm tradingwebserv
docker run -d -p 6789:6789 --name tradingwebserv -v $PWD/cfg:/home/tradingwebserv/cfg tradingwebserv