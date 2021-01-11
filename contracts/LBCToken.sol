// SPDX-License-Identifier: MIT


pragma solidity ^0.7.0;

import "../libraries/openzeppelin/GSN/Context.sol";
import "./ERC20CappedUnburnable.sol";
import "./ERC20Unburnable.sol";
import "./ERC20PausableUnburnable.sol";

contract LBCToken is Context, ERC20CappedUnburnable {

    /*
    ** Global State
    */
    bool public initialized; // default : false

    /*
    ** Addresses
    */
    address public _deployingAddress; // should remain the same as deployer's address
    address public _pauserAddress; // should be ico's address then deployer's
    address public _minterAddress; // should be ico's address then poe's
    address public _reserveAddress; // should be deployer then humble reserve

    /*
     * Events
     */
    event InitializedContract(address minterAddress, address pauserAddress, address reserveAddress, address indexed changerAddress, uint256 initialSupply);
    event ChangedMinterAddress(address indexed minterAddress, address indexed changerAddress);
    event ChangedPauserAddress(address indexed pauserAddress, address indexed changerAddress);
    event ChangedReserveAddress(address indexed reserveAddress, address indexed changerAddress);


    constructor(
        string memory name,
        string memory symbol
    )
    ERC20Unburnable(name, symbol)
    ERC20CappedUnburnable(300000000)
    {
        _deployingAddress = msg.sender;
    }

    function init(
        address minterAddress,
        address pauserAddress,
        address reserveAddress,
        uint256 initialSupply
    )
    public
    isNotInitialized
    onlyDeployingAddress
    {
        require(pauserAddress != address(0), "_pauserAddress cannot be 0x");
        require(reserveAddress != address(0), "_pauserAddress cannot be 0x");

        _pauserAddress = pauserAddress;
        _reserveAddress = reserveAddress;

        _mint(_reserveAddress, initialSupply);

        if(reserveAddress != address(0)) {
            _minterAddress = minterAddress;
        }

        initialized = true;

        emit InitializedContract(minterAddress, pauserAddress, reserveAddress, _msgSender(), initialSupply);
    }

    function mint(address to, uint256 amount)
    public
    onlyMinterAddress
    virtual {
        _mint(to, amount);
    }

    function pause()
    public
    onlyPauserAddress
    virtual {
        _pause();
    }

    function unpause()
    public
    onlyPauserAddress
    virtual {
        _unpause();
    }

    function changeMinter(address newMinterAddress)
    public
    onlyDeployingAddress
    whenNotPaused
    {
        _minterAddress = newMinterAddress;
        emit ChangedMinterAddress(newMinterAddress, _msgSender());
    }

    /*
    ** Checks if the sender is the minter controller address.
    */
    modifier onlyDeployingAddress() {
        require(msg.sender == _deployingAddress, "Only the deploying address can call this method.");
        _;
    }

    /*
    ** Checks if the sender is the minter controller address.
    */
    modifier onlyMinterAddress() {
        require(msg.sender == _minterAddress, "Only the minter address can call this method.");
        _;
    }

    /*
    ** Checks if the sender is the pauser controller address.
    */
    modifier onlyPauserAddress() {
        require(msg.sender == _pauserAddress, "Only the pauser address can call this method.");
        _;
    }

    /*
    ** Checks if the contract hasn't already been initialized.
    */
    modifier isNotInitialized() {
        require(initialized == false, "Contract is already initialized.");
        _;
    }

    function _beforeTokenTransfer(address from, address to, uint256 amount) internal virtual override(ERC20CappedUnburnable) {
        super._beforeTokenTransfer(from, to, amount);
    }
}