FROM python:3.8
LABEL maintainer="sujining13@keti.re.kr"

COPY . .

RUN pip3 install -r requirements.txt

EXPOSE 9010
ENTRYPOINT ["python3", "server.py"]
