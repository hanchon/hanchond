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
  ],
  description:
    "Hanchon's toolkit to avoid rewritting the same code one million times.",
  title: "Vivi",
});
