FROM arm32v6/golang:1.12-alpine as builder
# Force the go compiler to use modules
ENV GO111MODULE=on
# Create the user and group files to run unprivileged 
RUN mkdir /user 
RUN echo 'user:x:65534:65534:users:/:' > /user/passwd
RUN echo 'users:x:65534:' > /user/group
RUN apk update && apk add --no-cache git ca-certificates tzdata 
RUN mkdir /build 
COPY ./go.mod ./go.sum /build/
WORKDIR /build 
RUN go mod download
# Import the code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o motion-tracker-publisher .


# Minimal image, copy executable and system settings from build
FROM arm32v6/alpine AS final
LABEL author="feffe"
# Import the time zone files
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# Import the user and group files
COPY --from=builder /user/group /user/passwd /etc/
# Import binary requirements
COPY --from=builder /bin/uname /bin/uname
# Import the compiled go executable
COPY --from=builder /build/motion-tracker-publisher /
WORKDIR /
# Run as unpriveleged
#USER user:users
ENTRYPOINT ["/motion-tracker-publisher"]
# expose port
EXPOSE 8080
