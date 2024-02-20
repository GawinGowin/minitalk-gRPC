# Copyright 2015 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""The Python implementation of the GRPC helloworld.Greeter client."""

from __future__ import print_function

import logging
import sys

import grpc
import minitalk_pb2 # 生成された要求クラスと応答クラスが含まれる
import minitalk_pb2_grpc # 生成されたサービスクラスが含まれる

def run(msg:str=""):
  # NOTE(gRPC Python Team): .close() is possible on a channel and should be
  # used in circumstances in which the with statement does not fit the needs
  # of the code.
  print("Will try to greet world ...")
  with grpc.insecure_channel("localhost:50051") as channel:
    stub = minitalk_pb2_grpc.MinitalkStub(channel)
    response = stub.SendMsg(minitalk_pb2.MsgRequest(request=msg))
  print("Greeter client received: " + response.reply)

if __name__ == "__main__":
  logging.basicConfig()
  args = sys.argv
  run(args[1] if len(args) == 2 else args[0])
