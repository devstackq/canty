pragma solidity ^0.8.0;

contract VideoMonetization {
address public owner;
uint public adPrice = 1 ether;  // Цена за размещение рекламы

event AdPlaced(address indexed from, string adText, string adImage);

constructor() {
owner = msg.sender;
}

function placeAd(string memory adText, string memory adImage) public payable {
require(msg.value >= adPrice, "Insufficient payment");
emit AdPlaced(msg.sender, adText, adImage);

// Передача оплаты владельцу контракта
payable(owner).transfer(msg.value);
}

function updateAdPrice(uint newPrice) public {
require(msg.sender == owner, "Only owner can update ad price");
adPrice = newPrice;
}
}
abigen --bin=output/contract.bin --abi=output/VideoMonetization.abi --pkg=ads --out=VideoMonetization.go


