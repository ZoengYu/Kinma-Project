# syntax=docker/dockerfile:1
FROM python:3
LABEL maintainer Harrywang
ENV PYTHONUNBUFFERED=1
WORKDIR .
COPY requirements.txt /code/
RUN pip install -r requirements.txt
COPY . /
