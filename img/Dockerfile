# vim: ft=dockerfile

FROM fedora:37

RUN dnf makecache

# Install cobbler
RUN dnf install -y            \
    cobbler                   \
    supervisor                \
    wget

RUN mkdir -p /cmd 

COPY supervisord/supervisord.conf /etc/supervisord.conf
COPY supervisord/conf.d /etc/supervisord/conf.d
COPY import.sh /
#COPY samples /samples
#COPY ../niceshoes /

CMD ["/usr/bin/supervisord", "-n", "-c", "/etc/supervisord.conf"]

