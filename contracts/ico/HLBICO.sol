// SPDX-License-Identifier: MIT

pragma solidity ^0.7.0;

import "./CappedTimedCrowdsale.sol";
import "./RefundPostdevCrowdsale.sol";


/**
**  ICO Contract for the LBC crowdsale
*/
contract HLBICO is CappedTimedCrowdsale, RefundablePostDeliveryCrowdsale {

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
}