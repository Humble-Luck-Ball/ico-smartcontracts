// SPDX-License-Identifier: MIT

pragma solidity ^0.7.0;

import "../../libraries/openzeppelin/math/SafeMath.sol";
import "./CappedTimedCrowdsale.sol";
import "./RefundPostdevCrowdsale.sol";


/**
**  ICO Contract for the LBC crowdsale
*/
contract HLBICO is CappedTimedCrowdsale, RefundablePostDeliveryCrowdsale {
    using SafeMath for uint256;

    /*
    ** Global State
    */
    bool public initialized; // default : false

    /*
    ** Addresses
    */
    address public _deployingAddress; // should remain the same as deployer's address
    address public _whitelistingAddress; // should be oracle
    address public _reserveAddress; // should be deployer then humble reserve

    /*
    ** Events
    */
    event InitializedContract(address indexed changerAddress, address indexed whitelistingAddress);
    event ChangedWhitelisterAddress(address indexed whitelisterAddress, address indexed changerAddress);
    event ChangedReserveAddress(address indexed reserveAddress, address indexed changerAddress);
    event WhitelistedAdded(address indexed account);
    event WhitelistedRemoved(address indexed account);
    event UpdatedCaps(uint256 newGoal, uint256 newCap, uint256 newTranche, uint256 newMaxInvest, uint256 newRate, uint256 newRateCoef);

    /*
    ** Attrs
    */
    uint256 private _currentRate;
    uint256 private _rateCoef;
    mapping(address => uint8) private _whitelistedAddrs;
    mapping(address => uint256) private _investmentAddrs;
    uint256 private _weiMaxInvest;
    uint256 private _weiNoKYCMaxInvest;
    uint256 private _etherTranche;
    uint256 private _currentWeiTranche; // Holds the current invested value for a tranche

    /*
    * initialRateReceived : Number of token units a buyer gets per wei for the first investment slice. Should be 5000 (diving by 1000 for 3 decimals).
    * walletReceived : Wallet that will get the invested eth at the end of the crowdsale
    * tokenReceived : Address of the LBC token being sold
    * openingTimeReceived : Starting date of the ICO
    * closingtimeReceived : Ending date of the ICO
    * capReceived : Max amount of wei to be contributed
    * goalReceived : Funding goal
    * etherMaxInvestReceived : Maximum ether that can be invested
    */
    constructor(uint256 initialRateReceived,
        uint256 rateCoefficientReceived,
        address payable walletReceived,
        LBCToken tokenReceived,
        uint256 openingTimeReceived,
        uint256 closingTimeReceived,
        uint256 capReceived,
        uint256 goalReceived)
        CrowdsaleMint(initialRateReceived, walletReceived, tokenReceived)
        TimedCrowdsale(openingTimeReceived, closingTimeReceived)
        CappedTimedCrowdsale(capReceived)
        RefundableCrowdsale(goalReceived) {
        _deployingAddress = msg.sender;
        _etherTranche = 3000000000000000000; // 3eth For eth = 1000€; DANGER : Don't be a bottom and change it back to its previous value : 300000000000000000000 
        _weiMaxInvest = 10000000000000000000; // 10.000€; for eth = 1000 €
        _weiNoKYCMaxInvest = 1000000000000000000; // 1000€; for eth = 1000 €
        _currentRate = initialRateReceived;
        _rateCoef = rateCoefficientReceived;
        _currentWeiTranche = 0;
    }

    /*
    ** Initializes the contract address and affects addresses to their roles.
    */
    function init(
        address whitelistingAddress,
        address reserveAddress
    )
    public
    isNotInitialized
    onlyDeployingAddress
    {
        require(whitelistingAddress != address(0), "HLBICO: whitelistingAddress cannot be 0x");
        require(reserveAddress != address(0), "HLBICO: reserveAddress cannot be 0x");

        _whitelistingAddress = whitelistingAddress;
        _reserveAddress = reserveAddress;
        initialized = true;

        emit InitializedContract(_msgSender(), whitelistingAddress);
    }

    /**
     * @dev Returns the rate of tokens per wei at the present time and computes rate depending on tranche.
     * @param weiAmount The value in wei to be converted into tokens
     * @return The number of tokens a buyer gets per wei for a given tranche
     */
    function _getCustomAmount(uint256 weiAmount) internal returns (uint256) {
        if (!isOpen()) {
            return 0;
        }

        uint256 calculatedAmount = 0;

        _currentWeiTranche = _currentWeiTranche.add(weiAmount);

        if (_currentWeiTranche > _etherTranche) {
            _currentWeiTranche = _currentWeiTranche.sub(_etherTranche);

            //If we updated the tranche manually to a smaller one
            uint256 manualSkew = weiAmount.sub(_currentWeiTranche);

            if (manualSkew >= 0) {
                calculatedAmount = calculatedAmount.add(weiAmount.sub(_currentWeiTranche).mul(rate()));
                _currentRate -= 175; // coefficient for 35 tokens reduction for each tranche
                calculatedAmount = calculatedAmount.add(_currentWeiTranche.mul(rate()));
            }
            //If there is a skew between invested wei and calculated wei for a tranche
            else {
                _currentRate -= 175; // coefficient for 35 tokens reduction for each tranche
                calculatedAmount = calculatedAmount.add(weiAmount.mul(rate()));
            }
        }
        else
            calculatedAmount = calculatedAmount.add(weiAmount.mul(rate()));

        calculatedAmount = calculatedAmount.sub(calculatedAmount.mul(5).div(100));

        return calculatedAmount;
    }

    /*
    ** Adjusts all parameters influenced by Ether value based on a percentage coefficient
    ** coef is based on 4 digits for decimal representation with 1 precision
    ** i.e : 934 -> 93.4%; 1278 -> 127.8%
    */
    function adjustEtherValue(uint256 coef)
    public
    onlyDeployingAddress {
        require(coef > 0 && coef < 10000, "HLBICO: coef isn't within range of authorized values");

        changeGoal(goal().mul(coef).div(1000));
        changeCap(cap().mul(coef).div(1000));
        _etherTranche = _etherTranche.mul(coef).div(1000);
        _weiMaxInvest = _weiMaxInvest.mul(coef).div(1000);
        _weiNoKYCMaxInvest = _weiNoKYCMaxInvest.mul(coef).div(1000);
        _currentRate = _currentRate.mul(coef).div(1000);
        _rateCoef = _rateCoef.mul(coef).div(1000);

        emit UpdatedCaps(goal(), cap(), _etherTranche, _weiMaxInvest, _currentRate, _rateCoef);
    }

    function rate() public view override returns (uint256) {
       return _currentRate;
    }


    /*
    ** Changes the address of the token contract. Must only be callable by deployer
    */
    function changeToken(LBCToken newToken)
    public
    onlyDeployingAddress
    {
        _changeToken(newToken);
    }

    /*
    ** Changes the address with whitelisting role and can only be called by deployer
    */
    function changeWhitelister(address newWhitelisterAddress)
    public
    onlyDeployingAddress
    {
        _whitelistingAddress = newWhitelisterAddress;
        emit ChangedWhitelisterAddress(newWhitelisterAddress, _msgSender());
    }

    /*
    ** Changes the address with pause role and can only be called by deployer
    */
    function changeReserveAddress(address newReserveAddress)
    public
    onlyDeployingAddress
    {
        _reserveAddress = newReserveAddress;
        emit ChangedReserveAddress(newReserveAddress, _msgSender());
    }

    /**
     * @dev Escrow finalization task, called when finalize() is called.
     */
    function _finalization() override virtual internal {
        // Mints the 5% participation and sends it to humblereserve
        if (goalReached()) {
            _deliverTokens(_reserveAddress, weiRaised().mul(5).div(100));
        }

        super._finalization();
    }

    /**
     * @dev Overrides parent method taking into account variable rate.
     * @param weiAmount The value in wei to be converted into tokens
     * @return The number of tokens _weiAmount wei will buy at present time
     */
    function _getTokenAmount(uint256 weiAmount) internal override returns (uint256) {
       return _getCustomAmount(weiAmount);
    }

    function _forwardFunds() internal override(CrowdsaleMint, RefundablePostDeliveryCrowdsale) {
        RefundablePostDeliveryCrowdsale._forwardFunds();
    }

    function _preValidatePurchase(address beneficiary, uint256 weiAmount) internal override(TimedCrowdsale, CappedTimedCrowdsale) view {
        require(isWhitelisted(beneficiary) > 0, "HLBICO: Account is not whitelisted");
        _dontExceedAmount(beneficiary, weiAmount, isWhitelisted(beneficiary));
        CappedTimedCrowdsale._preValidatePurchase(beneficiary, weiAmount);
    }

    function _postValidatePurchase(address beneficiary, uint256 weiAmount) internal override {
        require(beneficiary != address(0), "HLBICO: _postValidatePurchase benificiary is the zero address");

        _investmentAddrs[beneficiary] = _investmentAddrs[beneficiary].add(weiAmount);        
    }

    function _processPurchase(address beneficiary, uint256 tokenAmount) internal override(CrowdsaleMint, RefundablePostDeliveryCrowdsale) {
        RefundablePostDeliveryCrowdsale._processPurchase(beneficiary, tokenAmount);
    }

    function hasClosed() public view override(TimedCrowdsale, CappedTimedCrowdsale) returns (bool) {
        // solhint-disable-next-line not-rely-on-time
        return CappedTimedCrowdsale.hasClosed();
    }

    function isWhitelisted(address account) public view returns (uint8) {
        require(account != address(0), "HLCICO: account is zero address");
        return _whitelistedAddrs[account];
    }

    function addWhitelistedRegistered(address account) public onlyWhitelistingAddress {
        _addWhitelisted(account, 1);
    }

    function addWhitelistedKYC(address account) public onlyWhitelistingAddress {
        _addWhitelisted(account, 2);
    }

    function removeWhitelisted(address account) public onlyWhitelistingAddress {
        _removeWhitelisted(account);
    }

    function _addWhitelisted(address account, uint8 flag) internal {
        require(flag == 1 || flag == 2, "HLBICO: whitelisting flag must be 1 or 2");
        require(isWhitelisted(account) < flag, "HLCICO: account already whitelisted");
        _whitelistedAddrs[account] = flag;
        emit WhitelistedAdded(account);
    }

    function _removeWhitelisted(address account) internal {
        require(isWhitelisted(account) > 0, "HLCICO: account is not whitelisted");
        _whitelistedAddrs[account] = 0;
        emit WhitelistedRemoved(account);
    }

    function _dontExceedAmount(address beneficiary, uint256 weiAmount, uint8 flag) internal view {
        require((_investmentAddrs[beneficiary].add(weiAmount) < _weiMaxInvest && flag == 2)
         || _investmentAddrs[beneficiary].add(weiAmount) <= _weiNoKYCMaxInvest,
          "HLBICO: Cannot invest more than NoKYC limit or weiMaxInvest. User needs to pass KYC screening first.");
    }

    function maxEtherToInvest() public view returns (uint256) {
        return _weiMaxInvest;
    }

    modifier onlyWhitelistingAddress() {
        require(_msgSender() == _whitelistingAddress, "HLBICO: caller does not have the Whitelisted role");
        _;
    }

    /*
    ** Checks if the contract hasn't already been initialized
    */
    modifier isNotInitialized() {
        require(initialized == false, "HLBICO: contract is already initialized.");
        _;
    }

    /*
    ** Checks if the sender is the minter controller address
    */
    modifier onlyDeployingAddress() {
        require(msg.sender == _deployingAddress, "HLBICO: only the deploying address can call this method.");
        _;
    }

}