FROM locustio/locust

ENV LOCUST_MASTER_HOST=localhost
ENV LOCUST_MASTER_PORT=8080
# ENV LOCUSTFILE_PATH=/mnt/locust/locustfile.py

COPY locustfiles ./locustfiles
# COPY locust.conf ./locust.conf

EXPOSE 8089

# ENTRYPOINT locust --headless --users 10 --spawn-rate 1 -H http://localhost:8080