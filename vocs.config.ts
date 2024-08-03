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
              items: [
                {
                  text: "Config",
                  link: "/hanchond/playground/config",
                },
                {
                  text: "Remove Data",
                  link: "/hanchond/playground/removeData",
                },
                {
                  text: "List Binaries",
                  link: "/hanchond/playground/listBinaries",
                },
              ],
            },
            {
              text: "Evmos",
              items: [
                {
                  text: "Build Evmos",
                  link: "/hanchond/playground/buildEvmos",
                },
                {
                  text: "Init Genesis",
                  link: "/hanchond/playground/initGenesis",
                },
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
              items: [
                {
                  text: "Build Hermes",
                  link: "/hanchond/playground/buildHermes",
                },
                {
                  text: "Add a New Channel",
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
              text: "Transactions",
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
              ],
            },
            {
              text: "Queries",
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
                  items: [
                    {
                      text: "Token Pairs",
                      link: "/hanchond/playground/queries/evmos/tokenPairs",
                    },
                    {
                      text: "IBC Rate Limit",
                      link: "/hanchond/playground/queries/evmos/rateLimit",
                    },
                  ],
                },
              ],
            },
            {
              text: "Examples",
              items: [
                {
                  text: "v19.1.0 Upgrade",
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
          items: [
            {
              text: "ERC20 Example",
              link: "/lib/examples/erc20",
            },
          ],
        },
      ],
    },
  ],
});
