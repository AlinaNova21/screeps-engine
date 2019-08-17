const common = require('@screeps/common')
const net = require('net')

const sock = net.connect(3000)
const client = new common.rpc.RpcClient(sock)

async function run() {
  await client.request('dbEnvSet', 'abc', '123')
  const ret = await client.request('dbEnvGet', 'abc')
  console.log(ret)
}
run().catch(console.error).then(() => {
  sock.end()
  sock.destroy()
})

// Object.assign(exports.env, {
//   get: resetInterceptor(rpcClient.request.bind(rpcClient, 'dbEnvGet')),
//   mget: resetInterceptor(rpcClient.request.bind(rpcClient, 'dbEnvMget')),
//   set: resetInterceptor(rpcClient.request.bind(rpcClient, 'dbEnvSet')),
//   setex: resetInterceptor(rpcClient.request.bind(rpcClient, 'dbEnvSetex')),
//   expire: resetInterceptor(rpcClient.request.bind(rpcClient, 'dbEnvExpire')),
//   ttl: resetInterceptor(rpcClient.request.bind(rpcClient, 'dbEnvTtl')),
//   del: resetInterceptor(rpcClient.request.bind(rpcClient, 'dbEnvDel')),
//   hmget: resetInterceptor(rpcClient.request.bind(rpcClient, 'dbEnvHmget')),
//   hmset: resetInterceptor(rpcClient.request.bind(rpcClient, 'dbEnvHmset')),
//   hget: resetInterceptor(rpcClient.request.bind(rpcClient, 'dbEnvHget')),
//   hset: resetInterceptor(rpcClient.request.bind(rpcClient, 'dbEnvHset'))
// });