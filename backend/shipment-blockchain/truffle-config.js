module.exports = {
  networks: {
      development: {
          host: "127.0.0.1", // локальний хост
          port: 7545,        // порт Ganache
          network_id: "*"    // будь-який network_id
      }
  },
  compilers: {
      solc: {
          version: "0.8.0" // версія компілятора Solidity
      }
  }
};
