1 Untar the package into your $GOPATH. 

2 You can run tests with: make test

3 If you want to build the api server just run make build

4 To simply run the API server:  make run

5 If some dependencies are missing run: make deps 

6 The dataset file location is the root of the project: go-graph-example/dataset.
You can set up the DATASETFILE env variable if you want to move the dataset
file in another place:

    cp dataset /etc
    export DATASETFILE=/etc/dataset

7 To use the API you can run the following curl command by default the server is listening on port 8080:

	curl -XGET -v http://127.0.0.1:8080/distance/34/23
	curl -XGET -v http://127.0.0.1:8080/friends/4017/3980

If you want to run the API server on another port you just need to setup the SNETPORT env variable with a valid port

	export SNETPORT=8888
