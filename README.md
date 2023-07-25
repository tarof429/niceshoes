# Steps

1. Download the Fedora 37 ISO

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

// may need to do: cobbler system edit --name=test-system --delete-interface --interface=default
	

# References

https://cobbler.readthedocs.io/en/latest/cobbler.html