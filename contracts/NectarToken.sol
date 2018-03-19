pragma solidity ^0.4.18;

import "zeppelin-solidity/contracts/token/ERC20/MintableToken.sol";


contract NectarToken is MintableToken {
    string public name = "Rektar";
    string public symbol = "RKT";
    uint8 public decimals = 18;

    bool public transfersEnabled = false;
    event TransfersEnabled();

    // Disable transfers until after the sale
    modifier whenTransfersEnabled() {
        require(transfersEnabled);
        _;
    }

    modifier whenTransfersNotEnabled() {
        require(!transfersEnabled);
        _;
    }

    function enableTransfers() onlyOwner whenTransfersNotEnabled public {
        transfersEnabled = true;
        TransfersEnabled();
    }

    function transfer(address to, uint256 value) public whenTransfersEnabled returns (bool) {
        return super.transfer(to, value);
    }

    function transferFrom(address from, address to, uint256 value) public whenTransfersEnabled returns (bool) {
        return super.transferFrom(from, to, value);
    }

    // Approves and then calls the receiving contract
    function approveAndCall(address _spender, uint256 _value, bytes _extraData) public returns (bool success) {
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);

        // call the receiveApproval function on the contract you want to be notified. This crafts the function signature manually so one doesn't have to include a contract in here just for this.
        // receiveApproval(address _from, uint256 _value, address _tokenContract, bytes _extraData)
        // it is assumed that when does this that the call *should* succeed, otherwise one would use vanilla approve instead.

        // solium-disable-next-line security/no-low-level-calls
        require(_spender.call(bytes4(bytes32(keccak256("receiveApproval(address,uint256,address,bytes)"))), msg.sender, _value, this, _extraData));
        return true;
    }
}
