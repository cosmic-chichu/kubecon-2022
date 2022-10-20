# kubecon-2022
Numaflow kubecon demo


### build and push images
- docker login --username <dockerhubusername> --email <email>
- podman login docker.io
- execute build.sh
### setup container env with podman and k3d k8s cluster
- brew upgrade
- brew upgrade podman or brew install podman
- brew install k3d
- podman machine rm (for removing old machine image)
- podman machine init --cpus 8 --memory 10000 --disk-size 60
- sudo /usr/local/Cellar/podman/4.2.1/bin/podman-mac-helper install
- podman machine set --rootful
- podman machine start
- k3d cluster create

### setup numaflow controller and server
- `kubectl create ns numaflow-system`
- `kubectl apply -n numaflow-system -f https://raw.githubusercontent.com/numaproj/numaflow/stable/config/install.yaml`
- `kubectl apply -n numaflow-system -f https://raw.githubusercontent.com/numaproj/numaflow/stable/examples/0-isbsvc-jetstream.yaml`

### start pipeline and export ports
- `kubectl apply -n numaflow-system -f go-numa-http-pipeline.yaml`
- `kubectl -n numaflow-system port-forward go-numa-cartoonizer-input-0-xxxx 8444:8443`
- `kubectl -n numaflow-system port-forward go-numa-cartoonizer-go-numa-stream-0-xxxx 9898:9898`
- `kubectl -n numaflow-system port-forward deployment/numaflow-server 8443:8443`

### run numa-webserver
- open localhost:8080 to see live streaming