FROM alpine:11
RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install -y apache2 
COPY ./index.html /var/www/html/
COPY ./requester.js /var/www/html/
EXPOSE 80
CMD ["apache2ctl", "-D", "FOREGROUND"]