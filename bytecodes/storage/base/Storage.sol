// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <0.8.0;

contract Storage {

    mapping (address => uint256) private addresses;

    function add(address _addr, uint256 _value) public returns(bool) {
        addresses[_addr] = _value;
        return true;
    }

    function get(address _addr) public view returns(uint256) {
        uint256 value = addresses[_addr];
        return value;
    }

}
