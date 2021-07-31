# BioretentionGardeningApp

## Required Dependencies
- git
- docker
- docker-compose

You must also allow inbound traffic on ports 80 for http connections.
This can be done by either manually adding a rule to your system's iptables, or by using a tool such as `ufw` for Ubuntu and `firewalld` for CentOS.

### Installing dependencies:
1. Install `git` using your operating system's package manager
ie: 
```
sudo apt-get install git
```

2. Install `docker` using the instructions for your operating system found at: https://docs.docker.com/engine/install/

3. Install `docker-compose` using the instructions for your operating system found at: https://docs.docker.com/compose/install/

## Running the Application
First you'll want to update the passwords and usernames found in the `.env` file. 
- `SERVER_PORT` is the port that the server will listen to within the docker network. By default this value is mapped to external port 80, but it is made available to make it easier to configure other services such as nginX.
- `MYSQL_ROOT_PASSWORD` is the root password for the newly created database.
- `MYSQL_DATABASE` is the name of the database.
- `MYSQL_USER` and `MYSQL_PASSWORD` are the username and password of the account that the server will use to interact with the database.

Once the dependencies are installed, to run the application simply clone down this git repository, navigate to the repository's root directory and run:
```
sudo docker-compose up --build -d
```
