const { assert } = require("chai");

const LBCToken = artifacts.require("LBCToken")
const HLBICO = artifacts.require("HLBICO")

contract("HLBICO", accounts => {
    describe('What the ICO should do', async () => {
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

        let isWhitelisted = await ico.isWhitelisted(accounts[0]);

        assert.equal(isWhitelisted.toString(), "0", `Account ${accounts[0]} was already whitelisted when it shouldn't be`);

        await ico.addWhitelistedRegistered(accounts[0]);

        isWhitelisted = await ico.isWhitelisted(accounts[0]);

        assert.equal(isWhitelisted.toString(), "1", `Account ${accounts[0]} wasn't whitelisted properly for registered`);

        await ico.addWhitelistedKYC(accounts[0]);

        isWhitelisted = await ico.isWhitelisted(accounts[0]);

        assert.equal(isWhitelisted.toString(), "2", `Account ${accounts[0]} wasn't whitelisted properly for KYC`);
    });

    it("should have a rate of 5000 (cap1)", async () => {
        let ico = await HLBICO.deployed();

        let rate = await ico.rate();

        assert.equal(rate.toString(), '5000', `First rate shoud be 5000`);
    });

    it("should be able to invest 1ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("1", "ether"), from: accounts[0]})

        let accountBalance = await ico.balanceOf(accounts[0]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("4750", "ether").toString(), "Account's balance should be 4750 LBC");
    });
    /*
    it("should be able to invest 2ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("2", "ether"), from: accounts[0]})

        let accountBalance = await ico.balanceOf(accounts[0]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("14250", "ether").toString(), "Account's balance should be 14250 LBC");
    });

    it("should reach Soft Cap goal (cap1)", async () => {
        let ico = await HLBICO.deployed();

        let goal = await ico.goalReached();

        assert.equal(goal, true, `Soft Cap goal is reached`);
    });

    it("should trigger another cap (cap2)", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[3]);

        let isWhitelisted = await ico.isWhitelisted(accounts[3]);

        assert.equal(isWhitelisted.toString(), "1", `Account ${accounts[3]} wasn't whitelisted properly`);

        await ico.buyTokens(accounts[3], {value: web3.utils.toWei("1", "ether"), from: accounts[3]})

        let accountBalance = await ico.balanceOf(accounts[3]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("4583750", "milliether").toString(), "Account's balance should be 4583,75 LBC");
    });

    it("should have an updated rate of 4825", async () => {
        let ico = await HLBICO.deployed();

        var rate = await ico.rate();

        assert.equal(rate.toString(), '4825', `Second rate shoud be 4825`)
    });*/

    it("should adjust Ether value to 96%", async () => {
    	let ico = await HLBICO.deployed();

    	let adjust900 = await ico.adjustEtherValue(960);

        adjust900;

        var rate2 = await ico.rate();

        assert.equal(rate2.toString(), '5200', "Rate should have been updated");
    });

    it("should update the etherTranche", async () => {
        let ico = await HLBICO.deployed();

        let tranche1 = await ico.etherTranche();

        assert.equal(tranche1.toString(), web3.utils.toWei("2880", "milliether").toString(), "etherTranche updated to 2,88ETH")
    })

    it("should update the cap", async () => {
        let ico = await HLBICO.deployed();

        let cap1 = await ico.cap();

        assert.equal(cap1.toString(), web3.utils.toWei("57600", "milliether").toString(), "Cap updated to 57,6ETH")
    })

    it("should be able to invest 1ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[1]);

        await ico.addWhitelistedKYC(accounts[1]);

        await ico.buyTokens(accounts[1], {value: web3.utils.toWei("1", "ether"), from: accounts[1]})

        let accountBalance = await ico.balanceOf(accounts[1]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("4940", "ether").toString(), "Account's balance should be 4940 LBC");
    });

    it("should adjust Ether value to 112%", async () => {
        let ico = await HLBICO.deployed();

        await ico.adjustEtherValue(1120);

        var rate2 = await ico.rate();

        assert.equal(rate2.toString(), '4576', "Rate should have been updated");
    });

    it("should update the etherTranche", async () => {
        let ico = await HLBICO.deployed();

        let tranche1 = await ico.etherTranche();

        assert.equal(tranche1.toString(), web3.utils.toWei("3225600", "microether").toString(), "etherTranche updated to 3,2256ETH")
    })

    it("should update the cap", async () => {
        let ico = await HLBICO.deployed();

        let cap1 = await ico.cap();

        assert.equal(cap1.toString(), web3.utils.toWei("64512", "milliether").toString(), "Cap updated to 64,512ETH")
    })

    it("should be able to invest 1ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[2]);

        await ico.addWhitelistedKYC(accounts[2]);

        await ico.buyTokens(accounts[2], {value: web3.utils.toWei("1", "ether"), from: accounts[2]})

        let accountBalance = await ico.balanceOf(accounts[2]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("4347200", "milliether").toString(), "Account's balance should be 4347,2 LBC");
    });

    it("should be able to invest 1ETH and trigger a new tranche", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[3]);

        await ico.addWhitelistedKYC(accounts[3]);

        await ico.buyTokens(accounts[3], {value: web3.utils.toWei("1", "ether"), from: accounts[3]})

        let accountBalance = await ico.balanceOf(accounts[3]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), '4228755520000000000000', "Account's balance should be 4347,2 LBC");
    });

    it("should have an updated rate of 4415", async () => {
        let ico = await HLBICO.deployed();

        var rate = await ico.rate();

        assert.equal(rate.toString(), '4415', `Updated rate shoud be 4415`)
    });

    it("should have a weiRaised of 4ETH", async () => {
        let ico = await HLBICO.deployed();

        var weiRaised = await ico.weiRaised();

        assert.equal(weiRaised.toString(), web3.utils.toWei("4", "ether").toString(), `weiRaised should be 4ETH`)
    });

    // REACHING HARD CAP

    it("should invest 3ETH + 3ETH + 3ETH + 1ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[4]);
        await ico.addWhitelistedKYC(accounts[4]);

        await ico.buyTokens(accounts[4], {value: web3.utils.toWei("3", "ether"), from: accounts[4]});
        await ico.buyTokens(accounts[4], {value: web3.utils.toWei("3", "ether"), from: accounts[4]});
        await ico.buyTokens(accounts[4], {value: web3.utils.toWei("3", "ether"), from: accounts[4]});
        await ico.buyTokens(accounts[4], {value: web3.utils.toWei("1", "ether"), from: accounts[4]});

        let accountBalance = await ico.balanceOf(accounts[4]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), '39958799680000000000000', "Account's balance should be 39.958,79968 LBC");
    });

    it("should invest 3ETH + 3ETH + 3ETH + 1ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[5]);
        await ico.addWhitelistedKYC(accounts[5]);

        await ico.buyTokens(accounts[5], {value: web3.utils.toWei("3", "ether"), from: accounts[5]});
        await ico.buyTokens(accounts[5], {value: web3.utils.toWei("3", "ether"), from: accounts[5]});
        await ico.buyTokens(accounts[5], {value: web3.utils.toWei("3", "ether"), from: accounts[5]});
        await ico.buyTokens(accounts[5], {value: web3.utils.toWei("1", "ether"), from: accounts[5]});

        let accountBalance = await ico.balanceOf(accounts[5]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), '35221999360000000000000', "Account's balance should be 35.221,99936 LBC");
    });

    it("should invest 3ETH + 3ETH + 3ETH + 1ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[6]);
        await ico.addWhitelistedKYC(accounts[6]);

        await ico.buyTokens(accounts[6], {value: web3.utils.toWei("3", "ether"), from: accounts[6]});
        await ico.buyTokens(accounts[6], {value: web3.utils.toWei("3", "ether"), from: accounts[6]});
        await ico.buyTokens(accounts[6], {value: web3.utils.toWei("3", "ether"), from: accounts[6]});
        await ico.buyTokens(accounts[6], {value: web3.utils.toWei("1", "ether"), from: accounts[6]});

        let accountBalance = await ico.balanceOf(accounts[6]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), '30485199040000000000000', "Account's balance should be 30.485,19904 LBC");
    });

    it("should invest 3ETH + 3ETH + 3ETH + 1ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[7]);
        await ico.addWhitelistedKYC(accounts[7]);

        await ico.buyTokens(accounts[7], {value: web3.utils.toWei("3", "ether"), from: accounts[7]});
        await ico.buyTokens(accounts[7], {value: web3.utils.toWei("3", "ether"), from: accounts[7]});
        await ico.buyTokens(accounts[7], {value: web3.utils.toWei("3", "ether"), from: accounts[7]});
        await ico.buyTokens(accounts[7], {value: web3.utils.toWei("1", "ether"), from: accounts[7]});

        let accountBalance = await ico.balanceOf(accounts[7]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), '25748398720000000000000', "Account's balance should be 25.748,39872 LBC");
    });

    it("should invest 3ETH + 3ETH + 3ETH + 1ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[8]);
        await ico.addWhitelistedKYC(accounts[8]);

        await ico.buyTokens(accounts[8], {value: web3.utils.toWei("3", "ether"), from: accounts[8]});
        await ico.buyTokens(accounts[8], {value: web3.utils.toWei("3", "ether"), from: accounts[8]});
        await ico.buyTokens(accounts[8], {value: web3.utils.toWei("3", "ether"), from: accounts[8]});
        await ico.buyTokens(accounts[8], {value: web3.utils.toWei("1", "ether"), from: accounts[8]});

        let accountBalance = await ico.balanceOf(accounts[8]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), '21011598400000000000000', "Account's balance should be 21.011,5984 LBC");
    });

    it("should invest 3ETH + 3ETH + 3ETH + 1ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[9]);
        await ico.addWhitelistedKYC(accounts[9]);

        await ico.buyTokens(accounts[9], {value: web3.utils.toWei("3", "ether"), from: accounts[9]});
        await ico.buyTokens(accounts[9], {value: web3.utils.toWei("3", "ether"), from: accounts[9]});
        await ico.buyTokens(accounts[9], {value: web3.utils.toWei("3", "ether"), from: accounts[9]});
        await ico.buyTokens(accounts[9], {value: web3.utils.toWei("1", "ether"), from: accounts[9]});

        let accountBalance = await ico.balanceOf(accounts[9]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), '16274798080000000000000', "Account's balance should be 16.274,49808 LBC");
    });

    it("should have a weiRaised of 64ETH", async () => {
        let ico = await HLBICO.deployed();

        var weiRaised = await ico.weiRaised();

        assert.equal(weiRaised.toString(), web3.utils.toWei("64", "ether").toString(), `weiRaised should be 64ETH`)
    });

    it("should be able to invest 1ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("512", "milliether"), from: accounts[0]})

        let accountBalance = await ico.balanceOf(accounts[0]);

        // for 1eth = 1000€
        assert.equal(accountBalance, '5487868800000000000000', "Account's balance should be 5487,xxxx LBC");
    });

    it("should be able to finalize", async () => {
        let ico = await HLBICO.deployed();

        await ico.finalize();

        let finalized = ico.finalized();

        assert.equal(finalized, true, "Should have finalized");
    })
});
})