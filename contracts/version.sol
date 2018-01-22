pragma solidity ^0.4.10;


contract Version {

    struct Metadata {
        string version; // Version of the contract
        string nodeVersion; // Version of the node program
        string monitorVersion; //Version of the monitor
        uint16 networkId; //Current network id
        uint256 lastUpdated; //Block number it was lastUpdated
    }

    address public owner;
    Metadata public metadata;

    function Version(string _version) public {
        owner = msg.sender; //To verify that Alastria admins deployed the contract
        metadata.version = _version;
    }

    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    function updateVersion(string _version, string _nodeVersion, string _monitorVersion) public onlyOwner {
        metadata.version = _version;
        metadata.nodeVersion = _nodeVersion;
        metadata.monitorVersion = _monitorVersion;
        metadata.lastUpdated = now;
    }

    function updateNetwork(uint16 _newtorkId) public onlyOwner {
        metadata.networkId = _newtorkId;
    }

}
