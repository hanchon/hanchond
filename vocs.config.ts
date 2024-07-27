import { defineConfig } from "vocs";

export default defineConfig({
  sidebar: [
    {
      text: "Home",
      link: "/",
    },
    {
      text: "Go Library",
      collapsed: false,
      items: [
        {
          text: "Requester",
          link: "/lib/requester/client",
          collapsed: true,
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
          text: "Smart Contracts",
          link: "/lib/smartcontract",
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
  description:
    "Hanchon's toolkit to avoid rewritting the same code one million times.",
  title: "Hanchond",
});
