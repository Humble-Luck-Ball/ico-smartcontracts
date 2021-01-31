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

    /*
    ** Attrs
    */
    uint256 private _currentRate;
    mapping(address => bool) whitelistedAddrs;
    uint256 _weiMaxInvest;
    uint256 _etherTranche;
    uint256 _currentWeiTranche; // Holds the current invested value for a tranche

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
        _etherTranche = 300000000000000000000; // For eth = 1000€
        _weiMaxInvest = 10000000000000000000; // 10.000€; for eth = 1000 €
        _currentRate = initialRateReceived;
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

        _currentWeiTranche.add(weiAmount);

        if (_currentWeiTranche > _etherTranche) {
            _currentWeiTranche = _currentWeiTranche.sub(_etherTranche);
            calculatedAmount = calculatedAmount.add(weiAmount.sub(_currentWeiTranche).mul(rate()).div(1000));
            _currentRate -= 175; // coefficient for 35 tokens reduction for each tranche
            calculatedAmount = calculatedAmount.add(_currentWeiTranche.mul(rate()).div(1000));
        }
        else
            calculatedAmount.add(weiAmount.mul(rate()).div(1000));

        calculatedAmount = calculatedAmount.sub(calculatedAmount.mul(5).div(100));

        return calculatedAmount;
    }

    function rate() public view override returns (uint256) {
       return _currentRate;
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
        _dontExceedAmount(weiAmount);
        CappedTimedCrowdsale._preValidatePurchase(beneficiary, weiAmount);
    }

    function _processPurchase(address beneficiary, uint256 tokenAmount) internal override(CrowdsaleMint, RefundablePostDeliveryCrowdsale) {
        RefundablePostDeliveryCrowdsale._processPurchase(beneficiary, tokenAmount);
    }

    function isWhitelisted(address account) public view returns (bool) {
        require(account != address(0), "HLCICO: account is zero address");
        return whitelistedAddrs[account];
    }

    function addWhitelisted(address account) public onlyWhitelistingAddress {
        _addWhitelisted(account);
    }

    function removeWhitelisted(address account) public onlyWhitelistingAddress {
        _removeWhitelisted(account);
    }

    function _addWhitelisted(address account) internal {
        require(!isWhitelisted(account), "HLCICO: account already whitelisted");
        whitelistedAddrs[account] = true;
        emit WhitelistedAdded(account);
    }

    function _removeWhitelisted(address account) internal {
        require(isWhitelisted(account), "HLCICO: account is not whitelisted");
        whitelistedAddrs[account] = false;
        emit WhitelistedRemoved(account);
    }

    function _dontExceedAmount(uint256 weiAmount) internal view {
        require(weiAmount <= maxEtherToInvest(), "HLBICO: Cannot invest more than limit");
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