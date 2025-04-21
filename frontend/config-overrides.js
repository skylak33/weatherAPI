module.exports = function override(config, env) {
    if (config.devServer) {
      config.devServer.allowedHosts = ["localhost"]; // Разрешить все хосты
    }
    return config;
  };