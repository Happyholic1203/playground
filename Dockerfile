FROM happyholic1203/devbox
MAINTAINER Yu-Cheng (Henry) Huang

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -yq mysql-server python-mysqldb && \
    pip install sqlalchemy

# Default entry point for happyholic1203/devbox
CMD ["/root/init"]
