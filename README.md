# Introduction

*niceshoes* is a tool to import cobbler systems described as JSON into cobbler. 

# Testing

This project builds a docker container to test the tool. 

## Steps

1. On your localhost, download the Fedora 37 ISO

```
cd iso
wget http://mirrors.sonic.net/centos/7.9.2009/isos/x86_64/CentOS-7-x86_64-Minimal-2009.iso
```

2. Mount the ISO

```
sudo mkdir -p /mnt/iso

sudo mount -o loop iso/CentOS-7-x86_64-Minimal-2009.iso /mnt/iso
```

3. Run the build

```
make docker
```

4. Run the container

```
make run-docker
```

5. Exec into the container

```
docker exec -ti niceshoes bash
```

6. Create the distro

```
cobbler import --path=/mnt --name=centos79
```

7. Import the systems

```
/usr/local/bin/niceshoes --file samples/system.json
```

# Usage

```
# /usr/local/bin/niceshoes --file samples/system.json
Successfully imported test-system
Successfully imported test-system2
Successfully imported test-system3
```

# References

https://cobbler.readthedocs.io/en/latest/cobbler.html