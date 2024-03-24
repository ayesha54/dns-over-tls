
## Getting started

This program is writen in Golang and depend on the [miekg/dns](https://github.com/miekg/dns) library. Miekg/dns library is used for great projects as coredns.

Also a Docker image is available on [DockerHub](https://hub.docker.com/repository/docker/ayesha306/dns-over-tls-proxy)

**UDP**

    docker run -it -p 8053:53/UDP ayesha306/dns-over-tls-proxy:latest /bin/app udp
  
  To test `dig +short  google.com @localhost -p 8053`

**TCP**

    docker run -it -p 8053:53 ayesha306/dns-over-tls-proxy:latest /bin/app tcp

to test `dig +short +tcp google.com @localhost -p 8053`
