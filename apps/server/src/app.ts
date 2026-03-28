import { Hono } from "hono";
import { getContainer } from "@cloudflare/containers";

export const App = new Hono<{ Bindings: Env }>();

App.get("/health", (c) => {
  // Pass the raw request to container
  const containerService = getContainer(c.env.ZabdocContainer); // uses ‘cf-singleton-container’ by default
  return containerService.fetch(c.req.raw);
});

App.post("/generate", (c) => {
  // Pass the raw request to container
  const containerService = getContainer(c.env.ZabdocContainer); // uses ‘cf-singleton-container’ by default
  return containerService.fetch(c.req.raw);
});

App.post("/scrape", (c) => {
  // Pass the raw request to container
  const containerService = getContainer(c.env.ZabdocContainer); // uses ‘cf-singleton-container’ by default
  return containerService.fetch(c.req.raw);
});

App.get();
