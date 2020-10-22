# build environment for Python

# final stage
FROM python:latest
WORKDIR /python
COPY . .
RUN python -m pip install --upgrade pip && \
    python -m pip install grpcio && \
    python -m pip install grpcio-tools && \
    python -m grpc_tools.protoc  --python_out=. --grpc_python_out=. *.proto -I .
ENTRYPOINT ["python", "main.py"]