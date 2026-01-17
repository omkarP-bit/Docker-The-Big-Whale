FROM ubuntu

WORKDIR /app

COPY requirements.txt /app/
COPY . /app/

RUN apt-get update && \
    apt-get install -y python3-pip && \
    pip3 install -r requirements.txt --break-system-packages

EXPOSE 8000

CMD ["python3", "manage.py", "runserver", "0.0.0.0:8000"]
