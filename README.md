

## install golang

> wget https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz


## install nvm

> curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash

## install yarn

> npm install --global yarn

> curl -o- -L https://yarnpkg.com/install.sh | bash


## install mongodb

> sudo wget -qO - https://www.mongodb.org/static/pgp/server-4.4.asc | sudo apt-key add -

> echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/4.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.4.list

> sudo apt update

> sudo apt install -y mongodb-org

> sudo curl -o /etc/init.d/mongod https://raw.githubusercontent.com/heidemn/mongod-sysv/master/mongod

> sudo chmod +x /etc/init.d/mongod 

> sudo service mongod start