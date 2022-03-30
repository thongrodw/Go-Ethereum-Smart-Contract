
//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

contract Greeter {
    string private greeting = "HelloWorld";

    function Greet() public view returns (string memory) {
        return greeting;
    }

    function SetGreeting(string memory _greeting) public {
        greeting = _greeting;
    }
}