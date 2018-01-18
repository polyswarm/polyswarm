pragma solidity ^0.4.18;

import "zeppelin-solidity/contracts/token/MintableToken.sol";
import "zeppelin-solidity/contracts/token/ERC827/ERC827Token.sol";


contract NectarToken is MintableToken, ERC827Token {
    string public name = "Nectar";
    string public symbol = "NCT";
    uint8 public decimals = 18;
}
