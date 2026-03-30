import pyetcd
import grpc
import os
from . import question_pb2
from . import question_pb2_grpc

class QuestionRpcClient:
    def __init__(self):
        address = self._get_service_address("microservice.question")
        self.channel = grpc.insecure_channel(address)
        self.stub = question_pb2_grpc.QuestionStub(self.channel)

    def _get_service_address(self, prefix):
        etcd_host = os.getenv("ETCD_HOST")
        etcd_port = int(os.getenv("ETCD_PORT"))
        etcd = pyetcd.client(host=etcd_host, port=etcd_port)
        addresses = []
        for value, _ in etcd.get_prefix(prefix):
            addresses.append(value.decode("utf-8"))
        return addresses[0]

    def get_question_info(self, id):
        response = self.stub.GetQuestionInfo(
            question_pb2.GetQuestionInfoReq(id=id, col=["title", "content", "tags"])
        )
        return {
            "title": response.question_info.title,
            "content": response.question_info.content,
            "tags": response.question_info.tags,
        }
