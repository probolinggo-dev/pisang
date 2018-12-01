# pisang
Fastest way to Create RESTFull API with Golang. 

## Installation #1

1. Clone
    
    ```
    git clone git@github.com:probolinggo-dev/pisang.git
    ```

2. Install dependencies
    ```
    cd $GOPATH/src/pisang
    go get -d -v
    ```
3. Settings
    ```
    cp config.example.json config.json
    ```
    Edit config.json!
4. Build or run it!
    ```
    go run *.go
    # or
    go build *.go -o app && ./app
    ```

## Installation #2 (Dockerize)

1. You have install docker in your machine
2. Goto directory and build the image
    ```
    cd $GOPATH/src/pisang
    docker build -t tunas:pisang .
    ```
3. Gotcha!