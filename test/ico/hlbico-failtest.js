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

    it("should reproduce part 1 of testnet (5ETH sent)", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("1", "ether"), from: accounts[0]});
        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("3", "ether"), from: accounts[0]});
        await ico.buyTokens(accounts[1], {value: web3.utils.toWei("1", "ether"), from: accounts[1]});

        await ico.adjustEtherValue(1040);

        var rate = await ico.rate();
        assert.equal(rate.toString(), '4632', "Rate should have been updated to 4632");

        var cap = await ico.cap();
        assert.equal(cap.toString(), '62400000000000000000', "cap should be 62.4ETH");

        var weiRaised = await ico.weiRaised();
        assert.equal(weiRaised.toString(), web3.utils.toWei("5", "ether"), "weiRaised should be 5ETH");

        var etherTranche = await ico.etherTranche();
        assert.equal(etherTranche.toString(), '3120000000000000000', "etherTranche should be 3.12ETH");
    });

    it("should reproduce part 2 of testnet (11ETH sent)", async () => {
    	let ico = await HLBICO.deployed();

    	await ico.buyTokens(accounts[1], {value: web3.utils.toWei("3", "ether"), from: accounts[1]});
        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("3", "ether"), from: accounts[0]});

        await ico.adjustEtherValue(920);

        var rate = await ico.rate();
        assert.equal(rate.toString(), '4639', "Rate should have been updated to 4639");

        var cap = await ico.cap();
        assert.equal(cap.toString(), '57408000000000000000', "cap should be 57.408ETH");

        var weiRaised = await ico.weiRaised();
        assert.equal(weiRaised.toString(), web3.utils.toWei("11", "ether"), "weiRaised should be 11ETH");

        var etherTranche = await ico.etherTranche();
        assert.equal(etherTranche.toString(), '2870400000000000000', "etherTranche should be 28.704ETH");
    });

    it("should reproduce part 3 of testnet (16ETH sent)", async () => {
    	let ico = await HLBICO.deployed();

    	await ico.buyTokens(accounts[1], {value: web3.utils.toWei("2", "ether"), from: accounts[1]});
        await ico.buyTokens(accounts[2], {value: web3.utils.toWei("3", "ether"), from: accounts[2]});

        await ico.adjustEtherValue(960);

        var rate = await ico.rate();
        assert.equal(rate.toString(), '4448', "Rate should have been updated to 4448");

        var cap = await ico.cap();
        assert.equal(cap.toString(), '55111680000000000000', "cap should be 55.11168ETH");

        var weiRaised = await ico.weiRaised();
        assert.equal(weiRaised.toString(), web3.utils.toWei("16", "ether"), "weiRaised should be 16ETH");

        var etherTranche = await ico.etherTranche();
        assert.equal(etherTranche.toString(), '2755584000000000000', "etherTranche should be 27.55584ETH");
    });

    it("should reproduce part 4 of testnet (22,5ETH sent)", async () => {
    	let ico = await HLBICO.deployed();

    	await ico.buyTokens(accounts[1], {value: web3.utils.toWei("1500", "milliether"), from: accounts[1]});
        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("1", "ether"), from: accounts[0]});
        await ico.buyTokens(accounts[1], {value: web3.utils.toWei("1", "ether"), from: accounts[1]});
        await ico.buyTokens(accounts[2], {value: web3.utils.toWei("3", "ether"), from: accounts[2]});

        await ico.adjustEtherValue(880);

        var rate = await ico.rate();
        assert.equal(rate.toString(), '4560', "Rate should have been updated to 4560");

        var cap = await ico.cap();
        assert.equal(cap.toString(), '48498278400000000000', "cap should be 48.4982784ETH");

        var weiRaised = await ico.weiRaised();
        assert.equal(weiRaised.toString(), web3.utils.toWei("22500", "milliether"), "weiRaised should be 22,5ETH");

        var etherTranche = await ico.etherTranche();
        assert.equal(etherTranche.toString(), '2424913920000000000', "etherTranche should be 2.42491392ETH");
    });

    it("should reproduce part 5 of testnet (44,99ETH sent)", async () => {
    	let ico = await HLBICO.deployed();

    	await ico.buyTokens(accounts[3], {value: web3.utils.toWei("2", "ether"), from: accounts[3]});
        await ico.buyTokens(accounts[3], {value: web3.utils.toWei("3", "ether"), from: accounts[3]});
        await ico.buyTokens(accounts[3], {value: '1498278400000000000', from: accounts[3]});
        await ico.buyTokens(accounts[2], {value: web3.utils.toWei("2", "ether"), from: accounts[2]});
        await ico.buyTokens(accounts[4], {value: web3.utils.toWei("3", "ether"), from: accounts[4]});
        await ico.buyTokens(accounts[4], {value: web3.utils.toWei("3", "ether"), from: accounts[4]});
        await ico.buyTokens(accounts[4], {value: web3.utils.toWei("2", "ether"), from: accounts[4]});
        await ico.buyTokens(accounts[5], {value: web3.utils.toWei("3", "ether"), from: accounts[5]});
        await ico.buyTokens(accounts[6], {value: web3.utils.toWei("3", "ether"), from: accounts[6]});

        var rate = await ico.rate();
        assert.equal(rate.toString(), '2670', "Rate should have been updated to 2670");

        var cap = await ico.cap();
        assert.equal(cap.toString(), '48498278400000000000', "cap should be XXX");

        var weiRaised = await ico.weiRaised();
        assert.equal(weiRaised.toString(), '44998278400000000000', "weiRaised should be 44,9982784ETH");

        var etherTranche = await ico.etherTranche();
        assert.equal(etherTranche.toString(), '2424913920000000000', "etherTranche should be XXX");
    });

    it("should Fail with error 'SafeMath: subtraction overflow'", async () => {
    	let ico = await HLBICO.deployed();

    	await ico.buyTokens(accounts[5], {value: web3.utils.toWei("3", "ether"), from: accounts[5]});

        var weiRaised = await ico.weiRaised();
        assert.equal(weiRaised.toString(), '47998278400000000000', "weiRaised should be 47,9982784ETH");
    });

});

})