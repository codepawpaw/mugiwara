FROM mysql:5.7

# Add a database
# ENV MYSQL_DATABASE company
ENV MYSQL_USER admin
ENV MYSQL_PASSWORD 12345
ENV MYSQL_ROOT_PASSWORD 12345
ENV MYSQL_DATABASE church-rundown

# Add the content of the sql-scripts/ directory to your image
# All scripts in docker-entrypoint-initdb.d/ are automatically
# executed during container startup
COPY ./sql-scripts/ /docker-entrypoint-initdb.d/


