const Demo = artifacts.require("LBCToken");

module.exports = function (deployer) {
  deployer.deploy(LBCToken, "Luck Ball Coin", "LBC");
};
