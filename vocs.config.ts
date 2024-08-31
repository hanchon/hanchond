import * as React from "react";

import { defineConfig } from "vocs";

export default defineConfig({
  title: "Hanchond",
  description:
    "Hanchon's toolkit to avoid rewritting the same code one million times.",
  head() {
    return React.createElement("script", {
      defer: true,
      "data-domain": "hanchond",
      src: "https://plausible.hanchon.me/js/script.js",
    });
  },
  sidebar: [
    {
      text: "Home",
      link: "/",
    },
    {
      text: "Hanchond",
      collapsed: false,
      items: [
        {
          text: "Playground",
          collapsed: false,
          items: [
            {
              text: "Utils",
              collapsed: true,
              items: [
                {
                  text: "Config",
                  link: "/hanchond/playground/config",
                },
                {
                  text: "Remove Data",
                  link: "/hanchond/playground/removeData",
                },
              ],
            },
            {
              text: "Binaries",
              collapsed: true,
              items: [
                {
                  text: "Build Evmosd",
                  link: "/hanchond/playground/buildEvmos",
                },
                {
                  text: "Build Gaiad",
                  link: "/hanchond/playground/buildGaiad",
                },
                {
                  text: "Build Hermes",
                  link: "/hanchond/playground/buildHermes",
                },
                {
                  text: "Build Solc",
                  link: "/hanchond/playground/buildSolc",
                },
                {
                  text: "List Binaries",
                  link: "/hanchond/playground/listBinaries",
                },
              ],
            },
            {
              text: "Chain",
              collapsed: true,
              items: [
                {
                  text: "Init Chain",
                  link: "/hanchond/playground/initChain",
                },
                {
                  text: "Start Chain",
                  link: "/hanchond/playground/startChain",
                },
                {
                  text: "Stop Chain",
                  link: "/hanchond/playground/stopChain",
                },
              ],
            },
            {
              text: "Nodes",
              collapsed: true,
              items: [
                {
                  text: "Start Node",
                  link: "/hanchond/playground/startNode",
                },
                {
                  text: "Get Node Info",
                  link: "/hanchond/playground/getNode",
                },
                {
                  text: "Stop Node",
                  link: "/hanchond/playground/stopNode",
                },
                {
                  text: "Change Binary Version",
                  link: "/hanchond/playground/changeBinaryVersion",
                },
              ],
            },
            {
              text: "Hermes",
              collapsed: true,
              items: [
                {
                  text: "Add a New Channel (only local networks)",
                  link: "/hanchond/playground/hermesAddChannel",
                },
                {
                  text: "Hermes Start",
                  link: "/hanchond/playground/hermesStart",
                },
                {
                  text: "Hermes Stop",
                  link: "/hanchond/playground/hermesStop",
                },
              ],
            },
            {
              text: "Relayer",
              collapsed: true,
              items: [
                {
                  text: "Add a Chain Configuration",
                  link: "/hanchond/playground/relayer/addChainConfig",
                },
                {
                  text: "Create a Channel",
                  link: "/hanchond/playground/relayer/createChannel",
                },
              ],
            },
            {
              text: "Transactions",
              collapsed: true,
              items: [
                {
                  text: "General Flags",
                  link: "/hanchond/playground/tx/flags",
                },
                {
                  text: "IBC Transfer",
                  link: "/hanchond/playground/tx/ibc",
                },
                {
                  text: "Vote",
                  link: "/hanchond/playground/tx/vote",
                },
                {
                  text: "Upgrade Proposal",
                  link: "/hanchond/playground/tx/upgrade",
                },
                {
                  text: "STRv1 Proposal",
                  link: "/hanchond/playground/tx/str",
                },
                {
                  text: "Rate Limit Proposal",
                  link: "/hanchond/playground/tx/ratelimit",
                },
                {
                  text: "Solidity",
                  collapsed: true,
                  items: [
                    {
                      text: "Compile a Contract",
                      link: "/hanchond/playground/tx/solidity/compile",
                    },
                    {
                      text: "Deploy a Contract",
                      link: "/hanchond/playground/tx/solidity/deploy",
                    },
                    {
                      text: "Deploy an ERC20",
                      link: "/hanchond/playground/tx/solidity/deployERC20",
                    },
                    {
                      text: "Call Contract View",
                      link: "/hanchond/playground/tx/solidity/contractView",
                    },
                    {
                      text: "Deploy UniswapV2",
                      link: "/hanchond/playground/tx/solidity/deployUniswapV2",
                    },
                    {
                      text: "ERC20 Transfer",
                      link: "/hanchond/playground/tx/solidity/erc20Transfer",
                    },
                  ],
                },
              ],
            },
            {
              text: "Queries",
              collapsed: true,
              items: [
                {
                  text: "General Flags",
                  link: "/hanchond/playground/queries/flags",
                },
                {
                  text: "Current Height",
                  link: "/hanchond/playground/queries/height",
                },
                {
                  text: "Bank Balances",
                  link: "/hanchond/playground/queries/balances",
                },
                {
                  text: "Transaction",
                  link: "/hanchond/playground/queries/transaction",
                },
                {
                  text: "ERC20",
                  collapsed: true,
                  items: [
                    {
                      text: "Supply",
                      link: "/hanchond/playground/queries/erc20/supply",
                    },
                    {
                      text: "Balance",
                      link: "/hanchond/playground/queries/erc20/balance",
                    },
                  ],
                },
                {
                  text: "Evmos",
                  collapsed: true,
                  items: [
                    {
                      text: "Token Pairs",
                      link: "/hanchond/playground/queries/evmos/tokenPairs",
                    },
                    {
                      text: "IBC Rate Limit",
                      link: "/hanchond/playground/queries/evmos/rateLimit",
                    },
                    {
                      text: "Smart-Contract Bytecode",
                      link: "/hanchond/playground/queries/evmos/ethCode",
                    },
                  ],
                },
              ],
            },
            {
              text: "Examples",
              collapsed: true,
              items: [
                {
                  text: "Spam the network with Web3 Txns (v0.4.0)",
                  link: "/hanchond/playground/examples/spamTheNetwork",
                },
                {
                  text: "Debug eth_code (v0.3.0)",
                  link: "/hanchond/playground/examples/debugEthCode",
                },
                {
                  text: "Debug eth_call (v0.2.0)",
                  link: "/hanchond/playground/examples/debugEthCall",
                },
                {
                  text: "v19.1.0 Upgrade (v0.1.0)",
                  link: "/hanchond/playground/examples/upgradev19",
                },
              ],
            },
          ],
        },
        {
          text: "Converter",
          collapsed: false,
          items: [
            {
              text: "Address",
              link: "/hanchond/convert/address",
            },
            {
              text: "Numbers",
              link: "/hanchond/convert/numbers",
            },
          ],
        },
      ],
    },
    {
      text: "Go Library",
      collapsed: false,
      items: [
        {
          text: "Requester",
          collapsed: false,
          items: [
            {
              text: "Client",
              link: "/lib/requester/client",
            },
            {
              text: "Web3 Requests",
              link: "/lib/requester/web3",
            },
            {
              text: "Cosmos Requests",
              link: "/lib/requester/cosmos",
            },
            {
              text: "Tendermint Requests",
              link: "/lib/requester/tendermint",
            },
          ],
        },
        {
          text: "Tx Builder",
          collapsed: false,
          items: [
            {
              text: "Builder",
              link: "/lib/txbuilder/txbuilder",
            },
            {
              text: "Transaction",
              link: "/lib/txbuilder/transaction",
            },
            {
              text: "Wallet",
              link: "/lib/txbuilder/wallet",
            },
            {
              text: "Simple Web3 Wallet",
              link: "/lib/txbuilder/simple_wallet",
            },
            {
              text: "Mnemonic",
              link: "/lib/txbuilder/mnemonic",
            },
            {
              text: "Contract",
              link: "/lib/txbuilder/contract",
            },
          ],
        },
        {
          text: "Smart Contracts",
          link: "/lib/smartcontract",
          collapsed: false,
          items: [
            {
              text: "ABI",
              link: "/lib/smartcontract/ABI",
            },
            {
              text: "ERC20",
              link: "/lib/smartcontract/erc20",
            },
          ],
        },
        {
          text: "Converter",
          link: "/lib/converter",
        },
        {
          text: "Proto Encoder",
          link: "/lib/encoder",
        },
        {
          text: "Examples",
          collapsed: false,
          items: [
            {
              text: "Spam the network with Web3 Txns (v0.4.0)",
              link: "/hanchond/playground/examples/spamTheNetwork",
            },
            {
              text: "ERC20 Example (v0.1.0)",
              link: "/lib/examples/erc20",
            },
          ],
        },
      ],
    },
  ],
});
