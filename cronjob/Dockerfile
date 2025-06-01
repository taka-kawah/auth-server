FROM ubuntu:25.10

RUN apt update
RUN apt install -y postgresql-client
RUN apt install -y cron

COPY ./task /etc/cron.d/task
COPY ./reserve_clean.sql /etc/cron.d/reserve_clean.sql
COPY ./.env.sh /etc/cron.d/.env.sh

RUN chmod 0644 /etc/cron.d/task

CMD [ "cron", "-f" ]