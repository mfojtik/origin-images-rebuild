# origin-images-rebuild

This tool allows to rebuild OpenShift Origin images using locally build binaries.
The result is an image set that can be use to run cluster up with local modifications
or to support fast development interation for OpenShift engineers.

## Usage

To install:

```
go get -u github.com/mfojtik/origin-images-rebuild
go install github.com/mfojtik/origin-images-rebuild
```

* Your current working directory must be `$GOPATH/src/github.com/openshift/origin`.
* First build the Origin using `make build`.
* Specify `--config` to point to 'default.yaml'
* Run: `origin-images-rebuild`

If you don't like the `:latest` tag, you can override it by `--tag`.

## Example:

```
[mfojtik@dev origin]$ GOPATH=$HOME/go ~/go/bin/origin-images-rebuild --config=$HOME/go/src/github.com/mfojtik/origin-images-rebuild/config/default.yaml
2018/08/23 13:40:19 Using "/home/mfojtik/go/src/github.com/mfojtik/origin-images-rebuild/config/default.yaml" as build configuration
2018/08/23 13:40:19 Building "openshift/origin-cli:latest" ...
2018/08/23 13:40:19 FROM openshift/origin-cli:latest as 0
2018/08/23 13:40:19 ADD _output/local/bin/linux/amd64/oc /usr/bin/oc
2018/08/23 13:40:19 Committing changes to openshift/origin-cli:latest ...
2018/08/23 13:40:19 Done
2018/08/23 13:40:19 Building "openshift/origin-control-plane:latest" ...
2018/08/23 13:40:19 FROM openshift/origin-control-plane:latest as 0
2018/08/23 13:40:19 ADD _output/local/bin/linux/amd64/oc /usr/bin/oc
2018/08/23 13:40:19 ADD _output/local/bin/linux/amd64/openshift /usr/bin/openshift
2018/08/23 13:40:19 Committing changes to openshift/origin-control-plane:latest ...
2018/08/23 13:40:19 Done
2018/08/23 13:40:19 Building "openshift/origin-hyperkube:latest" ...
2018/08/23 13:40:19 FROM openshift/origin-hyperkube:latest as 0
2018/08/23 13:40:19 ADD _output/local/bin/linux/amd64/hyperkube /usr/bin/hyperkube
2018/08/23 13:40:19 Committing changes to openshift/origin-hyperkube:latest ...
2018/08/23 13:40:19 Done
2018/08/23 13:40:19 Building "openshift/origin-hypershift:latest" ...
2018/08/23 13:40:19 FROM openshift/origin-hypershift:latest as 0
2018/08/23 13:40:19 ADD _output/local/bin/linux/amd64/hypershift /usr/bin/hypershift
2018/08/23 13:40:19 Committing changes to openshift/origin-hypershift:latest ...
2018/08/23 13:40:19 Done
2018/08/23 13:40:19 Building "openshift/origin-deployer:latest" ...
2018/08/23 13:40:19 FROM openshift/origin-deployer:latest as 0
2018/08/23 13:40:19 ADD _output/local/bin/linux/amd64/oc /usr/bin/oc
2018/08/23 13:40:20 ADD _output/local/bin/linux/amd64/openshift-deploy /usr/bin/openshift-deploy
2018/08/23 13:40:20 Committing changes to openshift/origin-deployer:latest ...
2018/08/23 13:40:20 Done
2018/08/23 13:40:20 Building "openshift/origin-recycler:latest" ...
2018/08/23 13:40:20 FROM openshift/origin-recycler:latest as 0
2018/08/23 13:40:20 ADD _output/local/bin/linux/amd64/oc /usr/bin/oc
2018/08/23 13:40:20 Committing changes to openshift/origin-recycler:latest ...
2018/08/23 13:40:20 Done
2018/08/23 13:40:20 Building "openshift/origin-haproxy-router:latest" ...
2018/08/23 13:40:20 FROM openshift/origin-haproxy-router:latest as 0
2018/08/23 13:40:20 ADD _output/local/bin/linux/amd64/openshift /usr/bin/openshift
2018/08/23 13:40:20 Committing changes to openshift/origin-haproxy-router:latest ...
2018/08/23 13:40:20 Done
2018/08/23 13:40:20 Building "openshift/origin-template-service-broker:latest" ...
2018/08/23 13:40:20 FROM openshift/origin-template-service-broker:latest as 0
2018/08/23 13:40:20 ADD _output/local/bin/linux/amd64/template-service-broker /usr/bin/template-service-broker
2018/08/23 13:40:20 Committing changes to openshift/origin-template-service-broker:latest ...
2018/08/23 13:40:20 Done
2018/08/23 13:40:20 Building "openshift/origin-docker-builder:latest" ...
2018/08/23 13:40:20 FROM openshift/origin-docker-builder:latest as 0
2018/08/23 13:40:20 ADD _output/local/bin/linux/amd64/oc /usr/bin/oc
2018/08/23 13:40:20 Committing changes to openshift/origin-docker-builder:latest ...
2018/08/23 13:40:20 Done
2018/08/23 13:40:20 Building "openshift/origin-keepalived-ipfailover:latest" ...
2018/08/23 13:40:20 FROM openshift/origin-keepalived-ipfailover:latest as 0
2018/08/23 13:40:20 ADD _output/local/bin/linux/amd64/openshift /usr/bin/openshift
2018/08/23 13:40:20 Committing changes to openshift/origin-keepalived-ipfailover:latest ...
2018/08/23 13:40:20 Done
```

## License

This tool is licensed under the Apache License, Version 2.0.