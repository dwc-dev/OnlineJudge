FROM docker.io/pgsty/minio:RELEASE.2026-03-25T00-00-00Z

COPY deploy/minio/entrypoint.sh /entrypoint.sh
COPY deploy/minio/default_avatar.jpg /tmp/default_avatar.jpg

RUN chmod +x /entrypoint.sh

EXPOSE 9000 9001

ENTRYPOINT ["/entrypoint.sh"]