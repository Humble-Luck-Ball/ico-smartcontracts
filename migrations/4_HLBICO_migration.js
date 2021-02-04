const HLBICO = artifacts.require("HLBICO");
const BN = require("bignumber.js");

/*
(uint256 initialRateReceived,
        address payable walletReceived,
        LBCToken tokenReceived,
        uint256 openingTimeReceived,
        uint256 closingTimeReceived,
        uint256 capReceived,
        uint256 goalReceived)
*/

module.exports = function (deployer) {

  var initDate = Math.floor(new Date().getTime() / 1000) + 5;
  var endDate = initDate + 60 * 120;

  var initialRateReceived = new BN(5 * 1000);
  var walletReceived = "0xFe771D9E6C364D5CbFb29455CfFe57cB0AC58252"; // To change to HLB's wallet
  var tokenReceived = "0x2481994096b5979bD610Ac761BcA8dc83b3b1d17"; // To change to LBC's contract address
  var openingTimeReceived = new BN(initDate);
  var closingTimeReceived = new BN(endDate); // 5 minutes closing time
  var capReceived = new BN("6000000000000000000000"); // 6Million for eth = 1000€
  var goalReceived = new BN("300000000000000000000"); // 500K for eth = 1000€

  deployer.deploy(HLBICO, initialRateReceived,
    walletReceived,
    tokenReceived,
    openingTimeReceived,
    closingTimeReceived,
    capReceived,
    goalReceived);
};
