rootdir = $(realpath .)

build/itami: 
	bazel build //packages/itami:build

build/nioh/example_image: 
	bazel build //packages/nioh/cmd/example:oss-hoshi-example

build/docker/local:
	docker run --rm -v /tmp/bazel:/tmp/bazel -v $(rootdir):/workspace -w /workspace gcr.io/cloud-marketplace-containers/google/bazel:3.5.0 --output_user_root=/tmp/bazel build //cmd/...:all

build/docker/remote/rw:
	docker run --rm -v $(rootdir):/workspace -w /workspace gcr.io/cloud-marketplace-containers/google/bazel:3.5.0 build --remote_cache=http://replace-with-your.host:port //cmd/...:all

build/docker/remote/ro:	
	# no changes pushed back
	docker run --rm -v $(rootdir):/workspace -w /workspace gcr.io/cloud-marketplace-containers/google/bazel:3.5.0 build --remote_upload_local_results=false //cmd/...:all

test/local/build:
	# example local testing 
	cat ~/github.com/go-zaru/oss-hoshi/bazel-bin/cmd/example/example_image.executable.runfiles/__main__/cmd/example/example_image-layer.tar | docker import - example_image:v0.2.0

gofmt:
	gofmt -w -s pkgs/ cmd/ internal/

gazelle/run:
	bazel run //:gazelle

gazelle/update:
	go mod tidy -v
	bazel run //:gazelle -- update-repos -from_file=go.mo
