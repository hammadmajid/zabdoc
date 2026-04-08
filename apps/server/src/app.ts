import { Hono } from "hono";
import { getContainer } from "@cloudflare/containers";
import { cors } from "hono/cors";
import { documentSchema, scrapeInitSchema, scrapedData } from "@repo/types";
import { validator } from "hono/validator";
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

App.post(
  "/scrape",
  validator("json", (value, c) => {
    const parsed = scrapeInitSchema.safeParse(value);
    if (!parsed.success) {
      return c.text(`${parsed.error}`, 400);
    }
    return parsed.data;
  }),
  async (c) => {
    const data = c.req.valid("json");
    const url = `https://${data.semester}zabdesk.szabist-${data.campus}.edu.pk`;

    const containerService = getContainer(c.env.ZabdocContainer); // uses ‘cf-singleton-container’ by default

    const response = await containerService.containerFetch(constructUrl(c.req.path), {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        username: data.username,
        password: data.password,
        url,
      }),
    });

    const parsed = await scrapedData.safeParseAsync(response.body);
    if (!parsed.success) {
      // if this doesn't work its on us
      return c.text(`${parsed.error}`, 500);
    }

    return c.json(parsed.data);
  },
);
