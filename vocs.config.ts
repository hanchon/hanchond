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
          text: "Converter",
          collapsed: false,
          items: [
            {
              text: "Address",
              link: "/hanchond/converter/address",
            },
            {
              text: "Numbers",
              link: "/hanchond/converter/numbers",
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
          link: "/lib/requester/client",
          collapsed: false,
          items: [
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
      ],
    },
  ],
});
