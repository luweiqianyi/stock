# 部署mysql
docker-compose -p stock up -d

# 如果部署文件的名字不是默认的docker-compose.yaml，比如mysql-docker-compose.yaml，则需要用-f标记来指定
# docker-compose -p stock -f ./mysql-docker-compose.yaml up -d