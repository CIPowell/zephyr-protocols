// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var auth_pb = require('./auth_pb.js');

function serialize_auth_LoginRequest(arg) {
  if (!(arg instanceof auth_pb.LoginRequest)) {
    throw new Error('Expected argument of type auth.LoginRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_auth_LoginRequest(buffer_arg) {
  return auth_pb.LoginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_LoginResult(arg) {
  if (!(arg instanceof auth_pb.LoginResult)) {
    throw new Error('Expected argument of type auth.LoginResult');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_auth_LoginResult(buffer_arg) {
  return auth_pb.LoginResult.deserializeBinary(new Uint8Array(buffer_arg));
}


var AuthService = exports.AuthService = {
  login: {
    path: '/auth.Auth/Login',
    requestStream: false,
    responseStream: false,
    requestType: auth_pb.LoginRequest,
    responseType: auth_pb.LoginResult,
    requestSerialize: serialize_auth_LoginRequest,
    requestDeserialize: deserialize_auth_LoginRequest,
    responseSerialize: serialize_auth_LoginResult,
    responseDeserialize: deserialize_auth_LoginResult,
  },
};

exports.AuthClient = grpc.makeGenericClientConstructor(AuthService);
