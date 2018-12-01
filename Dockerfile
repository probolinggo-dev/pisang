
################# Build Go Project ###############
FROM golang:alpine AS builder 

#   install git
RUN apk update && apk add --no-cache git

# copy project
COPY . /app/

#   set working dir
WORKDIR /app/

# instal dependencies
RUN go get -d -v

# build a binary

RUN go build -o /bin/app

########### Create new image for the  binary #######
FROM scratch

COPY --from=builder /bin/app /bin/app
COPY ./config.json /bin/config.json

ENTRYPOINT [ "/bin/app" ]
