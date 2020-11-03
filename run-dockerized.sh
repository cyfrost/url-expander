port=5000
prg="local_server.go"
build_out="url-expander"
container_tag="url-expander"

rm -rf "$build_out"
go build -ldflags "-linkmode external -extldflags -static" -o "$build_out" "$prg" || exit 1
docker stop $container_tag
docker rm $container_tag
docker build --no-cache -t "$container_tag:1.0" .
docker run --publish $port:$port --name $container_tag $container_tag:1.0
