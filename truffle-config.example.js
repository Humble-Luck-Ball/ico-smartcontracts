module.exports = {
  networks: {
    development: {
      host: "172.19.80.1",
      port: 7545,
      network_id: "*"
    }
  },

  compilers: {
    solc: {
      version: "^0.7",
    }
  }
};
