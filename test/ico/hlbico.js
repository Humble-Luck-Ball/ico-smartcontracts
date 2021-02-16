const { assert } = require("chai");

const LBCToken = artifacts.require("LBCToken")
const HLBICO = artifacts.require("HLBICO")

contract("HLBICO", accounts => {
    it("should initialize", async () => {
        let ico = await HLBICO.deployed();
        let lbc = await LBCToken.deployed();
        let addrToWhitelist = "0xd5F6A8eDb411Bb42411f532224e0CCA9a909e710"

        await lbc.init(
            ico.address,
            accounts[0],
            accounts[0]
        );

        await ico.init(
            addrToWhitelist,
            accounts[0]
        );

        await web3.eth.sendTransaction({value: web3.utils.toWei("0.5", "ether"), from: accounts[1], to: addrToWhitelist})

        await ico.changeToken(lbc.address);
    });

    it("should whitelist", async () => {
        let ico = await HLBICO.deployed();


        let isWhitelisted = await ico.isWhitelisted(accounts[0]);

        assert.equal(isWhitelisted.toString(), "0", `Account ${accounts[0]} was already whitelisted when it shouldn't be`);

        await ico.addWhitelistedRegistered(accounts[0]);

        isWhitelisted = await ico.isWhitelisted(accounts[0]);

        assert.equal(isWhitelisted.toString(), "1", `Account ${accounts[0]} wasn't whitelisted properly for registered`);

        await ico.addWhitelistedKYC(accounts[0]);

        isWhitelisted = await ico.isWhitelisted(accounts[0]);

        assert.equal(isWhitelisted.toString(), "2", `Account ${accounts[0]} wasn't whitelisted properly for KYC`);
    });

    it("should be able to invest", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("0.5", "ether"), from: accounts[0]})

        let accountBalance = await ico.balanceOf(accounts[0]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei((4750 * 0.5).toString(), "ether").toString(), "Account's balance should be 4750 LBC");
    });

    it("Should adjust by 4%", async () => {
        let ico = await HLBICO.deployed();

        await ico.adjustEtherValue(new web3.utils.BN("960"))

        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("0.5", "ether"), from: accounts[0]})

        let accountBalance = await ico.balanceOf(accounts[0]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei((4940 * 0.5 + (4750 * 0.5)).toString(), "ether").toString(), "Account's balance should be 9500 LBC");
    });
        
});