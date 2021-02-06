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

        assert.equal(rate.toString(), '5000', `First rate shoud be 5000`)
    });

    it("should be able to invest 1ETH", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("1", "ether"), from: accounts[0]})

        let accountBalance = await ico.balanceOf(accounts[0]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("4750", "ether").toString(), "Account's balance should be 4750 LBC");
    });

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

        let rate = await ico.rate();

        assert.equal(rate.toString(), '4825', `Second rate shoud be 4825`)
    });

    it("should have balance of 3 times the rate - 5%", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedKYC(accounts[3]);

        await ico.buyTokens(accounts[3], {value: web3.utils.toWei("2", "ether"), from: accounts[3]})

        let accountBalance = await ico.balanceOf(accounts[3]);

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("13751250", "milliether").toString(), "Account's balance should be 13751,25 LBC");
    });

    it("should trigger another cap with rate 4650 (cap3)", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[4]);

        await ico.addWhitelistedKYC(accounts[4]);

        await ico.buyTokens(accounts[4], {value: web3.utils.toWei("3", "ether"), from: accounts[4]})

        let accountBalance = await ico.balanceOf(accounts[4]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '4650', `Third rate shoud be 4650`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("13252500", "milliether").toString(), "Account's balance should be 13252,5 LBC");
    });

    it("should trigger another cap with rate 4475 (cap4)", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[5]);

        await ico.addWhitelistedKYC(accounts[5]);

        await ico.buyTokens(accounts[5], {value: web3.utils.toWei("3", "ether"), from: accounts[5]})

        let accountBalance = await ico.balanceOf(accounts[5]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '4475', `Fourth rate shoud be 4475`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("12753750", "milliether").toString(), "Account's balance should be 12753,75 LBC");
    });

    it("should trigger another cap with rate 4300 (cap5)", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[6]);

        await ico.addWhitelistedKYC(accounts[6]);

        await ico.buyTokens(accounts[6], {value: web3.utils.toWei("3", "ether"), from: accounts[6]})

        let accountBalance = await ico.balanceOf(accounts[6]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '4300', `Fifth rate shoud be 4300`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("12255000", "milliether").toString(), "Account's balance should be 12255 LBC");
    });

    it("should trigger another cap with rate 4125 (cap6)", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[7]);

        await ico.addWhitelistedKYC(accounts[7]);

        await ico.buyTokens(accounts[7], {value: web3.utils.toWei("3", "ether"), from: accounts[7]})

        let accountBalance = await ico.balanceOf(accounts[7]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '4125', `Sixth rate shoud be 4125`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("11756250", "milliether").toString(), "Account's balance should be 11756,25 LBC");
    });

    it("should trigger another cap with rate 3950 (cap7)", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[8]);

        await ico.addWhitelistedKYC(accounts[8]);

        await ico.buyTokens(accounts[8], {value: web3.utils.toWei("3", "ether"), from: accounts[8]})

        let accountBalance = await ico.balanceOf(accounts[8]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '3950', `Seventh rate shoud be 3950`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("11257500", "milliether").toString(), "Account's balance should be 11257,5 LBC");
    });

    it("should trigger another cap with rate 3775 (cap8)", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[9]);

        await ico.addWhitelistedKYC(accounts[9]);

        await ico.buyTokens(accounts[9], {value: web3.utils.toWei("3", "ether"), from: accounts[9]})

        let accountBalance = await ico.balanceOf(accounts[9]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '3775', `Eighth rate shoud be 3775`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("10758750", "milliether").toString(), "Account's balance should be 10758,75 LBC");
    });

    it("should trigger another cap with rate 3600 (cap9)", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[1]);

        await ico.addWhitelistedKYC(accounts[1]);

        await ico.buyTokens(accounts[1], {value: web3.utils.toWei("3", "ether"), from: accounts[1]})

        let accountBalance = await ico.balanceOf(accounts[1]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '3600', `Ninth rate shoud be 3600`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("10260000", "milliether").toString(), "Account's balance should be 10260 LBC");
    });

    it("should trigger another cap with rate 3425 (cap10)", async () => {
        let ico = await HLBICO.deployed();

        await ico.addWhitelistedRegistered(accounts[2]);

        await ico.addWhitelistedKYC(accounts[2]);

        await ico.buyTokens(accounts[2], {value: web3.utils.toWei("3", "ether"), from: accounts[2]})

        let accountBalance = await ico.balanceOf(accounts[2]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '3425', `Tenth rate shoud be 3425`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("9761250", "milliether").toString(), "Account's balance should be 9761,25 LBC");
    });

    it("should trigger another cap with rate 3250 (cap11)", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[3], {value: web3.utils.toWei("3", "ether"), from: accounts[3]})

        let accountBalance = await ico.balanceOf(accounts[3]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '3250', `Eleventh rate shoud be 3250`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("23013750", "milliether").toString(), "Account's balance should be 23013,75 LBC (13751,25 + 9262,5)"); 
    });

    it("should trigger another cap with rate 3075 (cap12)", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[4], {value: web3.utils.toWei("3", "ether"), from: accounts[4]})

        let accountBalance = await ico.balanceOf(accounts[4]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '3075', `Twelveth rate shoud be 3075`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("22016250", "milliether").toString(), "Account's balance should be 22016,25 LBC (13252,5 + 8763,75)"); 
    });

    it("should trigger another cap with rate 2900 (cap13)", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[5], {value: web3.utils.toWei("3", "ether"), from: accounts[5]})

        let accountBalance = await ico.balanceOf(accounts[5]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '2900', `Thirteenth rate shoud be 2900`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("21018750", "milliether").toString(), "Account's balance should be 21018,75 LBC (12753,75 + 8265)"); 
    });

    it("should trigger another cap with rate 2725 (cap14)", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[6], {value: web3.utils.toWei("3", "ether"), from: accounts[6]})

        let accountBalance = await ico.balanceOf(accounts[6]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '2725', `Fourteenth rate shoud be 2725`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("20021250", "milliether").toString(), "Account's balance should be 20021,25 LBC (12255 + 7766,25)"); 
    });

    it("should trigger another cap with rate 2550 (cap15)", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[7], {value: web3.utils.toWei("3", "ether"), from: accounts[7]})

        let accountBalance = await ico.balanceOf(accounts[7]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '2550', `Fifteenth rate shoud be 2550`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("19023750", "milliether").toString(), "Account's balance should be 19023,75 LBC (11756,25 + 7267,5)"); 
    });

    it("should trigger another cap with rate 2375 (cap16)", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[8], {value: web3.utils.toWei("3", "ether"), from: accounts[8]})

        let accountBalance = await ico.balanceOf(accounts[8]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '2375', `Sixteenth rate shoud be 2375`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("18026250", "milliether").toString(), "Account's balance should be 18026,25 LBC (11257,5 + 6768,75)"); 
    });

    it("should trigger another cap with rate 2200 (cap17)", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[9], {value: web3.utils.toWei("3", "ether"), from: accounts[9]})

        let accountBalance = await ico.balanceOf(accounts[9]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '2200', `Seventeenth rate shoud be 2200`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("17028750", "milliether").toString(), "Account's balance should be 17028,75 LBC (10758,75 + 6270)"); 
    });

    it("should trigger another cap with rate 2025 (cap18)", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[1], {value: web3.utils.toWei("3", "ether"), from: accounts[1]})

        let accountBalance = await ico.balanceOf(accounts[1]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '2025', `Eighteenth rate shoud be 2025`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("16031250", "milliether").toString(), "Account's balance should be 16031,25 LBC (10260 + 5771,25)"); 
    });

    it("should trigger another cap with rate 1850 (cap19)", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[2], {value: web3.utils.toWei("3", "ether"), from: accounts[2]})

        let accountBalance = await ico.balanceOf(accounts[2]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '1850', `Nineteenth rate shoud be 1850`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("15033750", "milliether").toString(), "Account's balance should be 15033,75 LBC (9761,25 + 5272,5)"); 
    });

    it("should trigger Hard Cap with rate 1675 (cap20)", async () => {
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("3", "ether"), from: accounts[0]})

        let accountBalance = await ico.balanceOf(accounts[0]);

        let rate = await ico.rate();
        assert.equal(rate.toString(), '1675', `Twenteeth rate shoud be 1675`)

        // for 1eth = 1000€
        assert.equal(accountBalance.toString(), web3.utils.toWei("19023750", "milliether").toString(), "Account's balance should be 19023,75 LBC (14250 + 4773,75)"); 
    });

    it("should reach Hard Cap goal", async () => {
        let ico = await HLBICO.deployed();

        let capReached = await ico.capReached();

        assert.equal(capReached, true, `Hard Cap goal shoud be reached`);
    });

    it("should finalize and send the promised LBC tokens", async () => {
        let ico = await HLBICO.deployed();

        await ico.finalize();

        let finalized = await ico.finalized()

        assert.equal(finalized, true, `Should have been able to finalize`);
    });

