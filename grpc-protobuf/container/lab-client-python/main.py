import grpc
from service_pb2 import *
from service_pb2_grpc import *

if __name__ == "__main__":
  channel = grpc.insecure_channel('10.7.14.14:8080')
  stub = MergeStub(channel)
  response = stub.Do(request(a='hello', b='Python'))
  print(response.message)