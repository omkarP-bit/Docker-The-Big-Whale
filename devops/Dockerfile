FROM python:3.9

WORKDIR /app

COPY requirements.txt /app/
RUN apt-get update && \
    apt-get install -y python3 python3-pip && \
    pip install -r requirements.txt --break-system-packages

COPY . /app/

EXPOSE 8000

CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]