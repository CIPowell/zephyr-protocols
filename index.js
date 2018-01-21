const authMessages = require('./node/auth_pb');
const authService = require('./node/auth_grpc_pb');

module.exports = {
    Auth: {
        messages: authMessages,
        service: authService
    }
}