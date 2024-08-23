import { FastifyInstance, FastifyServerOptions } from "fastify";

// Health check route
export default function (
  instance: FastifyInstance,
  _opts: FastifyServerOptions,
  done: any,
) {
  instance.route({
    method: "GET",
    url: "/health",
    schema: {
      response: {
        200: {
          type: "object",
          properties: {
            message: { type: "string" },
          },
        },
      },
    },
    handler: async (req, _) => {
      req.log.info("Health check");
      return { message: "Healthy!" };
    },
  });
  done();
}
