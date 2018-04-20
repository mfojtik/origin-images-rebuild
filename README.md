# origin-images-rebuild

This tool helps rebuilding OpenShift images from locally built [Origin]() binaries.
The final image should be squashed, so the resulting image should not grow in size
every time you rebuilt.

## Usage

To install:

```
go get -u github.com/mfojtik/origin-images-rebuild
```

Your current working directory must be `$GOPATH/src/github.com/openshift/origin`.
First build the Origin using `make build`. Then to rebuild the images just execute:

```
origin-images-rebuild
```

## Options

* `--all` - by default this tool just build minimal set of images needed to get cluster up working.
  If you need to rebuild all images (unlikely) you can use this option.
* `--tag` - the default tag is `latest`. You can use this with `oc cluster up --tag=latest`. If you want to use
  fancier tag, use this.

## Example output

```
$ origin-images-rebuild
INFO[0000] Building "openshift/origin-sti-builder:latest" ...
INFO[0000] FROM openshift/origin-sti-builder:latest as 0
INFO[0000] ADD _output/local/bin/linux/amd64/openshift /usr/bin/openshift
INFO[0000] Committing changes to openshift/origin-sti-builder:latest ...
INFO[0000] Done
INFO[0000] Building "openshift/origin-docker-builder:latest" ...
INFO[0000] FROM openshift/origin-docker-builder:latest as 0
INFO[0000] ADD _output/local/bin/linux/amd64/openshift /usr/bin/openshift
INFO[0000] Committing changes to openshift/origin-docker-builder:latest ...
INFO[0000] Done
INFO[0000] Building "openshift/origin-haproxy-router:latest" ...
INFO[0000] FROM openshift/origin-haproxy-router:latest as 0
INFO[0000] ADD _output/local/bin/linux/amd64/openshift /usr/bin/openshift
INFO[0000] Committing changes to openshift/origin-haproxy-router:latest ...
INFO[0000] Done
INFO[0000] Building "openshift/origin-template-service-broker:latest" ...
INFO[0000] FROM openshift/origin-template-service-broker:latest as 0
INFO[0000] ADD _output/local/bin/linux/amd64/template-service-broker /usr/bin/template-service-broker
INFO[0000] Committing changes to openshift/origin-template-service-broker:latest ...
INFO[0000] Done
INFO[0000] Building "openshift/origin-control-plane:latest" ...
INFO[0000] FROM openshift/origin-control-plane:latest as 0
INFO[0000] ADD _output/local/bin/linux/amd64/openshift /usr/bin/openshift
INFO[0000] ADD _output/local/bin/linux/amd64/oc /usr/bin/oc
INFO[0000] ADD _output/local/bin/linux/amd64/hypershift /usr/bin/hypershift
INFO[0000] ADD _output/local/bin/linux/amd64/hyperkube /usr/bin/hyperkube
INFO[0000] Committing changes to openshift/origin-control-plane:latest ...
INFO[0000] Done
INFO[0000] Building "openshift/origin-deployer:latest" ...
INFO[0000] FROM openshift/origin-deployer:latest as 0
INFO[0000] ADD _output/local/bin/linux/amd64/openshift /usr/bin/openshift
INFO[0000] Committing changes to openshift/origin-deployer:latest ...
INFO[0000] Done
```

## License

This tool is licensed under the Apache License, Version 2.0.