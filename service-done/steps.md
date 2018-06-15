
### Build and push

go build ./server.go
sudo docker build -t lawrencegripper/go-server .
sudo docker push lawrencegripper/go-server

### Create container:

az container create -g demo-go --name go-server --image lawrencegripper/go-server --ports 8087 --dns-name-label gripdev1

### Call instance

../client http://gripdev1.northeurope.azurecontainer.io:8087/jeff