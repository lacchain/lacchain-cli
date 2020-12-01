pragma solidity 0.5.9;



contract NodeRulesList {

    enum NodeType{
        Boot,
        Validator,
        Writer,
        Observer,
        Other
    }

    // struct size = 82 bytes
    struct enode {
        bytes32 enodeHigh;
        bytes32 enodeLow;
        bytes16 ip;
        uint16 port;
        NodeType nodeType;
        bytes6 geoHash;
        string name;
        string organization;
        string did;
        bytes32 group;
    }

    enode[] public allowlist;
    mapping (uint256 => uint256) internal indexOf; //1-based indexing. 0 means non-existent
    mapping (bytes32 => bool) private allowGroups;

    function calculateKey(bytes32 _enodeHigh, bytes32 _enodeLow, bytes16 _ip, uint16 _port) internal pure returns(uint256) {
        return uint256(keccak256(abi.encodePacked(_enodeHigh, _enodeLow, _ip, _port)));
    }

    function size() internal view returns (uint256) {
        return allowlist.length;
    }

    function exists(bytes32 _enodeHigh, bytes32 _enodeLow, bytes16 _ip, uint16 _port) internal view returns (bool) {
        return indexOf[calculateKey(_enodeHigh, _enodeLow, _ip, _port)] != 0;
    }

    function groupConnectionAllowed(
        bytes32 sourceEnodeHigh,
        bytes32 sourceEnodeLow,
        bytes16 sourceEnodeIp,
        uint16 sourceEnodePort,
        bytes32 destinationEnodeHigh,
        bytes32 destinationEnodeLow,
        bytes16 destinationEnodeIp,
        uint16 destinationEnodePort) internal view returns (bool){
        if (exists(sourceEnodeHigh,sourceEnodeLow,sourceEnodeIp,sourceEnodePort) && exists(destinationEnodeHigh,destinationEnodeLow,destinationEnodeIp,destinationEnodePort)){    
            enode memory source = allowlist[indexOf[calculateKey(sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort)]-1];
            enode memory destination = allowlist[indexOf[calculateKey(destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)]-1];        

            return allowGroups[keccak256(abi.encodePacked(source.group, destination.group))];
        }
        else{
            return false;
        }
    }

    function _addConnectionAllowed(bytes32 groupSource, bytes32 groupDestination)internal returns(bool){
        allowGroups[keccak256(abi.encodePacked(groupSource, groupDestination))]=true;
        allowGroups[keccak256(abi.encodePacked(groupDestination, groupSource))]=true;
        return true;
    }

    function _removeConnection(bytes32 groupSource, bytes32 groupDestination)internal returns(bool){
        allowGroups[keccak256(abi.encodePacked(groupSource, groupDestination))]=false;
        allowGroups[keccak256(abi.encodePacked(groupDestination, groupSource))]=false;
        return true;
    }

    function add(bytes32 _enodeHigh, bytes32 _enodeLow, bytes16 _ip, uint16 _port, NodeType _nodeType, bytes6 _geoHash, string memory _name, string memory _organization, string memory _did, bytes32 _group) internal returns (bool) {
        uint256 key = calculateKey(_enodeHigh, _enodeLow, _ip, _port);
        if (indexOf[key] == 0) {
            indexOf[key] = allowlist.push(enode(_enodeHigh, _enodeLow, _ip, _port, _nodeType, _geoHash, _name, _organization, _did, _group));
            return true;
        }
        return false;
    }

    function remove(bytes32 _enodeHigh, bytes32 _enodeLow, bytes16 _ip, uint16 _port) internal returns (bool) {
        uint256 key = calculateKey(_enodeHigh, _enodeLow, _ip, _port);
        uint256 index = indexOf[key];

        if (index > 0 && index <= allowlist.length) { //1 based indexing
            //move last item into index being vacated (unless we are dealing with last index)
            if (index != allowlist.length) {
                enode memory lastEnode = allowlist[allowlist.length - 1];
                allowlist[index - 1] = lastEnode;
                indexOf[calculateKey(lastEnode.enodeHigh, lastEnode.enodeLow, lastEnode.ip, lastEnode.port)] = index;
            }

            //shrink array
            allowlist.length -= 1;
            indexOf[key] = 0;
            return true;
        }

        return false;
    }
}
