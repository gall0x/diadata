# Oracle Documentation

{% content-ref url="deployed-contracts.md" %}
[deployed-contracts.md](deployed-contracts.md)
{% endcontent-ref %}

{% content-ref url="guest-quotations/coinmarketcap-quotations.md" %}
[coinmarketcap-quotations.md](guest-quotations/coinmarketcap-quotations.md)
{% endcontent-ref %}

{% content-ref url="guest-quotations/coingecko-quotations.md" %}
[coingecko-quotations.md](guest-quotations/coingecko-quotations.md)
{% endcontent-ref %}

{% content-ref url="defi-protocol-rates-and-states.md" %}
[defi-protocol-rates-and-states.md](defi-protocol-rates-and-states.md)
{% endcontent-ref %}

{% content-ref url="broken-reference" %}
[Broken link](broken-reference)
{% endcontent-ref %}

{% content-ref url="chart-points.md" %}
[chart-points.md](chart-points.md)
{% endcontent-ref %}

{% content-ref url="fiat-prices.md" %}
[fiat-prices.md](fiat-prices.md)
{% endcontent-ref %}

{% content-ref url="interest-rates.md" %}
[interest-rates.md](interest-rates.md)
{% endcontent-ref %}

{% content-ref url="crypto-assets.md" %}
[crypto-assets.md](crypto-assets.md)
{% endcontent-ref %}

{% content-ref url="guest-quotations/" %}
[guest-quotations](guest-quotations/)
{% endcontent-ref %}



Oracles are smart contract interfaces that bring data from an external source into a smart contract. Smart contracts run in an isolated environment in a virtual machine and can not "measure" any outside information by themselves. The only way to get external data into a smart contract is by executing a transaction with the external data as payload. With that data, a smart contract can perform calculations and operations that depend on this data, such as calculating an interest based on the reference interest rates published by a central bank.

DIA is capable of publishing financial data with such an oracle so that any smart contract can read and use it from within the Ethereum virtual machine. By generating verifiable oracle data, any user can use financial data in smart contracts. For each of our asset classes there is an example data set available in the oracle.

## Immutable Oracle Feeds

The DIA oracle data can be accessed in an Ethereum [smart contract](https://etherscan.io/address/0xD47FDf51D61c100C447E2D4747c7126F19fa23Ef). This smart contract holds mappings from an index to an asset's name, its price value, and – if applicable – supply data.

By using the read function `getCoinInfo(coin_name)` it is possible to retrieve this data. Along with the actual data there is a metadata field for the timestamp of the last update.

Updates are supplied by the DIA Oracle service that periodically supplies updates into the smart contract. Each update also generates an event, so that the latest updates can be seen in the event view on Etherscan.

## Example DApp: ECB FX Rates for Cryptos

As an example application, we implemented and deployed a [simple smart contract that converts crypto asset prices from our oracle from USD to EUR](https://etherscan.io/address/0xccb30bf12177705d41ac208802a6066482a76eaa).&#x20;

Call `getAssetEurRate()` with a crypto asset name as argument in order to request the current crypto asset price from our oracle which is then converted to EUR by using the ECB exchange rate for EUR-USD, also published in the oracle contract.  In the above link, you can use the "Read Contract" tab for easy web access. The rate is displayed in a fix comma format with five decimal digits. The source code of that example contract can be found at the "Code" tab. This shows how easy it is to integrate our oracle into any application: compile your application against the oracle interface and set the address where the oracle is deployed afterwards. Ideally, this should be done in a way that allows updates if the oracle address changes at any time in the future.

## Gas Estimation

All DIA oracles use the go-ethereum gas estimation to estimate the gas costs for an update contract call. To ensure timely delivery of the transaction, the gas estimate is scaled to 110% of the original estimate.
