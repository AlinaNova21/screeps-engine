module.exports = config => {
  if (!config.engine) return
  config.engine.on('init', type => {
    if (type === 'processor') {
      config.engine.driver.queue.create = () => ({
        fetch() {
          return new Promise((res,rej) => {})
        }
      })
    }
  })
}