/*
    it("should mean all accounts received a total of 19.023.750LBC", async () => {
        let ico = await HLBICO.deployed();

        let accountBalance0 = await ico.balanceOf(accounts[0]);
        let accountBalance1 = await ico.balanceOf(accounts[1]);
        let accountBalance2 = await ico.balanceOf(accounts[2]);
        let accountBalance3 = await ico.balanceOf(accounts[3]);
        let accountBalance4 = await ico.balanceOf(accounts[4]);
        let accountBalance5 = await ico.balanceOf(accounts[5]);
        let accountBalance6 = await ico.balanceOf(accounts[6]);
        let accountBalance7 = await ico.balanceOf(accounts[7]);
        let accountBalance8 = await ico.balanceOf(accounts[8]);
        let accountBalance9 = await ico.balanceOf(accounts[9]);

        let allAccounts = accountBalance0 += accountBalance1 += accountBalance2 += accountBalance3 += accountBalance4 += accountBalance5 += accountBalance6 += accountBalance7 += accountBalance8 += accountBalance9

        assert.equal(allAccounts.toString(), web3.utils.toWei("19023750", "ether").toString(), "Total distributed tokens should be 19023750")
    });

    it("should not accept any more Ether", async () => { // Returns "VM exception CappedCrowdsale : cap exceeded" as intended.
        let ico = await HLBICO.deployed();

        await ico.buyTokens(accounts[0], {value: web3.utils.toWei("3", "ether"), from: accounts[0]})

        let accountBalance = await ico.balanceOf(accounts[0]);

        assert.equal(accountBalance.toString(), web3.utils.toWei("19023750", "milliether").toString(), "Account's balance should be 19023,75 LBC (14250 + 4773,75)");
    });*/
});

    describe('LBC Token', () => {
    it('tracks the name', async () => {
        let lbc = await LBCToken.deployed();

        const name = await lbc.name()
        assert.equal(name, 'Luck Ball Coin', `Our token should definitely be named Luck Ball Coin`)
    })

    it('tracks the symbol', async ()  => {
        let lbc = await LBCToken.deployed();

        const symbol = await lbc.symbol()
        assert.equal(symbol, 'LBC', `Our token should definitely have the LBC symbol`)
    })

    it('tracks the decimals', async ()  => {
        let lbc = await LBCToken.deployed();

        const decimals = await lbc.decimals()
        assert.equal(decimals.toString(), '18', `Our token should have 18 decimals`)
    })

    it('tracks the total supply', async ()  => { // Total supply doesn't seem to be defined.
        let lbc = await LBCToken.deployed();

        const totalSupply = await lbc.totalSupply()
        result.equal(totalSupply.toString(), '300000000', `Our token should have 300000000 total supply`)
    })

    it("should let users trade their LBC tokens", async () => {

    })
});

})



