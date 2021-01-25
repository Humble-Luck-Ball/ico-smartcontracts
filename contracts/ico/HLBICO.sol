// SPDX-License-Identifier: MIT

pragma solidity ^0.7.0;

import "./CappedTimedCrowdsale.sol";
import "./RefundPostdevCrowdsale.sol";


/**
**  ICO Contract for the LBC crowdsale
*/
contract HLBICO is CappedTimedCrowdsale, RefundablePostDeliveryCrowdsale {

    /*
    ** Global State
    */
    bool public initialized; // default : false

    /*
    ** Addresses
    */
    address public _deployingAddress; // should remain the same as deployer's address
    address public _whitelistingAddress; // should be oracle

    /*
    ** Events
    */
    event InitializedContract(address indexed changerAddress, address indexed whitelistingAddress);
    event WhitelistedAdded(address indexed account);
    event WhitelistedRemoved(address indexed account);

    /*
    ** Attrs
    */
    mapping(address => bool) whitelistedAddrs;

    /*
    * rateReceived : Number of token units a buyer gets per wei for the first investment slice
    * walletReceived : Wallet that will get the invested eth at the end of the crowdsale
    * tokenReceived : Address of the LBC token being sold
    * openingTimeReceived : Starting date of the ICO
    * closingtimeReceived : Ending date of the ICO
    * capReceived : Max amount of wei to be contributed
    * goalReceived : Funding goal
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
    }

    /*
    ** Initializes the contract address and affects addresses to their roles.
    */
    function init(
        address whitelistingAddress
    )
    public
    isNotInitialized
    onlyDeployingAddress
    {
        require(whitelistingAddress != address(0), "whitelistingAddress cannot be 0x");

        _whitelistingAddress = whitelistingAddress;
        initialized = true;

        emit InitializedContract(_msgSender(), whitelistingAddress);
    }

    function _forwardFunds() internal override(CrowdsaleMint, RefundablePostDeliveryCrowdsale) {
        RefundablePostDeliveryCrowdsale._forwardFunds();
    }

    function _preValidatePurchase(address beneficiary, uint256 weiAmount) internal override(TimedCrowdsale, CappedTimedCrowdsale) view {
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

    modifier onlyWhitelistingAddress() {
        require(_msgSender() == _whitelistingAddress, "HLBICO: caller does not have the Whitelisted role");
        _;
    }

    /*
    ** Checks if the contract hasn't already been initialized
    */
    modifier isNotInitialized() {
        require(initialized == false, "Contract is already initialized.");
        _;
    }

    /*
    ** Checks if the sender is the minter controller address
    */
    modifier onlyDeployingAddress() {
        require(msg.sender == _deployingAddress, "Only the deploying address can call this method.");
        _;
    }
}