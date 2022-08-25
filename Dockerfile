FROM python:3-alpine

RUN adduser -D -s /bin/ash -u 1000 -g 1000 py
USER py

COPY get_data.py /get_data.py
COPY requirements.txt /requirements.txt
RUN pip install -r requirements.txt


CMD ["python", "/get_data.py"]
