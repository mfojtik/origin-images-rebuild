# origin-images-rebuild

This tool helps rebuilding OpenShift images from locally built [Origin]() binaries.
The final image should be squashed, so the resulting image should not grow in size
every time you rebuilt.

## Usage

To install:

```
go get -u github.com/mfojtik/origin-images-rebuild
go install github.com/mfojtik/origin-images-rebuild
```

Your current working directory must be `$GOPATH/src/github.com/openshift/origin`.
First build the Origin using `make build`. Then to rebuild the images just execute:

```
origin-images-rebuild
```


## License

This tool is licensed under the Apache License, Version 2.0.