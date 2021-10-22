# syntax=docker/dockerfile:1
FROM python:3
LABEL maintainer Harrywang
ENV PYTHONUNBUFFERED=1
WORKDIR .
COPY requirements.txt /server/
RUN pip install -r /server/requirements.txt
COPY . /
