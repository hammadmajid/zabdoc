import { Hono } from "hono";
import { getContainer } from "@cloudflare/containers";
import { cors } from "hono/cors";
import { documentSchema, scrapeSchema } from "@repo/types"
import { validator } from 'hono/validator'
import { constructUrl } from "./utils";

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

App.post("/document", validator('json', (value, c) => {
  const parsed = documentSchema.safeParse(value);
  if (!parsed.success) {
    return c.text(`${parsed.error}`, 400)
  }
  return parsed.data
}), (c) => {
  const data = c.req.valid('json');
  const containerService = getContainer(c.env.ZabdocContainer); // uses ‘cf-singleton-container’ by default

  return containerService.containerFetch(constructUrl(c.req.path), {
    method: "POST",
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  })
});

App.post("/scrape", validator('json', (value, c) => {
  const parsed = scrapeSchema.safeParse(value);
  if (!parsed.success) {
    return c.text(`${parsed.error}`, 400)
  }
  return parsed.data
}), (c) => {
  const data = c.req.valid('json');
  const containerService = getContainer(c.env.ZabdocContainer); // uses ‘cf-singleton-container’ by default

  return containerService.containerFetch(constructUrl(c.req.path), {
    method: "POST",
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  }, 8080)
});

App.get();
