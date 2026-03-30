import { Hono } from "hono";
import { getContainer } from "@cloudflare/containers";
import { cors } from "hono/cors";

export const App = new Hono<{ Bindings: Env }>();

App.use(
  "*",
  cors({
    origin: ["https://zabdoc.xyz", "http://localhost:5173"],
    allowMethods: ["GET", "POST", "OPTIONS"],
    allowHeaders: ["Content-Type", "Authorization"],
    credentials: true,
    maxAge: 43200, // 12 hours
  }),
);

App.get("/health", (c) => {
  // Pass the raw request to container
  const containerService = getContainer(c.env.ZabdocContainer); // uses ‘cf-singleton-container’ by default
  return containerService.fetch(c.req.raw);
});

App.post("/document", (c) => {
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
