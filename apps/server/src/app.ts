import { Hono } from "hono";
import { getContainer } from "@cloudflare/containers";

export const App = new Hono<{ Bindings: Env }>();

App.get("/container", (c) => {
  const containerService = getContainer(c.env.ZabdocContainer, "instance-1");
  return containerService.fetch(c.req.raw);
});

// App.get("/container/load-balance", async (c) => {
//   const containerService = await loadBalance(c.env.ZabdocContainer, 3);
//   return await containerService.fetch("https://container/load-balance");
// });
