const { assert } = require("chai");

const LBCToken = artifacts.require("LBCToken")
const HLBICO = artifacts.require("HLBICO")

contract("HLBICO", accounts => {
    it("should initialize", async () => {
        let ico = await HLBICO.deployed();
        let lbc = await LBCToken.deployed();

        await lbc.init(
            ico.address,
            accounts[0],
            accounts[0]
        );

        await ico.init(
            accounts[0],
            accounts[0]
        );

        await ico.changeToken(lbc.address);
    });

    it("should whitelist", async () => {
        let ico = await HLBICO.deployed();
        let lbc = await LBCToken.deployed();

        await ico.addWhitelisted(accounts[0]);

        let isWhitelisted = await ico.isWhitelisted(accounts[0]);

        assert.equal(isWhitelisted, true, `Account ${accounts[0]} wasn't whitelisted properly`);
    });

    it("should be able to invest", async () => {
        let ico = await HLBICO.deployed();
        let lbc = await LBCToken.deployed();

        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("1", "ether"), from: accounts[0]})

        let accountBalance = await ico.balanceOf(accounts[0]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("4750", "ether").toString(), "Account's balance should be 4750 LBC");
    });
        
});