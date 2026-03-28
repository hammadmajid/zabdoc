import { App } from "./app";
export { ZabdocContainer } from "./container";

export default {
  async fetch(request: Request, env: Env, ctx: ExecutionContext): Promise<Response> {
    return App.fetch(request, env, ctx);
  },
};
