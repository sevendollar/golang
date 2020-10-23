import grpc
from greet_pb2 import *
from greet_pb2_grpc import *

if __name__ == "__main__":
    channel = grpc.insecure_channel('10.168.99.118:32768')
    stub = GreetServiceStub(channel)
    response = stub.Hello(HelloRequest(PersonalInfo(first_name="Hello", last_name="Python")))
    print(response.message)
