const Demo = artifacts.require("HLBICO");
const BigNumber = require("bignumber.js");

module.exports = function (deployer) {
  var initialRateReceived = 5 * 1000;
  var walletReceived = "0x0"; // To change to HLB's wallet
  var tokenReceived = "0x1"; // To change to LBC's contract address
  var openingTimeReceived = new Date().getTime();
  var closingTimeReceived = openingTimeReceived + 5 * 60000; // 5 minutes closing time
  var capReceived = new BigNumber("6000000000000000000000"); // 6Million for eth = 1000€
  var goalReceived = new BigNumber("500000000000000000000"); // 500K for eth = 1000€

  deployer.deploy(HLBICO, initialRateReceived,
    walletReceived,
    tokenReceived,
    openingTimeReceived,
    closingTimeReceived,
    capReceived,
    goalReceived);
};
